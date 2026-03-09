# V3 Quality Gates

This document tracks the V3 quality automation added to the repository.

## Frontend

Workflow: `.github/workflows/frontend-quality.yml`

Checks:
- `pnpm i18n:check:strict`
- `pnpm i18n:sanitize:guard`
- `pnpm i18n:check:report`
- `pnpm i18n:summary`
- `pnpm i18n:compare`
- `pnpm ci:frontend`

Artifacts:
- `soybean-admin/docs/i18n-coverage-report.md`
- `soybean-admin/docs/i18n-quality-summary.json`
- `soybean-admin/docs/i18n-sanitize-report.json`

## Backend

Workflow: `.github/workflows/backend-quality.yml`

Checks:
- `go vet ./...`
- `go test ./...`

## Local Quick Commands

```bash
pnpm --dir soybean-admin run i18n:check
pnpm --dir soybean-admin run i18n:check:strict
pnpm --dir soybean-admin run i18n:sanitize:guard
pnpm --dir soybean-admin run i18n:compare
pnpm --dir soybean-admin run ci:frontend
go test ./... -C backend
```
