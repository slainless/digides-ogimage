import { Hono } from 'hono'

import { generateKeyHandler } from './handlers/generate-key'
import { exampleHandler } from './handlers/example'
import { mainHandler } from './handlers/main'

const app = new Hono<Env>()

app.get('/', mainHandler)
app.get("/generate-key", generateKeyHandler)
app.get("/example", exampleHandler)

export default app
