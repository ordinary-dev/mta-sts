# MTA-STS

A small server responding to the `/.well-known/mta-sts.txt` path.

Status: feature-complete, pull requests are accepted.


## Preparation

You need to set up DNS records correctly, otherwise this server will be useless.
The server must be available on the `mta-sts` subdomain.


## Environment variables

All configuration is done through environment variables.

- `PORT` - the port used by the server, 8080 by default
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
