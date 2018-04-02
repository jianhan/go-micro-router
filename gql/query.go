package gql

import (
	"github.com/graphql-go/graphql"
	pcourse "github.com/jianhan/go-micro-courses/proto/course"
)

type queryGenerator struct {
	coursesClient pcourse.CourseServiceClient
}

func NewQueryGenerator(coursesClient pcourse.CourseServiceClient) QueryMutationGenerator {
	return &queryGenerator{
		coursesClient: coursesClient,
	}
}

func (q *queryGenerator) Generate() graphql.Fields {
	return graphql.Fields{
		"courses": courses(q.coursesClient),
	}
}
