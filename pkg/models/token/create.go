package token

import (
	"github.com/taskalla/api/pkg/db"
	"github.com/taskalla/api/pkg/models"
	"github.com/taskalla/api/pkg/utils"
)

func CreateClientToken(user, description string) (*models.Token, error) {
	generated_token := utils.GenerateToken()
	token := models.Token{
		UserID:      user,
		Description: &description,
		Token:       generated_token,
		Type:        models.TokenTypeClient,
	}

	result := db.DB.Create(&token)

	if result.Error != nil {
		return nil, result.Error
	}

	return &token, nil
}
