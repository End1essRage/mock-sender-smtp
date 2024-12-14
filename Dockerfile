# Используем образ golang для сборки
FROM golang:1.22 AS builder

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY . .

# Собираем приложение
RUN go build -o /app/mock-sender-smtp ./cmd/main.go

# Запускаем приложение при запуске контейнера
CMD ["./mock-sender-smtp"]

EXPOSE 8080