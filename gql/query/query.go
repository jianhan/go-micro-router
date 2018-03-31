package query

import (
	"github.com/graphql-go/graphql"
	"github.com/jianhan/go-micro-courses/proto/course"
)

var QueryMap = map[string]*graphql.Field{}

func init() {
	QueryMap["course"] = getCoursesQuery(course.NewCoursesClient("", nil))
}
