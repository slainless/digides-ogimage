import { Hono } from 'hono'

import { randomKeyGenerator } from './handlers/random_key'
import { example } from './handlers/example'
import { main } from './handlers/main'

import '../.generated/wasm_exec.cjs'
import wasm from '../.generated/drawer.wasm'

const go = new Go()
await WebAssembly.instantiate(wasm, go.importObject).then((result) => {
  // @ts-expect-error
  if (result.instance != null) {
    // @ts-expect-error
    go.run(result.instance)
  } else {
    go.run(result)
  }
})

const app = new Hono<Env>()
app.get('/', main())

if (!navigator.userAgent.startsWith("Cloudflare")) {
  app.get("/random-key", randomKeyGenerator())
  app.get("/example", example())
}

export default app
