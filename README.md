# Cassandra

Welcome to the README for Cassandra, your new favourite fortune teller from Team Tarot.

Cassandra is a tarot card reading application. Tell her your name and your story, choose
a spread, and she deals your cards — each upright or reversed — then gives you a
personalised AI reading in the voice of a warm, old-world seer. Registered seekers can
revisit their past readings; anyone can read as a guest.

Originally built by Team Tarot (Alannah Lawlor, Josh Day, Fara Ifaturoti, Oli Kelly,
Giandomenico Loparco and Matt Doyle), now modernised with a Go + PostgreSQL backend and a
Svelte 5 + TypeScript frontend.

## Features

- **Guest or account** — read instantly as a guest, or sign up (email + password) to save your history.
- **Multiple spreads** — single card, Past · Present · Future, a five-card cross, and the full Celtic Cross.
- **AI readings** — each reading is interpreted by Claude (`claude-opus-4-8`) in Cassandra's persona, woven from your cards and your story.
- **Reading history** — signed-in seekers can revisit every reading they've drawn.
- **Card encyclopedia** — browse all 78 cards with upright/reversed meanings and descriptions.
- **Card of the day** — one shared card drawn for everyone each day.

## Tech stack

- **Backend:** Go + [Gin](https://gin-gonic.com/), [GORM](https://gorm.io/) over PostgreSQL, and the official [Anthropic Go SDK](https://github.com/anthropics/anthropic-sdk-go).
- **Frontend:** [SvelteKit](https://svelte.dev/docs/kit) (Svelte 5 runes) + TypeScript + Tailwind CSS. The SvelteKit server acts as a BFF, proxying the Go API so the browser never calls it directly.
- **Data:** the 78 cards are seeded once from [tarotapi.dev](https://tarotapi.dev/) into Postgres; nothing external is called at reading time.

## Prerequisites

- [Go](https://go.dev/) 1.24+
- [Node.js](https://nodejs.org/) 22+
- [Docker](https://www.docker.com/) (for local PostgreSQL) — or your own Postgres 16 instance
- An [Anthropic API key](https://console.anthropic.com/) for real interpretations

## Getting started

Clone the repo:

```sh
git clone https://github.com/pirateking92/teamtarot
cd teamtarot
```

### 1. Start PostgreSQL

```sh
docker compose up -d postgres
```

This runs Postgres 16 on `localhost:5432` with database/user/password all `cassandra`
(matches the backend defaults).

### 2. Configure & run the backend

```sh
cd backend
cp .env.example .env          # then set ANTHROPIC_API_KEY (or leave TESTING=True to stub readings)
go run . seed                 # one-time: import the 78 cards from tarotapi.dev
go run .                      # starts the API on http://localhost:8082
```

> Without an `ANTHROPIC_API_KEY` (or with `TESTING=True`), interpretations fall back to a
> stub string so you can develop offline.

### 3. Run the frontend

```sh
cd frontend
npm install
npm run dev                   # http://localhost:5173
```

Open <http://localhost:5173> and follow Cassandra.

## Production notes

- The Go API is an **internal service** behind the SvelteKit BFF — do not expose it publicly. The browser only ever talks to SvelteKit, which calls the API server-side (set `BACKEND_URL` for the SvelteKit server).
- The frontend builds with `@sveltejs/adapter-node`: `npm run build` then `node build`. In production, set `ORIGIN` to the site's public URL (required for SvelteKit's form-action CSRF protection) and `BACKEND_URL` to the API's address, e.g. `ORIGIN=https://cassandra.example BACKEND_URL=http://api:8082 node build`.

## Project structure

```
backend/            Go + Gin API
  controllers/      HTTP handlers (readings, cards, spreads, auth)
  services/         card seeding, spreads, draw, interpretation (Anthropic), auth
  models/           GORM models (Card, User, Session, Reading, ReadingCard, DailyCard)
  middleware/       optional/required auth
  db/               Postgres connection + AutoMigrate
  env/              env-based config
frontend/           SvelteKit (Svelte 5 + TS)
  src/lib/          types, server API helper (BFF), shared components
  src/routes/       landing, reading flow, encyclopedia, daily, history, auth
  static/cards/     78 card faces + card back + deck
  static/bg/        staged background scenes
docker-compose.yml  local PostgreSQL
```

## API overview (`/api`)

| Method | Path | Purpose |
|---|---|---|
| `POST` | `/readings` | Draw a spread and start a reading (guest or authed) |
| `GET`  | `/readings/:id` | Fetch a reading + its cards + interpretation (poll target) |
| `GET`  | `/readings` | Reading history (auth) |
| `GET`  | `/cards` `?arcana=&suit=` | Browse the deck |
| `GET`  | `/cards/:nameShort` | A single card |
| `GET`  | `/daily` | Card of the day |
| `GET`  | `/spreads` | Available spreads + positions |
| `POST` | `/auth/register` · `/auth/login` · `/auth/logout`, `GET /auth/me` | Accounts |

## Testing

```sh
cd backend && go test ./...     # in-memory SQLite; no Postgres or API key needed
cd frontend && npm run check    # type-check
```

## Acknowledgements

Card data from [tarotapi.dev](https://tarotapi.dev/). Readings by [Claude](https://www.anthropic.com/claude).
