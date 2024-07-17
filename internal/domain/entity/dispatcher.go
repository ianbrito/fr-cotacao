package entity

type Dispatcher struct {
	ID                         string
	RequestID                  string
	RegisteredNumberShipper    string
	RegisteredNumberDispatcher string
	ZipCode                    int
	Offers                     []*Offer
}
