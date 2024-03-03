# Dockerfile obtained from https://github.com/cosmtrek/air#installation-and-usage-for-docker-users-who-dont-want-to-use-air-image

# Choose whatever you want, version >= 1.16
FROM golang:1.22-alpine

WORKDIR /app

RUN go install github.com/cosmtrek/air@latest

COPY go.mod go.sum ./
RUN go mod download

CMD ["air", "-c", ".air.toml"]