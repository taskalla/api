package token

import (
	"time"

	"github.com/graphql-go/graphql"
)

const (
	TokenTypeOAuth  = "oauth"
	TokenTypeClient = "client"
)

type Token struct {
	ID         string    `graphql:"id"`
	Token      string    `graphql:"token"`
	Scopes     []string  `graphql:"scopes"`
	Valid      bool      `graphql:"valid"`
	CreatedOn  time.Time `graphql:"created_on"`
	TokenType  string    `graphql:"type"`
	User       string    `graphql:"user"`
	ClientType string    `graphql:"client_type"`
}

var TokenTypeObj = graphql.NewEnum(graphql.EnumConfig{
	Name: "TokenType",
	Values: graphql.EnumValueConfigMap{
		"oauth": &graphql.EnumValueConfig{
			Value: "oauth",
		},
		"client": &graphql.EnumValueConfig{
			Value: "client",
		},
	},
})

var ClientType = graphql.NewEnum(graphql.EnumConfig{
	Name: "ClientType",
	Values: graphql.EnumValueConfigMap{
		"mobile": &graphql.EnumValueConfig{
			Value: "mobile",
		},
		"web": &graphql.EnumValueConfig{
			Value: "web",
		},
	},
})

var TokenObj = graphql.NewObject(graphql.ObjectConfig{
	Name: "Token",
	Fields: graphql.Fields{
		"id": &graphql.Field{
			Type: graphql.NewNonNull(graphql.ID),
		},
		"token": &graphql.Field{
			Type: graphql.String,
		},
		"type": &graphql.Field{
			Type: TokenTypeObj,
		},
		"user": &graphql.Field{
			Type: graphql.String,
		},
		"client_type": &graphql.Field{
			Type: ClientType,
		},
	},
})
