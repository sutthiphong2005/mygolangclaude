---
name: go-best-practices
description: Go (golang) coding best practices and idioms for this helloworld HTTP server. Use when writing, reviewing, or refactoring Go code — adding routes/handlers, error handling, naming, formatting, concurrency, or tests. Enforces the project's message/handler/transport split.
---

# Go Coding Best Practices

Apply these when writing or reviewing Go in this repository. Project-specific
rules come first because they override generic advice.

## Project architecture (must preserve)

`main.go` keeps three layers separate so messages stay testable without HTTP:

1. **Message functions** (`greeting()`, `demo()`) — return response strings. No
   `http` types. Pure and trivially testable.
2. **Handlers** (`helloHandler`, `demoHandler`) — call a message function and
   write to `http.ResponseWriter`. No business strings inlined here.
3. **Transport** (`main()`) — registers handlers and calls `ListenAndServe`.

When adding a route: write a new message function **and** a handler that calls
it. Never inline response strings inside a handler. Mirror the test pattern —
test the message function directly, and test the handler via
`net/http/httptest`.

## Before you finish

Run these and make them clean:

```sh
gofmt -w .        # format (non-negotiable)
go vet ./...      # catch suspicious constructs
go test ./...     # all tests must pass
```

If you change a message string, update its test in the same change — they are
coupled by design.

## Naming

- Packages: short, lower-case, no underscores (`http`, not `http_util`).
- Exported identifiers need doc comments starting with the identifier name.
- Use mixedCaps / MixedCaps, never snake_case.
- Keep names short in small scopes (`i`, `r`, `w`); descriptive in wide scopes.
- Acronyms keep one case: `HTTPServer`, `userID`, `parseURL` — not `HttpServer`.

## Error handling

- Check every error; never discard with `_` unless intentional and obvious.
- Return errors, don't `panic`, in library/handler code.
- Wrap with context: `fmt.Errorf("loading config: %w", err)` to preserve the
  chain (`errors.Is` / `errors.As` still work).
- Handle the error where you have the most context; don't log **and** return the
  same error (that double-reports it).
- Keep the happy path at minimal indentation — return early on errors.

## Formatting & structure

- `gofmt` is the only style authority; never hand-format.
- Group imports stdlib / third-party / local, separated by blank lines.
- Prefer small functions with a single responsibility.
- Accept interfaces, return concrete types.
- Zero values should be useful where practical.

## HTTP handlers (this project)

- Write status codes explicitly when not 200 (`w.WriteHeader(...)`).
- Don't ignore the error from `w.Write` / `fmt.Fprintln` if it matters.
- Keep handlers thin — delegate logic to message/other functions.

## Concurrency (when it becomes relevant)

- Don't start a goroutine without knowing how it stops.
- Protect shared state with a mutex or pass ownership via channels — "share
  memory by communicating."
- Pass `context.Context` as the first arg for cancellable/IO work; never store
  it in a struct.
- Run tests with `-race` when touching concurrent code: `go test -race ./...`.

## Testing

- Table-driven tests for multiple cases; subtests via `t.Run`.
- Test message functions directly; test handlers with `httptest.NewRecorder`.
- Name tests `TestXxx`; keep `want`/`got` comparisons explicit with `t.Errorf`.
- No external network/filesystem dependencies in unit tests.

## Common pitfalls

- Loop variable capture in goroutines/closures (pre-Go 1.22) — shadow with
  `x := x` if targeting older toolchains.
- Appending to a slice shared with another reference can mutate it unexpectedly.
- Comparing or printing errors by string instead of `errors.Is`.
- Forgetting to `defer resp.Body.Close()` on HTTP responses.
