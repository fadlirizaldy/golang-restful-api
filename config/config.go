package config

import (
	"log"
	"os"
	"regexp"

	"project_alterra/model"

	"github.com/joho/godotenv"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

const projectDirName = "project_alterra"

func GoDotEnvVariable(key string) string {

	projectName := regexp.MustCompile(`^(.*` + projectDirName + `)`)
    currentWorkDirectory, _ := os.Getwd()
    rootPath := projectName.Find([]byte(currentWorkDirectory))
	// load .env file
	err := godotenv.Load(string(rootPath) + `/.env`)
  
	if err != nil {
	  log.Fatalf("Error loading .env file")
	}
  
	return os.Getenv(key)
  }

func InitDB() *gorm.DB{
	// refer https://github.com/go-sql-driver/mysql#dsn-data-source-name for details
	host := GoDotEnvVariable("DB_HOST")
	port := GoDotEnvVariable("DB_PORT")
	dbname := GoDotEnvVariable("DB_NAME")
	username := GoDotEnvVariable("DB_USER")
	password := GoDotEnvVariable("DB_PASSWORD")

	// dsn := "root:@tcp(localhost:3306)/alterra_golang?charset=utf8mb4&parseTime=True&loc=Local"
	dsn := username + ":" + password + "@tcp(" + host + ":" + port + ")/" + dbname + "?charset=utf8&parseTime=true&loc=Local"
	var err error
	DB, err := gorm.Open(mysql.Open(dsn), &gorm.Config{SkipDefaultTransaction: true})
	// DB, err = gorm.Open("mysql", dsn)

	if err != nil {
		panic(err.Error())
	}
	// fmt.Println("Connected to DB!")

	return DB
}

func InitMigrate() {
	DB := InitDB()
	DB.AutoMigrate(&model.Movie{})
}