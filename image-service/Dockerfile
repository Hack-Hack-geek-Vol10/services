FROM golang:1.21.4 AS builder

WORKDIR /app/image-service

COPY . .
COPY sa.json .

RUN go build -o image-service ./cmd/app/main.go

FROM apline AS runner
WORKDIR /app/image-service
COPY --from=builder /app/image-service/image-service /app/image-service/image-service

ENV SERVER_ADDR=8083
ENV FIREBASE_BUCKET=

EXPOSE ${SERVER_ADDR}

CMD ["./image-service"]