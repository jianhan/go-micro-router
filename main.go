package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	pcourse "github.com/jianhan/go-micro-courses/proto/course"
	"github.com/jianhan/go-micro-router/gql"
	"github.com/jianhan/go-micro-router/handler"
	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		panic("Error loading .env file")
	}
	r, err := handler.GetRouter(gql.NewGQLSchemaGenerator(
		gql.NewQueryGenerator(pcourse.NewCourseServiceClient("", nil)),
		gql.NewMutationGenerator(pcourse.NewCourseServiceClient("", nil)),
	))
	if err != nil {
		panic(err)
	}
	srv := &http.Server{
		Handler: r,
		Addr:    fmt.Sprintf("%s:%s", os.Getenv("ADDRESS"), os.Getenv("PORT")),
	}
	log.Fatal(srv.ListenAndServe())
}
