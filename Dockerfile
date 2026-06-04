# Build stage
FROM golang:1.26-alpine@sha256:f23e8b227fb4493eabe03bede4d5a32d04092da71962f1fb79b5f7d1e6c2a17f AS builder

WORKDIR /build

# Copy module manifests first to leverage Docker layer cache for `go mod download`.
# go.sum is optional — this project is currently stdlib-only.
COPY go.mod go.sum* ./
RUN go mod download

COPY . .
RUN CGO_ENABLED=0 GOOS=linux go build \
      -trimpath \
      -ldflags="-s -w" \
      -o /contributors-action ./cmd/main.go

# Runtime stage
FROM alpine:3.23@sha256:5b10f432ef3da1b8d4c7eb6c487f2f5a8f096bc91145e68878dd4a5019afde11

ARG VERSION=dev
ARG REVISION=unknown
ARG CREATED=unknown

LABEL org.opencontainers.image.source="https://github.com/somaz94/contributors-action" \
      org.opencontainers.image.description="Generate and update contributors list from GitHub repository data" \
      org.opencontainers.image.licenses="MIT" \
      org.opencontainers.image.version="${VERSION}" \
      org.opencontainers.image.revision="${REVISION}" \
      org.opencontainers.image.created="${CREATED}"

RUN apk add --no-cache git ca-certificates

COPY --from=builder /contributors-action /usr/local/bin/contributors-action

# Intentionally runs as root: GitHub Actions bind-mounts $GITHUB_WORKSPACE
# with the runner's uid; writing CONTRIBUTORS.md back to the workspace
# requires write access to that mount.
ENTRYPOINT ["contributors-action"]
