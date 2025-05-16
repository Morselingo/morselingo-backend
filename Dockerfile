# syntax=docker/dockerfile:1.4

FROM golang:1.24.3-alpine3.21

RUN apk add --no-cache git

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY . .

RUN go build -o morse-backend ./cmd/server
EXPOSE 8080

CMD ["./morse-backend"]