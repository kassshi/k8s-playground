init:
	if [ ! -f .env ]; then \
		cp .env.example .env; \
	fi

dev:
	air

sqlc-gen:
	sqlc generate
