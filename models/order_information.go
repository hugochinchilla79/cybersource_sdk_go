package models

type OrderInformation struct {
	AmountDetails *AmountDetails `json:"amountDetails,omitempty"`
	BillTo        *BillTo        `json:"billTo,omitempty"`
}

type AmountDetails struct {
	TotalAmount string `json:"totalAmount,omitempty"`
	Currency    string `json:"currency,omitempty"`
}

type BillTo struct {
	FirstName          string `json:"firstName,omitempty"`
	LastName           string `json:"lastName,omitempty"`
	Address1           string `json:"address1,omitempty"`
	Locality           string `json:"locality,omitempty"`
	AdministrativeArea string `json:"administrativeArea,omitempty"`
	PostalCode         string `json:"postalCode,omitempty"`
	Country            string `json:"country,omitempty"`
	Email              string `json:"email,omitempty"`
	PhoneNumber        string `json:"phoneNumber,omitempty"`
}
