package gql

import (
	"context"
	"errors"

	"time"

	"github.com/golang/protobuf/ptypes"
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
			"query": &graphql.ArgumentConfig{
				Type: graphql.String,
			},
			"start": &graphql.ArgumentConfig{
				Type: scalar.ProtoDateTime,
			},
			"end": &graphql.ArgumentConfig{
				Type: scalar.ProtoDateTime,
			},
			"inclusive": &graphql.ArgumentConfig{
				Type: graphql.Boolean,
			},
			"sort": &graphql.ArgumentConfig{
				Type: graphql.NewList(graphql.String),
			},
			"per_page": &graphql.ArgumentConfig{
				Type: graphql.Int,
			},
			"current_page": &graphql.ArgumentConfig{
				Type: graphql.Int,
			},
		},
		Resolve: func(params graphql.ResolveParams) (interface{}, error) {
			req := &pcourse.FindCoursesRequest{}
			if startStr, ok := params.Args["start"].(string); ok {
				t, err := time.Parse(time.RFC3339, startStr)
				if err != nil {
					return nil, errors.New("start time is not a valid RFC3339 format")
				}
				pt, err := ptypes.TimestampProto(t)
				if err != nil {
					return nil, errors.New("can not parse start time")
				}
				req.Start = pt
			}
			//govalidator.IsTime(params.Args["start"])
			courses, err := coursesClient.FindCourses(context.Background(), req)
			if err != nil {
				return nil, err
			}
			return courses.Courses, nil
		},
	}
}
