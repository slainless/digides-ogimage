#!/usr/bin/env bash

mkdir -p dist

echo "Building wasm..."
GOOS=js GOARCH=wasm go build -o .generated/drawer.wasm ./pkg/wasm/main/main.go
if [ $? -eq 0 ]; then
  echo "Done building wasm."
else
  echo "Failed building wasm."
  exit $?
fi