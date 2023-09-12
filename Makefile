dev:
	go run ./app/main.go

docker-up:
	docker-compose up -d

docker-down:
	docker-compose down

build-app:
	go build -o bin/app app/main.go

test:
	mkdir -p coverage
	go test -v -coverprofile ./coverage/cover.out ./...
	go tool cover -html=./coverage/cover.out -o ./coverage/cover.html
	open ./coverage/cover.html