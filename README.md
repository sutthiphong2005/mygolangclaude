# myapp

Full-stack monorepo: a React + TypeScript frontend and a Go HTTP backend.

## Layout

```
├── frontend/   # React + TypeScript (Vite) SPA
└── backend/    # Go HTTP server (helloworld)
```

## Backend

```sh
cd backend
go run .                       # serve on :8080
curl http://localhost:8080/    # -> Hello, World!
go test ./...                  # tests
```

## Frontend

Not yet scaffolded. To initialize:

```sh
cd frontend
npm create vite@latest . -- --template react-ts
npm install
npm run dev                    # Vite dev server on :5173
```

See `CLAUDE.md` (and the per-directory `CLAUDE.md` files) for development
guidance.
