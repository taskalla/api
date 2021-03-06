package tokenutils

import (
	"errors"

	"github.com/graphql-go/graphql"
	"github.com/taskalla/api/pkg/models"
)

func ExtractToken(p graphql.ResolveParams) (*models.Token, error) {
	root := p.Info.RootValue.(map[string]interface{})
	token, ok := root["token"].(*models.Token)
	if !ok || token == nil {
		return nil, errors.New("invalid auth")
	}

	return token, nil
}
