package database

import (
	"encoding/json"
	"fmt"
	"log"
	"os"

	"gorm.io/driver/mysql"
	"gorm.io/driver/postgres"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"

	"github.com/gunjanmistry08/diary-app/internal/models"
)

var DB *gorm.DB

type DBConfig struct {
	Driver   string `json:"driver"`
	Host     string `json:"host"`
	User     string `json:"user"`
	Password string `json:"password"`
	Name     string `json:"name"`
	Port     string `json:"port"`
}

var config DBConfig

// Add getter for config
func Config() DBConfig {
	return config
}

func LoadConfig(path string) {
	file, err := os.Open(path)
	if err != nil {
		log.Fatalf("Failed to open config file: %v", err)
	}
	defer file.Close()

	decoder := json.NewDecoder(file)
	if err := decoder.Decode(&config); err != nil {
		log.Fatalf("Failed to decode config file: %v", err)
	}
}

func Connect() {
	LoadConfig("configs/database.json") // database.json should be in the configs folder
	var dsn string
	var db *gorm.DB
	var err error

	switch config.Driver {
	case "postgres":
		dsn = fmt.Sprintf(
			"host=%s user=%s password=%s dbname=%s port=%s sslmode=disable",
			config.Host,
			config.User,
			config.Password,
			config.Name,
			config.Port,
		)
		db, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
	case "mysql":
		dsn = fmt.Sprintf(
			"%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local",
			config.User,
			config.Password,
			config.Host,
			config.Port,
			config.Name,
		)
		db, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
	case "sqlite":
		dsn = config.Name
		db, err = gorm.Open(sqlite.Open(dsn), &gorm.Config{})
	default:
		log.Fatalf("Unsupported DB_DRIVER: %s", config.Driver)
	}

	if err != nil {
		log.Fatal("Failed to connect to database: ", err)
	}

	log.Println("Connected to database!")

	db.AutoMigrate(&models.User{}, &models.DiaryEntry{})

	DB = db
}

// getEnv removed; config values are now loaded from config.json
