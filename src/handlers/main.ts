import { Handler } from 'hono'
import { decode } from '../core/payload'
import { cache } from '../core/cache'
import { cacheKey, consumeKey } from '../core/crypto'

export function main(): Handler<Env> {
  return async (c) => {
    const decryptionKey = await cache.get(cacheKey, () => consumeKey(c.env.PAYLOAD_ENCRYPTION_SECRET))
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