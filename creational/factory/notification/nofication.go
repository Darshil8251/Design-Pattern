package notification

import "errors"

type notifier interface {
	Send(message string) error
	// Other methods can be added here
}

type NotifierType string

const (
	EmailNotifierType   NotifierType = "email"
	SlackNotifierType   NotifierType = "slack"
	MessageNotifierType NotifierType = "message"
)

func CreateNotifier(notifierType NotifierType, contactInfo string) (notifier, error) {
	switch notifierType {
	case EmailNotifierType:
		return newEmailNotifier(contactInfo), nil
	case SlackNotifierType:
		return newSlackNotifier(contactInfo), nil
	case MessageNotifierType:
		return newMessageNotifier(contactInfo), nil
	default:
		ErrUnknownNotifierType := errors.New("unknown notifier type")
		return nil, ErrUnknownNotifierType
	}
}
