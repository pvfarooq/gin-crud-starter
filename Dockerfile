FROM golang:1.23-alpine

WORKDIR /app

RUN apk add --no-cache \
    build-base \
    inotify-tools \
    git \
    curl \
    bash

# Install golang-migrate
RUN go install -tags 'postgres' github.com/golang-migrate/migrate/v4/cmd/migrate@latest

COPY go.mod go.sum ./

RUN go mod download

COPY . .

