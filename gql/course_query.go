package gql

import (
	"context"

	"time"

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
		"hidden": &graphql.Field{
			Type: graphql.Boolean,
		},
		"description": &graphql.Field{
			Type: graphql.String,
		},
	},
})

func getCourseQuery(coursesClient pcourse.CoursesClient) *graphql.Field {
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
			// ids
			if ids, found := params.Args["ids"]; found {
				if idSlice, ok := ids.([]interface{}); ok {
					for _, v := range idSlice {
						req.Ids = append(req.Ids, v.(string))
					}
				}
			}
			// start
			startProtoTimestamp, found, err := GetProtoTimestamp(params, "start", time.RFC3339)
			if err != nil {
				return nil, err
			}
			if found {
				req.Start = startProtoTimestamp
			}
			// end
			endProtoTimestamp, found, err := GetProtoTimestamp(params, "end", time.RFC3339)
			if err != nil {
				return nil, err
			}
			if found {
				req.End = endProtoTimestamp
			}
			// inclusive
			if in, found := params.Args["inclusive"]; found {
				if in, ok := in.(bool); ok {
					req.Inclusive = in
				}
			}
			// sort
			if s, found := params.Args["sort"]; found {
				if s, ok := s.([]interface{}); ok {
					for _, v := range s {
						req.Sort = append(req.Sort, v.(string))
					}
				}
			}
			// per page
			if i, found := params.Args["per_page"]; found {
				if i, ok := i.(int); ok {
					req.PerPage = int64(i)
				}
			}
			// current page
			if i, found := params.Args["current_page"]; found {
				if i, ok := i.(int); ok {
					req.PerPage = int64(i)
				}
			}
			courses, err := coursesClient.FindCourses(context.Background(), req)
			if err != nil {
				return nil, err
			}
			return courses.Courses, nil
		},
	}
}
