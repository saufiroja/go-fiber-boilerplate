# Go Fiber Boilerplate

A Starter project with Golang and Fiber

```
routes -> handler -> service -> repository -> database
```

In folder infrastructure you can also add other methods such as gRPC, GraphQL, etc. </br>
You can add with folder name grpc, graphql, etc. </br>
In folder infrastructure you can also add other connection like aws, sendgrid, cloudinary, etc. </br>
You can add with folder name aws, sendgrid, cloudinary, etc. </br>
In Folder utils/constants you can add error code and validation message. </br>
In Folder utils you can add helper function like generate token, hash password, etc. </br>

# Getting Started

## Prerequisites

- You need to install [Golang](https://golang.org/doc/install) and [Docker](https://docs.docker.com/get-docker/) on your machine.
- You need to install Makefile on your machine. [Makefile](https://sourceforge.net/projects/mingw/)
- You need to install [PostgreSQL](https://www.postgresql.org/download/) on your machine.
- You need to install [Postman](https://www.postman.com/downloads/) on your machine.
- You need to install [Git](https://git-scm.com/downloads) on your machine.
- You need to install [Visual Studio Code](https://code.visualstudio.com/download) on your machine.

## Tech Stack

- Golang (Go Programming Language)
- Fiber (Web Framework)
- JWT (JSON Web Token)
- Docker
- PostgreSQL
- Validator (go-playground/validator)
- GORM (ORM)
- UUID (Universally Unique Identifier)

## Installation

Clone the repository and run the following commands:

```bash
git clone https://github.com/saufiroja/go-fiber-boilerplate.git
```

Go to the project directory:

```bash
cd go-fiber-boilerplate
```

Copy the .env_example file to .env:

```bash
copy .env.example .env
```

Install dependencies:

```bash
go mod tidy
```

Run the project:

```bash
make run
```

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
│       └── server/
│           └── fiber.go
├── initdb.d/
│   └── migrations/
│       ├── 0001_init.up.sql
│       └── 0002_init.users.sql
├── interfaces/
│   ├── auth_interfaces.go
│   └── user_interfaces.go
├── models/
│   ├── dto/
│   │   └── user_dto.go
│   └── entity/
│       └── user_entity.go
├── nginx/
│   ├── nginx.conf
│   └── Dockerfile
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
│   ├── constants/
│   │   ├── error.go
│   │   └── validation.go
│   ├── generate_token.go
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

## License

This project is licensed under the terms of the MIT license.
