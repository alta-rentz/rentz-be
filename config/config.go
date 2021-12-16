package config

import (
	"fmt"
	"log"
	"os"

	"project3/models"

	"github.com/joho/godotenv"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

func InitDB() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
	config := os.Getenv("CONNECTION_STRING")
	// config := os.Getenv("CONNECTION_DB")

	var e error

	DB, e = gorm.Open(mysql.Open(config), &gorm.Config{})
	if e != nil {
		panic(e)
	}
	InitMigrate()
}

func InitMigrate() {
	DB.AutoMigrate(&models.Users{})
}

// ===============================================================//

func InitDBTest() {
	config := map[string]string{
		"DB_Username": "root",
		"DB_Password": "root",
		"DB_Port":     "3306",
		"DB_Host":     "localhost",
		"DB_Name":     "db_test",
	}

	connectionString := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8&parseTime=True&loc=Local",
		config["DB_Username"],
		config["DB_Password"],
		config["DB_Host"],
		config["DB_Port"],
		config["DB_Name"],
	)

	var e error
	DB, e = gorm.Open(mysql.Open(connectionString), &gorm.Config{})
	if e != nil {
		panic(e)
	}
	InitMigrationTest()
}

func InitMigrationTest() {
	DB.Migrator().DropTable(&models.Users{})
	DB.AutoMigrate(&models.Users{})
}
