package model

type Response struct {
	Rcode int `json:"code"`
	Message string `json:"message"`
	Data interface{} `json:"data"`
	Status bool `json:"status"`
}
