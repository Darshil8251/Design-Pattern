package notification

import "fmt"

type slackNotifier struct {
	WebhookURL string
}

func (s *slackNotifier) Send(message string) error {
	fmt.Printf("Sending Slack notification: %s\n", message)
	return nil
}

func newSlackNotifier(webhookURL string) *slackNotifier {
	return &slackNotifier{
		WebhookURL: webhookURL,
	}
}
