package account

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/BrunoDM2943/pock_bank-api/domain"
	"github.com/gorilla/mux"
)

func NewAccountHandler(service Service) *personHandler {
	return &personHandler{
		service,
	}
}

type personHandler struct {
	service Service
}

func (handler personHandler) PostAccount(w http.ResponseWriter, r *http.Request) {
	account := &domain.Account{}
	json.NewDecoder(r.Body).Decode(account)
	err := handler.service.SaveAccount(account)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(err)
	} else {
		w.WriteHeader(http.StatusCreated)
		json.NewEncoder(w).Encode(map[string]interface{}{
			"id":      account.ID,
			"message": "Account created sucessfully",
		})
	}
}

func (handler personHandler) GetAccount(w http.ResponseWriter, r *http.Request) {
	id, _ := strconv.ParseInt(mux.Vars(r)["id"], 10, 64)
	account, err := handler.service.GetAccount(id)

	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(err)
	} else {
		w.WriteHeader(http.StatusCreated)
		json.NewEncoder(w).Encode(account)
	}
}
