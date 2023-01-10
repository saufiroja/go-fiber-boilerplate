package config

import "os"

func initPostgres(conf *AppConfig) {
	host := os.Getenv("DB_HOST")
	port := os.Getenv("DB_PORT")
	user := os.Getenv("DB_USER")
	pass := os.Getenv("DB_PASS")
	dbname := os.Getenv("DB_NAME")

	conf.Postgres.Host = host
	conf.Postgres.Port = port
	conf.Postgres.User = user
	conf.Postgres.Pass = pass
	conf.Postgres.Name = dbname
}
