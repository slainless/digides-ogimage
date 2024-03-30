import { Handler } from 'hono'
import { cache } from '../core/cache'
import { cacheKey, consumeKey } from '../core/crypto'
import { Payload, encode } from '../core/payload'

export const exampleHandler: Handler = async (c) => {
  if (navigator.userAgent.startsWith("Cloudflare"))
    return new Response("Not available in Worker!", { status: 501 })
  const key = await cache.get(cacheKey, () => consumeKey(c.env.PAYLOAD_ENCRYPTION_SECRET!))
  const payload: Payload = {
    title: "This is a title",
    subtitle: "This is subtitle",
    background: "uploads/background.jpg",
    icon: "uploads/icon.jpg"
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