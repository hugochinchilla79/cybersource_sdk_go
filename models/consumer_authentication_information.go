package models

type ConsumerAuthenticationInformation struct {
	Cavv                         string `json:"cavv,omitempty"`
	EciRaw                       string `json:"eciRaw,omitempty"`
	Xid                          string `json:"xid,omitempty"`
	UcafCollectionIndicator      string `json:"ucafCollectionIndicator,omitempty"`
	UcafAuthenticationData       string `json:"ucafAuthenticationData,omitempty"`
	DirectoryServerTransactionID string `json:"directoryServerTransactionId,omitempty"`
	PaSpecificationVersion       string `json:"paSpecificationVersion,omitempty"`
	AuthenticationTransactionID  string `json:"authenticationTransactionId,omitempty"`
}
