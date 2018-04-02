package gql

import (
	"github.com/graphql-go/graphql"
	pcourse "github.com/jianhan/go-micro-courses/proto/course"
)

type SchemaConfig int

const (
	RootQuery SchemaConfig = iota
	RootMutation
)

var configTypes = [...]string{
	"RootQuery",
	"RootMutation",
}

func (s SchemaConfig) String() string { return configTypes[s] }

type GQLField func(courseClient *pcourse.CoursesClient) (string, *graphql.Field)

type GQLSchemaGenerator interface {
	Generate() (schema graphql.Schema, err error)
}

type schemaGenerator struct {
	rootQueryGenerator    QueryMutationGenerator
	rootMutationGenerator QueryMutationGenerator
}

func NewGQLSchemaGenerator(rqg, rmg QueryMutationGenerator) GQLSchemaGenerator {
	return &schemaGenerator{
		rootQueryGenerator:    rqg,
		rootMutationGenerator: rmg,
	}
}

func (s *schemaGenerator) Generate() (schema graphql.Schema, err error) {
	rootQuery := graphql.ObjectConfig{Name: RootQuery.String(), Fields: s.rootQueryGenerator.Generate()}
	rootMutation := graphql.ObjectConfig{Name: RootQuery.String(), Fields: s.rootMutationGenerator.Generate()}
	schema, err = graphql.NewSchema(graphql.SchemaConfig{
		Query:    graphql.NewObject(rootQuery),
		Mutation: graphql.NewObject(rootMutation),
	})
	if err != nil {
		return
	}
	return schema, nil
}
