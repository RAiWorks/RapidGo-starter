package providers

import (
	"github.com/raiworks/rapidgo/v2/core/container"
	"github.com/raiworks/rapidgo/v2/core/notification"
	"gorm.io/gorm"
)

// NotificationProvider registers the notification system in the service container.
type NotificationProvider struct{}

// Register binds a *notification.Notifier singleton with a DatabaseChannel.
func (p *NotificationProvider) Register(c *container.Container) {
	c.Singleton("notifier", func(c *container.Container) interface{} {
		notifier := notification.NewNotifier()
		db := c.Make("db").(*gorm.DB)
		notifier.RegisterChannel("database", notification.NewDatabaseChannel(db))
		return notifier
	})
}

// Boot is a no-op.
func (p *NotificationProvider) Boot(c *container.Container) {}
