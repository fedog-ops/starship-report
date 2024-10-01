package database

import (
	"fmt"
	"gorm.io/driver/postgres"
	"log"
	"os"

	"api/models"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

type Dbinstance struct {
	Db *gorm.DB
}

var DB Dbinstance

func ConnectDb() {
	dsn := fmt.Sprintf("host=db user=%s password=%s dbname=%s port=5432 sslmode=disable TimeZone=Asia/Shanghai",
	os.Getenv("DB_USER"),
	os.Getenv("DB_PASSWORD"),
	os.Getenv("DB_NAME"),

)

	//without gorm
	// db, err = sql.Open("postgres", dsn)

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Info),
	})

	if err != nil {
		log.Fatal("failed to connect. \n", err)
	}

	log.Println("connected!")
	db.Logger = logger.Default.LogMode(logger.Info)

	log.Println("running migrations")
	db.AutoMigrate(&models.Officer{}, &models.CrewMember{}, &models.Report{})

	DB = Dbinstance{
		Db: db,
	}
}