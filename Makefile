KIND_CLUSTER_NAME := k8s-playground

run-backend:
	make -C backend run

dev-backend:
	make -C backend dev

dev-frontend:
	make -C frontend dev

build-backend:
	docker build -t golang-practice-api:latest backend/
	$(MAKE) k8s-load-images

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

k8s-load-images:
	kind load docker-image golang-practice-api:latest --name $(KIND_CLUSTER_NAME)

k8s-create-secret:
	kubectl create secret generic backend-credentials --from-env-file=backend/.env -n todo-api

k8s-apply:
	kubectl apply -Rf k8s/manifests/

k8s-helmfile-apply:
	helmfile apply -f k8s/helm/helmfile.yaml

k8s-init:
	$(MAKE) k8s-create-cluster
	$(MAKE) k8s-create-namespace
	$(MAKE) k8s-create-secret
	$(MAKE) k8s-load-images
	$(MAKE) k8s-apply
	helm plugin install https://github.com/databus23/helm-diff --verify=false
