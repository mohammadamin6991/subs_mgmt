PROJECT_NAME=subs_mgmt
AUTH_BINARY=auth
GATEWAY_BINARY=gateway
SVC_S3_BINARY=svcS3
SUBSCRIPTION_BINARY=subscription

## up: starts all containers in the background without forcing build
up:
	@echo "Starting Docker images..."
	docker-compose -p subs_mgmt up -d
	@echo "Docker images started!"

## up_build: stops docker-compose (if running), builds all projects and starts docker compose
up_build: build_svc_s3 build_subscription build_auth build_gateway
	@echo "Stopping docker images (if running...)"
	docker-compose -p subs_mgmt down

	@echo "Building (when required) and starting docker images..."
	docker-compose -p subs_mgmt up --build -d
	@echo "Docker images built and started!"

## down: stop docker compose
down:
	@echo "Stopping docker compose..."
	docker-compose -p subs_mgmt down
	@echo "Done!"

# build_auth: builds the broker binary as a linux executable
build_auth:
	@echo "Building auth binary..."
	cd ../authentication && env GOOS=linux CGO_ENABLED=0 go build -o bin/${AUTH_BINARY} ./cmd/api
	@echo "Done!"

# build_auth: builds the broker binary as a linux executable
build_gateway:
	@echo "Building auth binary..."
	cd ../gateway && env GOOS=linux CGO_ENABLED=0 go build -o bin/${GATEWAY_BINARY} ./cmd/api
	@echo "Done!"

## build_front: builds the frone end binary
build_subscription:
	@echo "Building front end binary..."
	cd ../subscription && env CGO_ENABLED=0 go build -o bin/${SUBSCRIPTION_BINARY} ./cmd/api
	@echo "Done!"

## build_broker: builds the broker binary as a linux executable
build_svc_s3:
	@echo "Building svc-s3 binary..."
	cd ../svc-s3 && env GOOS=linux CGO_ENABLED=0 go build -o bin/${SVC_S3_BINARY} ./cmd/api
	@echo "Done!"


migrate_up:
	@echo "Migrating Up..."
	docker-compose -p subs_mgmt up -d postgres
	migrate -path ../authentication/db/migration -database postgres://postgres:postgres@127.0.0.1:5432/subs_auth?sslmode=disable up
	docker-compose -p subs_mgmt down postgres
	@echo "Done!"

migrate_down:
	@echo "Migrating Down..."
	docker-compose -p subs_mgmt up -d postgres
	migrate -path ../authentication/db/migration -database postgres://postgres:postgres@localhost:5432/subs_auth?sslmode=disable down -all
	docker-compose -p subs_mgmt down postgres
	@echo "Done!"

migrate_clean:
	@echo "Cleaning Migration..."
	docker-compose -p subs_mgmt up -d postgres
	sleep 1
	migrate -path ../authentication/db/migration -database postgres://postgres:postgres@localhost:5432/subs_auth?sslmode=disable down -all
	migrate -path ../authentication/db/migration -database postgres://postgres:postgres@127.0.0.1:5432/subs_auth?sslmode=disable up
