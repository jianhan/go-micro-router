package gql

import (
	"context"

	"github.com/golang/protobuf/ptypes"
	google_protobuf "github.com/golang/protobuf/ptypes/timestamp"
	"github.com/graphql-go/graphql"
	"github.com/graphql-go/graphql/language/ast"
	pcourse "github.com/jianhan/go-micro-courses/proto/course"
	"github.com/jianhan/pkg/gql/scalar"
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

func serializeDateTime(value interface{}) interface{} {
	switch value := value.(type) {
	case google_protobuf.Timestamp:
		t, err := ptypes.Timestamp(&value)
		if err != nil {
			return nil
		}
		buff, err := t.MarshalText()
		if err != nil {
			return nil
		}
		return string(buff)
	case *google_protobuf.Timestamp:
		return serializeDateTime(*value)
	default:
		return nil
	}

	return nil
}

func unserializeDateTime(value interface{}) interface{} {
	q.Q("Unser")
	switch value := value.(type) {
	case []byte:
		t := google_protobuf.Timestamp{}
		tt, err := ptypes.Timestamp(&t)
		q.Q(t, err)
		if err != nil {
			return nil
		}
		err = tt.UnmarshalText(value)
		if err != nil {
			return nil
		}

		return t
	case string:
		return unserializeDateTime([]byte(value))
	case *string:
		return unserializeDateTime([]byte(*value))
	default:
		return nil
	}

	return nil
}

var PGraphqlDateTime = graphql.NewScalar(graphql.ScalarConfig{
	Name: "PDateTime",
	Description: "The `DateTime` scalar type represents a DateTime." +
		" The DateTime is serialized as an RFC 3339 quoted string",
	Serialize:  serializeDateTime,
	ParseValue: unserializeDateTime,
	ParseLiteral: func(valueAST ast.Value) interface{} {
		switch valueAST := valueAST.(type) {
		case *ast.StringValue:
			return valueAST.Value
		}
		return nil
	},
})
