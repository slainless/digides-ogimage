#!/usr/bin/env bash

export GOOS=js
export GOARCH=wasm

bun vitest bench --run --outputFile .test/benchmark_e2e.json --reporter=json
go test -benchmem -run=^$ -bench ^BenchmarkDirectDraw$ github.com/slainless/digides-ogimage/pkg/draw \
  -exec="$(go env GOROOT)/misc/wasm/go_js_wasm_exec" \
  -count=10 | tee .test/benchmark_direct_draw.txt
go test -benchmem -run=^$ -bench ^BenchmarkWASMDraw$ github.com/slainless/digides-ogimage/pkg/wasm \
  -exec="$(go env GOROOT)/misc/wasm/go_js_wasm_exec" \
  -count=10 | tee .test/benchmark_wasm_draw.txt