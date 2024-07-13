import { Handler } from 'hono'
import { decode } from '../core/payload'
import { cache } from '../core/cache'
import { decryptCacheKey, consumeKey } from '../core/crypto'
import { goKey } from '../core/wasm'

import '../../.generated/wasm_exec.cjs'
import wasm from '../../.generated/drawer.wasm'
import { Payload } from '../schema/generated/payload'
import { TypeGuardError } from 'typia'

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

    let parameters: Payload
    try {
      parameters = await decode(decryptionKey, data)
      const image = await godrawer.draw(parameters, c.env.R2_ASSETS)
      return new Response(image, {
        status: 200,
        headers: {
          'Content-Type': "image/jpeg"
        }
      })
    } catch (e) {
      if (e instanceof TypeGuardError) {
        return new Response(e.message, { status: 400 })
      }

      switch (e) {
        case godrawer.errors.ErrBucketNotFound:
          return new Response("Misconfigured: no bucket found", { status: 500 })
        case godrawer.errors.ErrInvalidCloudflareEnv:
          return new Response("Misconfigured: invalid cloudflare env", { status: 500 })
        case godrawer.errors.ErrFileNotFound:
          return new Response("Either logo or background not found:\n" + parameters!.icon + "\n" + parameters!.background, { status: 400 })
        case godrawer.errors.ErrInvalidReadingResult:
        case godrawer.errors.ErrInvalidStream:
          return new Response(e.toString(), { status: 500 })
        case godrawer.errors.ErrParametersInvalid:
        case godrawer.errors.ErrParametersInvalidField:
          return new Response(e.toString(), { status: 400 })
      }

      return new Response(e?.toString() ?? JSON.stringify(e), { status: 500 })
    }
  }
}