FROM golang:1.23-alpine AS builder

WORKDIR /app

COPY go.mod ./
COPY *.go ./

RUN go build -o main

FROM alpine:3.20

WORKDIR /app
COPY --from=builder /app/main /usr/local/bin/mta-sts

RUN adduser --disabled-password mtasts
USER mtasts

ENV PORT="8080"
EXPOSE 8080

ENTRYPOINT ["/usr/local/bin/mta-sts"]
