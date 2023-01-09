# Go Fiber Boilerplate

A Starter project with Golang and Fiber

# Boilerplate structure

```
.
├── config
│   ├── config.go
│   ├── app.go
│   ├── fiber.go
│   └── postgres.go
├── controllers
│   └── user_controllers.go
├── database
│   └── postgres
|       └──postgres.go
├── dto
│   └── user_dto.go
├── entity
│   └── user_entity.go
├── infrastructure
│    ├── routers
│      └── user_routers.go
│    ├── server
│        └── fiber.go
│    ├── middlewares
│        └── jwt_middlewares.go
├── interfaces
│   └── user_interfaces.go
├── repository
│   └── user_repository.go
├── service
│   └── user_service.go
├── utils
│    ├── password.go
│    └── handlerError.go
├── .env
├── .env.example
├── .gitignore
├── Dockerfile
├── docker-compose.yaml
├── README.md
├── main.go
├── go.mod
├── go.sum
```
