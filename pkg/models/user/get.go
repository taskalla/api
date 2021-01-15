package user

import (
	"github.com/taskalla/api/pkg/db"
	"github.com/taskalla/api/pkg/models"
)

func GetUserByEmail(email string) (*models.User, error) {
	var user models.User

	result := db.DB.Where("email = ?", email).First(&user)

	if result.Error != nil {
		return &models.User{}, result.Error
	}

	return &user, nil
}

func GetUserById(id string) (*models.User, error) {
	var user models.User

	result := db.DB.Where("id = ?", id).First(&user)

	if result.Error != nil {
		return &models.User{}, result.Error
	}

	return &user, nil
}
