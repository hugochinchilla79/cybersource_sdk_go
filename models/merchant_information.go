package models

type MerchantInformation struct {
	MerchantDescriptor  *MerchantDescriptor `json:"merchantDescriptor,omitempty"`
	SalesOrganizationID string              `json:"salesOrganizationId,omitempty"`
	CategoryCode        *int                `json:"categoryCode,omitempty"`
}

type MerchantDescriptor struct {
	Name            string `json:"name,omitempty"`
	Locality        string `json:"locality,omitempty"`
	CountryOfOrigin string `json:"countryOfOrigin,omitempty"`
}
