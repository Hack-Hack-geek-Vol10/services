FROM golang:1.21.4 AS builder

WORKDIR /app/user-service

COPY . .

RUN go build -o user-service ./cmd/app/main.go

FROM alpine:3.14.2 AS runner
WORKDIR /app/user-service
COPY --from=builder /app/user-service/user-service /app/user-service/user-service

ENV SERVER_ADDR=8081
ENV PSQL_HOST=
ENV PSQL_PORT=
ENV PSQL_USER=
ENV PSQL_PASSWORD=
ENV PSQL_DBNAME=
ENV PSQL_SSLMODE=
ENV PSQL_CONNECT_TIMEOUT=
ENV PSQL_CONNECT_WAIT_TIME=
ENV PSQL_CONNECT_ATTEMPTS=
ENV PSQL_CONNECT_BLOCKS=
ENV PSQL_CLOSE_TIMEOUT=

EXPOSE ${SERVER_ADDR}

CMD ["./user-service"]