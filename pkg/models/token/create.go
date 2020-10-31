package token

import (
	"context"
	"time"

	"github.com/taskalla/api/pkg/db"
	"github.com/taskalla/api/pkg/utils"
)

func CreateClientToken(user string) (*Token, error) {
	id := utils.GenerateUUID()
	createdOn := time.Now()

	token := utils.GenerateToken()

	_, err := db.DB.Exec(context.Background(), "INSERT INTO tokens (id, token, created_on, token_type, user_id) VALUES ($1, $2, $3, 'client', $4)", id, token, createdOn, user)
	if err != nil {
		return &Token{}, err
	}

	return &Token{
		ID:        id,
		User:      user,
		TokenType: TokenTypeClient,
		Valid:     true,
		CreatedOn: createdOn,
		Token:     token,
	}, nil
}
