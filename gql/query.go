package gql

import (
	"github.com/graphql-go/graphql"
	pcourse "github.com/jianhan/go-micro-courses/proto/course"
)

type QueryMutationGenerator interface {
	Generate() map[string]*graphql.Field
}

type queryGenerator struct {
	coursesClient pcourse.CoursesClient
}

func NewQueryGenerator(coursesClient pcourse.CoursesClient) QueryMutationGenerator {
	return &queryGenerator{
		coursesClient: coursesClient,
	}
}

func (q *queryGenerator) Generate() map[string]*graphql.Field {
	return map[string]*graphql.Field{
		// "courses": getCourseQuery(course.NewCoursesClient("", nil)),
		"courses": getCourseQuery(q.coursesClient),
	}
}
