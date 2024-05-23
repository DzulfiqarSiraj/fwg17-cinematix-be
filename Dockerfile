# Build Stage
FROM golang:1.22.3-alpine3.20 AS builder
WORKDIR /app
COPY . .
RUN go mod download
RUN go build -o ./cinematix-app ./main.go

# Run Stage
FROM alpine:latest
WORKDIR /app
COPY --from=builder /app/cinematix-app .
COPY .env .
EXPOSE 8080
ENTRYPOINT [ "./cinematix-app" ]