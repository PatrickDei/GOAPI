package app

import (
	"fmt"
	"github.com/gorilla/mux"
	"github.com/jmoiron/sqlx"
	"log"
	"net/http"
	"os"
	"staycation/domain"
	"staycation/service"
	"time"
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
	dbClient := getDbClient()

	suh := StaycationUserHandler{service: service.NewStaycationUserService(domain.NewStaycationRepositoryDb(dbClient))}
	ah := AccountHandler{service: service.NewAccountService(domain.NewAccountRepositoryDb(dbClient))}

	router.HandleFunc("/users", suh.getAllStaycationUsers).Methods(http.MethodGet)
	router.HandleFunc("/users/{user_id:[0-9]+}", suh.getStaycationUserById).Methods(http.MethodGet)
	router.HandleFunc("/users/{user_id:[0-9]+}/account", ah.createAccount).Methods(http.MethodPost)

	address := os.Getenv("SERVER_ADDRESS")
	port := os.Getenv("SERVER_PORT")
	log.Fatal(http.ListenAndServe(fmt.Sprintf("%s:%s", address, port), router))
}

func getDbClient() *sqlx.DB {
	user := os.Getenv("DB_USER")
	password := os.Getenv("DB_PASSWORD")
	address := os.Getenv("DB_ADDRESS")
	port := os.Getenv("DB_PORT")
	dbName := os.Getenv("DB_NAME")

	datasource := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s", user, password, address, port, dbName)

	db, err := sqlx.Open("mysql", datasource)

	if err != nil {
		panic(err)
	}

	db.SetConnMaxLifetime(time.Minute * 3)
	db.SetMaxOpenConns(10)
	db.SetMaxIdleConns(10)

	return db
}
