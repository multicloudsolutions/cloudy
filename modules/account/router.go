package account

import (
	"github.com/gorilla/mux"
)

// Router for /account
func Router() *mux.Router {
	accountSubRouter := mux.NewRouter().PathPrefix("/account").Subrouter()
	accountSubRouter.HandleFunc("/create", Create).Methods("POST")
	accountSubRouter.HandleFunc("/read", GetDetails).Methods("GET")
	accountSubRouter.HandleFunc("/update", Update).Methods("PUT")
	accountSubRouter.HandleFunc("/delete", Remove).Methods("DELETE")
	return accountSubRouter
}
