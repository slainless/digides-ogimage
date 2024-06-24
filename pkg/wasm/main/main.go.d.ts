import { R2Bucket } from '@cloudflare/workers-types'
import { Payload } from '../../../src/core/payload'

declare global {
  function go_draw(parameters: Payload, bucket: R2Bucket): Promise<ReadableStream>
}
