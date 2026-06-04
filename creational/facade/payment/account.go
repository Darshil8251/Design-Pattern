package payment

type PaymentProcessor interface {
	CreditMoney(amount int) (bool, error)
	PaymentStatus(transactionID string) (bool, error)
	DebitMoney(amount int) (bool, error)
	CheckBalance() int
	IsActive() bool
}
