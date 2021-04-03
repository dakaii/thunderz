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

	http.Handle("/graphql", controller.GraphqlHandlfunc(schema))

	fmt.Println("server is started at: http://localhost:/" + envvar.ServerPort() + "/")
	fmt.Println("graphql api server is started at: http://localhost:" + envvar.ServerPort() + "/graphql")
	http.ListenAndServe(":"+envvar.ServerPort(), nil)
}
