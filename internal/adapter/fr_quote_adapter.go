package adapter

import (
	"github.com/ianbrito/fr-cotacao/internal/domain/entity"
	fr "github.com/ianbrito/fr-cotacao/pkg/frete-rapido"
)

type FreteRapidoQuoteAdapter struct {
	Quotes *fr.QuoteResponse
}

func NewFreteRapidoQuoteAdapter(quotes *fr.QuoteResponse) *FreteRapidoQuoteAdapter {
	return &FreteRapidoQuoteAdapter{
		Quotes: quotes,
	}
}

func (a *FreteRapidoQuoteAdapter) ToEntity() []*entity.Dispatcher {
	var dispatchers []*entity.Dispatcher
	for _, d := range a.Quotes.Dispatchers {
		var offers []*entity.Offer
		for _, o := range d.Offers {
			carrier := &entity.Carrier{
				Reference:        int64(o.Carrier.Reference),
				Name:             o.Carrier.Name,
				RegisteredNumber: o.Carrier.RegisteredNumber,
				StateInscription: o.Carrier.StateInscription,
				LogoUrl:          o.Carrier.Logo,
			}
			weights := &entity.Weights{
				Real:  o.Weights.Real,
				Cubed: &o.Weights.Cubed,
				Used:  o.Weights.Cubed,
			}

			deliveryTime := &entity.DeliveryTime{
				Days:          o.DeliveryTime.Days,
				Hours:         o.DeliveryTime.Hours,
				Minutes:       o.DeliveryTime.Days,
				EstimatedDate: o.DeliveryTime.EstimatedDate,
			}
			originalDeliveryTime := &entity.DeliveryTime{
				Days:          o.OriginalDeliveryTime.Days,
				Hours:         o.OriginalDeliveryTime.Hours,
				Minutes:       o.OriginalDeliveryTime.Days,
				EstimatedDate: o.OriginalDeliveryTime.EstimatedDate,
			}

			offers = append(offers, entity.NewOffer(
				int32(o.Offer),
				int32(o.SimulationType),
				carrier,
				o.Service,
				o.ServiceCode,
				o.ServiceDescription,
				o.Identifier,
				o.DeliveryNote,
				o.HomeDelivery,
				o.CarrierNeedsToReturnToSender,
				o.Expiration,
				o.CostPrice,
				o.FinalPrice,
				o.Modal,
				weights,
				deliveryTime,
				originalDeliveryTime,
			))
		}
		dispatchers = append(dispatchers, &entity.Dispatcher{
			ID:                         d.ID,
			RequestID:                  d.RequestID,
			RegisteredNumberShipper:    d.RegisteredNumberShipper,
			RegisteredNumberDispatcher: d.RegisteredNumberDispatcher,
			ZipCode:                    d.ZipcodeOrigin,
			Offers:                     offers,
		})
	}
	return dispatchers
}
