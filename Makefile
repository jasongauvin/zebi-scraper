run: ## Start the app containers
	docker-compose up --build -d

init: ## refresh the env file, generate public and private RSA keys and start the app
	cp .env.dist .env
	openssl genrsa -des3 -out private.pem 2048
	openssl rsa -in private.pem -outform PEM -pubout -out public.pem
	make run
	
stop: ## stop all running containers of the application
	docker-compose stop