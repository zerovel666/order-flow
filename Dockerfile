# ---------- Stage 1: сборка ----------
FROM golang:1.27.0-alpine AS builder

WORKDIR /src

# Сначала зависимости — этот слой кэшируется, пока go.mod/go.sum не меняются
COPY go.mod ./
RUN --mount=type=cache,target=/go/pkg/mod \
    go mod download

COPY . .

# Статический бинарник без CGO, уменьшенный размер (-s -w)
RUN --mount=type=cache,target=/go/pkg/mod \
    --mount=type=cache,target=/root/.cache/go-build \
    CGO_ENABLED=0 GOOS=linux \
    go build -ldflags="-s -w" -o /bin/app ./cmd/app

# ---------- Stage 2: runtime ----------
FROM gcr.io/distroless/static-debian12:nonroot

COPY --from=builder /bin/app /app

EXPOSE 8080
USER nonroot:nonroot

ENTRYPOINT ["/app"]