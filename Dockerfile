FROM docker pull golang:1.21.0-bullseye AS builder

COPY . .

RUN go build ./app -o /app

FROM debian:11 AS runner

COPY --from=builder /app /app

ENTRYPOINT [ "/app" ]