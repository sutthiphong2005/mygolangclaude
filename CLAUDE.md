# CLAUDE.md

This file provides guidance to Claude Code (claude.ai/code) when working with code in this repository.

## Project

Single-package Go module (`helloworld`) — a minimal HTTP server with two routes: `/` returns "Hello, World!" and `/demo` returns "demo". `main.go` has three layers: message-producing functions (`greeting()`, `demo()`) return the response strings, handlers (`helloHandler`, `demoHandler`) write them to an `http.ResponseWriter`, and `main()` registers the handlers and listens on `:8080`. Tests cover each message function directly and each handler via `net/http/httptest`. This message/handler/transport split is what keeps messages testable independently of HTTP — preserve it when adding routes: write a new message function and a new handler that calls it, rather than inlining strings in handlers.

## Commands

```sh
go run .                  # run
go test ./...             # all tests
go test -run TestGreeting # single test by name
go build -o helloworld    # build binary
go vet ./...              # vet
gofmt -w .                # format
```
