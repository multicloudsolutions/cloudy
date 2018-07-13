package account

import (
	"github.com/gorilla/mux"
)

// Router for /account
func Router() *mux.Router {
	accountSubRouter := mux.NewRouter().PathPrefix("/account").Subrouter()
	accountSubRouter.HandleFunc("/", Create).Methods("POST")
	accountSubRouter.HandleFunc("/", GetDetails).Methods("GET")
	accountSubRouter.HandleFunc("/", Update).Methods("PUT")
	accountSubRouter.HandleFunc("/", Remove).Methods("DELETE")
	return accountSubRouter
}
