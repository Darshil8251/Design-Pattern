package main

import (
	"design-pattern/structural/bridge/paymentgateway"
	"fmt"
)

func main() {

	fmt.Printf("payment service\n")

	// create the payment gateway
	gatewayType := paymentgateway.UPIGateway

	gateway := paymentgateway.CreatePaymentGateway(gatewayType)

	gateway.InitiatePayment()
	gateway.VerifyPayment("12345")
	gateway.CancelPayment("12345")
	gateway.GetPaymentStatus("12345")
}
