package token

import (
	"github.com/graphql-go/graphql"
	"github.com/taskalla/api/pkg/models/token/token_struct"
	"github.com/taskalla/api/pkg/models/user"
)

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
		"personal": &graphql.EnumValueConfig{
			Value:       "personal",
			Description: "A personal token",
		},
		"other": &graphql.EnumValueConfig{
			Value: "other",
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
			Type: user.UserObj,
			Resolve: func(p graphql.ResolveParams) (interface{}, error) {
				source := p.Source.(*token_struct.Token)

				db_user, err := user.GetUserById(source.UserID)
				if err != nil {
					return &user.User{}, err
				}

				return db_user, nil
			},
		},
		"client_type": &graphql.Field{
			Type: ClientType,
		},
		"created_at": &graphql.Field{
			Type: graphql.DateTime,
		},
	},
})
