package upi

// define the UPI IDs

const (
	upiInterface   = "@upi"
	hdfcInterface  = "@hdfcbank"
	iciciInterface = "@icici"
	axisInterface  = "@axisbank"
)

type PaymentProcessor interface {
	SendMoney(amount float64, to string) error
}

type UPIPaymentProcessor struct {
	upiID    string
	userBank string
}

func NewUPIPaymentProcessor(upiID string, userBank string) *UPIPaymentProcessor {
	return &UPIPaymentProcessor{
		upiID:    upiID,
		userBank: userBank,
	}
}

func (u *UPIPaymentProcessor) SendMoney(amount float64, to string) error {

	//  called bank interface 

	// 
	return nil
}

