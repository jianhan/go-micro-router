package gql

import (
	"context"

	"github.com/graphql-go/graphql"
	pcourse "github.com/jianhan/go-micro-courses/proto/course"
	"github.com/jianhan/pkg/gql/scalar"
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
		"slug": &graphql.Field{
			Type: graphql.String,
		},
		"display_order": &graphql.Field{
			Type: graphql.Int,
		},
		"start": &graphql.Field{
			Type: scalar.ProtoDateTime,
		},
		"end": &graphql.Field{
			Type: scalar.ProtoDateTime,
		},
		"description": &graphql.Field{
			Type: graphql.String,
		},
	},
})

func getCoursesQuery(coursesClient pcourse.CoursesClient) *graphql.Field {
	return &graphql.Field{
		Type:        graphql.NewList(courseType),
		Description: "Get courses",
		Args: graphql.FieldConfigArgument{
			"ids": &graphql.ArgumentConfig{
				Type: graphql.NewList(graphql.String),
			},
		},
		Resolve: func(params graphql.ResolveParams) (interface{}, error) {
			courses, err := coursesClient.FindCourses(context.Background(), &pcourse.FindCoursesRequest{})
			if err != nil {
				return nil, err
			}
			return courses.Courses, nil
		},
	}
}
