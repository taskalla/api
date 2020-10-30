package schema

import (
	"github.com/graphql-go/graphql"
	"github.com/taskalla/api/pkg/models/token"
)

var RootMutation = graphql.NewObject(graphql.ObjectConfig{
	Name: "Mutation",
	Fields: graphql.Fields{
		"createTokenByPassword": token.CreateByPassword,
	},
})
