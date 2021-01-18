package item

import (
	"github.com/graphql-go/graphql"
	"github.com/taskalla/api/pkg/db"
	"github.com/taskalla/api/pkg/models"
	"github.com/taskalla/api/pkg/tokenutils"
)

var UpdateItemMutation = &graphql.Field{
	Type: graphql.NewNonNull(ItemObj),
	Args: graphql.FieldConfigArgument{
		"input": &graphql.ArgumentConfig{
			Type: graphql.NewNonNull(UpdateItemMutationInput),
		},
	},
	Resolve: func(p graphql.ResolveParams) (interface{}, error) {
		t, err := tokenutils.ExtractToken(p)
		if err != nil {
			return nil, err
		}

		input := p.Args["input"].(map[string]interface{})

		item := models.Item{}
		result := db.DB.First(&item, "id = ? AND user_id = ?", input["id"].(string), t.UserID)
		if result.Error != nil {
			return nil, result.Error
		}

		if input_description, ok := input["description"].(string); ok {
			item.Description = input_description
		}

		if input_done, ok := input["done"].(bool); ok {
			item.Done = input_done
		}

		result = db.DB.Save(&item)
		if result.Error != nil {
			return nil, result.Error
		}

		return item, nil
	},
}

var UpdateItemMutationInput = graphql.NewInputObject(graphql.InputObjectConfig{
	Name: "UpdateItemInput",
	Fields: graphql.InputObjectConfigFieldMap{
		"description": &graphql.InputObjectFieldConfig{
			Type: graphql.String,
		},
		"done": &graphql.InputObjectFieldConfig{
			Type: graphql.Boolean,
		},
		"id": &graphql.InputObjectFieldConfig{
			Type: graphql.NewNonNull(graphql.ID),
		},
	},
})
