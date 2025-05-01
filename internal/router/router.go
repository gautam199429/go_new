package router

import (
	"entitlements/internal/handler"

	"github.com/gorilla/mux"
)

func NewRouter() *mux.Router {

	r := mux.NewRouter()

	r.HandleFunc("/parse-graphql", handler.ParseGraphQLQuery).Methods("POST")

	r.HandleFunc("/parse-graphql/copy", handler.ParseGraphQLQueryCopy).Methods("POST")

	r.HandleFunc("/parse-graphql/test", handler.GetApolloPoliciesRequired).Methods("GET")

	return r

}
