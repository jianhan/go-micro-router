package gql

import (
	"github.com/graphql-go/graphql"
	pcourse "github.com/jianhan/go-micro-courses/proto/course"
)

type queryGenerator struct {
	coursesClient pcourse.CoursesClient
}

func NewQueryGenerator(coursesClient pcourse.CoursesClient) QueryMutationGenerator {
	return &queryGenerator{
		coursesClient: coursesClient,
	}
}

func (q *queryGenerator) Generate() graphql.Fields {
	return graphql.Fields{
		"courses": getCourseQuery(q.coursesClient),
	}
}
