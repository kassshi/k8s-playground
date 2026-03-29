include .env
DB_URL=postgres://$(DB_USER):$(DB_PASSWORD)@localhost:5432/$(DB_NAME)?sslmode=disable


run:
	go run cmd/server/main.go
migrate-up:
	migrate -path db/migrations -database "$(DB_URL)" up
migrate-down:
	migrate -path db/migrations -database "$(DB_URL)" down 1
migrate-force: 
	migrate -path db/migrations -database "$(DB_URL)" force $(VERSION)
migrate-version: 
	migrate -path db/migrations -database "$(DB_URL)" version 
migrate-create: 
	migrate create -ext sql -dir db/migrations -seq $(NAME)
