#!/usr/bin/env bash

mkdir -p dist

if [ ! -f .test/benchstat$(go env GOEXE) ]; then
  echo "Installing benchstat..."
  GOBIN="$(cd .test && pwd)" go install golang.org/x/perf/cmd/benchstat@latest
  echo "Done installing benchstat."
fi

echo "Copying wasm exec environment for testing..."
cp $(go env GOROOT)/misc/wasm/go_js_wasm_exec .test/go_js_wasm_exec
cp $(go env GOROOT)/misc/wasm/wasm_exec.js .test/wasm_exec.cjs
cp $(go env GOROOT)/misc/wasm/wasm_exec_node.js .test/wasm_exec_node.cjs
echo "Done copying wasm exec environment for testing."

echo "Copying wasm exec environment for building..."
cp $(go env GOROOT)/misc/wasm/wasm_exec.js dist/wasm_exec.cjs
echo "Done copying wasm exec environment for building"

echo "Creating esm stub..."
echo "// @ts-nocheck" > .test/wasm_exec_node.js
echo "import './wasm_exec_node.cjs'" >> .test/wasm_exec_node.js
echo "// @ts-nocheck" > .test/wasm_exec.js
echo "import './wasm_exec.cjs'" >> .test/wasm_exec.js
sed -i 's#"./wasm_exec"#"./wasm_exec.cjs"#' .test/wasm_exec_node.cjs
echo "Done creating esm stub..."

echo "Putting example R2 objects..."
./scripts/put_example_r2_objects.sh
echo "Done putting example R2 objects."