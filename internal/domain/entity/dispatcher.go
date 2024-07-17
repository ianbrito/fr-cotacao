package entity

import "os"

type Dispatcher struct {
	RegisteredNumber string
	ZipCode          string
	Volumes          []*Volume
}

func NewDispatcher() *Dispatcher {
	return &Dispatcher{
		RegisteredNumber: os.Getenv("TEST_CPNJ"),
		ZipCode:          os.Getenv("TEST_ZIP_CODE"),
	}
}

func (d *Dispatcher) AddVolume(volume *Volume) {
	d.Volumes = append(d.Volumes, volume)
}
