include .env

# ==================================================================================== #
# HELPERS
# ==================================================================================== #

## The help rule is only on linux
.PHONY: help
help:
	@echo 'Usage:'
	@sed -n 's/^##//p' ${MAKEFILE_LIST} | column -t -s ':' |  sed -e 's/^/ /'

.PHONY: confirm
confirm:
	@echo -n 'Are you sure? [y/N] ' && read ans && [ $${ans:-N} = y ]

# ==================================================================================== #
# DEVELOPMENT
# ==================================================================================== #

## run/web: run the cmd/web application
.PHONY: run/web
run/web:
	@go run ./cmd/api -db-dsn="${DB_USER}:${DB_PASSWORD}@tcp(localhost:${DB_PORT})/${DB_DATABASE}?parseTime=true"

.PHONY: db/connect
db/connect:
	mysql -h 127.0.0.1 -P ${DB_PORT} -u ${DB_USER} -p${DB_PASSWORD} ${DB_DATABASE}

.PHONY: db/connect/root
db/connect/root:
	mysql -h 127.0.0.1 -P ${DB_PORT} -u root -p${DB_ROOT_PASSWORD} ${DB_DATABASE}

.PHONY: db/migrations/new
db/migrations/new:
	@echo 'Creating migration files for ${name}...'
	migrate create -seq -ext=.sql -dir=./migrations ${name}

.PHONY: db/migrations/up
db/migrations/up: confirm
	@echo 'Running up migrations...'
	migrate -path ./migrations -database "${DB_DSN}" up


.PHONY: db/migrations/down
db/migrations/down: confirm
	@echo 'Running down migrations...'
	@read -p "How many steps to rollback? " steps; \
	migrate -path ./migrations -database "${DB_DSN}" down $$steps
