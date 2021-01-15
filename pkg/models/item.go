package models

import (
	"github.com/taskalla/api/pkg/utils"
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

func (t *Item) BeforeCreate(tx *gorm.DB) (err error) {
	t.ID = utils.GenerateUUID()

	return
}
