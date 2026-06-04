package payment

import "time"

type transaction struct {
	status          bool
	amount          int
	transactionType string
	date            time.Time
}
