# Go Fiber Boilerplate

A Starter project with Golang and Fiber

# Boilerplate structure

An example of implementing a hexagonal architecture backend using golang.

```
.
├── config/
│   ├── app.go
│   ├── config.go
│   ├── fiber.go
│   └── postgres.go
├── dto/
│   └── user_dto.go
├── entity/
│   └── user_entity.go
├── infrastructure/
│   ├── database/
│   │   ├── mongo/
│   │   │   └── mongo.go
│   │   └── postgres/
│   │       └── postgres.go
│   └── http/
│       ├── controllers/
│       │   └── user_controllers.go
│       ├── middlewares/
│       │   └── jwt_middlewares.go
│       ├── routes/
│       │   ├── auth_routes.go
│       │   └── user_routes.go
│       └── server/
│           └── server.go
├── interfaces/
│   └── user_interfaces.go
├── repository/
│   ├── auth/
│   │   └── auth_repository_postgres.go
│   └── user/
│       └── user_repository_postgres.go
├── service/
│   ├── auth/
│   │   └── auth_service.go
│   └── user/
│       └── user_service.go
├── utils/
│   ├── password.go
│   ├── handlerError.go
│   └── generate_token.go
├── .env
├── .env_example
├── .gitignore
├── Dockerfile
├── README.md
├── docker-compose.yaml
├── go.mod
├── go.sum
└── main.go
```
