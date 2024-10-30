FROM service-builder AS builder

ARG SERVICE

ENV CGO_ENABLED=0

RUN go build -v -o /app/$SERVICE /app/backend/service/$SERVICE

FROM alpine:3.20

WORKDIR /app

ARG SERVICE
ENV APP_PATH=/app/$SERVICE

COPY --from=builder /app/$SERVICE .

ENTRYPOINT ${APP_PATH}
