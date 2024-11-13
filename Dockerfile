# Использование образа Golang
FROM golang:1.19 AS builder
WORKDIR /app
COPY . .
RUN go build -o ServiceT

# Использование образа на основе Alpine для минимального размера
FROM alpine:latest
WORKDIR /root/
COPY --from=builder /app/ServiceT .
CMD ["./ServiceT"]