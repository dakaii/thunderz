package controller

import (
	"strconv"

	"github.com/graphql-go/graphql"
	"github.com/graphql-go/graphql/gqlerrors"
)

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
					strLatitude, _ := params.Args["latitude"].(string)
					strLongitude, _ := params.Args["longitude"].(string)
					strDistance, _ := params.Args["distance"].(string)

					latitude, err := strconv.ParseFloat(strLatitude, 64)
					longitude, err := strconv.ParseFloat(strLongitude, 64)
					distance, _ := strconv.Atoi(strDistance)

					res, err := contrs.scooterController.GetNearbyScooters(latitude, longitude, distance)
					if err != nil {
						return nil, gqlerrors.FormatError(err)
					}
					return res, nil
				},
			},
		},
	})

}
