package infrastructure

import (
	"fmt"
	"log"

	"github.com/Rajanhub/goapi/lib"
	"github.com/Rajanhub/goapi/models"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

// Database modal
type Database struct {
	*gorm.DB
}

// NewDatabase creates a new database instance
func NewDatabase(env *lib.Env) Database {

	username := env.DBUsername
	password := env.DBPassword
	host := env.DBHost
	port := env.DBPort
	dbname := env.DBName

	url := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8&parseTime=True&loc=Local", username, password, host, port, dbname)

	db, err := gorm.Open(mysql.Open(url), &gorm.Config{})

	if err != nil {
		log.Println("Url: ", url)
		log.Panic(err)
	}

	log.Println("Database connection established")
	db.AutoMigrate(&models.Post{})

	return Database{
		DB: db,
	}
}
