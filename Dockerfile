FROM golang:1.25-alpine AS builder

RUN apk update && apk add --no-cache \
    postgresql-client \
    build-base \
    git

WORKDIR /app

ENV GOBIN /usr/local/bin

COPY go.mod go.sum ./
RUN go mod download

COPY . .

RUN go build -o /app/main ./cmd/app/main.go

FROM alpine:3.22.1

RUN apk --no-cache add postgresql-client ca-certificates

WORKDIR /root/

COPY --from=builder /app/main /usr/local/bin/main

RUN chmod +x /usr/local/bin/main

EXPOSE 8080

ENTRYPOINT ["/main"]