package infrastructure

import (
	"fmt"
	"log"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type GormPostgres interface {
	GetConnection() *gorm.DB
}

type gormPostgresImpl struct {
	master *gorm.DB
}

func NewGormPostgres() GormPostgres {
	return &gormPostgresImpl{master: connect()}
}

func connect() *gorm.DB {
	host := "localhost"
	port := "5432"
	user := "postgres"
	password := "admin"
	dbname := "retail-admin-app"

	connectionStr := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable", host, port, user, password, dbname)
	db, err := gorm.Open(postgres.Open(connectionStr), &gorm.Config{})

	if err != nil {
		log.Println("DB Connect error: ", err)
	}

	return db
}

func (g *gormPostgresImpl) GetConnection() *gorm.DB {
	return g.master
}
