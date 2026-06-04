package main

import (
	"design-pattern/creational/facade/payment"
	"errors"
	"strings"
)

// define the UPI IDs
const (
	upiInterface   = "@upi"
	hdfcInterface  = "@hdfcbank"
	iciciInterface = "@icici"
	axisInterface  = "@axisbank"
)

// getAccountDetails retrieves account details based on UPI ID
func getAccountDetails(upiId string) (account payment.PaymentProcessor, err error) {
	if upiId == "" {
		return nil, errors.New("UPI ID cannot be empty")
	}

	parts := strings.Split(upiId, "@")
	if len(parts) != 2 {
		return nil, errors.New("invalid UPI ID format: expected username@bank")
	}

	username := parts[0]
	bankInterface := "@" + parts[1]

	switch bankInterface {
	case upiInterface:
		account, err = payment.GetHdfcAccountDetails(username)
		if err != nil {
			return nil, errors.Join(errors.New("error finding account"), err)
		}
		return account, nil
	case hdfcInterface:
		// TODO: Implement HDFC specific logic
		return nil, errors.New("HDFC interface not implemented yet")
	case iciciInterface:
		// TODO: Implement ICICI specific logic
		return nil, errors.New("ICICI interface not implemented yet")
	case axisInterface:
		// TODO: Implement Axis specific logic
		return nil, errors.New("Axis interface not implemented yet")
	default:
		return nil, errors.New("unsupported bank interface")
	}
}

// pay performs a UPI payment between sender and receiver
func pay(sender payment.PaymentProcessor, upiId string, amount int) (bool, error) {
	// Validate inputs
	if sender == nil {
		return false, errors.New("sender cannot be nil")
	}
	if amount <= 0 {
		return false, errors.New("amount must be positive")
	}

	// Check sender balance
	senderBalance := sender.CheckBalance()
	if senderBalance < amount {
		return false, errors.New("insufficient balance")
	}

	// Check if sender account is active
	if !sender.IsActive() {
		return false, errors.New("sender account is disabled")
	}

	// Get receiver account details
	receiverInterface, err := getAccountDetails(upiId)
	if err != nil {
		return false, err
	}

	// Check if receiver account is active
	if !receiverInterface.IsActive() {
		return false, errors.New("receiver account is disabled")
	}

	// Perform debit from sender
	_, err = sender.DebitMoney(amount)
	if err != nil {
		return false, errors.Join(errors.New("failed to debit money from sender account"), err)
	}

	// Perform credit to receiver
	_, err = receiverInterface.CreditMoney(amount)
	if err != nil {
		// Revert the sender debit transaction
		_, revertErr := sender.CreditMoney(amount)
		if revertErr != nil {
			// Log the revert error but return the original error
			// In production, you might want to implement retry logic here
			return false, errors.Join(
				errors.New("failed to send money to receiver and revert failed"),
				err,
				revertErr,
			)
		}
		return false, errors.Join(errors.New("failed to send money to receiver account"), err)
	}

	return true, nil
}
