package main

import (
	"net/http"
	"strconv"

	"github.com/taskalla/api/pkg/db"
	"github.com/taskalla/api/pkg/env"
	"github.com/taskalla/api/pkg/models"
	"github.com/taskalla/api/pkg/root_object"
	"github.com/taskalla/api/pkg/schema"

	"github.com/graphql-go/graphql"
	"github.com/graphql-go/handler"
	"github.com/taskalla/api/pkg/logging"
)

func corsMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Access-Control-Allow-Origin", "http://localhost:3002")
		w.Header().Set("Access-Control-Allow-Headers", "Accept, Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization")

		next.ServeHTTP(w, r)
	})
}

func main() {
	db.Connect()
	db.DB.AutoMigrate(&models.User{}, &models.Item{}, &models.Token{})

	schema, err := graphql.NewSchema(graphql.SchemaConfig{
		Query:    schema.RootQuery,
		Mutation: schema.RootMutation,
	})

	if err != nil {
		logging.Log("Error creating GraphQL schema: "+err.Error(), logging.LogLevelCritical)
		return
	}

	h := handler.New(&handler.Config{
		Playground:   true,
		Schema:       &schema,
		RootObjectFn: root_object.ResolveRootObject,
	})

	http.Handle("/graphql", corsMiddleware(h))

	logging.Info("Starting up...")

	err = http.ListenAndServe(":"+strconv.Itoa(env.Int("PORT", 3000)), nil)
	if err != nil {
		logging.Log("Error starting HTTP server: "+err.Error(), logging.LogLevelCritical)
	}
}
