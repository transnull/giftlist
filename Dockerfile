# ── Stage 1: Build frontend ──────────────────────────────────────────────────
FROM node:20-alpine AS frontend-builder

WORKDIR /app/frontend
COPY frontend/package*.json ./
RUN npm ci --silent
COPY frontend/ ./
RUN npm run build

# ── Stage 2: Build Go binary ──────────────────────────────────────────────────
FROM golang:1.22-alpine AS go-builder

# CGO requires gcc
RUN apk add --no-cache gcc musl-dev

WORKDIR /app
COPY go.mod go.sum ./
RUN go mod download

COPY . .
# Copy compiled frontend into static/ (overwrite any placeholder)
COPY --from=frontend-builder /app/static ./static

RUN CGO_ENABLED=1 GOOS=linux go build -ldflags="-s -w" -o giftlist .

# ── Stage 3: Minimal runtime image ────────────────────────────────────────────
FROM alpine:3.19

RUN apk add --no-cache ca-certificates tzdata \
    && cp /usr/share/zoneinfo/Asia/Shanghai /etc/localtime \
    && echo "Asia/Shanghai" > /etc/timezone

WORKDIR /app

COPY --from=go-builder /app/giftlist .

# Data directory for SQLite
RUN mkdir -p /data

ENV PORT=8080 \
    DB_PATH=/data/giftlist.db \
    ADMIN_PASSWORD=admin123

EXPOSE 8080

VOLUME ["/data"]

ENTRYPOINT ["./giftlist"]
