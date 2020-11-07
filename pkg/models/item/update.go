package item

import (
	"context"
	"errors"

	"github.com/graphql-go/graphql"
	"github.com/taskalla/api/pkg/db"
	"github.com/taskalla/api/pkg/tokenutils"
)

var UpdateItemMutation = &graphql.Field{
	Type: graphql.NewNonNull(graphql.String),
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

		var (
			description *string
			done        *bool
		)

		if input_description, ok := input["description"].(string); ok {
			description = &input_description
		}
		if input_done, ok := input["done"].(bool); ok {
			done = &input_done
		}

		new_item, err := UpdateItem(input["id"].(string), t.UserID, description, done)
		if err != nil {
			return nil, err
		}

		return new_item, nil
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

func UpdateItem(id, user_id string, description *string, done *bool) (string, error) {
	info, err := db.DB.Exec(context.Background(), "UPDATE items SET item_description = coalesce($1, item_description), done = coalesce($2, done) WHERE id = $3 AND user_id = $4", description, done, id, user_id)
	if err != nil {
		return "", err
	}

	if info.RowsAffected() == 0 {
		return "", errors.New("Item not found")
	}

	return id, nil
}
