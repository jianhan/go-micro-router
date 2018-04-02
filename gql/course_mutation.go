package gql

import (
	"context"

	"github.com/gosimple/slug"
	"github.com/graphql-go/graphql"
	pcourse "github.com/jianhan/go-micro-courses/proto/course"
	"github.com/jianhan/pkg/gql/scalar"
	"github.com/y0ssar1an/q"
)

func createCourse(coursesClient pcourse.CourseServiceClient) *graphql.Field {
	return &graphql.Field{
		Type:        courseType,
		Description: "Create new course",
		Args: graphql.FieldConfigArgument{
			"id": &graphql.ArgumentConfig{
				Type: graphql.String,
			},
			"name": &graphql.ArgumentConfig{
				Type: graphql.NewNonNull(graphql.String),
			},
			"slug": &graphql.ArgumentConfig{
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
			id, _ := params.Args["id"].(string)
			name, _ := params.Args["name"].(string)
			slug := slug.Make(name)
			displayOrder, _ := params.Args["display_order"].(int)
			start, _ := params.Args["start"].(string)
			end, _ := params.Args["end"].(string)
			q.Q(name, slug, displayOrder, start, end)
			rsp, err := coursesClient.UpsertCourse(
				context.Background(),
				&pcourse.UpsertCourseReq{
					Course: &pcourse.Course{
						ID:   id,
						Name: name,
						Slug: slug,
					},
				},
			)
			if err != nil {
				return nil, err
			}
			return rsp.Course, nil

		},
	}
}
