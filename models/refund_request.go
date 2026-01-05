package models

type RefundRequest struct {
	ClientReferenceInformation *ClientReferenceInformation `json:"clientReferenceInformation,omitempty"`
	OrderInformation           *OrderInformation           `json:"orderInformation,omitempty"`
}
