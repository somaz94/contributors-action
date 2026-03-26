# CLAUDE.md

<br/>

- Do not include `Co-Authored-By` lines in commit messages.
- Do not push to remote. Only commit. The user will push manually.
- Do not modify git config.

<br/>

## Project Structure

- Go-based GitHub Action (Docker container action)
- Generates and updates contributor lists from GitHub repository data
- Output formats: HTML table, Markdown list, image grid
- Features: bot filtering, section update markers, dry-run mode, cross-repo support

<br/>

## Build & Test

```bash
make build       # Build binary
make test        # Unit tests with coverage
make cover       # Generate coverage report
make fmt         # Format code
make lint        # Run go vet
```

<br/>

## Key Directories

- `cmd/` — Entry point (main.go)
- `internal/config/` — Configuration loading from env vars (INPUT_*)
- `internal/github/` — GitHub API client, contributor filtering, sorting
- `internal/formatter/` — Output formatting (table/list/image)
- `internal/writer/` — File writing, section updates with markers
- `docs/` — Usage guides

<br/>

## Action Inputs

Key inputs: `token`, `output_file`, `format` (table/list/image), `columns`, `max_contributors`, `exclude`, `include_bots`, `avatar_size`, `sort_by`, `update_section`, `dry_run`

<br/>

## CI

- `ci.yml` — Unit tests (80% coverage threshold), Docker build, 6 integration tests
- Docker: multi-stage build (golang:1.26-alpine → alpine:3.23)

<br/>

- Communicate with the user in Korean.
- All documentation and code comments must be written in English.
