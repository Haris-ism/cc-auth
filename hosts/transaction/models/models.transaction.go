package models

type TransactionItems struct{
	ItemID			int			`json:"item_id"`
	Discount		string		`json:"discount"`
	Quantity		int			`json:"quantity"`
	CCNumber		string		`json:"cc_number"`
	CVV				string		`json:"cvv"`
}

type ResponseItems struct{
	Voucher			string	`json:"voucher"`
}

type ResponseTransactionItems struct {
	Message string      `json:"message"`
	Code    int         `json:"code"`
	Data    []ResponseItems
}