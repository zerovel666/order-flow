# Development image: Air rebuilds the application when source files change.
FROM golang:1.27.0-alpine

RUN apk add --no-cache curl
RUN go install github.com/air-verse/air@latest
RUN go install github.com/swaggo/swag/cmd/swag@v1.16.6

WORKDIR /app
COPY go.mod ./
RUN go mod download
COPY . .

EXPOSE 8000
CMD ["air", "-c", ".air.toml"]
