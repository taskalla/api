package item

import (
	"time"

	"github.com/graphql-go/graphql"
	"github.com/taskalla/api/pkg/paginate"
)

type Item struct {
	Description string    `graphql:"description"`
	ID          string    `graphql:"id"`
	UserID      string    `graphql:"user_id"`
	Done        bool      `graphql:"done"`
	CreatedAt   time.Time `graphql:"created_at"`
}

var ItemObj = graphql.NewObject(graphql.ObjectConfig{
	Name: "Item",
	Fields: graphql.Fields{
		"id": &graphql.Field{
			Type:        graphql.NewNonNull(graphql.ID),
			Description: "The item's unique ID",
		},
		"description": &graphql.Field{
			Type:        graphql.NewNonNull(graphql.String),
			Description: "The item's description",
		},
		"done": &graphql.Field{
			Type: graphql.NewNonNull(graphql.Boolean),
		},
		"created_at": &graphql.Field{
			Type: graphql.NewNonNull(graphql.DateTime),
		},
	},
})

var ItemConnectionObj = paginate.NewConnectionObject("ItemConnection", ItemObj, graphql.FieldConfigArgument{
	"filter": &graphql.ArgumentConfig{
		Type: ItemFilterObj,
	},
})
