run:
	go run main.go

docker-up:
	docker-compose up -d

docker-down:
	docker-compose down

build-app:
	go build -o bin/app app/main.go