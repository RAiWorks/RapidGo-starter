# RapidGo Starter

A scaffold project for the [RapidGo](https://github.com/RAiWorks/RapidGo) framework.

## Getting Started

### Option 1: CLI (recommended)

```bash
go install github.com/RAiWorks/RapidGo/cmd/rapidgo@latest
rapidgo new myapp
```

### Option 2: Clone

```bash
git clone https://github.com/RAiWorks/RapidGo-starter myapp
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
cmd/main.go           ← Entry point with hook wiring
app/providers/        ← Service providers
app/helpers/          ← Utility functions
app/services/         ← Business logic
app/jobs/             ← Queue job handlers
app/schedule/         ← Scheduled tasks
routes/               ← Route definitions
http/controllers/     ← Request handlers
database/models/      ← GORM models
database/migrations/  ← Database migrations
database/seeders/     ← Seed data
resources/            ← Views, translations, static files
```

## Hook Wiring

See `cmd/main.go` for how hooks connect your app to the framework.

## License

MIT
