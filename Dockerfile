FROM golang:1.25-alpine AS builder
WORKDIR /src

COPY go.mod go.sum ./
RUN go mod download

COPY . .
RUN CGO_ENABLED=0 GOOS=linux go build -o /bin/atletismo-api ./cmd/atletismo-api

FROM alpine:3.20

RUN apk add --no-cache ca-certificates tzdata
RUN adduser -D -u 10001 appuser

WORKDIR /app
COPY --from=builder /bin/atletismo-api /app/atletismo-api

RUN mkdir -p /app/data && chown appuser:appuser /app/data

USER appuser
EXPOSE 8080

ENTRYPOINT ["/app/atletismo-api"]