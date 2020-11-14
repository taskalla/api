package item

import (
	"context"
	"time"

	"github.com/graphql-go/graphql"
	"github.com/taskalla/api/pkg/db"
	"github.com/taskalla/api/pkg/tokenutils"
	"github.com/taskalla/api/pkg/utils"
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

func CreateItem(params *CreateItemParams) (*Item, error) {
	id := utils.GenerateUUID()
	row := db.DB.QueryRow(context.Background(), "INSERT INTO items (id, item_description, user_id) VALUES ($1, $2, $3) RETURNING created_at, done", id, params.Description, params.UserID)

	var created_at time.Time
	var done bool

	err := row.Scan(&created_at, &done)
	if err != nil {
		return nil, err
	}

	return &Item{
		Description: params.Description,
		UserID:      params.UserID,
		ID:          id,
		CreatedAt:   created_at,
		Done:        done,
	}, nil
}
