#!/bin/bash
# Build script for Render

echo "Installing dependencies..."
go mod download

echo "Building application..."
go build -o bin/app main.go

echo "Build completed!"
