# Go Fiber Boilerplate

A Starter project with Golang and Fiber

# Getting Started

## Features

- [x] Authentication (Login, Register)
  - [ ] Forgot Password
  - [ ] Refresh Token
  - [ ] Logout
- [x] User (Get All User, Get User By ID, Update User By ID, Delete User By ID)
  - [x] Get All User
  - [x] Get User By ID
  - [x] Update User By ID
  - [x] Delete User By ID

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
.
├── app/
│   └── main.go
├── config/
│   ├── app.go
│   ├── config.go
│   ├── fiber.go
│   └── postgres.go
├── infratructure/
│   ├── database/
│   │   ├── postgres.go
│   │   └── mongo.go
│   └── http/
│       ├── handler/
│       │   ├── auth/
│       │   │   └── auth_handler.go
│       │   └── user/
│       │       └── user_handler.go
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
│       ├── auth_repository_postgres.go
│       └── user_repository_postgres.go
├── service/
│   ├── auth/
│   │   └── auth_service.go
│   └── user/
│       └── user_service.go
├── utils/
│   ├── generate_token.go
│   ├── handlerError.go
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
