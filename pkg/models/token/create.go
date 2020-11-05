package token

import (
	"context"
	"time"

	"github.com/taskalla/api/pkg/db"
	"github.com/taskalla/api/pkg/models/token/token_struct"
	"github.com/taskalla/api/pkg/utils"
)

func CreateClientToken(user, client_type string) (*token_struct.Token, error) {
	id := utils.GenerateUUID()
	createdOn := time.Now()

	token := utils.GenerateToken()

	_, err := db.DB.Exec(context.Background(), "INSERT INTO tokens (id, token, created_on, token_type, user_id, client_type) VALUES ($1, $2, $3, 'client', $4, $5)", id, token, createdOn, user, client_type)
	if err != nil {
		return &token_struct.Token{}, err
	}

	return &token_struct.Token{
		ID:         id,
		UserID:     user,
		TokenType:  token_struct.TokenTypeClient,
		Valid:      true,
		CreatedOn:  createdOn,
		Token:      token,
		ClientType: client_type,
	}, nil
}
