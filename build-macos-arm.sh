#!/bin/bash
# Build script for ledger-logger (macOS Apple Silicon)
set -e

GOARCH=arm64 GOOS=darwin go build -o ledger-logger-macos-arm ledger-logger.go

echo "Build complete: ./ledger-logger-macos-arm"
