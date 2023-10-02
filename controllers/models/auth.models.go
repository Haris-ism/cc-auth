package models

type Credentials struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type ReqHeader struct{
	Authorization	string	`json:"Authorization"`
}
