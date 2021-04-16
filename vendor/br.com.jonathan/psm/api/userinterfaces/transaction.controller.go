package controller

import (
	"time"
    "net/http"
	"encoding/json"
	"br.com.jonathan/psm/api/domain"
)

func PostTransactionHandler(w http.ResponseWriter, r *http.Request){
	w.Header().Set("Content-Type", "application/json")
	var transaction domain.Transaction
    _ = json.NewDecoder(r.Body).Decode(&transaction)
	switch {
		case transaction == domain.Transaction{}:
			w.WriteHeader(http.StatusBadRequest)
		default:
			transaction.ID = 1
			transaction.Date = time.Now().Format(time.RFC3339)
			w.WriteHeader(http.StatusCreated)
			json.NewEncoder(w).Encode(transaction)
	}
}
