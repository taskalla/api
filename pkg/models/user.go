package models

import (
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	ID           string `gorm:"type:uuid;primaryKey"`
	Email        string `gorm:"unique"`
	PasswordHash string
	Name         string
}
