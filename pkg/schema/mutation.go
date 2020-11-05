package schema

import (
	"github.com/graphql-go/graphql"
	"github.com/taskalla/api/pkg/models/item"
	"github.com/taskalla/api/pkg/models/token"
	"github.com/taskalla/api/pkg/models/user"
)

var RootMutation = graphql.NewObject(graphql.ObjectConfig{
	Name: "Mutation",
	Fields: graphql.Fields{
		"createTokenByPassword": token.CreateByPassword,
		"createUser":            user.CreateUserMutation,
		"createItem":            item.CreateItemMutation,
	},
})
