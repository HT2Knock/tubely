# Tubely — AGENTS.md

## Quick start

```bash
go mod download
cp .env.example .env       # edit as needed
./samplesdownload.sh        # download sample media to samples/
go run .                    # runs cmd/api-server, serves on :8091
```

Open http://localhost:8091/app/ — frontend at `/app/`, API at `/api/`.

## Key commands

| Action | Command |
|---|---|
| Run dev server | `go run .` (or `go run ./cmd/api-server/`) |
| Build | `go build ./...` |
| Regenerate sqlc | `sqlc generate` (config: `sqlc.yaml`, input: `sql/queries/`, output: `internal/database/`) |
| Reset DB | `POST /admin/reset` when `PLATFORM=dev` |

No tests, linters, CI, or Makefile exist.

## Architecture

- **Entrypoint:** `cmd/api-server/main.go` — reads `.env`, parses config, creates DB client, starts `net/http` server.
- **Database:** SQLite3 via `mattn/go-sqlite3`. Schema in `sql/schema/`. sqlc generates type-safe Go code (`internal/database/`). DB file is `tubely.db` (gitignored).
- **Auth:** JWT (HS256) + Argon2id hashing + refresh tokens. Token in `Authorization: Bearer <token>` header.
- **Static frontend:** `web/` dir served at `/app/`. Pure HTML/JS/CSS.
- **Uploads:** `assets/` dir (gitignored). Upload handlers are stubs — course exercise to implement S3/CloudFront uploads.
- **Thumbnail cache:** In-memory `map[uuid.UUID]thumbnail` in `cmd/api-server/server.go` — lost on restart.
- **Dep:** FFMPEG (`ffmpeg` + `ffprobe` must be in `PATH`).

## Conventions

- Go 1.25, standard `net/http` (no framework). Uses Go 1.22+ route patterns (`"POST /api/..."`, `{videoID}`).
- Logging via `rs/zerolog`.
- Config loaded from `.env` via `godotenv` + `caarlos0/env/v11`.
- UUIDs from `google/uuid`, passwords via `alexedwards/argon2id`.
- Edit `sql/queries/*.sql` then run `sqlc generate` to regenerate DB layer. Do not hand-edit `internal/database/*.sql.go` — they are auto-generated.
- Bruno collection at `bruno/` for API testing against `localhost:8091`.
