
docker_run:
	MIGRATOR_BUILD_PLATFORM=linux/amd64 docker compose up --build --remove-orphans

docker_down:
	docker compose down

swagger:
	# go install github.com/swaggo/swag/cmd/swag@latest
	swag init -g ./cmd/alchemy/main.go
	swagger generate markdown --output ./documentation/api.md -f docs/swagger.yaml