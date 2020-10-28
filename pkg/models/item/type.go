package item

import "github.com/graphql-go/graphql"

type Item struct {
	Title       string
	Description string
}

var ItemObj = graphql.NewObject(graphql.ObjectConfig{
	Name: "Item",
	Fields: graphql.Fields{
		"title": &graphql.Field{
			Type:        graphql.NewNonNull(graphql.String),
			Description: "The item's title",
		},
		"description": &graphql.Field{
			Type:        graphql.String,
			Description: "The item's description",
		},
	},
})
