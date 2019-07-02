package main

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/BrunoDM2943/pock_bank-api/infra"
	"github.com/BrunoDM2943/pock_bank-api/account"
)

func main() {
	r := mux.NewRouter()
	db := infra.InitDB()
	accountHandler := account.NewAccountHandler(
		account.NewAccountService(
			account.NewAccountRepo(
				db,
			),
		),
	)
	r.HandleFunc("/ping", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		w.Write([]byte("{pong}"))
	})
	r.HandleFunc("/person", accountHandler.PostAccount).Methods("POST")
	log.Fatal(http.ListenAndServe(":8000", r))
}