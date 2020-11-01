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

type ItemsConnection struct {
	Nodes       []Item `graphql:"nodes"`
	Number      int    `graphql:"number"`
	TotalNumber int    `graphql:"total_number"`
	FetchFunc   func() ([]Item, error)
}

var ItemsConnectionObj = graphql.NewObject(graphql.ObjectConfig{
	Name: "ItemsConnection",
	Fields: graphql.Fields{
		"nodes": &graphql.Field{
			Type: graphql.NewNonNull(
				graphql.NewList(ItemObj),
			),
			Resolve: func(p graphql.ResolveParams) (interface{}, error) {
				source := p.Source.(ItemsConnection)

				return source.FetchFunc()
			},
		},
		"number": &graphql.Field{
			Type:        graphql.NewNonNull(graphql.Int),
			Description: "The number of nodes on this page",
		},
		"total_number": &graphql.Field{
			Type:        graphql.NewNonNull(graphql.Int),
			Description: "The total number of items in the result set",
		},
	},
})
