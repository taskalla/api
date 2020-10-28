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
	},
})
