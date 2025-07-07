package config

import (
	"gorm.io/gorm"
)

var (
	db *gorm.DB
	logger *Logger
)

func Init() error {
  return nil
}

func GetLogger(loggerName string) *Logger {
	// Initialize Logger 
	logger :=  NewLogger(loggerName)
	return logger
}