package main

import (
	"net/http"
	"strconv"

	"github.com/taskalla/api/pkg/env"

	"github.com/graphql-go/graphql"
	"github.com/graphql-go/handler"
	"github.com/taskalla/api/pkg/logging"
	"github.com/taskalla/api/pkg/models/item"
)

func main() {
	rootQuery := graphql.NewObject(graphql.ObjectConfig{
		Name: "Query",
		Fields: graphql.Fields{
			"items": &graphql.Field{
				Type:    graphql.NewList(item.ItemObj),
				Resolve: item.ItemResolver,
			},
		},
	})

	schema, err := graphql.NewSchema(graphql.SchemaConfig{
		Query: rootQuery,
	})

	if err != nil {
		logging.Log("Error creating GraphQL schema: "+err.Error(), logging.LogLevelCritical)
		return
	}

	h := handler.New(&handler.Config{
		Playground: true,
		Schema:     &schema,
	})

	http.Handle("/graphql", h)

	err = http.ListenAndServe(":"+strconv.Itoa(env.Int("PORT", 3000)), nil)
	if err != nil {
		logging.Log("Error starting HTTP server: "+err.Error(), logging.LogLevelCritical)
	}
}
