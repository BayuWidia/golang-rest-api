package models

type Response struct {
	ResponseCode int `json:"ResponseCode"`

	ResponseDesc string `json:"ResponseDesc"`

	ResponseTime string `json:"ResponseTime"`

	Result interface{} `json:"Result"`
}
