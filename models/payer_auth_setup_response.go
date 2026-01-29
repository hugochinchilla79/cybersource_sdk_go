package models

import "encoding/json"

// PayerAuthSetupAPIResponse wraps the HTTP response and parsed data
type PayerAuthSetupAPIResponse struct {
	HTTPStatus int                    `json:"-"`
	Body       json.RawMessage        `json:"-"`
	Data       PayerAuthSetupResponse `json:"-"`
}

// PayerAuthSetupResponse contains the response from the payer auth setup service
type PayerAuthSetupResponse struct {
	ID            string `json:"id,omitempty"`
	SubmitTimeUTC string `json:"submitTimeUtc,omitempty"`
	Status        string `json:"status,omitempty"`

	ClientReferenceInformation *ClientReferenceInformation        `json:"clientReferenceInformation,omitempty"`
	ConsumerAuthenticationInfo *SetupAuthenticationInfoResponse   `json:"consumerAuthenticationInformation,omitempty"`
	ErrorInformation           *ErrorInformation                  `json:"errorInformation,omitempty"`
}

// SetupAuthenticationInfoResponse contains the authentication info from setup response
type SetupAuthenticationInfoResponse struct {
	// AccessToken is the JWT used to authenticate the payer
	AccessToken string `json:"accessToken,omitempty"`
	// ReferenceID to be used in the enrollment check request
	ReferenceID string `json:"referenceId,omitempty"`
	// DeviceDataCollectionURL is the URL used to collect device data
	DeviceDataCollectionURL string `json:"deviceDataCollectionUrl,omitempty"`
}

// ErrorInformation contains error details from the response
type ErrorInformation struct {
	Reason  string        `json:"reason,omitempty"`
	Message string        `json:"message,omitempty"`
	Details []FieldDetail `json:"details,omitempty"`
}
