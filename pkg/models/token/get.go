package token

import (
	"github.com/taskalla/api/pkg/db"
	"github.com/taskalla/api/pkg/models"
)

func GetToken(t string) (*models.Token, error) {

	token := &models.Token{}
	result := db.DB.Where("token = ?", t).First(&token)

	if result.Error != nil {
		return nil, result.Error
	}

	return token, nil
}
