# Cimeika Security Policy

## Key Guidelines
- No API keys are stored in public repositories or shared Google Drive folders.
- All sensitive configurations are loaded from secured `.env` files.
- Firebase access rules enforce authenticated access.

## Incident Response
If a key is compromised:
1. Revoke the key immediately in the source platform.
2. Rotate the key and update `.env`.
3. Record the incident in `SECURITY_LOG.md`.

## Contacts
ci@cimeika.com.ua
