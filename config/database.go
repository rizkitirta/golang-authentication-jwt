package config

import (
	"fmt"
	"os"

	"github.com/joho/godotenv"
	"gorm.io/gorm"
	"gorm.io/driver/mysql"
)

func SetUpDatabaseConnection() *gorm.DB {
	err := godotenv.Load()
	if err != nil {
		panic(err)
	}

	dbUser := os.Getenv("DB_USER")
	dbPassword := os.Getenv("DB_PASSWORD")
	dbHost := os.Getenv("DB_HOST")
	dbName := os.Getenv("DB_NAME")

	dsn := fmt.Sprintf("%s:%s@tcp(%s:3306)/%s?charset=utf8mb4&parseTime=True&loc=Local", dbUser, dbPassword, dbHost, dbName)
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic(err)
	}

	return db
}

func CloseDatabaseConnection(dB *gorm.DB) {
	dbSql, err := dB.DB()
	if err != nil {
		panic(err)
	}

	dbSql.Close()
}
