# Simple MTA-STS setup

This is an overly complicated way to set up MTA-STS. *Or maybe not?*

I wrote this server to make it easy to set up multiple mail servers.

## Configuration
You need to set up dns records correctly, otherwise this server will be useless.
This container needs to be launched at: mta-sts.your-domain.com.
As a result, a file for configuring mta-sts will be available at /.well-known/mta-sts.txt.
All configuration is done through environment variables.

## Environment variables
- `PORT` - the port used by the server, 8080 by default
- `MTA_STS_MODE` - one of the following values: enforce, testing or none (enforce by default)
- `MTA_STS_MAX_AGE` - the maximum lifetime of the policy in seconds, default - 604800 (7 days)
- `MTA_STS_MX` - list of mail servers separated by commas, for example: "mx1.domain.com,mx2.domain.com"

## Docker-compose example
```yml
services:
  mta-sts:
    image: ghcr.io/ordinary-dev/mta-sts
    environment:
      - MTA_STS_MX=mx.domain.com
```

## Learn more
- [Article on Dmarcian](https://dmarcian.com/mta-sts/)
