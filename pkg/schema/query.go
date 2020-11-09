package schema

import (
	"errors"

	"github.com/graphql-go/graphql"
	"github.com/taskalla/api/pkg/models/user"
	"github.com/taskalla/api/pkg/tokenutils"
	"github.com/taskalla/api/pkg/unsplash"
)

var RootQuery = graphql.NewObject(graphql.ObjectConfig{
	Name: "Query",
	Fields: graphql.Fields{
		"viewer": &graphql.Field{
			Type: user.UserObj,
			Resolve: func(p graphql.ResolveParams) (interface{}, error) {
				t, err := tokenutils.ExtractToken(p)
				if err != nil {
					return nil, err
				}

				db_user, err := user.GetUserById(t.UserID)
				if err != nil {
					return nil, errors.New("invalid auth")
				}

				// TODO
				return db_user, nil
			},
		},
		"randomImage": &graphql.Field{
			Description: "Only for use by official Taskalla clients. Using this field without written permission from Taskalla is a violation of our TOS.",
			Resolve:     unsplash.RandomImage,
			Type:        unsplash.UnsplashImageObj,
		},
	},
})
