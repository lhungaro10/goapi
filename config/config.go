package config

import (
	"fmt"

	"gorm.io/gorm"
)

var (
	db *gorm.DB
	logger *Logger
)

func Init() error {
	var err error
	//Initialize Sqlite
	db, err = InitializeSqlite()
	if err != nil {
		return fmt.Errorf("Error initializing sqlite: %v", err)
	}
  return nil
}

func GetDb() *gorm.DB {
	return db
}

func GetLogger(loggerName string) *Logger {
	// Initialize Logger 
	logger =  NewLogger(loggerName)
	return logger
}