package item

import (
	"github.com/graphql-go/graphql"
)

type Item struct {
	Title       string `graphql:"title"`
	Description string `graphql:"description"`
	ID          string `graphql:"id"`
	UserID      string `graphql:"user_id"`
}

var ItemObj = graphql.NewObject(graphql.ObjectConfig{
	Name: "Item",
	Fields: graphql.Fields{
		"id": &graphql.Field{
			Type:        graphql.NewNonNull(graphql.ID),
			Description: "The item's unique ID",
		},
		"title": &graphql.Field{
			Type:        graphql.NewNonNull(graphql.String),
			Description: "The item's title",
		},
		"description": &graphql.Field{
			Type:        graphql.NewNonNull(graphql.String),
			Description: "The item's description",
		},
	},
})
