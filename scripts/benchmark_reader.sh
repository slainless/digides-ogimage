#!/usr/bin/env bash

export GOOS=js
export GOARCH=wasm

go test -benchmem -run=^$ -bench ^BenchmarkDefaultReader512$ github.com/slainless/digides-ogimage/pkg/bridge \
  -exec="$(go env GOROOT)/misc/wasm/go_js_wasm_exec" \
  -count=10 | tee .test/benchmark_go_reader_512.txt
go test -benchmem -run=^$ -bench ^BenchmarkNativeDefaultReader512$ github.com/slainless/digides-ogimage/pkg/bridge \
  -exec="$(go env GOROOT)/misc/wasm/go_js_wasm_exec" \
  -count=10 | tee .test/benchmark_native_reader_512.txt
benchstat .test/benchmark_native_reader_512.txt .test/benchmark_go_reader_512.txt