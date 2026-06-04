package paymentgateway

import "fmt"

type upiGateway struct {
	id  string
	pin int
}

var _ PaymentGateway = (*upiGateway)(nil)

func (u *upiGateway) initiatePayment() (string, error) {
	// Logic to initiate payment using UPI
	fmt.Println("Initiating payment using UPI")
	return "UPI Transaction ID", nil
}
func (u *upiGateway) verifyPayment(transactionID string) (bool, error) {
	// Logic to verify payment using UPI
	fmt.Println("Verifying payment using UPI")
	return true, nil
}
func (u *upiGateway) cancelPayment(transactionID string) (bool, error) {
	// Logic to cancel payment using UPI
	fmt.Println("Canceling payment using UPI")
	return true, nil
}
func (u *upiGateway) getPaymentStatus(transactionID string) (string, error) {
	// Logic to get payment status using UPI
	fmt.Println("Getting payment status using UPI")
	return "Completed", nil
}
