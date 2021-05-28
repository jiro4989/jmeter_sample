FROM golang:1.16.4 AS base

RUN wget -O /usr/local/bin/air https://github.com/cosmtrek/air/releases/download/v1.27.3/air_1.27.3_linux_amd64
RUN chmod +x /usr/local/bin/air
WORKDIR /app
COPY . .
RUN go mod download
