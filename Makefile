run-backend:
	make -C backend run

dev-backend:
	make -C backend dev

dev-frontend:
	make -C frontend dev

build-backend:
	docker build -t golang-practice-api:latest backend/ && kind load docker-image golang-practice-api:latest

up:
	docker compose --env-file backend/.env up -d

down:
	docker compose --env-file backend/.env down

db-migrate:
	docker compose --env-file backend/.env run --rm db-migrate

gen-protobuf:
	cd proto && buf generate

k8s-create-secret:
	kubectl create secret generic backend-credentials --from-env-file=backend/.env -n todo-api

k8s-apply:
	kubectl apply -Rf k8s/
