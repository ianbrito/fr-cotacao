package entity

type Volume struct {
	Category      string
	Amount        int
	UnitaryWeight float64
	Price         float64
	Sku           string
	Height        float64
	Width         float64
	Length        float64
}

func NewVolume(category string, amount int, unitaryWeight float64, price float64, sku string, height float64, width float64, length float64) *Volume {
	return &Volume{
		Category:      category,
		Amount:        amount,
		UnitaryWeight: unitaryWeight,
		Price:         price,
		Sku:           sku,
		Height:        height,
		Width:         width,
		Length:        length,
	}
}
