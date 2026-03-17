# Build stage
FROM golang:1.26-alpine AS builder

WORKDIR /build

COPY go.mod ./
COPY . .
RUN CGO_ENABLED=0 GOOS=linux go build -o /contributors-action ./cmd/main.go

# Runtime stage
FROM alpine:3.23

RUN apk add --no-cache git ca-certificates

COPY --from=builder /contributors-action /usr/local/bin/contributors-action

ENTRYPOINT ["contributors-action"]
