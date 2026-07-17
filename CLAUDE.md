# CLAUDE.md

This file provides guidance to Claude Code (claude.ai/code) when working with code in this repository.

## What this is

Cassandra is a tarot-reading web app: a Go/Gin + PostgreSQL backend and a SvelteKit
(Svelte 5 + TypeScript) frontend. A user gives a name and story, picks a spread, draws
cards (each upright/reversed), and gets an AI reading from Claude in Cassandra's persona.
Guests can read without an account; optional email/password accounts save reading history.

## Repository layout

- `backend/` — Go + Gin API server (source of truth). Postgres via GORM.
- `frontend/` — SvelteKit app. Talks to the backend **only from server code** (BFF pattern).
- `docker-compose.yml` — local PostgreSQL 16.

## Commands

### Backend (`cd backend`)
- Run: `go run .` — serves `:8082` (needs Postgres up; env from `backend/.env`, optional).
- Seed the 78 cards (one-time, hits tarotapi.dev): `go run . seed`.
- Test: `go test ./...` — uses **in-memory SQLite** with a fixture deck; needs no Postgres, no network, no API key (interpretation is stubbed via `cfg.Testing`). Single test: `go test -run TestSuite/Test_AuthedReading_ShowsInHistory`.
- Build/vet: `go build ./...`, `go vet ./...`.

### Frontend (`cd frontend`)
- `npm install`, then `npm run dev` (port 5173), `npm run build` (adapter-node), `npm run check` (svelte-check/type-check).

### Local DB
- `docker compose up -d postgres` (db/user/password all `cassandra`, port 5432).

### CI
`.github/workflows/ci.yml` runs two jobs on push/PR to `main`: backend (build/vet/`go test`) and frontend (`npm ci` → `npm run check` → `npm run build`).

## Backend architecture

Go module is `github.com/pirateking92/teamtarot/backend`. `main.go` loads env, connects Postgres, `AutoMigrate`s, then either runs the `seed` subcommand or serves.

- **Config** (`env/`): `DATABASE_URL`, `ANTHROPIC_API_KEY`, `PORT`, `GIN_MODE`, `FRONTEND_ORIGIN`, `TESTING`. Missing `.env` is non-fatal.
- **DB** (`db/`): GORM + Postgres, `AutoMigrate` for schema. Models in `models/`: `Card`, `User`, `Session`, `Reading`, `ReadingCard`, `DailyCard` (UUIDs on User/Reading via `BeforeCreate`).
- **Controllers** hang off one `Handler{DB, Cfg}` struct (`controllers/handler.go`). Routes in `routes/routes.go`, all under `/api`.
- **Interpretation** (`services/interpret.go`): official Anthropic Go SDK, model `claude-opus-4-8`, Cassandra system prompt. Generated **asynchronously** in a goroutine that writes back into the `readings` row (so it survives restarts and powers history). Falls back to a stub when `TESTING` or no API key.
- **Spreads** (`services/spreads.go`) are static code config (single / past-present-future / five-card / celtic-cross), not DB rows.
- **Auth** (`services/auth.go`, `middleware/auth.go`): bcrypt passwords, opaque session tokens in a `sessions` table. `OptionalAuth` attaches the user if a session cookie/bearer is present (applied to all `/api`); `RequireAuth` gates history + `/auth/me`. Token read from the `session` cookie or `Authorization: Bearer`.

### Reading flow (two-phase, async)
1. `POST /api/readings` draws N unique cards (`ORDER BY RANDOM()`), persists the reading as `pending`, starts async interpretation, returns the reading + cards immediately.
2. Client polls `GET /api/readings/:id` until `status` is `complete`/`error`.

Cards are seeded into Postgres once; **tarotapi.dev is never called at reading time** (only by `go run . seed`).

## Frontend architecture (BFF)

The browser only ever talks to the SvelteKit origin. SvelteKit **server** code calls the Go API via `src/lib/server/api.ts` (`apiFetch`), reading `BACKEND_URL` (default `http://localhost:8082`) and forwarding the session token as a bearer.

- `hooks.server.ts` reads the `session` cookie → calls `/api/auth/me` → sets `event.locals.user`.
- Auth cookie is **first-party**, owned by SvelteKit: login/register form actions call the Go API, then `setSessionCookie` (in `lib/server/api.ts`); `logout` (`/logout/+server.ts`) revokes + clears it.
- `src/lib/types.ts` mirrors the Go JSON shapes (camelCase).
- Shared components in `src/lib/components/`: `OrnateButton`, `Typewriter` (completion-driven, skippable), `Flashcard` (single de-nested 3D flip), `CardThumb`.
- Client-side polling of a reading hits the same-origin proxy `src/routes/api/readings/[id]/+server.ts`.

### Routes
`/` (name) → `/reading/new` (spread + story; form action creates the reading, redirects) → `/reading/[id]` (draw/flip/poll). Plus `/cards` + `/cards/[nameShort]` (encyclopedia), `/daily`, `/history` (auth), `/login`, `/register`, `/account`, `/logout`.

## Styling

Bespoke "Cassandra" theme extracted into `src/app.css` (`.panel`, `.ornate-button`, `.input-field`, `.flip-card*`, staged `.scene-*` backgrounds) with design tokens in `tailwind.config.ts` (MedievalSharp font, maroon/amber palette). Card images are served from `static/cards/<nameShort>.jpg` and backgrounds from `static/bg/` — referenced by URL, not imported (this is deliberate; the old code's `src/lib/assets` runtime paths broke production builds).

## Conventions & gotchas

- The shell here is **zsh**: `status` is a read-only variable — don't use it as a loop variable in Bash tool commands.
- `redirect()` / `error()` from `@sveltejs/kit` throw internally (SvelteKit 2) — call them without `throw`.
- Scripted (fetch/enhance) POSTs to form actions get a JSON redirect envelope (`{type:"redirect",location:...}`), not a 303 header — relevant when testing actions with curl.
- Backend tests use SQLite (CGO); production uses Postgres. Both support `RANDOM()` and GORM `OnConflict`.
- Schema uses GORM `AutoMigrate` (no versioned migrations yet — a known future hardening step).
