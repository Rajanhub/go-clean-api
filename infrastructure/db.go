package infrastructure

import (
	"fmt"
	"log"

	"github.com/Rajanhub/goapi/lib"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

// Database modal
type Database struct {
	*gorm.DB
}

// NewDatabase creates a new database instance
func NewDatabase(logger lib.Logger, env *lib.Env) Database {

	username := env.DBUsername
	password := env.DBPassword
	host := env.DBHost
	port := env.DBPort
	dbname := env.DBName

	url := fmt.Sprintf("%s:%s@tcp(%s:%s)/?charset=utf8&parseTime=True&loc=Local", username, password, host, port)

	logger.Info("opening db connection")
	db, err := gorm.Open(mysql.Open(url), &gorm.Config{Logger: logger.GetGormLogger()})

	if err != nil {
		logger.Info("Url: ", url)
		logger.Panic(err)
	}
	logger.Info("creating database if it does't exist")
	if err = db.Exec("CREATE DATABASE IF NOT EXISTS " + dbname).Error; err != nil {
		logger.Info("couldn't create database")
		logger.Panic(err)
	}
	logger.Info("using given database")
	if err := db.Exec(fmt.Sprintf("USE %s", dbname)).Error; err != nil {
		logger.Info("cannot use the given database")
		logger.Panic(err)
	}

	logger.Info("database connection established")

	database := Database{
		DB: db,
	}

	err = RunMigration(logger, database)
	//db.AutoMigrate(&models.Post{})
	if err != nil {
		log.Println(err.Error())
		logger.Error(err)
	}
	logger.Info("currentDatabase:", db.Migrator().CurrentDatabase())

	return Database{
		DB: db,
	}
}
