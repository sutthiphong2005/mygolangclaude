# CLAUDE.md

This file provides guidance to Claude Code (claude.ai/code) when working with
code in this repository.

## Project

Full-stack monorepo with two independently runnable parts:

- **`frontend/`** — React + TypeScript single-page app, built with Vite. Calls
  the backend's HTTP API. See `frontend/CLAUDE.md`.
- **`backend/`** — Go HTTP server (`helloworld` module). Serves the API on
  `:8080`. See `backend/CLAUDE.md`.

Each part has its own `CLAUDE.md` with stack-specific guidance and commands;
this root file covers how they fit together. When working inside a part, the
nearest `CLAUDE.md` plus this one both apply.

## How the parts connect

- Backend listens on `:8080`; the frontend dev server (Vite) runs on `:5173` and
  proxies API calls to `:8080` during development.
- The API contract is plain HTTP returning text/JSON. Keep frontend types in
  sync with backend response shapes — when a response shape changes, update both
  sides in the same change.

## Skills

Reusable coding conventions live as on-demand skills under `.claude/skills/`:

- `go-best-practices` — Go idioms + the backend's message/handler/transport
  split. Triggers on Go work.
- `react-best-practices` — React + TypeScript + Vite conventions. Triggers on
  frontend work.

## Commands

Run commands from the relevant subdirectory (`frontend/` or `backend/`). See
each part's `CLAUDE.md` for the full list. Quick reference:

```sh
# backend (from backend/)
go run .          # serve API on :8080
go test ./...     # Go tests

# frontend (from frontend/)
npm run dev       # Vite dev server on :5173
npm test          # frontend tests
```
