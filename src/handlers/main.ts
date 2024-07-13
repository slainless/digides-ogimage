import { Handler } from 'hono'
import { decode } from '../core/payload'
import { cache } from '../core/cache'
import { decryptCacheKey, consumeKey } from '../core/crypto'
import { goKey } from '../core/wasm'

import '../../.generated/wasm_exec.cjs'
import wasm from '../../.generated/drawer.wasm'

const go = new Go()
const instance = await WebAssembly.instantiate(wasm, go.importObject)

export function main(): Handler<Env> {
  return async (c) => {
    cache.get(goKey, () => go.run(instance))
    const decryptionKey = await cache.get(decryptCacheKey, () => consumeKey(c.env.PAYLOAD_ENCRYPTION_SECRET))
    const data = c.req.query("d")

    if (data == null) {
      return new Response("Empty payload", {
        status: 400
      })
    }

    const parameters = await decode(decryptionKey, data)
    const image = await go_draw(parameters, c.env.R2_ASSETS)

    return new Response(image, {
      status: 200,
      headers: {
        'Content-Type': "image/jpeg"
      }
    })
  }
}