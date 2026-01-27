#!/bin/bash
# Build script for Render

set -e

echo "Installing dependencies..."
go mod download

echo "Building application..."
mkdir -p bin
go build -o bin/app .

echo "Build completed!"
ls -la bin/
