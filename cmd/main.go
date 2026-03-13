package main

import (
	"github.com/raiworks/rapidgo-starter/app/jobs"
	"github.com/raiworks/rapidgo-starter/app/providers"
	"github.com/raiworks/rapidgo-starter/app/schedule"
	"github.com/raiworks/rapidgo/v2/core/app"
	"github.com/raiworks/rapidgo/v2/core/cli"
	"github.com/raiworks/rapidgo/v2/core/container"
	"github.com/raiworks/rapidgo/v2/core/router"
	"github.com/raiworks/rapidgo/v2/core/service"
	"github.com/raiworks/rapidgo-starter/database/models"
	fwseeders "github.com/raiworks/rapidgo/v2/database/seeders"
	"github.com/raiworks/rapidgo-starter/routes"
	"gorm.io/gorm"
)

func main() {
	cli.SetBootstrap(func(a *app.App, mode service.Mode) {
		a.Register(&providers.ConfigProvider{})
		a.Register(&providers.LoggerProvider{})
		if mode.Has(service.ModeWeb) || mode.Has(service.ModeAPI) || mode.Has(service.ModeWS) {
			a.Register(&providers.DatabaseProvider{})
		}
		a.Register(&providers.RedisProvider{})
		a.Register(&providers.QueueProvider{})
		if mode.Has(service.ModeWeb) {
			a.Register(&providers.SessionProvider{})
		}
		a.Register(&providers.MiddlewareProvider{Mode: mode})
		a.Register(&providers.RouterProvider{Mode: mode})
		if mode.Has(service.ModeWeb) || mode.Has(service.ModeAPI) {
			a.Register(&providers.NotificationProvider{})
		}
	})

	cli.SetRoutes(func(r *router.Router, c *container.Container, mode service.Mode) {
		if mode.Has(service.ModeWeb) {
			routes.RegisterWeb(r)
		}
		if mode.Has(service.ModeAPI) {
			routes.RegisterAPI(r)
		}
		if mode.Has(service.ModeWS) {
			routes.RegisterWS(r)
		}
	})

	cli.SetJobRegistrar(jobs.RegisterJobs)
	cli.SetScheduleRegistrar(schedule.RegisterSchedule)

	cli.SetModelRegistry(models.All)

	cli.SetSeeder(func(db *gorm.DB, name string) error {
		if name != "" {
			return fwseeders.RunByName(db, name)
		}
		return fwseeders.RunAll(db)
	})

	cli.Execute()
}
