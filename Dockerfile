# Build stage
FROM golang:1.21-alpine3.20 AS builder
WORKDIR /app
COPY .  .
RUN go build -o main main.go

# Run stage
FROM alpine:3.20
WORKDIR /app
COPY --from=builder /app/main .
EXPOSE 9090
CMD [ "/app/main" ]