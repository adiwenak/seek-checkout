package entity

import "github.com/shopspring/decimal"

type ItemId int
type CustomerId string

type AdsItem struct {
	Id          ItemId
	Name        string
	Description string
	Price       decimal.Decimal
}

type PricingRule struct {
	Name            string
	AdsItemId       ItemId
	MinimumQuantity uint
	SpecialPrice    decimal.Decimal
}

type CustomerPricingRule struct {
	Id           CustomerId
	PricingRules map[ItemId]PricingRule
}
