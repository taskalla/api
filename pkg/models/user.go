package models

import (
	"github.com/taskalla/api/pkg/utils"
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	ID           string  `gorm:"primarykey" json:"id"`
	Email        string  `gorm:"unique;not null" json:"email"`
	PasswordHash string  `gorm:"not null" json:"password"`
	Name         *string `json:"name"`
}

func (u *User) BeforeCreate(tx *gorm.DB) (err error) {
	u.ID = utils.GenerateUUID()

	return
}
