package controller

import (
	"github.com/graphql-go/graphql"
	"github.com/graphql-go/graphql/gqlerrors"
)

func getRootQuery(contrs *Controllers) *graphql.Object {
	return graphql.NewObject(graphql.ObjectConfig{
		Name: "RootQuery",
		Fields: graphql.Fields{
			"scooter": &graphql.Field{
				Args: graphql.FieldConfigArgument{
					"latitude": &graphql.ArgumentConfig{
						Type: graphql.NewNonNull(graphql.Float),
					},
					"longitude": &graphql.ArgumentConfig{
						Type: graphql.NewNonNull(graphql.Float),
					},
					"distance": &graphql.ArgumentConfig{
						Type: graphql.NewNonNull(graphql.Int),
					},
					"limit": &graphql.ArgumentConfig{
						Type: graphql.NewNonNull(graphql.Int),
					},
				},
				Type: graphql.NewObject(graphql.ObjectConfig{
					Name: "Scooter",
					Fields: graphql.Fields{
						"latitude": &graphql.Field{
							Type: graphql.Float,
						},
						"longitude": &graphql.Field{
							Type: graphql.Float,
						},
					},
				}),
				Description: "Get scooters",
				Resolve: func(params graphql.ResolveParams) (interface{}, error) {
					latitude, _ := params.Args["latitude"].(float64)
					longitude, _ := params.Args["longitude"].(float64)
					distance, _ := params.Args["distance"].(int)
					limit, _ := params.Args["limit"].(int64)

					// latitude, err := strconv.ParseFloat(strLatitude, 64)
					// if err != nil {
					// 	return nil, gqlerrors.FormatError(err)
					// }
					// longitude, err := strconv.ParseFloat(strLongitude, 64)
					// if err != nil {
					// 	return nil, gqlerrors.FormatError(err)
					// }
					// distance, err := strconv.Atoi(strDistance)
					// if err != nil {
					// 	distance = 1000
					// }
					// limit, err := strconv.ParseInt(strLimit, 10, 64)
					// if err != nil {
					// 	limit = 20
					// }

					res, err := contrs.scooterController.GetNearbyScooters(latitude, longitude, distance, limit)
					if err != nil {
						return nil, gqlerrors.FormatError(err)
					}
					return res, nil
				},
			},
		},
	})

}
