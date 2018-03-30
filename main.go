package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strings"
	"errors"
	"log"
	"os"

	"github.com/codegangsta/negroni"
	"github.com/auth0/go-jwt-middleware"
	"github.com/dgrijalva/jwt-go"
	"github.com/gorilla/mux"
	"github.com/joho/godotenv"
	"github.com/y0ssar1an/q"
	"time"
	"github.com/jianhan/go-micro-router/handler"
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
	r, err := handler.GetRouter()
	if err != nil {
		panic(err)
	}
	srv := &http.Server{
		Handler:      r,
		Addr:         fmt.Sprintf("%s:%s", os.Getenv("ADDRESS"), os.Getenv("PORT")),
	}
	log.Fatal(srv.ListenAndServe())
}





