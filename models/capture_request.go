package models

type CaptureRequest struct {
	ClientReferenceInformation *ClientReferenceInformation `json:"clientReferenceInformation,omitempty"`
	OrderInformation           *OrderInformation           `json:"orderInformation,omitempty"`
}
