.PHONY: help
help:
	@grep -E '^[a-zA-Z_-]+:.*?## .*$$' $(firstword $(MAKEFILE_LIST)) | awk 'BEGIN {FS = ":.*?## "}; {printf "\033[36m%-30s\033[0m %s\n", $$1, $$2}'

start: ## Start the app containers
	docker-compose up --build -d

init: ## refresh the env file, generate public and private RSA keys and start the app
	make copy-dist-files
	make generate-keys
	make start

logs: ## Open logs for docker
	docker-compose logs -f

generate-keys: ## Generate public and private RSA keys
	openssl genrsa -des3 -out private.pem 2048
	openssl rsa -in private.pem -outform PEM -pubout -out public.pem

copy-dist-files: ## Copy dist files
	cp .env.dist .env

stop: ## stop all running containers of the application
	docker-compose stop

paf: ## Delete all docker containers and volumes on your computer
	docker system prune --volumes -a -f