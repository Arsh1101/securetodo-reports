FROM golang:1.22-alpine AS builder

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY . .

RUN go build -o securetodo ./cmd/web

FROM alpine:3.20

WORKDIR /app

COPY --from=builder /app/securetodo .
COPY static ./static
COPY templates ./templates

RUN mkdir -p data reports

EXPOSE 8080

CMD ["./securetodo"]
