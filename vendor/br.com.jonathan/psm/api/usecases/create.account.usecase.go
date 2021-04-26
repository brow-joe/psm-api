package usecase

import (
	"br.com.jonathan/psm/api/domain"
)

func CreateAccount(account domain.Account) domain.Account {
	account.ID = 1
	return account
}
