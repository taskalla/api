package item

import (
	"github.com/graphql-go/graphql"
	"github.com/taskalla/api/pkg/tokenutils"
)

var CreateItemMutation = &graphql.Field{
	Name: "createItem",
	Type: ItemObj,
	Resolve: func(p graphql.ResolveParams) (interface{}, error) {
		t, err := tokenutils.ExtractToken(p)
		if err != nil {
			return nil, err
		}

		return Item{
			Title: t.UserID,
		}, nil
	},
}
