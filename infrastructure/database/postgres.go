package database

import (
	"database/sql"
	"fmt"
	"log"
	"project/go-fiber-boilerplate/config"
	"time"

	_ "github.com/lib/pq"
)

func NewPostgres(conf *config.AppConfig) *sql.DB {
	host := conf.Postgres.Host
	port := conf.Postgres.Port
	user := conf.Postgres.User
	pass := conf.Postgres.Pass
	name := conf.Postgres.Name
	ssl := conf.Postgres.Ssl

	url := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=%s", host, port, user, pass, name, ssl)

	db, err := sql.Open("postgres", url)
	if err != nil {
		panic(err)
	}

	// connection pool
	db.SetMaxIdleConns(10)
	db.SetMaxOpenConns(100)
	db.SetConnMaxIdleTime(5 * time.Minute)
	db.SetConnMaxLifetime(60 * time.Minute)

	log.Println("connected to database")

	return db
}
