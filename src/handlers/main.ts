import { Handler } from 'hono'
import { cache } from '../core/cache'
import { cacheKey, consumeKey } from '../core/crypto'
import { decode } from '../core/payload'
import wasm from 'Go/main/main.go'

export const mainHandler: Handler<Env> = async (c) => {
  const key = await cache.get(cacheKey, () => consumeKey(c.env.PAYLOAD_ENCRYPTION_SECRET))
  const data = c.req.query("d")

  if (data == null) {
    return new Response("Empty payload", {
      status: 400
    })
  }

  const parameters = await decode(key, data)
  const image = await wasm.draw(parameters)

  return new Response("haha", {
    status: 200
  })
}