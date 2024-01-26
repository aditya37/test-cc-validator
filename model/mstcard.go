package model

// model data or field
type MSTCard struct {
	Number      string
	CardNetwork string
}

type ResponseValidateCard struct {
	Number      string `json:"number"`
	IsValid     bool   `json:"is_valid"`
	CardNetwork string `json:"card_network"`
}

type RequestValidateCard struct {
	CardNumber string `json:"card_number"`
}

type RequestAddCard struct {
	Number      string `json:"number"`
	CardNetwork string `json:"card_network"`
}
type ResponseAddCard struct {
	CardNetwork string `json:"card_network"`
}
