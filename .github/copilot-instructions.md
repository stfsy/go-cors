# go-cors — Copilot / AI contributor instructions

This small repo implements CORS middleware and adapters for several Go HTTP frameworks. Keep this file as the single-page guide for automated coding agents and contributors.

Essentials (what to know):
- Purpose: library package that exposes CORS middleware and helpers. See `cors.go` for the public API and `utils.go` for helpers.
- Adapters: integration examples and framework wrappers live under `examples/` (per-framework servers) and `wrapper/` (framework-specific adapter code, e.g. `wrapper/gin/gin.go`).
- Internal helpers: `internal/sortedset.go` contains the internal implementation used by tests and middleware.

Quick commands (repo root):
- Fetch deps: `go mod download`
- Run tests: `go test ./...` (unit tests live at root and `internal/`, wrapper tests under `wrapper/`)
- Run benchmarks: `go test -bench . -benchmem`
- Run an example server: `go run ./examples/nethttp` or `cd examples/nethttp && go run .` (each example is a small module under `examples/`)

Project-specific patterns and conventions:
- Public surface: `cors.go` contains the exported types/config (middleware constructors). Avoid changing exported signatures without adding tests.
- Examples-as-integration-tests: `examples/*/server.go` show how to wire the middleware into different routers (net/http, chi, gin, gorilla, httprouter, etc.). Use them as living documentation.
- Wrapper adapters: `wrapper/<framework>/` contains small adapter code and tests demonstrating the mapping between this package and framework middleware hooks (see `wrapper/gin/gin.go`, `wrapper/gin/gin_test.go`).
- internal package: anything under `internal/` is private API — don't make it public. Tests target behavior rather than internal details when possible.

Testing notes and expectations:
- Tests are deterministic and fast; prefer adding unit tests alongside code (`*_test.go`) and exercising public APIs.
- When adding functionality, include at least one unit test and, if relevant, an updated example under `examples/`.
- Use `go test ./...` to validate changes. Benchmarks live in `bench_test.go`.

When modifying dependencies or adding new modules:
- Update top-level `go.mod` when library changes required packages.
- Examples may have their own `go.mod` files; run `go mod download` inside the example directory if you modify its imports.

Files to look at immediately:
- `cors.go` — core middleware and config
- `utils.go` — helpers used by middleware and tests
- `internal/sortedset.go` — internal datastructure used by header parsing/ordering
- `wrapper/gin/gin.go` and `wrapper/gin/gin_test.go` — example adapter + tests
- `examples/*/server.go` — runnable examples for each router

Commit and PR guidance for agents:
- Keep changes minimal and focused. Add/adjust tests for behavior changes.
- Semantic commits are required. Use Conventional Commits (feat|fix|docs|chore|refactor|perf|test). Example: `feat(cors): allow wildcard origin for specific adapters`.
- Follow Go idioms used in repo: small, well-tested functions; prefer table-driven tests when appropriate.