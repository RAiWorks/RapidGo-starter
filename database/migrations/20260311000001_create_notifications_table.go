package migrations

import (
	"time"

	fwmigrations "github.com/raiworks/rapidgo/v2/database/migrations"
	"gorm.io/gorm"
)

func init() {
	fwmigrations.Register(fwmigrations.Migration{
		Version: "20260311000001_create_notifications_table",
		Up: func(db *gorm.DB) error {
			type NotificationRecord struct {
				ID           uint      `gorm:"primaryKey"`
				NotifiableID uint      `gorm:"index"`
				Type         string    `json:"type"`
				Data         string    `gorm:"type:text"`
				ReadAt       *time.Time
				CreatedAt    time.Time
			}
			return db.Table("notifications").AutoMigrate(&NotificationRecord{})
		},
		Down: func(db *gorm.DB) error {
			return db.Migrator().DropTable("notifications")
		},
	})
}
