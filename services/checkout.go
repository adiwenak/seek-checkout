package services

import (
	"github.com/adiwenak/seek-checkout-system/entity"
	"github.com/adiwenak/seek-checkout-system/repository"
	"github.com/shopspring/decimal"
)

type CheckoutServiceInterface interface {
	TotalPrice(customerid entity.CustomerId, items []CheckoutItem) TotalPriceResponse
	CalculateTotalPrice(customerid entity.CustomerId, items []CheckoutItem) string
}

// Ideally
type CheckoutService struct {
	repository repository.RepositoryInterface
}

func NewCheckoutService() CheckoutServiceInterface {
	return &CheckoutService{
		repository: repository.NewRepository(),
	}
}

func (serv *CheckoutService) TotalPrice(customerid entity.CustomerId, items []CheckoutItem) TotalPriceResponse {
	total := serv.CalculateTotalPrice(customerid, items)
	return TotalPriceResponse{
		TotalPrice: total,
	}
}

func (serv *CheckoutService) CalculateTotalPrice(customerid entity.CustomerId, items []CheckoutItem) string {
	var totalPrice decimal.Decimal = decimal.NewFromFloat(0.0)
	customerPricingRules := serv.repository.GetCustomerPricingRules(customerid)

	if len(items) == 0 {
		return totalPrice.String()
	}

	for _, item := range items {
		if customerPricingRules != nil {
			if specialRule, ok := customerPricingRules.PricingRules[item.Id]; ok {
				numberOfPackage := decimal.NewFromFloat(float64(item.Quantity / specialRule.MinimumQuantity))
				reminder := decimal.NewFromFloat(float64(item.Quantity % specialRule.MinimumQuantity))
				totalPackage := numberOfPackage.Mul(specialRule.SpecialPrice)

				if reminder != decimal.NewFromInt(0) {
					var itemPrice decimal.Decimal = decimal.NewFromFloat(0.0)
					if adsItem := serv.repository.GetAdsItem(item.Id); adsItem != nil {
						itemPrice = adsItem.Price
					}
					totalItems := reminder.Mul(itemPrice)
					totalPrice = totalPrice.Add(totalPackage.Add(totalItems))
				} else {
					totalPrice = totalPrice.Add(totalPackage)
				}

				continue
			}
		}

		if adsItem := serv.repository.GetAdsItem(item.Id); adsItem != nil {
			quantity := decimal.NewFromInt(int64(item.Quantity))
			totalItemPrice := quantity.Mul(adsItem.Price)
			totalPrice = totalPrice.Add(totalItemPrice)
		}
	}

	return totalPrice.String()
}
