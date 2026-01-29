package models

// PayerAuthSetupRequest is used to set up payer authentication with Cardinal.
// This service must be called before the enrollment check.
// POST /risk/v1/authentication-setups
type PayerAuthSetupRequest struct {
	ClientReferenceInformation *ClientReferenceInformation `json:"clientReferenceInformation,omitempty"`
	PaymentInformation         *PaymentInformation         `json:"paymentInformation,omitempty"`
	TokenInformation           *TokenInformation           `json:"tokenInformation,omitempty"`
}

// TokenInformation contains transient token data for payer authentication
type TokenInformation struct {
	TransientToken string `json:"transientToken,omitempty"`
}
