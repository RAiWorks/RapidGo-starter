# RapidGo Starter

A scaffold project for the [RapidGo](https://github.com/raiworks/rapidgo) framework (v2.1.0).

## Getting Started

### Option 1: CLI (recommended)

```bash
go install github.com/raiworks/rapidgo/cmd/rapidgo@latest
rapidgo new myapp
```

### Option 2: Clone

```bash
git clone https://github.com/raiworks/rapidgo-starter myapp
cd myapp
# Update module name in go.mod and all .go files
go mod tidy
```

### Configure

```bash
cp .env.example .env
# Edit .env with your database credentials
```

### Run

```bash
go run cmd/main.go serve
go run cmd/main.go migrate
go run cmd/main.go db:seed
```

## Project Structure

```
cmd/main.go             ← Entry point with hook wiring
app/providers/          ← Service providers
app/helpers/            ← Utility functions (pagination, crypto, etc.)
app/services/           ← Business logic
app/jobs/               ← Queue job handlers
app/schedule/           ← Scheduled tasks
app/notifications/      ← Notification definitions
plugins/                ← Plugin implementations
routes/                 ← Route definitions (web, api, ws)
http/controllers/       ← Request handlers
http/responses/         ← Standard JSON response helpers
database/models/        ← GORM models
database/migrations/    ← Database migrations
database/seeders/       ← Seed data
resources/              ← Views, translations, static files
storage/                ← Logs, cache, uploads, sessions
```

## Included Features (v2.1.0)

- **Service Providers** — Config, Logger, Database, Redis, Queue, Session, Middleware, Router, Notification
- **Notifications** — Multi-channel (database, mail) with `Notifiable` interface on User model
- **Cursor Pagination** — Offset and cursor-based pagination helpers + response wrappers
- **Queue Jobs** — Database/Redis/Memory/Sync drivers with backoff retry support
- **Scheduled Tasks** — Cron-based task scheduling
- **Health Checks** — `/health` and `/health/ready` endpoints
- **Prometheus Metrics** — Opt-in via `METRICS_ENABLED=true`
- **GORM Query Logging** — Auto-enabled in development, configurable via `DB_LOG`
- **Plugin System** — Extensible plugin architecture (see `plugins/example/`)
- **TOTP / 2FA** — Two-factor auth fields on User model
- **Audit Logging** — Change tracking via `audit_logs` table
- **WebSocket** — Route registration placeholder with heartbeat support

## Hook Wiring

See `cmd/main.go` for how hooks connect your app to the framework.

## Docker

```bash
docker compose up -d
```

## License

MIT
