FROM golang:1.24 AS builder

WORKDIR /app

COPY go.mod go.sum ./

RUN go mod tidy

COPY . .

RUN CGO_ENABLED=0 GOOS=linux go build -o link_service ./cmd/link_service/main.go

FROM alpine:3.21

WORKDIR /app

COPY --from=builder /app/link_service .

CMD ["./link_service"]