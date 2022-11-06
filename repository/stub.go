package repository

import (
	"github.com/adiwenak/seek-checkout-system/entity"
	"github.com/shopspring/decimal"
)

var allAdsItem []entity.AdsItem = []entity.AdsItem{
	{
		Id:          1,
		Name:        "Classic Ad",
		Description: "Bla",
		Price:       decimal.NewFromFloat(269.99),
	},
	{
		Id:          2,
		Name:        "Stand out Ad",
		Description: "Bla",
		Price:       decimal.NewFromFloat(322.99),
	},
	{
		Id:          3,
		Name:        "Premium Ad",
		Description: "Bla",
		Price:       decimal.NewFromFloat(394.99),
	},
}

var prPackageClassicAds entity.PricingRule = entity.PricingRule{
	Name:            "3 for 2 classic ads",
	AdsItemId:       1,
	MinimumQuantity: 3,
	SpecialPrice:    decimal.NewFromFloat(539.98),
}

var prDiscountStandoutAds entity.PricingRule = entity.PricingRule{
	Name:            "discount standout ads",
	AdsItemId:       2,
	MinimumQuantity: 1,
	SpecialPrice:    decimal.NewFromFloat(299.99),
}

var prPackageStandoutAds entity.PricingRule = entity.PricingRule{
	Name:            "discount 5 for 4 standout ads",
	AdsItemId:       2,
	MinimumQuantity: 5,
	SpecialPrice:    decimal.NewFromFloat(1291.96),
}

var prDiscountPremiumAds entity.PricingRule = entity.PricingRule{
	Name:            "discount premium ads",
	AdsItemId:       3,
	MinimumQuantity: 1,
	SpecialPrice:    decimal.NewFromFloat(389.99),
}

var customers map[entity.CustomerId]entity.CustomerPricingRule = map[entity.CustomerId]entity.CustomerPricingRule{
	"2": {
		PricingRules: map[entity.ItemId]entity.PricingRule{
			1: prPackageClassicAds,
		},
	},
	"3": {
		PricingRules: map[entity.ItemId]entity.PricingRule{
			2: prDiscountStandoutAds,
		},
	},
	"4": {
		PricingRules: map[entity.ItemId]entity.PricingRule{
			2: prPackageStandoutAds,
			3: prDiscountPremiumAds,
		},
	},
}

type RepositoryInterface interface {
	GetAdsItem(id entity.ItemId) *entity.AdsItem
	GetCustomerPricingRules(id entity.CustomerId) *entity.CustomerPricingRule
}

type Repository struct{}

func NewRepository() RepositoryInterface {
	return &Repository{}
}

func (rep *Repository) GetAdsItem(id entity.ItemId) *entity.AdsItem {
	var item *entity.AdsItem

	for _, i := range allAdsItem {
		if i.Id == id {
			item = &i
			break
		}
	}

	return item
}

func (rep *Repository) GetCustomerPricingRules(id entity.CustomerId) *entity.CustomerPricingRule {
	if cust, ok := customers[id]; ok {
		return &cust
	} else {
		return nil
	}
}
