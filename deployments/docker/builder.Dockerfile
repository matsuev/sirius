FROM golang:1.22.8-alpine3.20 AS service-builder

WORKDIR /app

COPY ./go.mod ./go.sum ./

RUN go mod download

COPY ./backend ./backend
