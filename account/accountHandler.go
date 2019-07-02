package account

import (
	"encoding/json"
	"net/http"

	"github.com/BrunoDM2943/pock_bank-api/domain"
)

func NewAccountHandler(service IAccountService) *personHandler {
	return &personHandler{
		service,
	}
}

type personHandler struct {
	service IAccountService
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
		json.NewEncoder(w).Encode(account)
	}

}