import { Hono } from 'hono'

import { randomKeyGenerator } from './handlers/random_key'
import { example } from './handlers/example'
import { main } from './handlers/main'

const app = new Hono<Env>()
app.get('/', main())

if (!navigator.userAgent.startsWith("Cloudflare")) {
  app.get("/random-key", randomKeyGenerator())
  app.get("/example", example())
}

export default app
