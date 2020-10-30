package token

import (
	"time"

	"github.com/graphql-go/graphql"
)

type TokenType string

const (
	TokenTypeOAuth  TokenType = "oauth"
	TokenTypeClient TokenType = "client"
)

type Token struct {
	Token     string    `graphql:"token"`
	Scopes    []string  `graphql:"scopes"`
	Valid     bool      `graphql:"valid"`
	CreatedOn time.Time `graphql:"created_on"`
	TokenType TokenType `graphql:"token_type"`
}

var TokenTypeObj = graphql.NewEnum(graphql.EnumConfig{
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
