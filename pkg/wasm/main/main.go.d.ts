import { Payload } from '../../../src/core/payload'

declare global {
  function go_draw(parameters: Payload, bucketName: string, env: any): Promise<ReadableStream>
}
