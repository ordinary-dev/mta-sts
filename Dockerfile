FROM golang:1.20.4-alpine AS builder

WORKDIR /app

COPY go.mod ./
COPY *.go ./

RUN go build -o main

FROM alpine:3.17.1

WORKDIR /app
COPY --from=builder /app/main /app/main

ENV PORT="8080"
EXPOSE 8080

ENTRYPOINT ["/app/main"]
