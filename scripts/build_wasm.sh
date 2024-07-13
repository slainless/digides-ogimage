#!/usr/bin/env bash

mkdir -p dist

echo "Building wasm..."
GOOS=js GOARCH=wasm go build -o .generated/drawer.wasm ./pkg/wasm/main/main.go
echo "Done building wasm."
