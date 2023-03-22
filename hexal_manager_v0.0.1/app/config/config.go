package config

import (
	"fmt"
	"github.com/joho/godotenv"
	"os"
)

type Config struct {
	DBUser     string
	DBPassword string
	DBHost     string
	DBPort     string
	DBName     string
}

func LoadConfig() (*Config, error) {
	err := godotenv.Load(".env")
	if err != nil {
		fmt.Println("Error loading .env file")
		return nil, err
	}

	cfg := &Config{
		DBUser:     os.Getenv("DBUser"),
		DBPassword: os.Getenv("DBPassword"),
		DBHost:     os.Getenv("DBHost"),
		DBPort:     os.Getenv("DBPort"),
		DBName:     os.Getenv("DBName"),
	}
	return cfg, nil
}
func (c *Config) DBURI() string {
	return fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local", c.DBUser, c.DBPassword, c.DBHost, c.DBPort, c.DBName)
}
