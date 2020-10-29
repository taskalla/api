package item

import "github.com/graphql-go/graphql"

func ItemResolver(p graphql.ResolveParams) (interface{}, error) {
	rootObject := p.Info.RootValue.(map[string]interface{})
	rootObject["item_id"] = "BLAHBLAH"
	return []Item{
		{
			Title: "cool title",
			ID:    "1",
		},
		{
			Title:       "super cool title",
			Description: "with a description!",
			ID:          "2",
		},
	}, nil
}
