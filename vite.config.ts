import { defineConfig } from 'vite'
import goWasm from 'vite-plugin-golang-wasm'
import devServer from '@hono/vite-dev-server'
import tsconfigPaths from 'vite-tsconfig-paths'
import { vitePluginViteNodeMiniflare } from "@hiogawa/vite-node-miniflare"
import type { WorkerOptions } from 'miniflare'

export default defineConfig({
  plugins: [
    goWasm({
      wasmExecPath: "./src/wasm/wasm_exec.js"
    }),
    tsconfigPaths(),
    vitePluginViteNodeMiniflare({
      entry: "./src/index.ts",
      hmr: true,
      miniflareOptions(options) {
        const opts = options as WorkerOptions
        opts.compatibilityFlags = ["nodejs_compat"]
        opts.r2Buckets = ["R2_ASSETS"]
        opts.bindings!.PAYLOAD_ENCRYPTION_SECRET = "ThlIVaRD1kwBGpdeyutAx6yKwb4ZKuIMqqhrNoeA9X0"
      },
    }),
  ],
  build: {
    manifest: true,
    rollupOptions: {
      input: "./src/index.ts",
    },
    ssr: true,
    ssrEmitAssets: true,
  },
  ssr: {
    target: "webworker",
    noExternal: ["node:buffer"],
  },
})