package gql

import (
	"context"
	"errors"
	"time"

	"github.com/golang/protobuf/ptypes"
	slugMaker "github.com/gosimple/slug"
	"github.com/graphql-go/graphql"
	pcourse "github.com/jianhan/go-micro-courses/proto/course"
	"github.com/jianhan/pkg/gql/scalar"
)

func upsertCourse(coursesClient pcourse.CourseServiceClient) *graphql.Field {
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
			"category_ids": &graphql.ArgumentConfig{
				Type: graphql.NewList(graphql.String),
			},
		},
		Resolve: func(params graphql.ResolveParams) (interface{}, error) {
			id, _ := params.Args["id"].(string)
			name, _ := params.Args["name"].(string)
			slug, _ := params.Args["slug"].(string)
			if slug == "" {
				slug = slugMaker.Make(name)
			}
			displayOrder, _ := params.Args["display_order"].(int)
			start, _ := params.Args["start"].(string)
			startTime, err := time.Parse(time.RFC3339, start)
			if err != nil {
				return nil, errors.New("start time is not a valid RFC3339 format")
			}
			startTimeProto, err := ptypes.TimestampProto(startTime)
			if err != nil {
				return nil, err
			}
			end, _ := params.Args["end"].(string)
			endTime, err := time.Parse(time.RFC3339, end)
			if err != nil {
				return nil, errors.New("end time is not a valid RFC3339 format")
			}
			endTimeProto, err := ptypes.TimestampProto(endTime)
			if err != nil {
				return nil, err
			}
			description, _ := params.Args["description"].(string)
			hidden, _ := params.Args["hidden"].(bool)
			cids, _ := params.Args["category_ids"].([]interface{})
			categoryIDs := make([]string, 0, len(cids))
			for _, v := range cids {
				if cid, ok := v.(string); ok {
					categoryIDs = append(categoryIDs, cid)
				}
			}
			rsp, err := coursesClient.UpsertCourse(
				context.Background(),
				&pcourse.UpsertCourseReq{
					Course: &pcourse.Course{
						ID:           id,
						Name:         name,
						Slug:         slug,
						Description:  description,
						DisplayOrder: uint64(displayOrder),
						Hidden:       hidden,
						Start:        startTimeProto,
						End:          endTimeProto,
						CategoryIds:  categoryIDs,
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
