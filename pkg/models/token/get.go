package token

import (
	"context"

	"github.com/taskalla/api/pkg/db"
	"github.com/taskalla/api/pkg/models/token/token_struct"
)

func GetToken(t string) (*token_struct.Token, error) {
	row := db.DB.QueryRow(context.Background(), "SELECT id, token, scopes, valid, created_on, token_type, user_id, client_type FROM tokens WHERE token = $1 AND valid IS TRUE", t)
	token := &token_struct.Token{}

	err := row.Scan(&token.ID, &token.Token, &token.Scopes, &token.Valid, &token.CreatedOn, &token.TokenType, &token.UserID, &token.ClientType)
	if err != nil {
		return &token_struct.Token{}, err
	}

	return token, nil
}
