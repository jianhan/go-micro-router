package gql

import (
	"github.com/gosimple/slug"
	"github.com/graphql-go/graphql"
	pcourse "github.com/jianhan/go-micro-courses/proto/course"
	"github.com/jianhan/pkg/gql/scalar"
	"github.com/y0ssar1an/q"
)

func getCourseMutation(coursesClient pcourse.CoursesClient) *graphql.Field {
	return &graphql.Field{
		Type:        courseType,
		Description: "Create new course",
		Args: graphql.FieldConfigArgument{
			"name": &graphql.ArgumentConfig{
				Type: graphql.NewNonNull(graphql.String),
			},
			"display_order": &graphql.ArgumentConfig{
				Type: graphql.Int,
			},
			"start": &graphql.ArgumentConfig{
				Type: graphql.NewNonNull(scalar.ProtoDateTime),
			},
			"end": &graphql.ArgumentConfig{
				Type: graphql.NewNonNull(scalar.ProtoDateTime),
			},
			"hidden": &graphql.ArgumentConfig{
				Type: graphql.Boolean,
			},
			"description": &graphql.ArgumentConfig{
				Type: graphql.NewNonNull(graphql.String),
			},
		},
		Resolve: func(params graphql.ResolveParams) (interface{}, error) {
			name, _ := params.Args["name"].(string)
			slug := slug.Make(name)
			displayOrder, _ := params.Args["display_order"].(int)
			q.Q(name, slug, displayOrder)
			return nil, nil

		},
	}
}
