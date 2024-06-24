#!/usr/bin/env bash

mkdir -p dist

echo "Building wasm..."
GOOS=js GOARCH=wasm go build -o dist/drawer.wasm ./pkg/wasm/main/main.go
echo "Done building wasm."

echo "Copying wasm_exec.js..."
cp $(go env GOROOT)/misc/wasm/wasm_exec.js dist/wasm_exec.cjs
echo "Done copying wasm_exec.js."