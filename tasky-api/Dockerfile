FROM golang:1.25-alpine AS builder

RUN apk add --no-cache build-base

WORKDIR /app

COPY go.mod go.sum ./

RUN go mod download

COPY . .


# -ldflags "-linkmode external -extldflags '-static'" is required for a statically linked binary for scratch
RUN CGO_ENABLED=1 GOOS=linux go build -ldflags "-linkmode external -extldflags '-static'" -o taskyapi cmd/taskyapi/main.go

# Create /tmp directory for SQLite to store temporary files in the scratch image
RUN mkdir -p /tmp_dir

FROM scratch

WORKDIR /

COPY --from=builder /app/taskyapi .

COPY --from=builder /tmp_dir /tmp

EXPOSE 3000

# Set environment variables
ENV DB_PATH=/data/tasks.db

ENTRYPOINT ["./taskyapi"]
