package models

import (
	"gorm.io/gorm"
	"hexal_manager_v0.0.1/app/utils"
)

type User struct {
	gorm.Model
	Name     string `gorm:"not null" json:"name"`
	Email    string `gorm:"unique;not null" json:"email"`
	Password string `gorm:"not null" json:"password"`
}

func (u *User) BeforeCreate(tx *gorm.DB) error {
	return tx.Callback().Create().Before("gorm:create").Register("hash_password", func(tx *gorm.DB) {
		if len(u.Password) > 0 {
			hashedPassword, err := utils.HashPassword(u.Password)
			if err == nil {
				u.Password = hashedPassword
			}
		}
	})
}

func (u *User) CheckPassword(password string) bool {
	return utils.CheckPasswordHash(password, u.Password)
}
