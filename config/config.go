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
var API_KEY string

func InitDB() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal(err)
		log.Fatal("Error loading .env file")
	}
	// ------------
	config := os.Getenv("CONNECTION_STRING")
	// config := os.Getenv("CONNECTION_LOCAL")
	API_KEY = os.Getenv("API_KEY")

	// viper.SetConfigFile(".env")
	// err := viper.ReadInConfig()

	// if err != nil {
	// 	log.Fatalf("Error while reading config file %s", err)
	// }
	// config, ok := viper.Get("CONNECTION_STRING").(string)

	// if !ok {
	// 	log.Fatalf("Invalid type assertion")
	// }

	var e error

	DB, e = gorm.Open(mysql.Open(config), &gorm.Config{})
	if e != nil {
		panic(e)
	}
	InitMigrate()
}

func InitMigrate() {
	DB.AutoMigrate(&models.Category{})
	DB.AutoMigrate(&models.Subcategory{})
	DB.AutoMigrate(&models.Province{})
	DB.AutoMigrate(&models.City{})
	DB.AutoMigrate(&models.Users{})
	DB.AutoMigrate(&models.Photos{})
	DB.AutoMigrate(&models.Products{})
	DB.AutoMigrate(&models.Guarantee{})
	DB.AutoMigrate(&models.ProductsGuarantee{})
	DB.AutoMigrate(&models.Cart{})
	// DB.AutoMigrate(&models.CheckoutMethodType{})
	// DB.AutoMigrate(&models.CheckoutMethod{})
	DB.AutoMigrate(&models.Transaction{})
	DB.AutoMigrate(&models.Booking{})
	DB.AutoMigrate(&models.Reviews{})
}

// ===============================================================//

func InitDBTest() {
	config := map[string]string{
		"DB_Username": "root",
		"DB_Password": "12345678",
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
	DB.Migrator().DropTable(&models.Reviews{})
	DB.Migrator().DropTable(&models.Booking{})
	DB.Migrator().DropTable(&models.Transaction{})
	DB.Migrator().DropTable(&models.CheckoutMethod{})
	DB.Migrator().DropTable(&models.CheckoutMethodType{})
	DB.Migrator().DropTable(&models.Cart{})
	DB.Migrator().DropTable(&models.ProductsGuarantee{})
	DB.Migrator().DropTable(&models.Guarantee{})
	DB.Migrator().DropTable(&models.Products{})
	DB.Migrator().DropTable(&models.Photos{})
	DB.Migrator().DropTable(&models.Users{})
	DB.Migrator().DropTable(&models.City{})
	DB.Migrator().DropTable(&models.Province{})
	DB.Migrator().DropTable(&models.Subcategory{})
	DB.Migrator().DropTable(&models.Category{})
	DB.AutoMigrate(&models.Category{})
	DB.AutoMigrate(&models.Subcategory{})
	DB.AutoMigrate(&models.Province{})
	DB.AutoMigrate(&models.City{})
	DB.AutoMigrate(&models.Users{})
	DB.AutoMigrate(&models.Photos{})
	DB.AutoMigrate(&models.Products{})
	DB.AutoMigrate(&models.Guarantee{})
	DB.AutoMigrate(&models.ProductsGuarantee{})
	DB.AutoMigrate(&models.Cart{})
	DB.AutoMigrate(&models.CheckoutMethodType{})
	DB.AutoMigrate(&models.CheckoutMethod{})
	DB.AutoMigrate(&models.Transaction{})
	DB.AutoMigrate(&models.Booking{})
	DB.AutoMigrate(&models.Reviews{})
}
