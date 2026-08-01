# Release Process

Use the original tag-based workflows for releases:

- `.github/workflows/build-release.yml`
- `.github/workflows/ghcr-docker.yml`

Do not create one-off release workflows unless absolutely necessary.

## Continuous Delivery (main branch)

Every push to `main` starts the online quality workflows. After all push-triggered workflows for the same commit complete successfully, `.github/workflows/ghcr-docker.yml` automatically:

1. Bumps the PATCH version in `VERSION` file (e.g., `1.1.16` → `1.1.17`)
2. Builds and pushes Docker image to GHCR with tags: `latest`, `main`, `sha-xxxxxxx`
3. Injects `APP_VERSION` and `BUILD_TIME` via ldflags (backend) and Vite env vars (frontend)

If any online workflow for that commit fails, GHCR publishing is blocked. No manual steps are needed when the quality gates are green.

## Standard Release Steps (tag-based)

1. Ensure `main` is clean and all intended changes are merged.
2. Confirm the `VERSION` file reflects the desired release version.
3. Confirm compliance files are present:
   - `LICENSE`
   - `NOTICE`
   - `THIRD_PARTY_NOTICES.md`
   - `DISCLAIMER.md`
   - `SECURITY.md`
   - `PRIVACY.md`
   - `COMPLIANCE.md`
   - `CONTRIBUTING.md`
   - `CODE_OF_CONDUCT.md`
   - `SUPPORT.md`
4. Confirm no secrets or runtime data are tracked.
5. Push a semantic tag matching `v*.*.*`.

Example:

```bash
git checkout main
git pull origin main
git tag v1.1.17
git push origin v1.1.17
```

## Expected Workflows

Tag push should trigger:

- `build-release.yml`: builds Linux, Windows and macOS packages for amd64 and arm64, then uploads zip assets to GitHub Release.
- `ghcr-docker.yml`: builds and pushes GHCR image tags, including `latest`, version tag (e.g., `1.1.17`) and `sha-*`.

## Do Not Include in Release Assets

- production databases;
- production config files;
- OAuth client secrets;
- SMTP credentials;
- user recordings;
- private logs;
- private keys or certificates.

## If a Release Fails

Prefer fixing the original workflow instead of adding a new one-off workflow.

If a tag or release is partially created:

1. Delete the failed GitHub Release if present.
2. Delete the failed tag both remotely and locally.
3. Fix the workflow or source issue.
4. Re-create the tag.

Example:

```bash
git tag -d v1.1.17 || true
git push origin :refs/tags/v1.1.17 || true
git tag v1.1.17
git push origin v1.1.17
```
