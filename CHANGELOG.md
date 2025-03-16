# Changelog


## [0.2.0] - 2025-03-16

### New features

- Add the ability to listen on a unix socket (see `SOCKET_PATH`).
- Added the ability to specify an IP address for binding (see `LISTEN_ADDRESS`).

### Changed

- Added a version of the program without Docker (x86-64 linux only).
- Handle interrupts and wait for connections to complete.
- Initialize response on startup.
- Update github actions and alpine image.


## [0.1.2] - 2024-10-25

### Changed

- Update readme.md to fix MTA_STS_MX typo.
- Update golang and alpine versions.
- Make documentation more useful.
- Remove unnecessary comments, rename the request handler.
- Update github actions.
