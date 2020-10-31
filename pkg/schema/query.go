package schema

import (
	"github.com/graphql-go/graphql"
	"github.com/taskalla/api/pkg/models/user"
)

var RootQuery = graphql.NewObject(graphql.ObjectConfig{
	Name: "Query",
	Fields: graphql.Fields{
		"viewer": &graphql.Field{
			Type: user.UserObj,
			Resolve: func(p graphql.ResolveParams) (interface{}, error) {
				// TODO
				return nil, nil
			},
		},
	},
})
