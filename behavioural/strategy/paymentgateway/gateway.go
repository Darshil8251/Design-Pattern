package paymentgateway

import (
	"fmt"
)

// Define gateway type
type GatewayType string

const (
	PaypalGateway     GatewayType = "PayPal"
	StripeGateway     GatewayType = "Stripe"
	UPIGateway        GatewayType = "UPI"
	CreditCardGateway GatewayType = "CreditCard"
)

// Define the methods for payment gateway
type PaymentGateway interface {
	initiatePayment() (string, error)
	verifyPayment(transactionID string) (bool, error)
	cancelPayment(transactionID string) (bool, error)
	getPaymentStatus(transactionID string) (string, error)
}

// gateway is bridge between multiple payment gateways
type gateway struct {
	myGateway PaymentGateway
}

func CreatePaymentGateway(gatewayType GatewayType) (*gateway, error) {
	var pg PaymentGateway
	switch gatewayType {
	case PaypalGateway:
		pg = &paypalGateway{}
	case StripeGateway:
		pg = &stripeGateway{}
	case UPIGateway:
		pg = &upiGateway{}
	case CreditCardGateway:
		pg = &creditCardGateway{}
	default:
		return nil, fmt.Errorf("unknown gateway type: %s", gatewayType)
	}
	return &gateway{myGateway: pg}, nil
}

// implement the gateway

func (g *gateway) InitiatePayment() (string, error) {
	if g.myGateway == nil {
		return "", fmt.Errorf("no payment gateway configured")
	}
	return g.myGateway.initiatePayment()
}
func (g *gateway) VerifyPayment(transactionID string) (bool, error) {
	pg := g.myGateway
	if pg == nil {
		return false, fmt.Errorf("no payment gateway configured")
	}
	return pg.verifyPayment(transactionID)
}
func (g *gateway) CancelPayment(transactionID string) (bool, error) {
	if g.myGateway == nil {
		return false, fmt.Errorf("no payment gateway configured")
	}
	return g.myGateway.cancelPayment(transactionID)
}
func (g *gateway) GetPaymentStatus(transactionID string) (string, error) {
	if g.myGateway == nil {
		return "", fmt.Errorf("no payment gateway configured")
	}
	return g.myGateway.getPaymentStatus(transactionID)
}

func (g *gateway) UpdatePaymentProcessor(gatewayType GatewayType) error {
	newG, err := CreatePaymentGateway(gatewayType)
	if err != nil {
		return err
	}
	g.myGateway = newG.myGateway
	return nil
}
