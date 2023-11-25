build:
	CGO_ENABLED=0 GOOS=linux go build -o bell
	docker compose up --build

remove:
	docker stop bell-service-postgres-1
	docker stop bell-service-server-1
	docker rm bell-service-postgres-1
	docker rm bell-service-server-1

