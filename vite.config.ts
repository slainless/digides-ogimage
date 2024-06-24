import { ResolvedConfig, defineConfig } from 'vite'
import devServer, { defaultOptions } from '@hono/vite-dev-server'
import cloudflareAdapter from '@hono/vite-dev-server/cloudflare'
import { readFile } from 'node:fs/promises'

export default defineConfig({
  plugins: [
    devServer({
      entry: "./src/index.ts",
      adapter: cloudflareAdapter,
      exclude: [/.*\.wasm$/, ...defaultOptions.exclude],
    }),
    (function wasmPlugin() {
      let cfg: ResolvedConfig
      return {
        name: "wasm",
        configResolved(config) {
          cfg = config
        },
        async load(this, id) {
          if (!id.endsWith(".wasm")) return
          return ""
        },
        async transform(this, code, id, options) {
          if (!id.endsWith(".wasm")) return

          let content: string
          const data = await readFile(id)

          if (cfg.command == 'serve') {
            content = `"data:application/wasm;base64,` + Buffer.from(data).toString("base64") + `"`
          } else {
            const emittedFile = this.emitFile({
              type: "asset",
              source: data
            })
            content = `import.meta.ROLLUP_FILE_URL_` + emittedFile
          }

          return `const wasm = await fetch(${content}).then(r => r.arrayBuffer());` +
            `export default wasm`
        },
      }
    })(),
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