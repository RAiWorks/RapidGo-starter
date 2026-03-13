# Changelog

All notable changes to the **RapidGo Starter** project will be documented in this file.

The format is based on [Keep a Changelog](https://keepachangelog.com/en/1.1.0/),
and this project adheres to [Semantic Versioning](https://semver.org/spec/v2.0.0.html).

## [1.1.0] - 2026-03-13

### Changed

- **Repository Rename**: Renamed GitHub repository from `RAiWorks/RapidGo-starter` to `raiworks/rapidgo-starter` ([#org-rename])
- **Organization Rename**: GitHub organization changed from `RAiWorks` to `raiworks`
- **Go Module Path**: Updated module path from `github.com/RAiWorks/RapidGo-starter` to `github.com/raiworks/rapidgo-starter`
- **Framework Dependency**: Updated RapidGo import from `github.com/RAiWorks/RapidGo/v2` to `github.com/raiworks/rapidgo/v2`
- **Framework Version**: Bumped RapidGo framework dependency from v2.1.0 to v2.2.0
- **Git Remote URL**: Updated origin from `https://github.com/RAiWorks/RapidGo-starter.git` to `https://github.com/raiworks/rapidgo-starter.git`
- Updated all Go import paths across 34 source files to use new lowercase module paths
- Updated `go.sum` dependency checksums for renamed framework module
- Updated `README.md` and `resources/views/home.html` with new repository references

### Files Affected

#### Go Source (34 files)
- `app/helpers/pagination.go`
- `app/jobs/example_job.go`
- `app/notifications/welcome_notification.go`
- `app/plugins.go`
- `app/providers/config_provider.go`, `app/providers/database_provider.go`
- `app/providers/logger_provider.go`, `app/providers/middleware_provider.go`
- `app/providers/notification_provider.go`, `app/providers/providers_test.go`
- `app/providers/queue_provider.go`, `app/providers/redis_provider.go`
- `app/providers/router_provider.go`, `app/providers/session_provider.go`
- `app/schedule/schedule.go`
- `app/services/user_service.go`, `app/services/user_service_test.go`
- `cmd/main.go`
- `database/migrations/20260307000001_create_jobs_tables.go`
- `database/migrations/20260308000001_add_soft_deletes.go`
- `database/migrations/20260308000002_add_totp_fields.go`
- `database/migrations/20260308000003_create_audit_logs_table.go`
- `database/migrations/20260311000001_create_notifications_table.go`
- `database/models/audit_log.go`, `database/models/notification.go`
- `database/models/post.go`, `database/models/user.go`
- `database/seeders/user_seeder.go`
- `http/controllers/controllers_test.go`
- `plugins/example/example.go`
- `routes/api.go`, `routes/web.go`, `routes/ws.go`

#### Configuration (2 files)
- `go.mod`
- `go.sum`

#### Other (2 files)
- `README.md`
- `resources/views/home.html`
