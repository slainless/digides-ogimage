import { Handler } from 'hono'
import { cache } from '../core/cache'
import { cacheKey, consumeKey } from '../core/crypto'
import { decode } from '../core/payload'
import wasm from 'Go/main/main.go'
import { Buffer } from 'node:buffer'

export const mainHandler: Handler<Env> = async (c) => {
  const key = await cache.get(cacheKey, () => consumeKey(c.env.PAYLOAD_ENCRYPTION_SECRET))
  const data = c.req.query("d")

  if (data == null) {
    return new Response("Empty payload", {
      status: 400
    })
  }

  const parameters = await decode(key, data)
  const [
    icon,
    background
  ] = await Promise.all([
    c.env.R2_ASSETS.get(parameters.icon).then(v => v?.arrayBuffer()),
    c.env.R2_ASSETS.get(parameters.background).then(v => v?.arrayBuffer()),
  ])

  if (icon == null || background == null)
    return new Response("No icon or background found", {
      status: 400
    })

  const image = await wasm.draw({
    ...parameters,
    icon: Buffer.from(icon).toString("base64"),
    background: Buffer.from(background).toString("base64")
  })
  return new Response(Buffer.from(image, "base64"), {
    status: 200,
    headers: {
      'Content-Type': "image/jpeg"
    }
  })
}