package model
type BaseResponse struct {
	Status  bool         `json:"status"`
	Message string      `json:"message"`
	Value   interface{} `json:"value"`
}