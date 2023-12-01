FROM golang:1.21.4

WORKDIR /app/image-service

COPY . .
COPY sa.json .

RUN go build -o image-service ./cmd/app/main.go

ENV SERVER_ADDR=8083
ENV FIREBASE_BUCKET=

EXPOSE ${SERVER_ADDR}

CMD ["./image-service"]