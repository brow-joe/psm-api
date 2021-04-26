package usecase

import (
	"br.com.jonathan/psm/api/domain"
)

func RetrieveAccount(id int) domain.Account {
	return domain.Account{ID: id, DocumentNumber: "12345678900", AvailableCreditLimit: 5000}
}
