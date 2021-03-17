package config

import (
	"fmt"
	"log"

	"code.qburst.com/navaneeth.k/DynamoDB-example/models"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
)

const (
	
	host     = "localhost"
	port     = 5432
	user     = "postgres"
	password = "user123"
	dbname   = "datb3"
)

var (
	db  *gorm.DB
	err error
)

func PsqlConnect() *gorm.DB {

	psqlconn := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable",
		host, port, user, password, dbname)

	db, err = gorm.Open("postgres", psqlconn)
	if err != nil {
		log.Fatal(err)
	}
	db.AutoMigrate(&models.Records{})

	return db

}
