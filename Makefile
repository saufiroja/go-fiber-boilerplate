run:
	go run app/main.go

docker-up:
	docker-compose up -d

docker-down:
	docker-compose down