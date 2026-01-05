package models

type PaymentInformation struct {
	Card *Card `json:"card,omitempty"`
}

type Card struct {
	Number          string `json:"number,omitempty"`
	ExpirationMonth string `json:"expirationMonth,omitempty"`
	ExpirationYear  string `json:"expirationYear,omitempty"`
	SecurityCode    string `json:"securityCode,omitempty"`
}
