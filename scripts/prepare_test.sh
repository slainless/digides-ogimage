#!/usr/bin/env bash

mkdir -p dist

echo "Installing benchstat..."
go install golang.org/x/perf/cmd/benchstat@latest
echo "Done installing benchstat."

echo "Copying wasm exec environment..."
cp $(go env GOROOT)/misc/wasm/go_js_wasm_exec .test/go_js_wasm_exec
cp $(go env GOROOT)/misc/wasm/wasm_exec.js .test/wasm_exec.cjs
cp $(go env GOROOT)/misc/wasm/wasm_exec_node.js .test/wasm_exec_node.cjs
echo "Done copying wasm exec environment."

echo "Creating esm stub..."
echo "// @ts-nocheck" > .test/wasm_exec_node.js
echo "import './wasm_exec_node.cjs'" >> .test/wasm_exec_node.js
echo "// @ts-nocheck" > .test/wasm_exec.js
echo "import './wasm_exec.cjs'" >> .test/wasm_exec.js
sed -i 's#"./wasm_exec"#"./wasm_exec.cjs"#' .test/wasm_exec_node.cjs
echo "Done creating esm stub..."