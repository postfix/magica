#!/bin/bash

set -e  # Exit immediately if a command exits with a non-zero status

echo "🔧 Setting up environment variables..."
export MAGIKA_ASSETS_DIR=$(pwd)/assets
export MAGIKA_MODEL=standard_v2_1
export LD_LIBRARY_PATH=/opt/onnxruntime/lib

# Ensure ONNX Runtime is installed
if [ ! -d "/opt/onnxruntime" ]; then
    echo "❌ ONNX Runtime not found in /opt/onnxruntime. Please install it before proceeding."
    exit 1
fi

# Set CGO flags explicitly
export CGO_ENABLED=1
export CGO_CFLAGS="-I/opt/onnxruntime/include"
export CGO_LDFLAGS="-L/opt/onnxruntime/lib"

echo "🚀 Running Magika unit tests..."
go test -tags onnxruntime -ldflags="-linkmode=external -extldflags=-L$LD_LIBRARY_PATH" ./...

echo "✅ Unit tests passed!"

echo "🛠 Compiling Magika CLI..."
cd cli
go build -tags onnxruntime -ldflags="-linkmode=external -extldflags=-L$LD_LIBRARY_PATH" .

echo "🎉 Magika CLI built successfully!"
