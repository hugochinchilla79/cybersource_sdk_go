package models

// CheckEnrollmentRequest is used to verify that a card is enrolled in a payer authentication program.
// POST /risk/v1/authentications
type CheckEnrollmentRequest struct {
	ClientReferenceInformation        *ClientReferenceInformation        `json:"clientReferenceInformation,omitempty"`
	OrderInformation                  *OrderInformation                  `json:"orderInformation,omitempty"`
	PaymentInformation                *PaymentInformation                `json:"paymentInformation,omitempty"`
	BuyerInformation                  *BuyerInformation                  `json:"buyerInformation,omitempty"`
	ConsumerAuthenticationInformation *EnrollmentAuthenticationInfo      `json:"consumerAuthenticationInformation,omitempty"`
	DeviceInformation                 *DeviceInformation                 `json:"deviceInformation,omitempty"`
	RiskInformation                   *RiskInformation                   `json:"riskInformation,omitempty"`
}

// BuyerInformation contains buyer contact information
type BuyerInformation struct {
	MobilePhone string `json:"mobilePhone,omitempty"`
	WorkPhone   string `json:"workPhone,omitempty"`
}

// EnrollmentAuthenticationInfo contains authentication info for enrollment check
type EnrollmentAuthenticationInfo struct {
	// ReferenceID from the setup payer auth response
	ReferenceID string `json:"referenceId,omitempty"`
	// ReturnURL where the issuer will redirect after authentication
	ReturnURL string `json:"returnUrl,omitempty"`
	// TransactionMode: S = eCommerce, M = MOTO, R = Recurring/Installment
	TransactionMode string `json:"transactionMode,omitempty"`
	// DeviceChannel: Browser, SDK, or 3RI
	DeviceChannel string `json:"deviceChannel,omitempty"`
	// MessageCategory: 01 = Payment Authentication, 02 = Non-Payment Authentication
	MessageCategory string `json:"messageCategory,omitempty"`
	// AuthenticationType: 01=Static, 02=Dynamic, 03=OOB, 04=Decoupled
	AuthenticationType string `json:"authenticationType,omitempty"`
	// ChallengeCode for Visa Secure and Mastercard Identity Check
	ChallengeCode string `json:"challengeCode,omitempty"`
	// AcsWindowSize: 01=250x400, 02=390x400, 03=500x600, 04=600x400, 05=Full page
	AcsWindowSize string `json:"acsWindowSize,omitempty"`
}

// DeviceInformation contains browser and device data
type DeviceInformation struct {
	HTTPBrowserColorDepth        string `json:"httpBrowserColorDepth,omitempty"`
	HTTPBrowserJavaEnabled       bool   `json:"httpBrowserJavaEnabled,omitempty"`
	HTTPBrowserJavaScriptEnabled bool   `json:"httpBrowserJavaScriptEnabled,omitempty"`
	HTTPBrowserLanguage          string `json:"httpBrowserLanguage,omitempty"`
	HTTPBrowserScreenHeight      string `json:"httpBrowserScreenHeight,omitempty"`
	HTTPBrowserScreenWidth       string `json:"httpBrowserScreenWidth,omitempty"`
	HTTPBrowserTimeDifference    string `json:"httpBrowserTimeDifference,omitempty"`
	UserAgentBrowserValue        string `json:"userAgentBrowserValue,omitempty"`
	HTTPAcceptContent            string `json:"httpAcceptContent,omitempty"`
	IPAddress                    string `json:"ipAddress,omitempty"`
}

// RiskInformation contains risk assessment data
type RiskInformation struct {
	BuyerHistory *BuyerHistory `json:"buyerHistory,omitempty"`
}

// BuyerHistory contains buyer transaction history
type BuyerHistory struct {
	CustomerAccount       *CustomerAccount `json:"customerAccount,omitempty"`
	TransactionCountYear  int              `json:"transactionCountYear,omitempty"`
	TransactionCountDay   int              `json:"transactionCountDay,omitempty"`
	AccountPurchases      int              `json:"accountPurchases,omitempty"`
	AddCardAttempts       int              `json:"addCardAttempts,omitempty"`
}

// CustomerAccount contains customer account information
type CustomerAccount struct {
	CreateDate     string `json:"createDate,omitempty"`
	ModifyDate     string `json:"modifyDate,omitempty"`
	PasswordChange string `json:"passwordChange,omitempty"`
}

// ValidateAuthResultsRequest is used to validate authentication results after challenge.
// POST /risk/v1/authentication-results
type ValidateAuthResultsRequest struct {
	ClientReferenceInformation        *ClientReferenceInformation       `json:"clientReferenceInformation,omitempty"`
	OrderInformation                  *OrderInformation                 `json:"orderInformation,omitempty"`
	PaymentInformation                *PaymentInformation               `json:"paymentInformation,omitempty"`
	ConsumerAuthenticationInformation *ValidationAuthenticationInfo     `json:"consumerAuthenticationInformation,omitempty"`
}

// ValidationAuthenticationInfo contains the transaction ID for validation
type ValidationAuthenticationInfo struct {
	// AuthenticationTransactionID from the enrollment check response
	AuthenticationTransactionID string `json:"authenticationTransactionId,omitempty"`
}
