# MTA-STS

A small server responding to the `/.well-known/mta-sts.txt` path.

Status: feature-complete, pull requests are accepted.


## Preparation

You need to set up DNS records correctly, otherwise this server will be useless.
The server must be available on the `mta-sts` subdomain.


## Quick start

```sh
docker run -it -p 8080:8080 ghcr.io/ordinary-dev/mta-sts:v0.2.0
```


## Configuration

All configuration is done through environment variables.

### Server settings

Select one option:

- `SOCKET_PATH` - Unix socket path, for example: `/run/mta-sts.sock`.
- `LISTEN_ADDRESS` - Address for the server, for example: `127.0.0.1:8080`.
- `PORT` - The port used by the server, `8080` by default (deprecated, prefer `LISTEN_ADDRESS`).

### MTA-STS settings

- `MTA_STS_MODE` - one of the following values: enforce, testing or none (enforce by default)
- `MTA_STS_MAX_AGE` - the maximum lifetime of the policy in seconds, default - 604800 (7 days)
- `MTA_STS_MX` - list of mail servers separated by commas, for example: "mx1.example.com,mx2.example.com"


## Docker-compose example

```yml
services:
  mta-sts:
    image: ghcr.io/ordinary-dev/mta-sts
    environment:
      MTA_STS_MX: mx.example.com
    ports:
      - 80:8080
```
