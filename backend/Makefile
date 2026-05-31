init:
	if [ ! -f .env ]; then \
		cp .env.example .env; \
	fi

dev:
	air

sqlc-gen:
	sqlc generate

lint:
	golangci-lint run

test:
	go test -v ./...

build:
	go build -o ./tmp/server ./cmd/server

