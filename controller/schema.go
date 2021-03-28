package controller

import (
	"graphyy/model"

	"github.com/graphql-go/graphql"
	"github.com/graphql-go/graphql/gqlerrors"
)

var authType = graphql.NewObject(graphql.ObjectConfig{
	Name: "Auth",
	Fields: graphql.Fields{
		"tokenType": &graphql.Field{
			Type: graphql.String,
		},
		"token": &graphql.Field{
			Type: graphql.String,
		},
		"expiresIn": &graphql.Field{
			Type: graphql.Int,
		},
	},
})

func getRootMutation(contrs *Controllers) *graphql.Object {
	return graphql.NewObject(graphql.ObjectConfig{
		Name: "RootMutation",
		Fields: graphql.Fields{
			"signup": &graphql.Field{
				Type:        authType, // the return type for this field
				Description: "Signup",
				Args: graphql.FieldConfigArgument{
					"username": &graphql.ArgumentConfig{
						Type: graphql.NewNonNull(graphql.String),
					},
					"password": &graphql.ArgumentConfig{
						Type: graphql.NewNonNull(graphql.String),
					},
				},
				Resolve: func(params graphql.ResolveParams) (interface{}, error) {
					// username, _ := params.Args["username"].(string)
					// password, _ := params.Args["password"].(string)
					// res, err := contrs.userController.Signup(model.User{Username: username, Password: password})
					// if err != nil {
					// 	return nil, gqlerrors.FormatError(err)
					// }
					// return res, nil
					// return
				},
			},
			"login": &graphql.Field{
				Type:        authType, // the return type for this field
				Description: "Login",
				Args: graphql.FieldConfigArgument{
					"username": &graphql.ArgumentConfig{
						Type: graphql.NewNonNull(graphql.String),
					},
					"password": &graphql.ArgumentConfig{
						Type: graphql.NewNonNull(graphql.String),
					},
				},
				Resolve: func(params graphql.ResolveParams) (interface{}, error) {
					// username, _ := params.Args["username"].(string)
					// password, _ := params.Args["password"].(string)
					// res, err := contrs.userController.Login(model.User{Username: username, Password: password})
					// if err != nil {
					// 	return nil, gqlerrors.FormatError(err)
					// }
					// return res, nil
				},
			},
		},
	})
}

func getRootQuery(contrs *Controllers) *graphql.Object {
	return graphql.NewObject(graphql.ObjectConfig{
		Name: "RootQuery",
		Fields: graphql.Fields{
			"scooter": &graphql.Field{
				Type: graphql.NewObject(graphql.ObjectConfig{
					Name: "Scooter",
					Fields: graphql.Fields{
						"latitude": &graphql.Field{
							Type: graphql.String,
						},
						"longitude": &graphql.Field{
							Type: graphql.String,
						},
					},
				}),
				Description: "Get scooters",
				Resolve: func(params graphql.ResolveParams) (interface{}, error) {
					latitude, _ := params.Args["latitude"].(string)
					longitude, _ := params.Args["longitude"].(string)
					res, err := contrs.scooterController.Login(model.Scooter{Latitude: latitude, Longitude: longitude})
					if err != nil {
						return nil, gqlerrors.FormatError(err)
					}
					return res, nil
				},
			},
		},
	})

}
