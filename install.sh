#!/bin/bash

set -e

echo "🚀 Installing Claude Usage Menu Bar App..."

# Detect architecture
ARCH=$(uname -m)
if [[ "$ARCH" == "arm64" ]]; then
    BINARY_URL="https://github.com/redsubmarine/clause-usage/releases/download/v1.0.0/claude-usage-macos-arm64"
    echo "✅ Detected Apple Silicon Mac"
else
    echo "❌ This installer currently supports Apple Silicon Macs only"
    echo "Please download the appropriate binary from: https://github.com/redsubmarine/clause-usage/releases"
    exit 1
fi

# Download and install
echo "📥 Downloading claude-usage..."
curl -L "$BINARY_URL" -o /tmp/claude-usage

echo "🔐 Making executable..."
chmod +x /tmp/claude-usage

echo "📁 Installing to /usr/local/bin..."
sudo mv /tmp/claude-usage /usr/local/bin/claude-usage

echo "🎉 Installation complete!"
echo ""
echo "Run with: claude-usage"
echo "Test with: claude-usage test"
echo ""
echo "⚠️  Make sure you have 'ccusage' CLI installed first!"