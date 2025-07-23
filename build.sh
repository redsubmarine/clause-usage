#!/bin/bash

echo "Building Claude Usage Menu Bar App..."

# Build for macOS
go build -o claude-usage-app .

if [ $? -eq 0 ]; then
    echo "Build successful! Executable: claude-usage-app"
    echo "Run with: ./claude-usage-app"
    echo "Test with: ./claude-usage-app test"
else
    echo "Build failed!"
    exit 1
fi 