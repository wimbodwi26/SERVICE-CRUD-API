package structs

type SuccessResponse struct {
	Succes bool `json:"success"`
	Message string `json:"massage"`
	Data any `json:"data"`
}