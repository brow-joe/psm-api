package controller

import (
	"strconv"
    "net/http"
	"encoding/json"
	"github.com/gorilla/mux"
	"br.com.jonathan/psm/api/domain"
)

func PostAccountHandler(w http.ResponseWriter, r *http.Request){
	w.Header().Set("Content-Type", "application/json")
	var account domain.Account
    _ = json.NewDecoder(r.Body).Decode(&account)
	switch {
		case account == domain.Account{}:
			w.WriteHeader(http.StatusBadRequest)
		default:
			account.ID = 1
			w.WriteHeader(http.StatusCreated)
			json.NewEncoder(w).Encode(account)
	}
}

func GetAccountHandler(w http.ResponseWriter, r *http.Request){
	w.Header().Set("Content-Type", "application/json")
	var params = mux.Vars(r)
	var id, _ = strconv.Atoi(params["accountId"])
	switch {
		case id < 1:
			w.WriteHeader(http.StatusNotFound)
		default:
			var account = domain.Account{ID: id, DocumentNumber: "12345678900"}
			w.WriteHeader(http.StatusOK)
			json.NewEncoder(w).Encode(account)
	}
}
