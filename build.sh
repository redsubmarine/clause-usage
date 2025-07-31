#!/bin/bash

echo "Building Claude Usage Menu Bar App..."

# Create release directory
mkdir -p release

# Build for macOS (Apple Silicon) - current architecture
echo "Building for macOS (Apple Silicon)..."
GOOS=darwin GOARCH=arm64 CGO_ENABLED=1 go build -o release/claude-usage-macos-arm64 .

# Build for local development (default)
echo "Building for local development..."
go build -o claude-usage-app .

if [ $? -eq 0 ]; then
    echo "Build successful!"
    echo "Release binaries created in ./release/"
    echo "Local executable: claude-usage-app"
    echo "Run with: ./claude-usage-app"
    echo "Test with: ./claude-usage-app test"
    echo ""
    echo "Note: Intel macOS build requires cross-compilation setup with CGO"
else
    echo "Build failed!"
    exit 1
fi 