package main

import (
	"design-pattern/behavioural/strategy/paymentgateway"
	"fmt"
)

func main() {
	fmt.Println("Welcome to strategy pattern")

	// Create a payment gateway
	gateway, err := paymentgateway.CreatePaymentGateway(paymentgateway.PaypalGateway)
	if err != nil {
		fmt.Println("Error creating payment gateway:", err)
		return
	}

	// Use the payment gateway
	transactionID, err := gateway.InitiatePayment()
	if err != nil {
		fmt.Println("Error initiating payment:", err)
		return
	}
	fmt.Println("Payment initiated, transaction ID:", transactionID)

	status, err := gateway.GetPaymentStatus(transactionID)
	if err != nil {
		fmt.Println("Error getting payment status:", err)
		return
	}
	fmt.Println("Payment status:", status)

	// Update the payment gateway to Stripe
	err = gateway.UpdatePaymentProcessor(paymentgateway.StripeGateway)
	if err != nil {
		fmt.Println("Error updating payment gateway:", err)
		return
	}
	fmt.Println("Payment gateway updated to Stripe")

	// Use the new payment gateway
	transactionID, err = gateway.InitiatePayment()
	if err != nil {
		fmt.Println("Error initiating payment:", err)
		return
	}
	fmt.Println("Payment initiated with Stripe, transaction ID:", transactionID)

	status, err = gateway.GetPaymentStatus(transactionID)
	if err != nil {
		fmt.Println("Error getting payment status:", err)
		return
	}
	fmt.Println("Payment status with Stripe:", status)
}
