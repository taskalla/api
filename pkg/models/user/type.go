package user

import (
	"crypto/md5"
	"errors"
	"fmt"
	"net/url"
	"strconv"
	"strings"

	"github.com/graphql-go/graphql"
	"github.com/taskalla/api/pkg/logging"
	"github.com/taskalla/api/pkg/models"
	"github.com/taskalla/api/pkg/models/item"
)

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
		"gravatar": &graphql.Field{
			Args: graphql.FieldConfigArgument{
				"size": &graphql.ArgumentConfig{
					Type:         graphql.Int,
					DefaultValue: 200,
				},
			},
			Type: graphql.NewNonNull(graphql.String),
			Resolve: func(p graphql.ResolveParams) (interface{}, error) {
				if user, ok := p.Source.(*models.User); ok {
					hashed_email := md5.Sum([]byte(strings.ToLower(user.Email)))
					size := p.Args["size"].(int)

					gravatar_url := url.URL{
						Host:   "www.gravatar.com",
						Scheme: "https",
						Path:   fmt.Sprintf("/avatar/%x", hashed_email),
						RawQuery: url.Values{
							"d": {"mp"},
							"s": {strconv.Itoa(size)},
						}.Encode(),
					}

					return gravatar_url.String(), nil
				} else {
					logging.Info(p.Source)
					return nil, errors.New("Error reading user's email")
				}
			},
		},
	},
})
