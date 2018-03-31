package gql

import (
	"errors"

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
	rootQuery    map[string]*graphql.Field
	rootMutation map[string]*graphql.Field
}

func NewGQLSchemaGenerator(rootQueries, rootMutations map[string]*graphql.Field) GQLSchemaGenerator {
	return &schemaGenerator{
		rootQuery:    rootQueries,
		rootMutation: rootMutations,
	}
}

func (s *schemaGenerator) Generate() (schema graphql.Schema, err error) {
	// generate root query
	rootQueryFields, err := s.fields(RootQuery.String())
	if err != nil {
		return
	}
	rootQuery := graphql.ObjectConfig{Name: RootQuery.String(), Fields: rootQueryFields}
	// generate root mutation
	rootMutationFields, err := s.fields(RootMutation.String())
	if err != nil {
		return
	}
	rootMutation := graphql.ObjectConfig{Name: RootMutation.String(), rootMutationFields}
	// generate schema
	schemaConfig := graphql.SchemaConfig{Query: graphql.NewObject(rootQuery), Mutation: graphql.NewObject(rootMutation)}
	schema, err = graphql.NewSchema(schemaConfig)
	if err != nil {
		return
	}
	return schema, nil
}

func (s *schemaGenerator) fields(fieldType string) (*graphql.Fields, error) {
	target := map[string]*graphql.Field{}
	if fieldType == RootQuery.String() {
		target = s.rootQuery
	} else if fieldType == RootMutation.String() {
		target = s.rootMutation
	} else {
		return nil, errors.New("field type must be queries or mutations")
	}
	fields := graphql.Fields{}
	for k, v := range target {
		fields[k] = v
	}
	return &fields, nil
}
