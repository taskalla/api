package user

import (
	"github.com/graphql-go/graphql"
	"github.com/taskalla/api/pkg/db"
	"github.com/taskalla/api/pkg/logging"
	"github.com/taskalla/api/pkg/models"
	"github.com/taskalla/api/pkg/utils"
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
		input_name, ok := input["name"].(string)

		var name *string = nil
		if ok {
			name = &input_name
		}
		return CreateUser(input["email"].(string), input["password"].(string), name)
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

func CreateUser(email, password string, name *string) (*models.User, error) {
	passwordHash, err := utils.HashPassword(password)
	if err != nil {
		logging.Error(err)
		return &models.User{}, err
	}

	user := models.User{
		Email:        email,
		PasswordHash: passwordHash,
		Name:         name,
	}

	result := db.DB.Create(&user)

	if result.Error != nil {
		// logging.Error(result.Error)
		return nil, result.Error
	}

	return &models.User{
		ID:    user.ID,
		Email: email,
		Name:  name,
	}, nil
}
