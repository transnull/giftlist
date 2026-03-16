#!/bin/bash
set -e

echo "[1/3] Building frontend..."
cd frontend
npm install
npm run build
cd ..

echo "[2/3] Downloading Go dependencies..."
go mod tidy

echo "[3/3] Building Go binary..."
CGO_ENABLED=1 go build -o giftlist .

echo ""
echo "Build complete! Run with:"
echo "  ADMIN_PASSWORD=your_password ./giftlist"
