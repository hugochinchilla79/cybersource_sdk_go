package models

type ProcessingInformation struct {
	ActionList        []string `json:"actionList,omitempty"`
	Capture           *bool    `json:"capture,omitempty"`
	CommerceIndicator string   `json:"commerceIndicator,omitempty"`
}
