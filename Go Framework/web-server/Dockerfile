# ── Stage 1: Build ──────────────────────────────────────────────────────────
FROM golang:1.25-alpine AS builder

# Install git and swag (needed by some go modules + swagger doc generation)
RUN apk add --no-cache git && \
    go install github.com/swaggo/swag/cmd/swag@latest

WORKDIR /app

# Cache dependency downloads separately from source
COPY go.mod go.sum ./
RUN go mod download

# Copy the rest of the source
COPY . .

# Generate swagger docs
RUN swag init

# Build a statically linked binary
RUN CGO_ENABLED=0 GOOS=linux go build -o web-server ./main.go

# ── Stage 2: Run ────────────────────────────────────────────────────────────
FROM alpine:3.20

# Add CA certificates for HTTPS calls and timezone data
RUN apk add --no-cache ca-certificates tzdata

WORKDIR /app

# Copy only the binary and migration files from the builder
COPY --from=builder /app/web-server .
COPY --from=builder /app/database/migration ./database/migration

# The app reads .env at runtime; it will be mounted via docker-compose
EXPOSE 8000

CMD ["./web-server"]
