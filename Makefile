test: docker-up-test
	docker-compose -f docker-compose-test.yaml down -v
	docker-compose -f docker-compose-test.yaml --env-file .env up -dker-compose -f docker-compose-test.yaml up -d --env-file .env.test
	go clean -testcache && go test -v ./test/e2e/...
	docker-compose -f docker-compose-test.yaml down -v

dev:
	go run main.go

docker-up-test:
	docker-compose -f docker-compose-test.yaml up -d

docker-down-test:
	docker-compose -f docker-compose-test.yaml down

docker-up:
	docker-compose up -d

docker-down:
	docker-compose down

build-app:
	go build -o bin/app app/main.go

# test:
# 	mkdir -p coverage
# 	go test -v -coverprofile ./coverage/cover.out ./...
# 	go tool cover -html=./coverage/cover.out -o ./coverage/cover.html
# 	open ./coverage/cover.html
