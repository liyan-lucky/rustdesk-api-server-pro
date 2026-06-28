# Compliance Checklist

This checklist records the repository-level compliance baseline for RustDesk API Server Pro.

It is not legal advice. It is intended to reduce obvious open-source, trademark, release, security, and privacy risks.

## Current Baseline

| Area | Status | Notes |
|---|---:|---|
| Repository license | Present | `LICENSE` contains AGPL-3.0. |
| Network copyleft notice | Added | `NOTICE` and README explain AGPL network-source obligations. |
| Third-party attribution | Added | `THIRD_PARTY_NOTICES.md` records Soybean Admin and dependency inventory process. |
| Trademark disclaimer | Added | `NOTICE` and `DISCLAIMER.md` clarify non-affiliation and descriptive use. |
| Privacy notice | Added | `PRIVACY.md` lists common self-hosted data categories. |
| Security policy | Added | `SECURITY.md` adds reporting and hardening guidance. |
| Release artifacts | Needs operator check | Do not include secrets, databases, logs, recordings, or production config. |
| Dependency license report | Recommended | Generate formal Go/npm license inventory before major public releases. |
| SBOM | Recommended | Generate container/binary SBOM before production release. |

## Mandatory Rules for Releases

Before pushing a release tag:

1. Confirm `LICENSE`, `NOTICE`, `THIRD_PARTY_NOTICES.md`, `DISCLAIMER.md`, `SECURITY.md`, and `PRIVACY.md` are present.
2. Confirm release artifacts do not contain:
   - `server.db`;
   - production `server.yaml`;
   - OAuth secrets;
   - SMTP passwords;
   - tokens or private keys;
   - uploaded recordings;
   - logs containing personal data.
3. Confirm Release text does not imply official RustDesk, Huawei, HarmonyOS, OpenHarmony, GitHub, Google, Docker, GHCR, or OpenWrt endorsement.
4. Confirm Docker images are built from repository source and do not include hidden binary-only modifications.
5. Confirm modified network deployments can provide corresponding source under AGPL-3.0.

## Recommended Commands

Go module inventory:

```bash
cd backend
go list -m all
```

Frontend dependency inventory:

```bash
cd soybean-admin
pnpm install --frozen-lockfile
pnpm licenses list || true
```

Secret scanning examples:

```bash
git grep -nE "(clientSecret|password|secret|token|private_key|BEGIN RSA|BEGIN OPENSSH|AKIA|ghp_)" -- . ':!COMPLIANCE.md' ':!SECURITY.md' ':!PRIVACY.md'
```

Docker image SBOM example:

```bash
syft ghcr.io/liyan-lucky/rustdesk-api-server-pro:latest -o table
```

## Known Project-Specific Risks

- The project uses RustDesk-compatible names, routes, and compatibility behavior. Keep all statements descriptive and avoid implying official affiliation.
- The frontend is based on Soybean Admin. Preserve MIT attribution.
- AGPL-3.0 obligations are important for hosted network services.
- OAuth, SMTP, database, and recording-upload features can process personal or sensitive data.
- OpenWrt deployment scripts may remove old containers; keep backups and communicate this behavior clearly.

## Maintainer Notes

If a future release changes dependency sets, bundled frontend assets, Docker base images, or third-party service integrations, update this checklist and `THIRD_PARTY_NOTICES.md` accordingly.
