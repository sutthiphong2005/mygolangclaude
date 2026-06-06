---
name: react-best-practices
description: React + TypeScript + Vite coding best practices for the frontend. Use when writing, reviewing, or refactoring frontend code — components, hooks, state, data fetching, types, or tests. Enforces the project's api/component/hook separation.
---

# React + TypeScript Best Practices

Apply these when writing or reviewing frontend code. Project-specific rules come
first because they override generic advice.

## Project architecture (must preserve)

Mirror the backend's discipline — keep data, logic, and presentation separate:

1. **API layer** (`src/api/`) — functions that call the backend and return typed
   data. No React, no JSX. Trivially testable and swappable.
2. **Hooks** (`src/hooks/`) — stateful logic (loading, caching, side effects)
   that bridges the API layer and components.
3. **Components** (`src/components/`) — presentational, props-driven. No direct
   `fetch`; consume hooks or receive data via props.

When adding a feature: add a typed API function, a hook that uses it, and a
component that renders the hook's result — don't inline `fetch` inside a
component.

## Before you finish

Run these and make them clean (from `frontend/`):

```sh
npm run lint      # ESLint
npx tsc --noEmit  # type-check
npm test          # tests pass
npm run build     # production build succeeds
```

## TypeScript

- No `any`. Use `unknown` + narrowing when a type is genuinely open.
- Type API responses explicitly; keep them in sync with backend shapes.
- Prefer `type` aliases for props/data; `interface` is fine too — be consistent.
- Enable and respect `strict` mode in `tsconfig.json`.
- Type component props; avoid `React.FC` (prefer explicit prop typing).

## Components

- Function components only; no class components.
- One component per file; name the file after the component (`UserCard.tsx`).
- Keep components small and focused; extract subcomponents over deep JSX.
- Derive state during render instead of duplicating it in `useState`.
- Lists need a stable, unique `key` — never the array index for dynamic lists.
- Don't put non-rendering logic in components — move it to hooks or the API layer.

## Hooks

- Obey the Rules of Hooks: call them unconditionally at the top level.
- Custom hooks start with `use` and encapsulate one concern.
- `useEffect` is for synchronizing with external systems, NOT for deriving data.
  If you can compute it during render, don't use an effect.
- Always specify the dependency array correctly; don't silence the linter.
- Clean up subscriptions/timers/aborts in the effect's return function.

## State management

- Start with local `useState`; lift state only as high as needed.
- For server data, prefer a data-fetching hook with loading/error/data states
  (or a library like TanStack Query) over hand-rolled effects.
- Reach for Context for genuinely global, low-churn values (theme, auth) — not
  as a substitute for prop passing in small trees.

## Data fetching

- All network calls live in `src/api/`, returning typed results.
- Handle the three states explicitly: loading, error, success.
- Use `AbortController` to cancel in-flight requests on unmount.
- Use the Vite dev proxy (`/api` → `:8080`) instead of hardcoding the backend
  origin.

## Performance

- Don't reach for `useMemo`/`useCallback`/`memo` preemptively — profile first.
- A genuine reason to memoize: stable references passed to memoized children, or
  expensive computations on every render.
- Avoid creating new object/array/function literals in props on hot paths.

## Testing

- Test behavior, not implementation, with React Testing Library + Vitest.
- Query by accessible role/text, not by test IDs where avoidable.
- Test API-layer functions directly (mock `fetch`).
- Test hooks via a component or `renderHook`.

## Common pitfalls

- Stale closures in effects/callbacks — check the dependency array.
- Missing/duplicate `key` causing remount or state bleed across list items.
- Mutating state directly instead of replacing it (`arr.push` vs `[...arr, x]`).
- `useEffect` chains that could be a single derived value or event handler.
- Forgetting to handle the error/loading state of an async call.
