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
│   ├── auth
│   │   └── auth_controller.go
│   ├── user
│   │   └── user_controller.go
├── dto
│   └── user_dto.go
├── entity
│   └── user_entity.go
├── infrastructure
│   ├── database
│   │   ├── mongo
│   │   │   └── mongo.go
│   │   └── postgres
│   │       └── postgres.go
│   └── http
│       ├── middleware
│       │   └── jwt_middleware.go
│       ├── routes
│       │   ├── auth_routes.go
│       │   └── user_routes.go
|       └── server
│           └── server.go
├── interfaces
│   └── user_interfaces.go
├── repository
│   ├── auth
│   │   └── auth_repository.go
│   ├── user
│       └── user_repository.go
├── service
│   ├── auth
│   │   └── auth_service.go
│   ├── user
│       └── user_service.go
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
