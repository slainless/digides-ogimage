import { defineConfig } from 'vite'
import goWasm from 'vite-plugin-golang-wasm'
import devServer from '@hono/vite-dev-server'
import cloudflareAdapter from '@hono/vite-dev-server/cloudflare'
import tsconfigPaths from 'vite-tsconfig-paths'
import type { WorkerOptions } from 'miniflare'

export default defineConfig({
  plugins: [
    goWasm(),
    tsconfigPaths(),
    devServer({
      entry: "./src/index.ts",
      adapter: cloudflareAdapter
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