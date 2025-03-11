FROM golang:1.21-alpine AS builder

WORKDIR /app

COPY . .

RUN go mod init github.com/nilsonvieira/devops-go-genpass || true
RUN go build -o genpass .

FROM alpine:latest

WORKDIR /app

COPY --from=builder /app/genpass .

ENTRYPOINT ["/app/genpass"]