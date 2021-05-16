package main

import (
	"fmt"
	"graphyy/controller"
	"graphyy/internal"
	"graphyy/migration"
	"graphyy/repository"
	"graphyy/storage"
	"net/http"
	"os"
)

func main() {
	if internal.Migrate {
		migration.DataMigration()
		os.Exit(0)
	}
	db := storage.InitDatabase()
	repos := repository.InitRepositories(db)
	controllers := controller.InitControllers(repos)
	schema := controller.Schema(controllers)
	handler := controller.GraphqlHandlfunc(schema)

	http.Handle("/graphql", corsMiddleware(handler))
	fmt.Println("graphql api server is started at: http://localhost:" + internal.ServerPort + "/graphql")
	http.ListenAndServe(":"+internal.ServerPort, nil)
}

func corsMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Access-Control-Allow-Origin", "*")
		w.Header().Set("Access-Control-Allow-Headers", "Origin, X-Requested-With, Content-Type, Accept")
		next.ServeHTTP(w, r)
	})
}
