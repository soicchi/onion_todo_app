FROM golang:1.23.2 as local
WORKDIR /app
RUN --mount=type=cache,target=/go/pkg/mod \
    --mount=type=bind,source=app/go.mod,target=/app/go.mod \
    --mount=type=bind,source=app/go.sum,target=/app/go.sum \
    go mod download

EXPOSE 8080
CMD ["go", "run", "cmd/api/main.go"]
