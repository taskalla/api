package token

import (
	"github.com/graphql-go/graphql"
	"github.com/taskalla/api/pkg/logging"
)

var CreateByPassword = &graphql.Field{
	Name: "createTokenByPassword",
	Type: graphql.String,
	Args: graphql.FieldConfigArgument{
		"input": &graphql.ArgumentConfig{
			Type: CreateByPasswordInput,
		},
	},
	Resolve: func(p graphql.ResolveParams) (interface{}, error) {
		logging.Info(p.Args["input"])
		return "cool token", nil
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
