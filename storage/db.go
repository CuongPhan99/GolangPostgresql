package storage

import (
	"log"

	config "github.com/CuongPhan99/GolangPostgresql/config"
	model "github.com/CuongPhan99/GolangPostgresql/models"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
)

var DB *gorm.DB

func NewDb(params ...string) *gorm.DB {
	var err error
	conString := config.GetPostgresConnectionString()

	log.Printf(conString)

	DB, err = gorm.Open(config.GetDBType(), conString)

	if err != nil {
		log.Panic(err)
	} else {
		DB.AutoMigrate(&model.Accounts{})
	}

	return DB
}

func GetDBInstance() *gorm.DB {
	return DB
}
