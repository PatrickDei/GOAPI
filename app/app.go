package app

import (
	"github.com/gorilla/mux"
	"log"
	"net/http"
	"staycation/domain"
	"staycation/service"
)

func Start() {

	router := mux.NewRouter()

	// wiring
	suh := StaycationUserHandler{service: service.NewStaycationUserService(domain.NewStaycationRepositoryDb())}

	router.HandleFunc("/users", suh.getAllStaycationUsers).Methods(http.MethodGet)

	router.HandleFunc("/users/{user_id:[0-9]+}", suh.getStaycationUserById).Methods(http.MethodGet)

	/*
		router.HandleFunc("/customers/{customer_id:[0-9]+}", getCustomer).Methods(http.MethodGet)

		router.HandleFunc("/customers", createCustomer).Methods(http.MethodPost)*/

	log.Fatal(http.ListenAndServe("localhost:8000", router))
}
