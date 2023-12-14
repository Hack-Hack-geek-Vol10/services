FROM golang:1.21.4 AS builder

WORKDIR /app/image-service

COPY . .
COPY sa.json .

RUN go build -o image-service ./cmd/app/main.go

FROM ubuntu AS runner
WORKDIR /app/image-service
COPY --from=builder /app/image-service/image-service /app/image-service
COPY ./sa.json /app/image-service

ENV SERVER_ADDR=8083
ENV FIREBASE_BUCKET=

ENV NEW_RELIC_LICENSE_KEY=
ENV NEW_RELIC_APP_NAME=
ENV NEW_RELIC_CODE_LEVEL_METRICS_ENABLED=
ENV NEW_RELIC_DISTRIBUTED_TRACING_ENABLED=
ENV NEW_RELIC_PROCESS_HOST_DISPLAY_NAME=
ENV NEW_RELIC_LABELS=


EXPOSE ${SERVER_ADDR}

CMD ["./image-service"]