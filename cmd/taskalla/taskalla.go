package main

import (
	"net/http"
	"strconv"

	"github.com/taskalla/api/pkg/env"
	"github.com/taskalla/api/pkg/schema"

	"github.com/graphql-go/graphql"
	"github.com/graphql-go/handler"
	"github.com/taskalla/api/pkg/logging"
)

func main() {
	schema, err := graphql.NewSchema(graphql.SchemaConfig{
		Query:    schema.RootQuery,
		Mutation: schema.RootMutation,
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
