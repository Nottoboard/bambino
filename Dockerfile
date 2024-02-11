# syntax=docker/dockerfile:1
FROM golang:1.22.0-alpine3.19

WORKDIR /bambino

COPY . .

RUN go mod download

RUN go build -o /bambino_server

EXPOSE 4444

CMD ["/bambino_server"]