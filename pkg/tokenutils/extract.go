package tokenutils

import (
	"errors"

	"github.com/graphql-go/graphql"
	"github.com/taskalla/api/pkg/models/token/token_struct"
)

func ExtractToken(p graphql.ResolveParams) (*token_struct.Token, error) {
	root := p.Info.RootValue.(map[string]interface{})
	token, ok := root["token"].(*token_struct.Token)
	if !ok || token == nil {
		return nil, errors.New("invalid auth")
	}

	return token, nil
}
