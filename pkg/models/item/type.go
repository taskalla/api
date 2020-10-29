package item

import (
	"github.com/Matt-Gleich/logoru"
	"github.com/graphql-go/graphql"
)

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
			Resolve: func(p graphql.ResolveParams) (interface{}, error) {
				logoru.Debug(p.Info.VariableValues)

				return "DESCRIPTION", nil
			},
		},
	},
})
