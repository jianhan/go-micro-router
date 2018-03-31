package gql

import (
	"github.com/graphql-go/graphql"
	"github.com/jianhan/go-micro-courses/proto/course"
)

var QueryMap = map[string]*graphql.Field{}

func init() {
	QueryMap["courses"] = getCoursesQuery(course.NewCoursesClient("", nil))
}