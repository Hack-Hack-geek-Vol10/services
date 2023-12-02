FROM golang:1.21.4 AS builder

WORKDIR /app/member-service

COPY . .

RUN go build -o member-service ./cmd/app/main.go

FROM ubuntu AS runner
WORKDIR /app/member-service
COPY --from=builder /app/member-service/member-service /app/member-service/member-service

ENV SERVER_ADDR=8084
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

CMD ["./member-service"]