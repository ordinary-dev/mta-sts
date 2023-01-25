FROM golang:1.19.5-alpine AS builder

WORKDIR /app

COPY go.mod ./
COPY *.go ./

RUN go build -o main

FROM alpine:3.17.1

WORKDIR /app
COPY --from=builder /app/main /app/main

ENTRYPOINT ["/app/main"]
