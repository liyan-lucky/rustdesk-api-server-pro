# Security Policy

## Supported Scope

This repository is a community-maintained compatibility implementation. Security fixes are handled on a best-effort basis.

Operators should validate every release before production deployment and should not expose the service to the public Internet without TLS, strong credentials, backups, and access controls.

## Reporting a Vulnerability

Please do not publish exploit details publicly before a fix or mitigation is available.

Recommended report content:

- affected version or commit;
- deployment mode, such as binary, Docker, OpenWrt, reverse proxy;
- reproduction steps;
- logs with secrets removed;
- impact assessment;
- whether the issue affects authentication, authorization, file upload, OAuth, database access, or remote code execution.

Use GitHub issues only for non-sensitive reports. For sensitive issues, contact the maintainer through a private channel available on the maintainer's GitHub profile or repository settings.

## Operator Hardening Checklist

- Change and permanently keep a strong `signKey`.
- Use strong `ADMIN_PASS` values and rotate default credentials immediately.
- Never commit OAuth client secrets, SMTP passwords, database files, tokens, private keys, or production `server.yaml` files.
- Use HTTPS behind a trusted reverse proxy for public deployments.
- Pass `Host`, `X-Forwarded-Host`, `X-Forwarded-Proto`, and `X-Real-IP` correctly when using a reverse proxy.
- Restrict access to the admin interface where possible.
- Back up `/app/data` or the mapped host data directory regularly.
- Keep Docker base images and host packages updated.
- Review logs before sharing them; remove credentials, tokens, email addresses, IP addresses, and user identifiers when necessary.

## Known Sensitive Data Locations

Common sensitive locations include:

- `server.yaml` and `/app/data/server.yaml`;
- SQLite database files such as `server.db`;
- `record_uploads/`;
- container logs;
- OAuth callback and token exchange logs;
- reverse proxy access logs.

## AGPL and Security Fixes

If you modify and run this project for users over a network, AGPL-3.0 may require you to offer those users the corresponding source code of the modified version, including security fixes and local patches.
