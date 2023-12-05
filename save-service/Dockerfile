FROM golang:1.21.4 As builder

WORKDIR /app/save-service

COPY . .

RUN go build -o save-service ./cmd/app/main.go

FROM ubuntu AS runner
WORKDIR /app/save-service
COPY --from=builder /app/save-service/save-service /app/save-service/save-service

ENV SERVER_ADDR=8086
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

CMD ["./save-service"]