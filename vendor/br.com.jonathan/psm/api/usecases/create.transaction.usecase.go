package usecase

import (
	"time"
	"net/http"
	"br.com.jonathan/psm/api/domain"
)

func CreateTransaction(w http.ResponseWriter, account domain.Account, transaction domain.Transaction) domain.Transaction {
	transaction.ID = 1
	transaction.Date = time.Now().Format(time.RFC3339)
	if transaction.OperationType == 1 {
		if account.AvailableCreditLimit >= transaction.Amount {
			account.AvailableCreditLimit = account.AvailableCreditLimit - transaction.Amount;
		} else {
			w.WriteHeader(http.StatusBadRequest)
		}
	} else if transaction.OperationType == 4 {
		account.AvailableCreditLimit = account.AvailableCreditLimit + transaction.Amount;
	}
	return transaction
}
