package item

import "github.com/graphql-go/graphql"

func ItemResolver(p graphql.ResolveParams) (interface{}, error) {
	return []Item{
		{
			Title: "cool title",
		},
		{
			Title:       "super cool title",
			Description: "with a description!",
		},
	}, nil
}
