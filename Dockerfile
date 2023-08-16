FROM golang:1.21.0-bullseye AS builder

COPY . /api

WORKDIR /api/

ARG CGO_ENABLED=0 GOOS=linux GOARCH=amd64

RUN go build -o /app /api/app/...

FROM scratch

COPY --from=builder app app

EXPOSE 8080

ENTRYPOINT [ "/app" ]