package controller

import (
	"graphyy/model"

	"github.com/graphql-go/graphql"
	"github.com/graphql-go/graphql/gqlerrors"
)

var pointInterface = graphql.NewObject(graphql.ObjectConfig{
	Name: "Point",
	Fields: graphql.Fields{
		"latitude": &graphql.Field{
			Type: graphql.NewNonNull(graphql.Float),
			Resolve: func(p graphql.ResolveParams) (interface{}, error) {
				if point, ok := p.Source.(model.Point); ok {
					return point.Location.Coordinates[1], nil
				}
				return nil, nil
			},
		},
		"longitude": &graphql.Field{
			Type: graphql.NewNonNull(graphql.Float),
			Resolve: func(p graphql.ResolveParams) (interface{}, error) {
				if point, ok := p.Source.(model.Point); ok {
					return point.Location.Coordinates[0], nil
				}
				return nil, nil
			},
		},
	},
})

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
				Type:        graphql.NewList(pointInterface),
				Description: "Get the scooters within the specified certain distance",
				Resolve: func(params graphql.ResolveParams) (interface{}, error) {
					latitude, _ := params.Args["latitude"].(float64)
					longitude, _ := params.Args["longitude"].(float64)
					distance, _ := params.Args["distance"].(int)
					limit, _ := params.Args["limit"].(int)

					res, err := contrs.scooterController.GetNearbyScooters(latitude, longitude, int64(distance), int64(limit))
					if err != nil {
						return nil, gqlerrors.FormatError(err)
					}
					return res, nil
				},
			},
		},
	})

}
