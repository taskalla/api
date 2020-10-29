package schema

import (
	"github.com/graphql-go/graphql"
	"github.com/taskalla/api/pkg/models/item"
)

var RootQuery = graphql.NewObject(graphql.ObjectConfig{
	Name: "Query",
	Fields: graphql.Fields{
		"items": &graphql.Field{
			Type:    graphql.NewList(item.ItemObj),
			Resolve: item.ItemResolver,
		},
		"item": &graphql.Field{
			Type: item.ItemObj,
			Args: graphql.FieldConfigArgument{
				"title": &graphql.ArgumentConfig{
					Type: graphql.String,
				},
			},
			Resolve: func(p graphql.ResolveParams) (interface{}, error) {
				return item.Item{
					Title:       p.Args["title"].(string),
					Description: "test",
				}, nil
			},
		},
	},
})
