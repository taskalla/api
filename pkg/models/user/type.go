package user

import (
	"github.com/graphql-go/graphql"
)

type User struct {
	Email        string
	PasswordHash string `graphql:"password_hash"`
	ID           int
	Name         string
}

var UserObj = graphql.NewObject(graphql.ObjectConfig{
	Name: "User",
	Fields: graphql.Fields{
		"email": &graphql.Field{
			Type: graphql.String,
		},
		"id": &graphql.Field{
			Type: graphql.ID,
		},
		"name": &graphql.Field{
			Type: graphql.String,
		},
	},
})
