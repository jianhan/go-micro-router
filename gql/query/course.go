package query

import (
	"github.com/graphql-go/graphql"
	"github.com/y0ssar1an/q"
)

var courseType = graphql.NewObject(graphql.ObjectConfig{
	Name: "Course",
	Fields: graphql.Fields{
		"id": &graphql.Field{
			Type: graphql.String,
		},
		"name": &graphql.Field{
			Type: graphql.String,
		},
		"description": &graphql.Field{
			Type: graphql.Boolean,
		},
	},
})

func getCoursesQuery() *graphql.Field {
	return &graphql.Field{
		Type:        courseType,
		Description: "Get courses",
		Args: graphql.FieldConfigArgument{
			"ids": &graphql.ArgumentConfig{
				Type: graphql.NewList(graphql.String),
			},
		},
		Resolve: func(params graphql.ResolveParams) (interface{}, error) {
			q.Q(params)
			return nil, nil
		},
	}
}
