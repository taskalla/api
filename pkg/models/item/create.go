package item

import (
	"context"

	"github.com/graphql-go/graphql"
	"github.com/taskalla/api/pkg/db"
	"github.com/taskalla/api/pkg/tokenutils"
)

var CreateItemMutation = &graphql.Field{
	Name: "createItem",
	Type: ItemObj,
	Resolve: func(p graphql.ResolveParams) (interface{}, error) {
		t, err := tokenutils.ExtractToken(p)
		if err != nil {
			return nil, err
		}

		return Item{
			Title: t.UserID,
		}, nil
	},
}

var CreateItemInput = graphql.NewInputObject(graphql.InputObjectConfig{
	Fields: graphql.InputObjectConfigFieldMap{
		"title": &graphql.InputObjectFieldConfig{
			Type: graphql.String,
		},
		"description": &graphql.InputObjectFieldConfig{
			Type: graphql.NewNonNull(graphql.String),
		},
	},
})

type CreateItemParams struct {
	UserID      string
	Title       string
	Description string
}

func CreateItem(params CreateItemParams) (*Item, error) {
	db.DB.Exec(context.Background(), "INSERT INTO items (id, title, item_description, user_id) VALUES ()")

	return nil, nil
}
