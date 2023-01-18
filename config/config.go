package config

import (
	"fmt"
	"log"
	"os"

	"project_alterra/model"

	"github.com/joho/godotenv"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

func initMigrate() {
	DB.AutoMigrate(&model.Movie{})
	DB.AutoMigrate(&model.Cast{})
	DB.AutoMigrate(&model.Movie_cast{})
}

func goDotEnvVariable(key string) string {

	// load .env file
	err := godotenv.Load(".env")
  
	if err != nil {
	  log.Fatalf("Error loading .env file")
	}
  
	return os.Getenv(key)
  }

func InitDB() {
	// refer https://github.com/go-sql-driver/mysql#dsn-data-source-name for details
	host := goDotEnvVariable("DB_HOST")
	port := goDotEnvVariable("DB_PORT")
	dbname := goDotEnvVariable("DB_NAME")
	username := goDotEnvVariable("DB_USER")
	password := goDotEnvVariable("DB_PASSWORD")

	// dsn := "root:@tcp(localhost:3306)/alterra_golang?charset=utf8mb4&parseTime=True&loc=Local"
	dsn := username + ":" + password + "@tcp(" + host + ":" + port + ")/" + dbname + "?charset=utf8&parseTime=true&loc=Local"
	var err error
	DB, err = gorm.Open(mysql.Open(dsn), &gorm.Config{SkipDefaultTransaction: true})
	// DB, err = gorm.Open("mysql", dsn)

	fmt.Println("Connected to DB!")
	if err != nil {
		panic(err.Error())
	}

	initMigrate()
}