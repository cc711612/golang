package config

import (
	"fmt"
	"github.com/joho/godotenv"
	"os"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type DatabaseConfig struct {
	Driver   string
	Host     string
	Port     string
	User     string
	Password string
	Name     string
}

func NewDatabaseConfig() *DatabaseConfig {
	// 讀取環境變數
	err := godotenv.Load()
	if err != nil {
		fmt.Println("Error loading .env file")
	}
	dbConfig := &DatabaseConfig{
		Driver:   os.Getenv("DB_DRIVER"),
		Host:     os.Getenv("DB_HOST"),
		Port:     os.Getenv("DB_PORT"),
		User:     os.Getenv("DB_USER"),
		Password: os.Getenv("DB_PASSWORD"),
		Name:     os.Getenv("DB_NAME"),
	}

	return dbConfig
}

func (config *DatabaseConfig) ConnectionString() string {
	return fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		config.User,
		config.Password,
		config.Host,
		config.Port,
		config.Name,
	)
}

func (config *DatabaseConfig) DB() (*gorm.DB, error) {
	db, err := gorm.Open(mysql.Open(config.ConnectionString()), &gorm.Config{})
	if err != nil {
		return nil, err
	}

	return db, nil
}

func ConnectDatabase() (*gorm.DB, error) {
	dbConfig := NewDatabaseConfig()
	return dbConfig.DB()
}
