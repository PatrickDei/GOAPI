package app

import (
	"encoding/json"
	"github.com/gorilla/mux"
	"net/http"
	"staycation/service"
)

type StaycationUserHandler struct {
	service service.StaycationUserService
}

func (suh StaycationUserHandler) getAllStaycationUsers(w http.ResponseWriter, r *http.Request) {
	staycationUsers, _ := suh.service.GetAllStaycationUsers()

	w.Header().Add("Content-Type", "application/json")

	json.NewEncoder(w).Encode(staycationUsers)
}

func (suh StaycationUserHandler) getStaycationUserById(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)

	id := vars["user_id"]

	su, err := suh.service.GetStaycationUserById(id)

	if err != nil {
		writeResponse(w, err.Code, err.AsMessage())
	} else {
		writeResponse(w, http.StatusOK, su)
	}
}

func writeResponse(w http.ResponseWriter, code int, data interface{}) {
	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(code)
	if err := json.NewEncoder(w).Encode(data); err != nil {
		panic(err)
	}
}
