FROM golang:1.22-alpine AS builder

WORKDIR /app

COPY . .

RUN if [ ! -f "go.mod" ]; then \
      go mod init github.com/nilsonvieira/devops-go-genpass; \
    fi

RUN go mod edit -go=1.22 || true
RUN go mod tidy

RUN go build -o genpass .

FROM alpine:latest

WORKDIR /app

COPY --from=builder /app/genpass .

ENTRYPOINT ["/app/genpass"]