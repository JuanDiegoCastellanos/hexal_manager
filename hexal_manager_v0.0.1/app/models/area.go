package models

import (
	"gorm.io/gorm"
)

type Area struct {
	gorm.Model
	Name        string `json:"name" gorm:"not null"`
	Description string `json:"description" gorm:"not null"`
	Category    string `json:"category" gorm:"not null"`
	//CreatedAt   time.Time `gorm:"autoCreateTime" json:"created_at"`
	//UpdatedAt   time.Time `gorm:"autoUpdateTime" json:"updated_at"`
	//DeletedAt   gorm.DeletedAt `gorm:"index" json:"deleted_at,omitempty"`
}
