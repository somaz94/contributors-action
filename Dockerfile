# Build stage
# Pinned to the BUILD platform, not the target: the image is published
# multi-arch, and letting the toolchain stage run under QEMU emulation to
# produce an arm64 binary is minutes of emulated compilation for no reason. Go
# cross-compiles natively instead, driven by the TARGET* args buildx injects.
FROM --platform=$BUILDPLATFORM golang:1.26-alpine@sha256:0178a641fbb4858c5f1b48e34bdaabe0350a330a1b1149aabd498d0699ff5fb2 AS builder

WORKDIR /build

# Supplied per target platform. The defaults keep a plain `docker build` (a
# local one-off, no --platform) working: TARGETARCH resolves empty there, and
# an empty GOARCH means "host default", which is what that build wants anyway.
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
FROM alpine:3.24@sha256:28bd5fe8b56d1bd048e5babf5b10710ebe0bae67db86916198a6eec434943f8b

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
