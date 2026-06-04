package paymentgateway

import "fmt"

type creditCardGateway struct {
	cardNo int
	cvv    int
	expiry string
	pin    int
}

var _ PaymentGateway = (*creditCardGateway)(nil)

func (c *creditCardGateway) initiatePayment() (string, error) {
	// Logic to initiate payment using Credit Card
	fmt.Println("Initiating payment using Credit Card")
	return "Credit Card Transaction ID", nil
}
func (c *creditCardGateway) verifyPayment(transactionID string) (bool, error) {
	// Logic to verify payment using Credit Card
	fmt.Println("Verifying payment using Credit Card")
	return true, nil
}
func (c *creditCardGateway) cancelPayment(transactionID string) (bool, error) {
	// Logic to cancel payment using Credit Card
	fmt.Println("Canceling payment using Credit Card")
	return true, nil
}
func (c *creditCardGateway) getPaymentStatus(transactionID string) (string, error) {
	// Logic to get payment status using Credit Card
	fmt.Println("Getting payment status using Credit Card")
	return "Completed", nil
}
