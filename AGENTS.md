# AGENTS.md

Default theme for [Resume CLI](https://github.com/ansidev/resume-cli): Handlebars templates, Tailwind CSS, and a Go helper file.

## Layout

- `theme.yaml`: theme manifest (`name`, `templateEngine: handlebars`, `helpers: helpers.go`).
- `helpers.go`: template helpers, entry point `Funcs() map[string]interface{}`. Package `theme`.
- `helpers_test.go`: Go tests for the helpers.
- `resources/resume.hbs`: main template; `resources/partials/*.hbs`: partials.
- `resources/assets/css/resume.css`: Tailwind source. `styles.css` is **generated output** (do not hand-edit).
- `go.mod` / `go.sum`: Go module (Go 1.27).
- `package.json` / `pnpm-lock.yaml`: Tailwind CLI only.

## Commands

```bash
go build ./...          # build
go test ./...           # run tests
go mod tidy             # after changing dependencies

pnpm install            # install Tailwind CLI
pnpm build              # regenerate resources/assets/css/styles.css (minified)
pnpm dev                # watch mode
```

## Rules

- `helpers.go` is interpreted at runtime by Resume CLI via [yaegi](https://github.com/traefik/yaegi). Keep it a **single self-contained file** using only the standard library, `github.com/mailgun/raymond/v2` and `github.com/yuin/goldmark`. Do not add other imports or split it into multiple files.
- Helpers run with the full privileges of the CLI process. Escape untrusted input (e.g. `raymond.Escape`) and be careful with `raymond.SafeString`.
- When adding, removing or renaming a helper, update `Funcs()`, the name list in `TestFuncs`, add a case in `TestHelpers`, and the helper list in `README.md`.
- After changing templates or classes used in `resources/`, run `pnpm build` and commit the regenerated `styles.css`.
- Use `pnpm`, not npm or yarn. `node_modules/` is git-ignored.
- Dependency upgrades: `go get -u ./... && go mod tidy`, then `go build ./... && go test ./...` must pass.
- Keep `README.md` in sync with structural changes (theme structure, helper list).

## Git

- Commit only files relevant to the change. Never push unless asked.
- License is AGPL-3.0.
