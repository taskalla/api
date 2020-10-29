package schema

import (
	"github.com/graphql-go/graphql"
	"github.com/taskalla/api/pkg/models/item"
	"github.com/taskalla/api/pkg/models/token"
)

var RootMutation = graphql.NewObject(graphql.ObjectConfig{
	Name: "Mutation",
	Fields: graphql.Fields{
		"createItem": &graphql.Field{
			Type: item.ItemObj,
			Args: graphql.FieldConfigArgument{
				"item": &graphql.ArgumentConfig{
					Type: item.ItemInput,
				},
			},
			Resolve: func(p graphql.ResolveParams) (interface{}, error) {
				return item.Item{
					Title: p.Args["item"].(map[string]interface{})["title"].(string),
				}, nil
			},
		},
		"createTokenByPassword": token.CreateByPassword,
	},
})
