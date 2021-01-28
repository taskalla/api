package models

import (
	"time"

	"github.com/taskalla/api/pkg/utils"
	"gorm.io/gorm"
)

type Item struct {
	ID          string         `gorm:"primarykey" json:"id"`
	CreatedAt   time.Time      `json:"created_at"`
	UpdatedAt   time.Time      `json:"updated_at"`
	DeletedAt   gorm.DeletedAt `gorm:"index"`
	Description string         `gorm:"not null"`
	UserID      string
	User        User
	Done        bool `gorm:"default:false;not null"`
}

func (t *Item) BeforeCreate(tx *gorm.DB) (err error) {
	t.ID = utils.GenerateUUID()

	return
}
