KIND_CLUSTER_NAME := k8s-playground

.DEFAULT_GOAL := help

help: ## List commands
	@echo '使い方: make [Target]'
	@grep -E '^[a-zA-Z_-]+:.*?## .*$$' $(MAKEFILE_LIST) \
		| sort \
		| awk 'BEGIN {FS = ":.*?## "}; {printf "  \033[36m%-20s\033[0m %s\n", $$1, $$2}'

init: ## Initialize project
	mise install
	${MAKE} -C backend init
	${MAKE} -C frontend init

setup: ## Setup project
	$(MAKE) init
	$(MAKE) gen

gen: ## Generate code
	$(MAKE) -C proto gen
	$(MAKE) -C backend sqlc-gen

backend-dev: ## Run backend in development mode
	make -C backend dev

frontend-dev: ## Run frontend in development mode
	make -C frontend dev

db-dev: ## Run database in development mode
	docker compose --env-file backend/.env up -d db

backend-docker-build: ## Build backend Docker image
	docker build -t todo-api:latest backend/

frontend-docker-build: ## Build frontend Docker image
	docker build -t todo-spa:latest frontend/ --build-arg VITE_API_BASE_URL="http://localhost:3001"

backend-lint: ## Lint backend code
	$(MAKE) -C backend lint

frontend-lint: ## Lint frontend code
	$(MAKE) -C frontend lint

proto-lint: ## Lint proto files
	$(MAKE) -C proto lint

backend-test: ## Run backend tests
	$(MAKE) -C backend test

frontend-test: ## Run frontend tests
	$(MAKE) -C frontend test

backend-build: ## Build backend
	$(MAKE) -C backend build

frontend-build: ## Build frontend
	$(MAKE) -C frontend build

proto-build: ## Build proto files
	$(MAKE) -C proto build

compose-up: ## Run all services with Docker Compose
	docker compose --env-file backend/.env up -d

compose-down: ## Stop all services with Docker Compose
	docker compose --env-file backend/.env down

db-migrate: ## Run database migrations
	docker compose --env-file backend/.env run --rm db-migrate

k8s-create-cluster: ## Create a Kubernetes cluster using kind
	kind create cluster --config cluster.yaml --name $(KIND_CLUSTER_NAME)

k8s-create-namespace: ## Create a Kubernetes namespace
	kubectl apply -f k8s/manifests/namespace.yaml

k8s-load-backend-image: ## Load backend Docker image into the kind cluster
	kind load docker-image todo-api:latest --name $(KIND_CLUSTER_NAME)

k8s-load-frontend-image: ## Load frontend Docker image into the kind cluster
	kind load docker-image todo-spa:latest --name $(KIND_CLUSTER_NAME)

k8s-load-images: ## Load both backend and frontend Docker images into the kind cluster
	$(MAKE) k8s-load-backend-image
	$(MAKE) k8s-load-frontend-image

k8s-create-secret: ## Create a Kubernetes secret for backend credentials
	kubectl create secret generic backend-credentials --from-env-file=backend/.env -n todo-api

k8s-apply-all: ## Apply all Kubernetes configurations
	$(MAKE) k8s-apply-manifests
	$(MAKE) k8s-helmfile-apply

k8s-apply-manifests: ## Apply Kubernetes manifests
	kubectl apply -Rf k8s/manifests

k8s-helmfile-apply: ## Apply Helm charts using Helmfile
	helmfile apply -f k8s/helm/helmfile.yaml

k8s-delete-cluster: ## Delete the kind Kubernetes cluster
	kind delete cluster --name $(KIND_CLUSTER_NAME)

k8s-init: ## Initialize Kubernetes cluster and deploy application
	$(MAKE) k8s-create-cluster
	$(MAKE) k8s-create-namespace
	$(MAKE) k8s-create-secret
	$(MAKE) backend-docker-build
	$(MAKE) frontend-docker-build
	$(MAKE) k8s-load-images
	@if helm plugin list | grep -q '^diff[[:space:]]'; then \
		echo "helm-diff plugin already installed"; \
	else \
		helm plugin install https://github.com/databus23/helm-diff --verify=false; \
	fi
	$(MAKE) k8s-apply-manifests
	$(MAKE) k8s-helmfile-apply

