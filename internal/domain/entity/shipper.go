package entity

import "os"

type Shipper struct {
	RegisteredNumber string
	Token            string
	PlatformCode     string
}

func NewShipper() *Shipper {
	return &Shipper{
		RegisteredNumber: os.Getenv("TEST_CPNJ"),
		Token:            os.Getenv("TEST_TOKEN"),
		PlatformCode:     os.Getenv("TEST_PLATFORM_CODE"),
	}
}
