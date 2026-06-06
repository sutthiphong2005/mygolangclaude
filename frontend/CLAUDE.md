# CLAUDE.md — frontend

Guidance for the React frontend. The root `CLAUDE.md` covers the monorepo as a
whole; this file applies when working under `frontend/`.

## Project

React + TypeScript single-page app built with **Vite**. Talks to the Go backend
(`:8080`) over HTTP. Not yet scaffolded — when initializing, use:

```sh
npm create vite@latest . -- --template react-ts
npm install
```

## Architecture (intended)

Keep a clear separation, mirroring the backend's discipline:

- **API layer** (`src/api/`) — functions that call the backend and return typed
  data. No React here. This keeps data-fetching testable and swappable.
- **Components** (`src/components/`) — presentational, props-driven, no direct
  `fetch`.
- **Hooks** (`src/hooks/`) — stateful logic (data loading, etc.) that bridges
  the API layer and components.

Keep response types in `src/api/` in sync with the backend's response shapes.

For detailed React/TypeScript conventions, the `react-best-practices` skill
loads on demand.

## Commands

Run these from the `frontend/` directory:

```sh
npm run dev       # Vite dev server on :5173
npm run build     # production build
npm run preview   # preview the production build
npm run lint      # ESLint
npm test          # tests (configure Vitest)
```

## Dev proxy

Configure Vite to proxy API calls to the backend during development, so the
frontend can call `/api/...` without CORS issues. In `vite.config.ts`:

```ts
server: {
  proxy: { '/api': 'http://localhost:8080' },
}
```
