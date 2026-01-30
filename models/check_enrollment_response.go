package models

import "encoding/json"

// CheckEnrollmentAPIResponse wraps the HTTP response and parsed data
type CheckEnrollmentAPIResponse struct {
	HTTPStatus int                     `json:"-"`
	Body       json.RawMessage         `json:"-"`
	Data       CheckEnrollmentResponse `json:"-"`
}

// CheckEnrollmentResponse contains the response from the enrollment check service
type CheckEnrollmentResponse struct {
	ID            string `json:"id,omitempty"`
	SubmitTimeUTC string `json:"submitTimeUtc,omitempty"`
	// Status: PENDING_AUTHENTICATION, AUTHENTICATION_SUCCESSFUL, AUTHENTICATION_FAILED
	Status string `json:"status,omitempty"`

	ClientReferenceInformation *ClientReferenceInformation              `json:"clientReferenceInformation,omitempty"`
	ConsumerAuthenticationInfo *EnrollmentAuthenticationInfoResponse    `json:"consumerAuthenticationInformation,omitempty"`
	PaymentInformation         *PaymentInfoResponse                     `json:"paymentInformation,omitempty"`
	ErrorInformation           *ErrorInformation                        `json:"errorInformation,omitempty"`
}

// EnrollmentAuthenticationInfoResponse contains the authentication info from enrollment response
type EnrollmentAuthenticationInfoResponse struct {
	// VeResEnrolled: Y = Enrolled, N = Not Enrolled, U = Unable to verify
	VeResEnrolled string `json:"veresEnrolled,omitempty"`
	// PaResStatus: Y = Authenticated, N = Not Authenticated, U = Unknown
	PaResStatus string `json:"paresStatus,omitempty"`
	// ChallengeRequired: Y or N
	ChallengeRequired string `json:"challengeRequired,omitempty"`
	// SpecificationVersion: 2.1.0, 2.2.0, etc.
	SpecificationVersion string `json:"specificationVersion,omitempty"`

	// AcsURL is the URL of the Access Control Server
	AcsURL string `json:"acsUrl,omitempty"`
	// StepUpURL is the URL for step-up authentication
	StepUpURL string `json:"stepUpUrl,omitempty"`
	// AccessToken is the JWT for the step-up
	AccessToken string `json:"accessToken,omitempty"`
	// Token for Cardinal Commerce
	Token string `json:"token,omitempty"`
	// PAReq is the Payer Authentication Request (Base64 encoded)
	PAReq string `json:"pareq,omitempty"`

	// AuthenticationTransactionID for the authentication
	AuthenticationTransactionID string `json:"authenticationTransactionId,omitempty"`
	// DirectoryServerTransactionID from the directory server
	DirectoryServerTransactionID string `json:"directoryServerTransactionId,omitempty"`
	// ThreeDSServerTransactionID from the 3DS server
	ThreeDSServerTransactionID string `json:"threeDSServerTransactionId,omitempty"`
	// AcsTransactionID from the ACS
	AcsTransactionID string `json:"acsTransactionId,omitempty"`

	// CAVV - Cardholder Authentication Verification Value
	CAVV string `json:"cavv,omitempty"`
	// ECI - Electronic Commerce Indicator
	ECI string `json:"eci,omitempty"`
	// ECIRaw - Raw ECI value from processor
	ECIRaw string `json:"eciRaw,omitempty"`
	// EcommerceIndicator - Commerce indicator for authorization (e.g., "vbv", "spa")
	EcommerceIndicator string `json:"ecommerceIndicator,omitempty"`
	// Indicator - Commerce indicator returned specifically by ValidateAuthenticationResults (e.g., "vbv", "spa")
	Indicator string `json:"indicator,omitempty"`
	// XID - Transaction identifier for 3DS 1.0
	XID string `json:"xid,omitempty"`

	// UcafCollectionIndicator for Mastercard SecureCode
	UcafCollectionIndicator string `json:"ucafCollectionIndicator,omitempty"`
	// UcafAuthenticationData for Mastercard
	UcafAuthenticationData string `json:"ucafAuthenticationData,omitempty"`

	// AuthenticationResult for the final result
	AuthenticationResult string `json:"authenticationResult,omitempty"`
	// AuthenticationStatusMsg for result message
	AuthenticationStatusMsg string `json:"authenticationStatusMsg,omitempty"`
}
