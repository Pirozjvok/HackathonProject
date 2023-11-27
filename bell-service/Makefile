push:
	git add .
	git commit -m "COMMIT_NAME"
	git push -u origin golang

merge:
	git checkout main
	git merge golang
	git push -u origin main

build:
	CGO_ENABLED=0 GOOS=linux go build -o bell
	docker compose up --build

remove:
	docker stop bell-service-postgres-1
	docker stop bell-service-server-1
	docker rm bell-service-postgres-1
	docker rm bell-service-server-1

