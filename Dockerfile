# syntax=docker/dockerfile:1

FROM golang:1.16-alpine

WORKDIR /app
COPY go.mod ./
COPY go.sum ./

RUN go mod download

COPY . .
COPY /servers/product_server.go ./

RUN go build -o /waste-saver-grpc-server

EXPOSE 50051

CMD [ "/waste-saver-grpc-server" ]