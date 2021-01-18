package models

import (
	"github.com/taskalla/api/pkg/utils"
	"gorm.io/gorm"
)

type TokenType string

const (
	TokenTypeOAuth    TokenType = "oauth"
	TokenTypeClient   TokenType = "client"
	TokenTypePersonal TokenType = "personal"
)

type Token struct {
	gorm.Model
	ID    string `gorm:"primarykey" json:"id"`
	Token string `gorm:"unique;not null" json:"token"`
	// Scopes      []string
	UserID      string
	User        User
	Type        TokenType `gorm:"type:varchar;not null"`
	Description *string
}

func (t *Token) BeforeCreate(tx *gorm.DB) (err error) {
	t.ID = utils.GenerateUUID()

	return
}
