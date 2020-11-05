package schema

import (
	"errors"

	"github.com/graphql-go/graphql"
	"github.com/taskalla/api/pkg/models/user"
	"github.com/taskalla/api/pkg/tokenutils"
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
	},
})
