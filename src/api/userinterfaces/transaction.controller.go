package controller

import (
    "net/http"
	"encoding/json"
	"br.com.jonathan/psm/api/domain"
	"br.com.jonathan/psm/api/usecase"
)

func PostTransactionHandler(w http.ResponseWriter, r *http.Request){
	w.Header().Set("Content-Type", "application/json")
	var transaction domain.Transaction
    _ = json.NewDecoder(r.Body).Decode(&transaction)
	switch {
		case transaction == domain.Transaction{}:
			w.WriteHeader(http.StatusBadRequest)
		default:
			w.WriteHeader(http.StatusCreated)
			json.NewEncoder(w).Encode(
				usecase.CreateTransaction(transaction)
			)
	}
}
