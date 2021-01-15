package models

import (
	"gorm.io/gorm"
)

type Item struct {
	gorm.Model
	ID          string `gorm:"primarykey" json:"id"`
	Description string `gorm:"not null"`
	UserID      string
	User        User
	Done        bool `gorm:"default:false;not null"`
}
