package dto

type Transaction struct {
	Id           string  `json:"id"`
	Status       string  `json:"status"`
	TransactedAt string  `json:"transacted_at"`
	Type         string  `json:"type"`
	Amount       float64 `json:"amount"`
	ReferenceId  string  `json:"reference_id"`
	DepositedBy  *string `json:"deposited_by,omitempty"`
	DepositedAt  *string `json:"deposited_at,omitempty"`
	WithdrawnBy  *string `json:"withdrawn_by,omitempty"`
	WithdrawnAt  *string `json:"withdrawn_at,omitempty"`
}

type TransactionResponseDto struct {
	Deposit    *Transaction `json:"deposit,omitempty"`
	Withdrawal *Transaction `json:"withdrawal,omitempty"`
}

type TransactionListResponseDto struct {
	Transactions []Transaction `json:"transactions"`
}
