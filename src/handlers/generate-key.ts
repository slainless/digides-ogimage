import { Handler } from 'hono'
import { generateKey } from '../core/crypto'

export const generateKeyHandler: Handler = async (c) => {
  if (navigator.userAgent.startsWith("Cloudflare"))
    return new Response("Not available in Worker!", { status: 501 })
  return new Response(await generateKey(), {
    status: 201,
    headers: {
      'Content-Type': "text/html"
    }
  })
}