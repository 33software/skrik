# Используем официальный образ Go в качестве базового
FROM golang:1.23-alpine

# Устанавливаем рабочую директорию внутри контейнера
WORKDIR /app

RUN go install github.com/swaggo/swag/cmd/swag@latest
RUN go install github.com/air-verse/air@latest
# Копируем файлы go.mod и go.sum и устанавливаем зависимости
COPY go.mod go.sum ./
RUN go mod download

# Копируем все файлы проекта в контейнер
COPY . .

# Запускаем приложение с использованием go run
CMD ["air", "-c", ".air.toml"]