package handler

import "github.com/gorilla/mux"

func GetRouter() (*mux.Router, error) {
	r := mux.NewRouter()
	r.Handle("/agql", GetAdminHandler())
	return r, nil
}
