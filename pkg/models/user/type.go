package user

import (
	"github.com/graphql-go/graphql"
	"github.com/taskalla/api/pkg/models/item"
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
			Type:    graphql.NewNonNull(item.ItemConnectionObj.Object),
			Args:    item.ItemConnectionObj.Args,
			Resolve: item.ItemConnectionObj.ResolveFunc(item.UserItemsResolver),
		},
	},
})
