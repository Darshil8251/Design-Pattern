package payment

import (
	"errors"
	"time"
)

// hdfcBank represents HDFC bank account details
type hdfcBank struct {
	id           string
	name         string
	phoneNumber  string
	balance      int
	status       bool // status represents the status of account i.e block,lock,etc
	ifsc         string
	transactions map[string]transaction
}

var hdfcAccountList = make(map[string]*hdfcBank)

var _ PaymentProcessor = (*hdfcBank)(nil)

// NewHdfcAccount creates a new HDFC bank account
func NewHdfcAccount(id, name, phoneNumber, ifsc string) *hdfcBank {
	account := &hdfcBank{
		id:           id,
		name:         name,
		phoneNumber:  phoneNumber,
		balance:      0,
		status:       true,
		ifsc:         ifsc,
		transactions: make(map[string]transaction),
	}
	hdfcAccountList[id] = account
	return account
}

// CheckBalance returns the current balance
func (h *hdfcBank) CheckBalance() int {
	return h.balance
}

// GetHdfcAccountDetails retrieves account details by ID
func GetHdfcAccountDetails(accountID string) (PaymentProcessor, error) {
	accountDetails, found := hdfcAccountList[accountID]
	if !found {
		return nil, errors.New("account not found")
	}
	return accountDetails, nil
}

// CreditMoney adds money to the account
func (h *hdfcBank) CreditMoney(amount int) (bool, error) {
	if amount <= 0 {
		return false, errors.New("invalid amount: must be positive")
	}

	h.balance += amount

	// Record transaction
	transactionID := generateTransactionID()
	h.transactions[transactionID] = transaction{
		status:          true,
		amount:          amount,
		transactionType: "credit",
		date:            time.Now(),
	}

	return true, nil
}

// PaymentStatus checks the status of a transaction
func (h *hdfcBank) PaymentStatus(transactionID string) (bool, error) {
	transactionDetails, found := h.transactions[transactionID]
	if !found {
		return false, errors.New("transaction not found")
	}
	return transactionDetails.status, nil
}

// DebitMoney removes money from the account
func (h *hdfcBank) DebitMoney(amount int) (bool, error) {
	if amount <= 0 {
		return false, errors.New("invalid amount: must be positive")
	}

	if h.balance < amount {
		return false, errors.New("insufficient balance")
	}

	h.balance -= amount

	// Record transaction
	transactionID := generateTransactionID()
	h.transactions[transactionID] = transaction{
		status:          true,
		amount:          amount,
		transactionType: "debit",
		date:            time.Now(),
	}

	return true, nil
}

// IsActive checks if the account is active
func (h *hdfcBank) IsActive() bool {
	return h.status
}

// generateTransactionID creates a unique transaction ID
func generateTransactionID() string {
	return time.Now().Format("20060102150405")
}
