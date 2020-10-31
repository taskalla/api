package schema

import (
	"errors"

	"github.com/graphql-go/graphql"
	"github.com/taskalla/api/pkg/models/token"
	"github.com/taskalla/api/pkg/models/user"
)

var RootQuery = graphql.NewObject(graphql.ObjectConfig{
	Name: "Query",
	Fields: graphql.Fields{
		"viewer": &graphql.Field{
			Type: user.UserObj,
			Resolve: func(p graphql.ResolveParams) (interface{}, error) {
				root := p.Info.RootValue.(map[string]interface{})
				token, ok := root["token"].(*token.Token)
				if !ok || token == nil {
					return nil, errors.New("invalid auth")
				}

				db_user, err := user.GetUserById(token.UserID)
				if err != nil {
					return nil, errors.New("invalid auth")
				}

				// TODO
				return db_user, nil
			},
		},
	},
})
