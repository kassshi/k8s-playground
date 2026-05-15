KIND_CLUSTER_NAME := k8s-playground

run-backend:
	make -C backend run

dev-backend:
	make -C backend dev

dev-frontend:
	make -C frontend dev

build-backend:
	docker build -t todo-api:latest backend/
	$(MAKE) k8s-load-backend-image

build-frontend:
	docker build -t todo-spa:latest frontend/ --build-arg VITE_API_BASE_URL="http://localhost:3001"
	$(MAKE) k8s-load-frontend-image

up:
	docker compose --env-file backend/.env up -d

down:
	docker compose --env-file backend/.env down

db-migrate:
	docker compose --env-file backend/.env run --rm db-migrate

gen-protobuf:
	cd proto && buf generate

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
	$(MAKE) k8s-load-images
	@if helm plugin list | grep -q '^diff[[:space:]]'; then \
		echo "helm-diff plugin already installed"; \
	else \
		helm plugin install https://github.com/databus23/helm-diff --verify=false; \
	fi
	$(MAKE) k8s-helmfile-apply
	$(MAKE) k8s-apply
