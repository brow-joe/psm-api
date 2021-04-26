package controller

import (
	"strconv"
    "net/http"
	"encoding/json"
	"github.com/gorilla/mux"
	"br.com.jonathan/psm/api/domain"
	"br.com.jonathan/psm/api/usecases"
)

func PostAccountHandler(w http.ResponseWriter, r *http.Request){
	w.Header().Set("Content-Type", "application/json")
	var account domain.Account
    _ = json.NewDecoder(r.Body).Decode(&account)
	switch {
		case account == domain.Account{}:
			w.WriteHeader(http.StatusBadRequest)
		default:
			var response = usecase.CreateAccount(account)
			w.WriteHeader(http.StatusCreated)
			json.NewEncoder(w).Encode(response)
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
			var response = usecase.RetrieveAccount(id)
			w.WriteHeader(http.StatusOK)
			json.NewEncoder(w).Encode(response)
	}
}
