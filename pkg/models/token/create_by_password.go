package token

import (
	"errors"

	"github.com/graphql-go/graphql"
	"github.com/taskalla/api/pkg/models/user"
	"github.com/taskalla/api/pkg/utils"
)

var CreateByPassword = &graphql.Field{
	Name: "createTokenByPassword",
	Type: TokenObj,
	Args: graphql.FieldConfigArgument{
		"input": &graphql.ArgumentConfig{
			Type: graphql.NewNonNull(CreateByPasswordInput),
		},
	},
	Resolve: func(p graphql.ResolveParams) (interface{}, error) {
		input := p.Args["input"].(map[string]interface{})

		db_user, err := user.GetUserByEmail(input["email"].(string))
		if err != nil {
			return nil, err
		}

		if !utils.PasswordIsCorrect(input["password"].(string), db_user.PasswordHash) {
			return nil, errors.New("Incorrect password")
		}

		token, err := CreateClientToken(db_user.ID, input["client_type"].(string))
		if err != nil {
			return nil, err
		}

		return token, nil
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
		"client_type": &graphql.InputObjectFieldConfig{
			Type: graphql.NewNonNull(ClientType),
		},
	},
})
