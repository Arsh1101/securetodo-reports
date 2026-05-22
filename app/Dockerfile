FROM golang:1.24-alpine AS builder

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY . .

RUN go build -o securetodo ./cmd/web

FROM alpine:3.20

WORKDIR /app

RUN addgroup -S appgroup && adduser -S appuser -G appgroup

COPY --from=builder /app/securetodo .
COPY static ./static
COPY templates ./templates

RUN mkdir -p data reports && chown -R appuser:appgroup /app

USER appuser

EXPOSE 8080

CMD ["./securetodo"]