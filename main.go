package main

import (
	"github.com/adiwenak/seek-checkout-system/rest"
	"github.com/adiwenak/seek-checkout-system/rest/swagger"
	"github.com/gin-gonic/gin"
)

func NewRouter() *gin.Engine {
	router := gin.New()
	checkoutRest := rest.NewCheckoutRest()
	router.GET("/swaggerui/*any", swagger.SwaggerUI())
	router.POST("/checkout/totalprice", checkoutRest.TotalPrice)

	return router
}

func main() {
	router := NewRouter()
	router.Run(":8080")
}
