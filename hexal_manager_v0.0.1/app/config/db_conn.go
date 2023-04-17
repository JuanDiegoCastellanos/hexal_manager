package config

import (
	"fmt"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

// NewDatabase crea una nueva conexión a la base de datos
func NewDatabase(cfg *Config) (*gorm.DB, error) {
	dsn := cfg.DBURI()
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		return nil, fmt.Errorf("failed to connect database: %w", err)
	}

	return db, nil
}

// Migrate realiza migraciones en la base de datos para modelos específicos
func Migrate(db *gorm.DB, models ...interface{}) error {
	for _, model := range models {
		if err := db.AutoMigrate(model); err != nil {
			return fmt.Errorf("failed to auto migrate: %w", err)
		}
	}
	return nil
}
