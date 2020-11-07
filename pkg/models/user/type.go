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
				"filter": &graphql.ArgumentConfig{
					Type: item.ItemFilterObj,
				},
			},
			Resolve: func(p graphql.ResolveParams) (interface{}, error) {
				t, err := tokenutils.ExtractToken(p)
				if err != nil {
					return nil, err
				}

				filter, filter_ok := p.Args["filter"].(map[string]interface{})
				filter_obj := item.ItemFilter{}
				if filter_ok {
					if filter_done, ok := filter["done"].(bool); ok {
						filter_obj.Done = &filter_done
					}
				}

				count, err := item.GetItemCountOnPage(t.UserID, p.Args["count"].(int), p.Args["page"].(int), filter_obj)
				if err != nil {
					return nil, err
				}

				total_count, err := item.GetTotalItemCount(t.UserID, filter_obj)
				if err != nil {
					return nil, err
				}

				return item.ItemsConnection{
					Count:      count,
					TotalCount: total_count,
					FetchFunc: func() ([]*item.Item, error) {
						return item.GetUserItems(t.UserID, p.Args["count"].(int), p.Args["page"].(int), filter_obj)
					},
				}, nil
			},
		},
	},
})
