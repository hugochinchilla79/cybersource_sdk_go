package models

type ReversalRequest struct {
	ClientReferenceInformation *ClientReferenceInformation `json:"clientReferenceInformation,omitempty"`
	ReversalInformation        *ReversalInformation        `json:"reversalInformation,omitempty"`
}

type ReversalInformation struct {
	AmountDetails *ReversalAmountDetails `json:"amountDetails,omitempty"`
	Reason        string                 `json:"reason,omitempty"`
}

type ReversalAmountDetails struct {
	TotalAmount string `json:"totalAmount,omitempty"`
}
