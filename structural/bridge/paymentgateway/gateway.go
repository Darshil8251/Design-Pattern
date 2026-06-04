package paymentgateway

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

func CreatePaymentGateway(gatewayType GatewayType) *gateway {
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
	}
	return &gateway{myGateway: pg}
}

// implement the gateway

func (g *gateway) InitiatePayment() (string, error) {
	return g.myGateway.initiatePayment()
}
func (g *gateway) VerifyPayment(transactionID string) (bool, error) {
	return g.myGateway.verifyPayment(transactionID)
}
func (g *gateway) CancelPayment(transactionID string) (bool, error) {
	return g.myGateway.cancelPayment(transactionID)
}
func (g *gateway) GetPaymentStatus(transactionID string) (string, error) {
	return g.myGateway.getPaymentStatus(transactionID)
}
