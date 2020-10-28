package item

import "github.com/graphql-go/graphql"

var ItemInput = graphql.NewInputObject(graphql.InputObjectConfig{
	Name: "ItemInput",
	Fields: graphql.InputObjectConfigFieldMap{
		"title": &graphql.InputObjectFieldConfig{
			Type: graphql.String,
		},
	},
})
