# Используем официальный образ Go в качестве базового
FROM golang:1.23-alpine

# Устанавливаем рабочую директорию внутри контейнера
WORKDIR /app/backend

RUN go install github.com/swaggo/swag/cmd/swag@latest
RUN go install github.com/air-verse/air@latest
# Копируем файлы go.mod и go.sum и устанавливаем зависимости
COPY go.mod go.sum ./
RUN go mod download

COPY . .


RUN chmod +x entrypoint.sh

# Копируем все файлы проекта в контейнер
ENTRYPOINT ["/app/backend/entrypoint.sh"]



# Запускаем приложение с использованием go run
CMD ["air", "-c", "/app/.air.toml"]
