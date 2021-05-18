package controller

import (
	"context"
	"fmt"
	"graphyy/controller/auth"
	"graphyy/model"
	"net/http"
	"strings"

	"github.com/graphql-go/graphql"
	"github.com/graphql-go/handler"
)

// Schema builds a graphql schema and returns it
func Schema(controllers *Controllers) graphql.Schema {
	schema, err := graphql.NewSchema(graphql.SchemaConfig{
		Query:    getRootQuery(controllers),
		Mutation: getRootMutation(controllers),
	})
	if err != nil {
		panic(err)
	}

	return schema
}

// GraphqlHandlfunc is a handler for the graphql endpoint.
func GraphqlHandlfunc(schema graphql.Schema) *handler.Handler {
	return handler.New(&handler.Config{
		Schema:     &schema,
		Pretty:     true,
		GraphiQL:   false,
		Playground: true,
		RootObjectFn: func(ctx context.Context, req *http.Request) map[string]interface{} {
			var user model.User
			var err error
			bearerToken := req.Header.Get("Authorization")
			splitToken := strings.Split(bearerToken, "Bearer ")
			if len(splitToken) >= 2 {
				token := splitToken[1]
				user, err = auth.VerifyJWT(token)
				if err != nil {
					fmt.Printf("%+v\n", err)
				}
			}
			return map[string]interface{}{
				"currentUser": user,
			}
		},
	})
}
