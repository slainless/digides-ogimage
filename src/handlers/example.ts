import { Handler } from 'hono'
import { cache } from '../core/cache'
import { cacheKey, consumeKey } from '../core/crypto'
import { Payload, encode } from '../core/payload'

export function example(): Handler<Env> {
  return async (c) => {
    if (navigator.userAgent.startsWith("Cloudflare"))
      return new Response("Not available in Worker!", { status: 501 })

    const key = await cache.get(cacheKey, () => consumeKey(c.env.PAYLOAD_ENCRYPTION_SECRET!))
    const payload: Payload = {
      title: "This is a title",
      subtitle: "This is subtitle",
      background: "uploads/profil/73.11.02.2006/common/35ade0a022c7b566dbffdc934f4cb174.png",
      icon: "uploads/online/73.11.02.2006/common/300_barru.png"
    }

    const encrypted = await encode(key, payload)
    if (c.req.query("redirect") == "false")
      return new Response(encrypted, {
        status: 201,
        headers: {
          "Content-Type": "text/html"
        }
      })

    return new Response(null, {
      status: 302,
      headers: {
        Location: `/?d=${encrypted}`
      }
    })
  }
}  