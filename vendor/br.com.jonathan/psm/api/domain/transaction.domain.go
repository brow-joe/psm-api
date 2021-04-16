package domain

type Transaction struct {
	ID        			int 	  	`json:"id,omitempty"`
    AccountId 			int 		`json:"account_id,omitempty"`
	OperationType 		int 		`json:"operation_type_id,omitempty"`
	Amount	 			float64 	`json:"amount,omitempty"`
	Date	 			string	 	`json:"date,omitempty"`
}
