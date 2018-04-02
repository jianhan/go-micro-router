package gql

import (
	"github.com/graphql-go/graphql"
	pcourse "github.com/jianhan/go-micro-courses/proto/course"
)

type mutationGenerator struct {
	coursesClient pcourse.CoursesClient
}

func NewMutationGenerator(coursesClient pcourse.CoursesClient) QueryMutationGenerator {
	return &mutationGenerator{
		coursesClient: coursesClient,
	}
}

func (q *mutationGenerator) Generate() graphql.Fields {
	return graphql.Fields{
		"courses": getCourseMutation(q.coursesClient),
	}
}
