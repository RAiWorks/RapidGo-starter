package migrations

import (
	"time"

	fwmigrations "github.com/raiworks/rapidgo/v2/database/migrations"
	"gorm.io/gorm"
)

func init() {
	fwmigrations.Register(fwmigrations.Migration{
		Version: "20260307000001_create_jobs_tables",
		Up: func(db *gorm.DB) error {
			type Job struct {
				ID          uint64     `gorm:"primaryKey;autoIncrement"`
				Queue       string     `gorm:"size:255;not null;index:idx_jobs_queue"`
				Type        string     `gorm:"size:255;not null"`
				Payload     string     `gorm:"type:text;not null"`
				Attempts    uint       `gorm:"not null;default:0"`
				MaxAttempts uint       `gorm:"not null;default:3"`
				AvailableAt time.Time  `gorm:"not null;index:idx_jobs_available_at"`
				ReservedAt  *time.Time `gorm:"index:idx_jobs_reserved_at"`
				CreatedAt   time.Time  `gorm:"not null"`
			}

			type FailedJob struct {
				ID       uint64    `gorm:"primaryKey;autoIncrement"`
				Queue    string    `gorm:"size:255;not null"`
				Type     string    `gorm:"size:255;not null"`
				Payload  string    `gorm:"type:text;not null"`
				Error    string    `gorm:"type:text;not null"`
				FailedAt time.Time `gorm:"not null"`
			}

			if err := db.AutoMigrate(&Job{}); err != nil {
				return err
			}
			return db.Table("failed_jobs").AutoMigrate(&FailedJob{})
		},
		Down: func(db *gorm.DB) error {
			if err := db.Migrator().DropTable("failed_jobs"); err != nil {
				return err
			}
			return db.Migrator().DropTable("jobs")
		},
	})
}