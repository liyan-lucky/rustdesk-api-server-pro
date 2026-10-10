# RustDesk compatibility update detected

A new upstream RustDesk release was detected and the compatibility target files were updated automatically.

- Current target before update: `1.4.9`
- Latest upstream: `1.5.0`

Generated reports:

- `docs/compat/reports/api-hints-1.5.0.md`

Suggested review checklist:

- [ ] Review upstream release changelog.
- [ ] Compare generated API hints with `docs/API_MATRIX.md`.
- [ ] Run Docker image build.
- [ ] Run `bash scripts/compat/smoke-api.sh http://127.0.0.1:12345`.
- [ ] Run login/currentUser/logout smoke tests with `TOKEN` if available.
- [ ] Run address book smoke tests.
- [ ] Run audit/record smoke tests.
- [ ] Confirm whether any new upstream API behavior needs real code changes beyond version target update.
