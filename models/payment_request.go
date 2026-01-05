package models

type PaymentRequest struct {
	ClientReferenceInformation *ClientReferenceInformation `json:"clientReferenceInformation,omitempty"`
	ProcessingInformation      *ProcessingInformation      `json:"processingInformation,omitempty"`
	PaymentInformation         *PaymentInformation         `json:"paymentInformation,omitempty"`
	OrderInformation           *OrderInformation           `json:"orderInformation,omitempty"`

	MerchantInformation               *MerchantInformation               `json:"merchantInformation,omitempty"`
	AggregatorInformation             *AggregatorInformation             `json:"aggregatorInformation,omitempty"`
	ConsumerAuthenticationInformation *ConsumerAuthenticationInformation `json:"consumerAuthenticationInformation,omitempty"`
}

type ClientReferenceInformation struct {
	Code string `json:"code,omitempty"`
}
