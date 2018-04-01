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
	SetRootQueryFields(fields map[string]*graphql.Field)
	SetMutationField(fields map[string]*graphql.Field)
}

type schemaGenerator struct {
	rootQuery    map[string]*graphql.Field
	rootMutation map[string]*graphql.Field
}

func NewGQLSchemaGenerator(rootQuery, rootMutation map[string]*graphql.Field) GQLSchemaGenerator {
	return &schemaGenerator{
		rootQuery:    rootQuery,
		rootMutation: rootMutation,
	}
}

func (s *schemaGenerator) SetRootQueryFields(fields map[string]*graphql.Field) {
	s.rootQuery = fields
}

func (s *schemaGenerator) SetMutationField(fields map[string]*graphql.Field) {
	s.rootMutation = fields
}

func (s *schemaGenerator) Generate() (schema graphql.Schema, err error) {
	// generate root query
	rootQueryFields, err := s.fields(RootQuery.String())
	if err != nil {
		return
	}
	rootQuery := graphql.ObjectConfig{Name: RootQuery.String(), Fields: *rootQueryFields}
	// generate root mutation
	//rootMutationFields, err := s.fields(RootMutation.String())
	//if err != nil {
	//	return
	//}
	//rootMutation := graphql.ObjectConfig{Name: RootMutation.String(), Fields: *rootMutationFields}
	// generate schema
	schema, err = graphql.NewSchema(graphql.SchemaConfig{
		Query: graphql.NewObject(rootQuery),
		//Mutation: graphql.NewObject(rootMutation),
	})
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
