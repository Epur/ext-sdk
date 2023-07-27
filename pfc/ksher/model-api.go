package ksher

type GetMerchantResponse struct {
	Items []interface{} `json:"items"`
	Limit int           `json:"limit"`
	Page  int           `json:"page"`
	Total int           `json:"total"`
}
