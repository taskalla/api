package item

import (
	"github.com/graphql-go/graphql"
	"github.com/taskalla/api/pkg/paginate"
)

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
		"createdat": &graphql.Field{
			Type: graphql.NewNonNull(graphql.DateTime),
		},
	},
})

var ItemConnectionObj = paginate.NewConnectionObject("ItemConnection", ItemObj, graphql.Int, graphql.FieldConfigArgument{
	"filter": &graphql.ArgumentConfig{
		Type: ItemFilterObj,
	},
})
