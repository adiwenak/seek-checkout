package services

import (
	"github.com/adiwenak/seek-checkout-system/entity"
)

type CheckoutItem struct {
	Id       entity.ItemId `json:"id"`
	Quantity uint          `json:"quantity"`
}

type TotalPriceResponse struct {
	TotalPrice string `json:"totalPrice"`
}
