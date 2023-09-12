FROM golang:1.21-alpine AS builder

RUN apk add --no-cache make git gcc libc-dev

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY . .

RUN CGO_ENABLED=0 GOOS=linux go build -o myapp

# Path: Dockerfile
FROM alpine:latest

RUN apk update && apk add --no-cache ca-certificates

WORKDIR /app

COPY --from=builder /app/myapp /app/app

CMD ["./app"]