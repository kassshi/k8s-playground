KIND_CLUSTER_NAME := k8s-playground

init:
	mise install
	${MAKE} -C backend init
	${MAKE} -C frontend init

setup:
	$(MAKE) init
	$(MAKE) gen

gen:
	$(MAKE) -C proto gen
	$(MAKE) -C backend sqlc-gen

backend-dev:
	make -C backend dev

frontend-dev:
	make -C frontend dev

db-dev:
	docker compose --env-file backend/.env up -d db

backend-docker-build:
	docker build -t todo-api:latest backend/

frontend-docker-build:
	docker build -t todo-spa:latest frontend/ --build-arg VITE_API_BASE_URL="http://localhost:3001"

backend-lint:
	$(MAKE) -C backend lint

frontend-lint:
	$(MAKE) -C frontend lint

proto-lint:
	$(MAKE) -C proto lint

backend-test:
	$(MAKE) -C backend test

frontend-test:
	$(MAKE) -C frontend test

backend-build:
	$(MAKE) -C backend build

frontend-build:
	$(MAKE) -C frontend build

proto-build:
	$(MAKE) -C proto build

compose-up:
	docker compose --env-file backend/.env up -d

compose-down:
	docker compose --env-file backend/.env down

db-migrate:
	docker compose --env-file backend/.env run --rm db-migrate

k8s-create-cluster:
	kind create cluster --config cluster.yaml --name $(KIND_CLUSTER_NAME)

k8s-create-namespace:
	kubectl apply -f k8s/manifests/namespace.yaml

k8s-load-backend-image:
	kind load docker-image todo-api:latest --name $(KIND_CLUSTER_NAME)

k8s-load-frontend-image:
	kind load docker-image todo-spa:latest --name $(KIND_CLUSTER_NAME)

k8s-load-images:
	$(MAKE) k8s-load-backend-image
	$(MAKE) k8s-load-frontend-image

k8s-create-secret:
	kubectl create secret generic backend-credentials --from-env-file=backend/.env -n todo-api

k8s-apply:
	kubectl apply -Rf k8s/manifests/

k8s-helmfile-apply:
	helmfile apply -f k8s/helm/helmfile.yaml

k8s-delete-cluster:
	kind delete cluster --name $(KIND_CLUSTER_NAME)

k8s-init:
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
	$(MAKE) k8s-helmfile-apply
	$(MAKE) k8s-apply

