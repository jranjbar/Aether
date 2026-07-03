# Feature 001 - Foundation

## Goal

Build the production-grade foundation for the Smart Factory API.

This is infrastructure only.

Do NOT implement business logic.

---

## Stack

- Go 1.26
- Standard Library
- slog
- Docker
- Docker Compose
- PostgreSQL (future)
- Redis (future)

---

## Repository

Do not change existing repository layout.

Work inside:

apps/smart-factory-api

---

## Required Structure

apps/smart-factory-api/

cmd/server/

internal/

config/

logger/

server/

middleware/

health/

---

## Deliverables

Implement:

- HTTP Server
- Graceful Shutdown
- Config Loader
- Environment Variables
- Structured Logger
- Router
- Health Endpoint
- Ready Endpoint
- Live Endpoint

---

## API

GET /health

returns

200 OK

---

GET /ready

returns

200 OK

---

GET /live

returns

200 OK

---

## Requirements

Use:

context.Context

http.Server

signal.NotifyContext

log/slog

No global variables.

Dependency Injection.

---

## Code Quality

Small packages.

Idiomatic Go.

Document exported functions.

No TODO.

No panic.

No hardcoded configuration.

---

## Acceptance Criteria

Project builds successfully.

go test ./...

passes.

go run ./apps/smart-factory-api/cmd/server

starts server.

---

## Git

Create logical commits.

Example:

feat(config)

feat(logger)

feat(server)

feat(health)

Do not squash commits.