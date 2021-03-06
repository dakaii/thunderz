package main

import (
	"coldhongdae/controller"
	"coldhongdae/database"
	"coldhongdae/repository"
	"net/http"
	"os"

	"github.com/samsarahq/thunder/graphql"
	"github.com/samsarahq/thunder/graphql/graphiql"
	"github.com/samsarahq/thunder/graphql/introspection"
)

func main() {
	port, exists := os.LookupEnv("PORT")
	if !exists {
		port = "3030"
	}
	collectionName, exists := os.LookupEnv("MONGODB_COLLECTION_NAME")
	if !exists {
		collectionName = "testingCollection"
	}

	ctx, db := database.GetDatabase(collectionName)
	userRepo := repository.NewUserRepo(ctx, db, db.Collection(collectionName))
	h := controller.NewBaseHandler(userRepo)

	schema := h.Schema()
	introspection.AddIntrospectionToSchema(schema)

	// Expose schema and graphiql.
	http.Handle("/graphql", graphql.Handler(schema))
	http.Handle("/graphiql/", http.StripPrefix("/graphiql/", graphiql.Handler()))
	http.ListenAndServe(":"+port, nil)
}
