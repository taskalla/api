package user

import (
	"github.com/graphql-go/graphql"
	"github.com/taskalla/api/pkg/models/item"
	"github.com/taskalla/api/pkg/tokenutils"
)

type User struct {
	Email        string `graphql:"email"`
	PasswordHash string `graphql:"password_hash"`
	ID           string `graphql:"id"`
	Name         string `graphql:"name"`
}

var UserObj = graphql.NewObject(graphql.ObjectConfig{
	Name: "User",
	Fields: graphql.Fields{
		"email": &graphql.Field{
			Type: graphql.NewNonNull(graphql.String),
		},
		"id": &graphql.Field{
			Type: graphql.NewNonNull(graphql.ID),
		},
		"name": &graphql.Field{
			Type: graphql.String,
		},
		"items": &graphql.Field{
			Type: graphql.NewNonNull(item.ItemsConnectionObj),
			Args: graphql.FieldConfigArgument{
				"count": &graphql.ArgumentConfig{
					Description: "The number of items to fetch per page",
					Type:        graphql.NewNonNull(graphql.Int),
				},
				"page": &graphql.ArgumentConfig{
					Description:  "The page to fetch (1-indexed)",
					Type:         graphql.Int,
					DefaultValue: 1,
				},
			},
			Resolve: func(p graphql.ResolveParams) (interface{}, error) {
				t, err := tokenutils.ExtractToken(p)
				if err != nil {
					return nil, err
				}

				return item.ItemsConnection{
					Count:      10,
					TotalCount: 10,
					FetchFunc: func() ([]*item.Item, error) {
						return item.GetUserItems(t.UserID, p.Args["count"].(int), p.Args["page"].(int))
					},
				}, nil
			},
		},
	},
})
