import { R2Bucket } from '@cloudflare/workers-types'
import { Payload } from '../../../src/schema/generated/payload'

declare global {
  namespace godrawer {
    function draw(parameters: Payload, bucket: R2Bucket, quality: number): Promise<ReadableStream>

    namespace errors {
      const ErrInvalidStream: TypeError
      const ErrInvalidReadingResult: TypeError
      const ErrFileNotFound: Error
      const ErrBucketNotFound: TypeError
      const ErrInvalidCloudflareEnv: TypeError
      const ErrParametersInvalid: TypeError
      const ErrParametersInvalidField: TypeError
    }
  }
}
