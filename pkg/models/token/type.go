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
	Token     string    `graphql:"token"`
	Scopes    []string  `graphql:"scopes"`
	Valid     bool      `graphql:"valid"`
	CreatedOn time.Time `graphql:"created_on"`
	TokenType string    `graphql:"type"`
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

var TokenObj = graphql.NewObject(graphql.ObjectConfig{
	Name: "Token",
	Fields: graphql.Fields{
		"token": &graphql.Field{
			Type: graphql.String,
		},
		"type": &graphql.Field{
			Type: TokenTypeObj,
		},
	},
})
