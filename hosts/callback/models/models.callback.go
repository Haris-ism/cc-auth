package models

type InquiryItems struct{
	ID				int			`json:"id"`
	Name			string		`json:"name"`
	Type			string		`json:"type"`
	Price			int			`json:"price"`
	Quantity		int			`json:"quantity"`
}
type InquiryDiscounts struct{
	ID				int			`json:"id"`
	Name			string		`json:"name"`
	Type			string		`json:"type"`
	Percentage		int			`json:"percentage"`
}

type ResponseMerchantItems struct {
	Message string      `json:"message"`
	Code    int         `json:"code"`
	Data    []InquiryItems
}
type ResponseMerchantDiscounts struct {
	Message string      `json:"message"`
	Code    int         `json:"code"`
	Data    []InquiryDiscounts
}