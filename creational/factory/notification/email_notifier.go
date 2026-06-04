package notification

import (
	"fmt"
)

type emailNotifier struct {
	emailAddress string
}

func newEmailNotifier(email string) notifier {
	return &emailNotifier{
		emailAddress: email,
	}
}

func (e *emailNotifier) Send(message string) error {
	fmt.Printf("Sending email to %s: %s\n", e.emailAddress, message)
	return nil
}
