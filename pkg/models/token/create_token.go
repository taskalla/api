package token

import (
	"github.com/graphql-go/graphql"
)

var CreateByPassword = &graphql.Field{
	Name: "createTokenByPassword",
	Type: TokenObj,
	Args: graphql.FieldConfigArgument{
		"input": &graphql.ArgumentConfig{
			Type: CreateByPasswordInput,
		},
	},
	Resolve: func(p graphql.ResolveParams) (interface{}, error) {
		return Token{
			Token:     "cool token!",
			TokenType: TokenTypeClient,
		}, nil
	},
}

var CreateByPasswordInput = graphql.NewInputObject(graphql.InputObjectConfig{
	Name: "CreateTokenByPasswordInput",
	Fields: graphql.InputObjectConfigFieldMap{
		"email": &graphql.InputObjectFieldConfig{
			Type: graphql.NewNonNull(graphql.String),
		},
		"password": &graphql.InputObjectFieldConfig{
			Type: graphql.NewNonNull(graphql.String),
		},
	},
})
