package main

import (
	"coldhongdae/controllers"
	"coldhongdae/database"
	"log"
	"net/http"
	"os"

	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
	"github.com/samsarahq/thunder/graphql"
	"github.com/samsarahq/thunder/graphql/graphiql"
	"github.com/samsarahq/thunder/graphql/introspection"
)

func main() {
	port, exists := os.LookupEnv("PORT")
	if !exists {
		port = "8080"
	}
	collectionName, exists := os.LookupEnv("MONGODB_COLLECTION_NAME")
	if !exists {
		collectionName = "testingCollection"
	}

	ctx, db := database.GetDatabase(collectionName)
	userRepo := database.NewUserRepo(db, ctx, db.Collection(collectionName))
	h := controllers.NewBaseHandler(userRepo)

	r := mux.NewRouter()
	schema := h.Schema()
	introspection.AddIntrospectionToSchema(schema)

	r.Handle("/graphql", graphql.Handler(schema))
	r.Handle("/graphiql", graphiql.Handler())
	headersOk := handlers.AllowedHeaders([]string{"X-Requested-With", "content-type"})
	originsOk := handlers.AllowedOrigins([]string{"*"})
	methodsOk := handlers.AllowedMethods([]string{"GET", "HEAD", "POST", "PUT", "OPTIONS"})
	log.Println(":" + port)

	err := http.ListenAndServe(":"+port, handlers.CORS(originsOk, headersOk, methodsOk)(r))
	if err != nil {
		log.Fatal(err)
	}
}
