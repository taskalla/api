package user

import (
	"context"

	"github.com/taskalla/api/pkg/db"
)

func GetUserByEmail(email string) (*User, error) {
	row := db.DB.QueryRow(context.Background(), "SELECT id, email, password_hash, name FROM users WHERE email = $1", email)

	user := &User{}
	err := row.Scan(&user.ID, &user.Email, &user.PasswordHash, &user.Name)

	if err != nil {
		return &User{}, err
	}

	return user, nil
}
