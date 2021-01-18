package item

import (
	"github.com/graphql-go/graphql"
	"github.com/taskalla/api/pkg/db"
	"github.com/taskalla/api/pkg/models"
	"github.com/taskalla/api/pkg/tokenutils"
)

var CreateItemMutation = &graphql.Field{
	Name: "createItem",
	Type: ItemObj,
	Args: graphql.FieldConfigArgument{
		"input": &graphql.ArgumentConfig{
			Type: graphql.NewNonNull(CreateItemInput),
		},
	},
	Resolve: func(p graphql.ResolveParams) (interface{}, error) {
		t, err := tokenutils.ExtractToken(p)
		if err != nil {
			return nil, err
		}

		input := p.Args["input"].(map[string]interface{})

		item, err := CreateItem(&CreateItemParams{
			Description: input["description"].(string),
			UserID:      t.UserID,
		})
		if err != nil {
			return nil, err
		}

		return item, nil
	},
}

var CreateItemInput = graphql.NewInputObject(graphql.InputObjectConfig{
	Name: "CreateItemInput",
	Fields: graphql.InputObjectConfigFieldMap{
		"description": &graphql.InputObjectFieldConfig{
			Type: graphql.NewNonNull(graphql.String),
		},
	},
})

type CreateItemParams struct {
	UserID      string
	Description string
}

func CreateItem(params *CreateItemParams) (*models.Item, error) {
	item := models.Item{
		Description: params.Description,
		UserID:      params.UserID,
	}

	result := db.DB.Create(&item)

	if result.Error != nil {
		return nil, result.Error
	}

	return &item, nil
}
