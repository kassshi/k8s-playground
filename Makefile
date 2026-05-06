run-backend:
	make -C backend run

dev-backend:
	make -C backend dev

dev-frontend:
	make -C frontend dev

up:
	docker compose --env-file backend/.env up -d

down:
	docker compose --env-file backend/.env down

db-migrate:
	docker compose --env-file backend/.env run --rm db-migrate

gen-protobuf:
	cd proto && buf generate
