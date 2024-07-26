package PSQL

import (
	"fmt"
	"os"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var dbinstance *gorm.DB

func Connect() *gorm.DB {
	if dbinstance == nil {
		host := os.Getenv("DB_HOST")
		user := os.Getenv("DB_USER")
		password := os.Getenv("DB_PASSWORD")
		dbname := os.Getenv("DB_NAME")
		port := os.Getenv("DB_PORT")
		dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable TimeZone=Asia/Kolkata", host, user, password, dbname, port)
		// db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{
		// 	Logger: logger.Default.LogMode(logger.Info),
		// })
		db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
		if err != nil {
			panic("failed to connect database")
		}
		dbinstance = db
	}
	return dbinstance
}
