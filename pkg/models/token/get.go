package token

import (
	"context"

	"github.com/taskalla/api/pkg/db"
)

func GetToken(t string) (*Token, error) {
	row := db.DB.QueryRow(context.Background(), "SELECT id, token, scopes, valid, created_on, token_type, user_id, client_type FROM tokens WHERE token = $1 AND valid IS TRUE", t)
	token := &Token{}

	err := row.Scan(&token.ID, &token.Token, &token.Scopes, &token.Valid, &token.CreatedOn, &token.TokenType, &token.UserID, &token.ClientType)
	if err != nil {
		return &Token{}, err
	}

	return token, nil
}
