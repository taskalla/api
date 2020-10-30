package user

import (
	"context"

	"github.com/graphql-go/graphql"
	"github.com/taskalla/api/pkg/db"
	"github.com/taskalla/api/pkg/logging"
)

var CreateUserMutation = &graphql.Field{
	Name: "createUser",
	Type: UserObj,
	Args: graphql.FieldConfigArgument{
		"input": &graphql.ArgumentConfig{
			Type: CreateUserInput,
		},
	},
	Resolve: func(p graphql.ResolveParams) (interface{}, error) {
		input := p.Args["input"].(map[string]interface{})
		return CreateUser(input["email"].(string), input["password"].(string), input["name"].(string)), nil
	},
}

var CreateUserInput = graphql.NewInputObject(graphql.InputObjectConfig{
	Name: "CreateUserInput",
	Fields: graphql.InputObjectConfigFieldMap{
		"email": &graphql.InputObjectFieldConfig{
			Type: graphql.NewNonNull(graphql.String),
		},
		"password": &graphql.InputObjectFieldConfig{
			Type: graphql.NewNonNull(graphql.String),
		},
		"name": &graphql.InputObjectFieldConfig{
			Type: graphql.String,
		},
	},
})

func CreateUser(email, password, name string) *User {
	_, err := db.DB.Exec(context.Background(), "INSERT INTO users (email, password_hash, name, id) VALUES ($1, $2, $3, 'blah')", email, password, name)

	if err != nil {
		logging.Error(err)
	}

	return &User{
		Email: email,
		Name:  name,
	}
}