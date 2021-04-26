package domain

type Account struct {
	ID        				int 	  	`json:"id,omitempty"`
    DocumentNumber 			string 		`json:"document_number,omitempty"`
	AvailableCreditLimit	 float64 	`json:"available_credit_limit,omitempty"`
}
