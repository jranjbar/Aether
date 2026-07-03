# Feature 002 - Persistence Layer

## Goal

Implement the persistence layer for Smart Factory API.

## Requirements

Use PostgreSQL.

Use Docker Compose.

No ORM.

Use database/sql.

Use pgx/v5 driver.

Implement:

- Connection Pool
- Ping
- Graceful Close
- Health Check

Repository pattern.

Create:

internal/database

internal/repository

internal/migrations

Docker Compose

.env.example

Health endpoint must verify DB connectivity.

Use context.Context everywhere.

No global variables.

Write unit tests when possible.

Everything must compile.

Run:

go test ./...

successfully.

Do not modify unrelated files.

Create logical commits.