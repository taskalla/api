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
	token := utils.GenerateToken()

	row := db.DB.QueryRow(context.Background(), "INSERT INTO tokens (id, token, token_type, user_id, client_type) VALUES ($1, $2, 'client', $3, $4) RETURNING created_at", id, token, user, client_type)

	var createdAt time.Time

	err := row.Scan(&createdAt)
	if err != nil {
		return &token_struct.Token{}, err
	}

	return &token_struct.Token{
		ID:         id,
		UserID:     user,
		TokenType:  token_struct.TokenTypeClient,
		Valid:      true,
		CreatedAt:  createdAt,
		Token:      token,
		ClientType: client_type,
	}, nil
}
