package usecase

import (
	"time"
	"br.com.jonathan/psm/api/domain"
)

func CreateTransaction(transaction domain.Transaction) domain.Transaction {
	transaction.ID = 1
	transaction.Date = time.Now().Format(time.RFC3339)
	return transaction
}
