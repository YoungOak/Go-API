FROM golang:1.21.0-bullseye AS builder

COPY . /api

WORKDIR /api/

RUN go build -o /app /api/app/ 

FROM debian:11 AS runner

COPY --from=builder /app /app

ENTRYPOINT [ "/app" ]