package item

import "github.com/graphql-go/graphql"

var ItemFilterObj = graphql.NewInputObject(graphql.InputObjectConfig{
	Name: "ItemFilter",
	Fields: graphql.InputObjectConfigFieldMap{
		"done": &graphql.InputObjectFieldConfig{
			Type: graphql.Boolean,
		},
	},
})

type ItemFilter struct {
	Done *bool `json:"done"`
}
