FROM golang:1.26.4-alpine AS builder

WORKDIR /build

COPY go.mod go.sum ./
RUN go mod download

COPY . .

RUN CGO_ENABLED=0 GOOS=linux go build -ldflags="-s -w" -o terminal-app .

FROM alpine:latest

WORKDIR /app

COPY --from=builder /build/terminal-app .

EXPOSE 22

CMD ["/app/terminal-app", "--prod", "--port", "22"]
