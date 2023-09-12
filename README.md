# Go Fiber Boilerplate

A Starter project with Golang and Fiber

```
routes -> handler -> service -> repository -> database
```

In folder infrastructure you can also add other methods such as gRPC, GraphQL, etc. </br>
You can add with folder name grpc, graphql, etc. </br>
</br>
In folder infrastructure you can also add other connection like aws, sendgrid, cloudinary, etc. </br>
You can add with folder name aws, sendgrid, cloudinary, etc. </br>

# Getting Started

## Prerequisites

- You need to install [Golang](https://golang.org/doc/install) and [Docker](https://docs.docker.com/get-docker/) on your machine.
- You need to install Makefile on your machine. [Makefile](https://sourceforge.net/projects/mingw/)
- You need to install [PostgreSQL](https://www.postgresql.org/download/) on your machine.
- You need to install [Postman](https://www.postman.com/downloads/) on your machine.
- You need to install [Git](https://git-scm.com/downloads) on your machine.
- You need to install [Visual Studio Code](https://code.visualstudio.com/download) on your machine.

## Installation

Clone the repository and run the following commands:

```
git clone https://github.com/saufiroja/go-fiber-boilerplate.git
```

Go to the project directory:

```
cd go-fiber-boilerplate
```

Copy the .env_example file to .env:

```
copy .env.example .env
```

Install dependencies:

```
go mod tidy
```

Run the project:

```
make run
```

# Tech Stack

- Golang (Go Programming Language)
- Fiber (Web Framework)
- JWT (JSON Web Token)
- Docker
- PostgreSQL
- Validator (go-playground/validator)
- GORM (ORM)
- UUID (Universally Unique Identifier)

# Structure Directory

An example of implementing a hexagonal architecture backend using golang.

```
.
├── .github/
│   └── workflows/
│       └── go.yaml
├── app/
│   └── main.go
├── config/
│   ├── app.go
│   ├── config.go
│   ├── fiber.go
│   └── postgres.go
├── infrastructure/
│   ├── database/
│   │   ├── mongo.go
│   │   └── postgres.go
│   └── http/
│       ├── handler/
│       │   ├── auth/
│       │   │   ├── auth_handler.go
│       │   │   └── routes.go
│       │   └── user/
│       │       ├── user_handler.go
│       │       └── routes.go
│       ├── middlewares/
│       │   └── jwt_middlewares.go
│       ├── routes/
│       │   └── routes.go
│       └── server/
│           └── fiber.go
├── interfaces/
│   ├── auth_interfaces.go
│   └── user_interfaces.go
├── models/
│   ├── dto/
│   │   └── user_dto.go
│   └── entity/
│       └── user_entity.go
├── repository/
│   └── postgres/
│       ├── auth/
│       │   └── auth_repository.go
│       └── user/
│           └── user_repository.go
├── service/
│   ├── auth/
│   │   └── auth_service.go
│   └── user/
│       └── user_service.go
├── utils/
│   ├── generate_token.go
│   ├── handler_error.go
│   └── password.go
├── .env
├── .env.example
├── .gitignore
├── docker-compose.yaml
├── Dockerfile
├── go.mod
├── go.sum
├── LICENCE
├── Makefile
└── README.md
```

# Example or Usage

## Interface

```go

type UserRepository interface {
	InsertUser(user *dto.Register) error
	FindUserByEmail(email string) (*entity.User, error)
}

type AuthService interface {
	Register(user *dto.Register) error
	Login(user *dto.Login) (*dto.LoginResponse, error)
}

type NewAuthHandler interface {
	Register(c *fiber.Ctx) error
	Login(c *fiber.Ctx) error
}

```

## Repository

```go

type authRepository struct {
	DB *gorm.DB
}

func NewAuthRepository(db *gorm.DB) interfaces.AuthRepository {
	return &authRepository{
		DB: db,
	}
}

```

## Service

```go
type authService struct {
	repoAuth interfaces.UserRepository
	validate *validator.Validate
}

func NewAuthService(repoAuth interfaces.AuthRepository) interfaces.AuthService {
	return &authService{
		repoAuth: repoAuth,
		validate: validator.New(),
	}
}
```

## Handler

```go
type authHandler struct {
	authService interfaces.AuthService
}

func NewAuthHandler(authService interfaces.AuthService) interfaces.NewAuthHandler {
	return &authHandler{
		authService: authService,
	}
}
```
