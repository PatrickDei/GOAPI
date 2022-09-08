package app

import (
	"encoding/json"
	"github.com/gorilla/mux"
	"net/http"
	"staycation/dto"
	"staycation/service"
)

type AccountHandler struct {
	service service.AccountService
}

func (ah AccountHandler) createAccount(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	userId := vars["user_id"]

	var request dto.AccountRequest
	err := json.NewDecoder(r.Body).Decode(&request)
	if err != nil {
		writeResponse(w, http.StatusBadRequest, err.Error())
	} else {
		request.UserId = userId
		account, err := ah.service.NewAccount(request)
		if err != nil {
			writeResponse(w, err.Code, err.Message)
		}

		writeResponse(w, http.StatusCreated, account)
	}
}
