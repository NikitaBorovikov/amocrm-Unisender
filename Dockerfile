FROM golang:1.24-alpine AS builder

LABEL maintainer="Nikita Borovikov"

RUN apk update && apk add --no-cache git

WORKDIR /app

COPY go.mod go.sum ./

RUN go mod download

COPY . .

RUN CGO_ENABLED=0 GOOS=linux go build -ldflags="-s -w" -o server ./cmd/server/main.go

FROM alpine:latest

RUN apk --no-cache add ca-certificates

WORKDIR /app

COPY --from=builder /app/server .
COPY --from=builder /app/config /app/config
COPY --from=builder /app/.env .

EXPOSE 8083

CMD ["./server"]