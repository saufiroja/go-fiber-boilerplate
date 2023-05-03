package database

import (
	"fmt"
	"log"
	"project/go-fiber-boilerplate/config"
	"project/go-fiber-boilerplate/models/entity"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func NewPostgres(conf *config.AppConfig) *gorm.DB {
	host := conf.Postgres.Host
	port := conf.Postgres.Port
	user := conf.Postgres.User
	pass := conf.Postgres.Pass
	name := conf.Postgres.Name

	url := fmt.Sprintf("host=%s port=%s user=%s dbname=%s password=%s sslmode=disable TimeZone=Asia/Jakarta", host, port, user, name, pass)

	db, err := gorm.Open(postgres.Open(url), &gorm.Config{})
	if err != nil {
		log.Panic(err)
	}

	err = db.AutoMigrate(&entity.User{})
	if err != nil {
		log.Panic(err)
	}

	log.Println("success connect to postgresql database!")

	return db
}
