package schedule

import (
	"context"
	"log/slog"

	"github.com/RAiWorks/RapidGo/v2/core/app"
	"github.com/RAiWorks/RapidGo/v2/core/scheduler"
)

// RegisterSchedule defines all scheduled tasks for the application.
// Add your tasks here using s.Add(cronExpr, name, taskFunc).
func RegisterSchedule(s *scheduler.Scheduler, application *app.App) {
	// Example: heartbeat task that logs every minute.
	s.Add("@every 1m", "heartbeat", func(ctx context.Context) error {
		slog.Info("scheduler heartbeat")
		return nil
	})
}