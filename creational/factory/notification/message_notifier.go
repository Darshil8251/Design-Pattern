package notification

type messageNotifier struct {
	phoneNumber string
}

func newMessageNotifier(phone string) *messageNotifier {
	return &messageNotifier{
		phoneNumber: phone,
	}
}

func (m *messageNotifier) Send(message string) error {
	// Here you would implement the logic to send a message, e.g., using an SMS API.
	// For demonstration purposes, we'll just print the message.
	println("Sending message to " + m.phoneNumber + ": " + message)
	return nil
}
