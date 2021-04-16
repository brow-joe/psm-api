package application

import (
    "net/http"
	"br.com.jonathan/psm/api/userinterfaces"
)

type Endpoint struct {
    path string
    method string
    handler func(http.ResponseWriter, *http.Request)
}

func getEndpoints() []Endpoint {
	return []Endpoint {
		Endpoint{path: "/accounts", method: "POST", handler: controller.PostAccountHandler},
		Endpoint{path: "/accounts/{accountId}", method: "GET", handler: controller.GetAccountHandler},
		Endpoint{path: "/transactions", method: "POST", handler: controller.PostTransactionHandler},
    }
}
