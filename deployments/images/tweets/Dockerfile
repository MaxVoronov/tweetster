# Build image
# ToDo: AS builder
FROM golang:1.15-alpine

ENV CGO_ENABLED=0
ENV GOFLAGS="-mod=readonly"

RUN apk add --no-cache --update ca-certificates git

RUN mkdir -p /app
WORKDIR /app

COPY go.* ./
RUN go mod download

COPY . .

RUN go build -o bin/service ./cmd/tweets/main.go

CMD ["/app/bin/service"]
