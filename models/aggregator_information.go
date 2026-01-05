package models

type AggregatorInformation struct {
	AggregatorID string       `json:"aggregatorId,omitempty"`
	SubMerchant  *SubMerchant `json:"subMerchant,omitempty"`
}

type SubMerchant struct {
	ID string `json:"id,omitempty"`
}
