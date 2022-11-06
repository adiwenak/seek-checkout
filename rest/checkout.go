package rest

import (
	"net/http"

	"github.com/adiwenak/seek-checkout-system/entity"
	"github.com/adiwenak/seek-checkout-system/services"
	"github.com/gin-gonic/gin"
)

type CheckoutRest struct {
	CheckoutService services.CheckoutServiceInterface
}

func NewCheckoutRest() *CheckoutRest {
	return &CheckoutRest{
		CheckoutService: services.NewCheckoutService(),
	}
}

func (controller *CheckoutRest) TotalPrice(c *gin.Context) {
	customerId := c.Request.Header.Get("customerid")

	if len(customerId) == 0 {
		c.Status(http.StatusBadRequest)
		return
	}

	var reqBody []services.CheckoutItem
	err := c.Bind(&reqBody)

	if err != nil {
		c.Status(http.StatusInternalServerError)
		return
	}

	resp := controller.CheckoutService.TotalPrice(entity.CustomerId(customerId), reqBody)

	c.JSON(http.StatusOK, resp)
}
