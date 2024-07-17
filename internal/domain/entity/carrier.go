package entity

type Carrier struct {
	Reference        int64
	Name             string
	RegisteredNumber string
	StateInscription string
	LogoUrl          string
}

func NewCarrier(reference int64, name string, registeredNumber string, stateInscription string, logoUrl string) *Carrier {
	return &Carrier{
		Reference:        reference,
		Name:             name,
		RegisteredNumber: registeredNumber,
		StateInscription: stateInscription,
		LogoUrl:          logoUrl,
	}
}
