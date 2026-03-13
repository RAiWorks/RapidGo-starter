package notifications

import "github.com/raiworks/rapidgo/v2/core/notification"

// WelcomeNotification is sent when a new user registers.
type WelcomeNotification struct {
	UserName string
}

// Channels returns which channels this notification uses.
func (n *WelcomeNotification) Channels() []string {
	return []string{"database"}
}

// ToDatabase returns the database payload for this notification.
func (n *WelcomeNotification) ToDatabase(notifiable notification.Notifiable) (notification.DatabaseMessage, error) {
	return notification.DatabaseMessage{
		Type: "welcome",
		Data: map[string]interface{}{
			"message": "Welcome to RapidGo, " + n.UserName + "!",
		},
	}, nil
}
