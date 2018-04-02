package gql

import (
	"github.com/graphql-go/graphql"
	pcourse "github.com/jianhan/go-micro-courses/proto/course"
)

type mutationGenerator struct {
	coursesClient pcourse.CourseServiceClient
}

func NewMutationGenerator(coursesClient pcourse.CourseServiceClient) QueryMutationGenerator {
	return &mutationGenerator{
		coursesClient: coursesClient,
	}
}

func (q *mutationGenerator) Generate() graphql.Fields {
	return graphql.Fields{
		"createCourse": createCourse(q.coursesClient),
	}
}
