package models

import "encoding/json"

type PaymentAPIResponse struct {
	HTTPStatus int             `json:"-"`
	Body       json.RawMessage `json:"-"`
	Data       PaymentResponse `json:"-"`
}

type PaymentResponse struct {
	Links *PaymentLinks `json:"_links,omitempty"`

	ClientReferenceInformation *ClientReferenceInformation `json:"clientReferenceInformation,omitempty"`

	ID            string `json:"id,omitempty"`
	Status        string `json:"status,omitempty"`
	Reason        string `json:"reason,omitempty"`
	Message       string `json:"message,omitempty"`
	SubmitTimeUTC string `json:"submitTimeUtc,omitempty"`

	ReconciliationID string `json:"reconciliationId,omitempty"`

	OrderInformation *PaymentOrderInformation `json:"orderInformation,omitempty"`

	PaymentAccountInformation *PaymentAccountInformation `json:"paymentAccountInformation,omitempty"`
	PaymentInformation        *PaymentInfoResponse       `json:"paymentInformation,omitempty"`

	PointOfSaleInformation *PointOfSaleInformation `json:"pointOfSaleInformation,omitempty"`
	ProcessorInformation   *ProcessorInformation   `json:"processorInformation,omitempty"`

	Details []FieldDetail `json:"details,omitempty"`
}

type PaymentLinks struct {
	Self         *Link `json:"self,omitempty"`
	Capture      *Link `json:"capture,omitempty"`
	AuthReversal *Link `json:"authReversal,omitempty"`
}

type Link struct {
	Method string `json:"method,omitempty"`
	Href   string `json:"href,omitempty"`
}

type PaymentOrderInformation struct {
	AmountDetails *PaymentAmountDetails `json:"amountDetails,omitempty"`
}

type PaymentAmountDetails struct {
	AuthorizedAmount string `json:"authorizedAmount,omitempty"`
	TotalAmount      string `json:"totalAmount,omitempty"` // por si aparece en otras respuestas
	Currency         string `json:"currency,omitempty"`
}

type PaymentAccountInformation struct {
	Card *CardTypeOnly `json:"card,omitempty"`
}

type PaymentInfoResponse struct {
	TokenizedCard *CardTypeOnly `json:"tokenizedCard,omitempty"`
	Card          *CardTypeOnly `json:"card,omitempty"`
}

type CardTypeOnly struct {
	Type string `json:"type,omitempty"`
}

type PointOfSaleInformation struct {
	TerminalID string `json:"terminalId,omitempty"`
}

type ProcessorInformation struct {
	ApprovalCode         string `json:"approvalCode,omitempty"`
	NetworkTransactionID string `json:"networkTransactionId,omitempty"`
	TransactionID        string `json:"transactionId,omitempty"`
	ResponseCode         string `json:"responseCode,omitempty"`
	Avs                  *AVS   `json:"avs,omitempty"`
}

type AVS struct {
	Code    string `json:"code,omitempty"`
	CodeRaw string `json:"codeRaw,omitempty"`
}

type FieldDetail struct {
	Field  string `json:"field,omitempty"`
	Reason string `json:"reason,omitempty"`
}
