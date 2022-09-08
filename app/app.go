package app

import (
	"fmt"
	"github.com/gorilla/mux"
	"log"
	"net/http"
	"os"
	"staycation/domain"
	"staycation/service"
)

func sanityCheck() {
	// and so on
	if os.Getenv("SERVER_ADDRESS") == "" ||
		os.Getenv("SERVER_PORT") == "" {
		log.Fatal("Environment variable not defined...")
	}
}

func Start() {

	sanityCheck()

	router := mux.NewRouter()

	// wiring
	suh := StaycationUserHandler{service: service.NewStaycationUserService(domain.NewStaycationRepositoryDb())}

	router.HandleFunc("/users", suh.getAllStaycationUsers).Methods(http.MethodGet)

	router.HandleFunc("/users/{user_id:[0-9]+}", suh.getStaycationUserById).Methods(http.MethodGet)

	/*
		router.HandleFunc("/customers/{customer_id:[0-9]+}", getCustomer).Methods(http.MethodGet)

		router.HandleFunc("/customers", createCustomer).Methods(http.MethodPost)*/

	address := os.Getenv("SERVER_ADDRESS")
	port := os.Getenv("SERVER_PORT")
	log.Fatal(http.ListenAndServe(fmt.Sprintf("%s:%s", address, port), router))
}
