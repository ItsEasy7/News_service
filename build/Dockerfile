FROM golang:latest AS builder

COPY go.mod go.sum ./
# RUN go mod tidy

WORKDIR /app
COPY . .


RUN CGO_ENABLED=0 GOOS=linux go build -o main ./cmd/main.go

FROM alpine:latest

WORKDIR /app
COPY --from=builder /app/main .
COPY --from=builder /app/.env .

EXPOSE 8080

CMD ["./main"]