# Contributing Guidelines

Thank you for contributing to RustDesk API Server Pro.

This project is an independent compatibility-oriented implementation. Contributions must avoid implying official endorsement by RustDesk, Huawei, HarmonyOS, OpenHarmony, or any other third-party vendor.

## Before Submitting Changes

Please check:

- Code builds successfully.
- Docker image still serves `/`, `/admin/*`, and `/api/*` through the same service.
- No production secrets, databases, logs, uploaded recordings, tokens, private keys, or OAuth credentials are committed.
- New dependencies are necessary and have licenses compatible with AGPL-3.0 distribution.
- New documentation uses descriptive compatibility language and does not imply official affiliation.
- If you modify behavior exposed over a network, AGPL-3.0 source availability obligations are considered.

## Legal and License Requirements

By contributing, you agree that your contribution can be distributed under the repository license, AGPL-3.0-only.

Do not submit code, assets, documentation, translations, or generated files unless you have the right to contribute them.

Do not copy proprietary SDKs, private API documentation, paid assets, non-redistributable binaries, screenshots containing private data, or code from incompatible licenses.

## Dependency Changes

For Go dependency changes:

```bash
cd backend
go mod tidy
go list -m all
```

For frontend dependency changes:

```bash
cd soybean-admin
pnpm install --frozen-lockfile
pnpm build
```

Update `THIRD_PARTY_NOTICES.md` when adding major bundled dependencies, templates, base images, or third-party service integrations.

## Security Reports

Do not open public issues with exploit details or secrets. See `SECURITY.md`.

## Pull Request Checklist

- [ ] I did not commit secrets or runtime data.
- [ ] I preserved license and third-party notices.
- [ ] I updated documentation where behavior changed.
- [ ] I avoided misleading trademark or affiliation claims.
- [ ] I tested Docker or binary deployment path relevant to my change.
