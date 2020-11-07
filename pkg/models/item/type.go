package item

import (
	"github.com/graphql-go/graphql"
)

type Item struct {
	Description string `graphql:"description"`
	ID          string `graphql:"id"`
	UserID      string `graphql:"user_id"`
	Done        bool   `graphql:"done"`
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
	},
})

type ItemsConnection struct {
	Nodes      []Item `graphql:"nodes"`
	Count      int    `graphql:"count"`
	TotalCount int    `graphql:"total_count"`
	FetchFunc  func() ([]*Item, error)
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
		"count": &graphql.Field{
			Type:        graphql.NewNonNull(graphql.Int),
			Description: "The number of nodes on this page",
		},
		"total_count": &graphql.Field{
			Type:        graphql.NewNonNull(graphql.Int),
			Description: "The total number of items in the result set",
		},
	},
})
