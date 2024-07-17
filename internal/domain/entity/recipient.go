package entity

import "os"

type Address struct {
	Country string
	ZipCode string
}

type Recipient struct {
	RegisteredNumber string
	Type             int
	Address          *Address
}

func NewRecipient(zipCode string) *Recipient {
	return &Recipient{
		RegisteredNumber: os.Getenv("TEST_CPNJ"),
		Type:             1,
		Address: &Address{
			Country: "BRA",
			ZipCode: zipCode,
		},
	}
}
