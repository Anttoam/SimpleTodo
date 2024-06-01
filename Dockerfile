FROM golang:1.22.2-alpine AS builder
WORKDIR /app
COPY . .
RUN go build -o main cmd/main.go
RUN apk add curl
RUN curl -sSf https://atlasgo.sh | sh

FROM alpine
WORKDIR /app
COPY --from=builder /app/main .
COPY config ./config
COPY migration ./migration

EXPOSE 8080
CMD ["/app/main"]
