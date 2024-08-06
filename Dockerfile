FROM golang:1.20 AS builder

WORKDIR /app

COPY . .

RUN GOOS=linux GOARCH=amd64 go build -o clamd-api ./cmd/clamd-api/main.go

FROM alpine:latest

WORKDIR /app

COPY --from=builder /app/clamd-api .
COPY ./configs/clamav-api.yaml ./configs/

# 暴露应用运行端口（根据实际情况设置）
EXPOSE 8080

CMD ["./clamd-api", "-f", "./configs/clamav-api.yaml"]
