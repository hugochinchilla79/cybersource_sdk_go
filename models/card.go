package models

type PaymentInformation struct {
	Card *Card `json:"card,omitempty"`
}

type Card struct {
	Number          string `json:"number,omitempty"`
	ExpirationMonth string `json:"expirationMonth,omitempty"`
	ExpirationYear  string `json:"expirationYear,omitempty"`
	SecurityCode    string `json:"securityCode,omitempty"`
	// Type is the card type code. Required for Cartes Bancaires and UPI.
	// Common values: 001=Visa, 002=Mastercard, 003=Amex, 004=Discover
	Type                  string `json:"type,omitempty"`
	SecurityCodeIndicator string `json:"securityCodeIndicator,omitempty"`
}
