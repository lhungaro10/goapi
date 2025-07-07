package config

import (
	"os"

	"github.com/lhungaro10/goapi/schemas"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

const DB_PATH = "./db/main.db"

func InitializeSqlite() (*gorm.DB, error) {
	logger := GetLogger("sqlite")
	
	// check if database file exists
	_, err := os.Stat(DB_PATH)
	if os.IsNotExist(err) {
		logger.Info("sqlite database not exists. Creating new database...")
		err = os.MkdirAll("./db", os.ModePerm)

		if err != nil {
			logger.ErrorF("db dir not created error: %v", err)
			return nil, err
		}
		file, err := os.Create(DB_PATH)

		if err != nil {
			logger.ErrorF("db dir not created error: %v", err)
			return nil, err
		}
		file.Close()
	}
	
	// create database connection
	db, err:= gorm.Open(sqlite.Open(DB_PATH), &gorm.Config{})
	if err != nil {
		logger.ErrorF("sqlite Opening error: %v", err)
		return nil, err
	}
	logger.Info("sqlite database opened successfully.")
	// runinng migrations
	err = db.AutoMigrate(&schemas.Opening{})
	if err != nil {
		logger.ErrorF("sqlite AutoMigrate error: %v", err)
		return nil, err
	}
	logger.Info("sqlite database AutoMigrate success.")


	return db, nil
}