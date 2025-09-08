#!/bin/bash
# Build script for ledger-logger
set -e

go build -o ledger-logger ledger-logger.go

echo "Build complete: ./ledger-logger"
