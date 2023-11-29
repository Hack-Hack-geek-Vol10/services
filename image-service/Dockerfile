FROM golang:1.21.4-alpine3.14

WORKDIR /app/image-service

COPY . .

RUN go build -o image-service

ENV SERVER_ADDR=8084
ENV FIREBASE_BUCKET=

EXPOSE ${SERVER_ADDR}

CMD ["./image-service"]