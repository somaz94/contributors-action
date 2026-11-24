# Build stage
# Pinned to the BUILD platform, not the target: the image is published
# multi-arch, and letting the toolchain stage run under QEMU emulation to
# produce an arm64 binary is minutes of emulated compilation for no reason. Go
# cross-compiles natively instead, driven by the TARGET* args buildx injects.
FROM --platform=$BUILDPLATFORM golang:1.27-alpine@sha256:70bd7ddea8b6cbe5d6943cde279d6b67feb10cc924fb9100d1d8757c1575b280 AS builder

WORKDIR /build

# BuildKit sets both for the target platform, the host's for a plain
# `docker build` without --platform; the TARGETOS default just restates linux.
#
# BuildKit is required either way — `$BUILDPLATFORM` on the FROM above is a
# BuildKit-only variable, and the legacy builder fails to parse the line rather
# than ignoring it. That is the default builder since Docker 23, so this only
# bites an explicit DOCKER_BUILDKIT=0.
ARG TARGETOS=linux
ARG TARGETARCH


# Copy module manifests first to leverage Docker layer cache for `go mod download`.
# go.sum is optional — this project is currently stdlib-only.
COPY go.mod go.sum* ./
RUN go mod download

COPY . .
RUN CGO_ENABLED=0 GOOS=${TARGETOS} GOARCH=${TARGETARCH} \
    go build \
      -trimpath \
      -ldflags="-s -w" \
      -o /contributors-action ./cmd/main.go

# Runtime stage
FROM alpine:3.24@sha256:f59abe656b692d901e184be3445e084786602108599d2a51d229c0f5a48c71bf

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
