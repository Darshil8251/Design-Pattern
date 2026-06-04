package main

import (
	"design-pattern/creational/factory/notification"
	"fmt"
)

func main() {
	println("Factory Pattern")

	// Create an email notifier
	emailNotifier, err := notification.CreateNotifier(notification.EmailNotifierType, "email@example.com")
	if err != nil {
		println("Error creating email notifier:", err)
		return
	}
	emailNotifier.Send("Hello via Email!")

	// Create a Slack notifier
	slackNotifier, err := notification.CreateNotifier(notification.SlackNotifierType, "#general")
	if err != nil {
		println("Error creating Slack notifier:", err)
		return
	}
	slackNotifier.Send("Hello via Slack!")

	// Create a message notifier
	messageNotifier, err := notification.CreateNotifier(notification.MessageNotifierType, "123-456-7890")
	if err != nil {
		println("Error creating message notifier:", err)
		return
	}
	messageNotifier.Send("Hello via SMS!")

	fmt.Println("Notifiers created and messages sent successfully!")
}
