package handler

import (
	"github.com/gorilla/mux"
	"github.com/graphql-go/handler"
	"github.com/jianhan/go-micro-router/gql"
)

func GetRouter(g gql.GQLSchemaGenerator) (*mux.Router, error) {
	r := mux.NewRouter()
	adminGQLHandler, err := adminGQLHandler(g)
	if err != nil {
		return nil, err
	}
	r.Handle("/agql", adminGQLHandler)
	return r, nil
}

func adminGQLHandler(g gql.GQLSchemaGenerator) (*handler.Handler, error) {
	schema, err := g.Generate()
	if err != nil {
		return nil, err
	}
	h := handler.New(&handler.Config{
		Schema:   &schema,
		Pretty:   true,
		GraphiQL: true,
	})
	return h, nil
}
