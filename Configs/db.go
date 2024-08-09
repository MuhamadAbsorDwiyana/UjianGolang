package configs

// Import modules gorm dan gotenv
import (
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

// Set variable for gorm DB
var DB *gorm.DB

// Function for init Database
func InitDB() {
	// get .env content use godotenv
	godotenv.Load()

	// Set err as error type data
	var err error

	// Set env info + set database
	dsn := fmt.Sprintf(
		"%s:%s@tcp(%s)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		os.Getenv("DB_USER"),
		os.Getenv("DB_PASSWORD"),
		os.Getenv("DB_HOST"),
		os.Getenv("DB_DBNAME"),
	)

	// Check database connecton
	DB, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
	// If error return sql error log
	if err != nil {
		return
	}

	// If success connect to database
	log.Printf("Akses ke database %s berhasil", os.Getenv("DB_DBNAME"))
}
