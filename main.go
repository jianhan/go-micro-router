package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/jianhan/go-micro-router/gql"
	"github.com/jianhan/go-micro-router/handler"
	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		panic("Error loading .env file")
	}

	//r := mux.NewRouter()
	//
	//// This route is always accessible
	//r.Handle("/api/public", http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
	//	message := "Hello from a public endpoint! You don't need to be authenticated to see this."
	//	responseJSON(message, w, http.StatusOK)
	//}))
	//
	//// This route is only accessible if the user has a valid access_token
	//// We are chaining the jwtmiddleware middleware into the negroni handler function which will check
	//// for a valid token.
	//r.Handle("/api/private", negroni.New(
	//	negroni.HandlerFunc(jwtMiddleware.HandlerWithNext),
	//	negroni.Wrap(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
	//		message := "Hello from a private endpoint! You need to be authenticated to see this."
	//		responseJSON(message, w, http.StatusOK)
	//	}))))
	//
	//// This route is only accessible if the user has a valid access_token with the read:messages scope
	//// We are chaining the jwtmiddleware middleware into the negroni handler function which will check
	//// for a valid token and and scope.
	//r.Handle("/api/private-scoped", negroni.New(
	//	negroni.HandlerFunc(jwtMiddleware.HandlerWithNext),
	//	negroni.Wrap(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
	//		authHeaderParts := strings.Split(r.Header.Get("Authorization"), " ")
	//		token := authHeaderParts[1]
	//
	//		hasScope := checkScope("read:messages", token)
	//
	//		if !hasScope {
	//			message := "Insufficient scope."
	//			responseJSON(message, w, http.StatusForbidden)
	//			return
	//		}
	//		message := "Hello from a private endpoint! You need to be authenticated to see this."
	//		responseJSON(message, w, http.StatusOK)
	//	}))))
	r, err := handler.GetRouter(gql.NewGQLSchemaGenerator(gql.QueryMap, nil))
	if err != nil {
		panic(err)
	}
	srv := &http.Server{
		Handler: r,
		Addr:    fmt.Sprintf("%s:%s", os.Getenv("ADDRESS"), os.Getenv("PORT")),
	}
	log.Fatal(srv.ListenAndServe())
}
