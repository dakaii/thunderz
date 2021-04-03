package main

import (
	"fmt"
	"graphyy/controller"
	"graphyy/database"
	"graphyy/internal/envvar"
	"graphyy/migration"
	"graphyy/repository"
	"net/http"
	"os"
)

func main() {
	if envvar.Migrate() {
		migration.DataMigration()
		os.Exit(0)
	}
	db := database.InitDatabase()
	repos := repository.InitRepositories(db)
	controllers := controller.InitControllers(repos)
	schema := controller.Schema(controllers)
	handler := controller.GraphqlHandlfunc(schema)

	http.Handle("/graphql", corsMiddleware(handler))
	fmt.Println("graphql api server is started at: http://localhost:" + envvar.ServerPort() + "/graphql")
	http.ListenAndServe(":"+envvar.ServerPort(), nil)
}

func corsMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Access-Control-Allow-Origin", "*")
		w.Header().Set("Access-Control-Allow-Headers", "Origin, X-Requested-With, Content-Type, Accept")
		next.ServeHTTP(w, r)
	})
}
