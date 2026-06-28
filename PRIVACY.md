# Privacy Notice

This project is self-hosted software. The repository maintainer does not receive data from deployments unless an operator intentionally sends logs, reports, screenshots, databases, configuration files, or other materials.

This notice describes common data categories that an operator may process when deploying this project. Operators are responsible for adapting this notice to their own environment and legal obligations.

## Data the Software May Process

Depending on configuration and usage, a deployment may process or store:

- administrator account names and password hashes;
- user account information;
- device identifiers and device metadata;
- address book entries and notes;
- audit logs and login records;
- OAuth provider identifiers, callback state, and email/profile claims;
- SMTP configuration and email metadata;
- uploaded recording files in `record_uploads/`;
- IP addresses and user-agent data in application or reverse proxy logs.

## Data Storage

By default, Docker deployments should persist `/app/data`, which may contain:

- `server.db`;
- runtime `server.yaml`;
- `.init.lock`;
- `record_uploads/`;
- other runtime data created by the application.

Operators should protect this directory, back it up securely, and avoid sharing it publicly.

## Third-Party Services

If OAuth, SMTP, reverse proxy, telemetry, monitoring, or external storage is configured, data may be sent to those third-party services according to the operator's configuration and the third party's terms.

This repository does not provide legal compliance guarantees for GDPR, CCPA, PIPL, HIPAA, financial regulation, telecom regulation, export controls, or other sector-specific requirements.

## Operator Responsibilities

Operators should:

- publish their own privacy policy if the service is offered to other users;
- minimize collected logs and retention periods;
- secure OAuth secrets and SMTP credentials;
- restrict access to admin functions;
- use HTTPS for public deployments;
- define backup retention and deletion procedures;
- provide user data export/deletion workflows where required by law.

## Sharing Logs or Issues

Before opening public issues or sharing logs, remove:

- passwords and password hashes;
- OAuth client secrets and tokens;
- SMTP credentials;
- database files;
- user emails and identifiers;
- IP addresses where not necessary;
- uploaded recordings or personal files.
