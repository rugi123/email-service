FROM golang:1.25-alpine AS builder
WORKDIR /app
COPY go.mod go.sum ./
RUN go mod download
COPY . .
RUN go build -o main ./cmd/worker

# Создаем новый этап для финального образа
FROM alpine:latest

RUN apk --no-cache add ca-certificates

WORKDIR /root/

# Копируем бинарник из builder этапа
COPY --from=builder /app/main .

# Копируем конфигурационный файл в текущую рабочую директорию
COPY --from=builder /app/internal/config/config.yaml ./config.yaml

CMD ["./main"]