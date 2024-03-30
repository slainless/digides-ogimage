declare const navigator: import('@cloudflare/workers-types/experimental').Navigator

declare interface Env {
  Bindings: {
    PAYLOAD_ENCRYPTION_SECRET: string
    R2_ASSETS: R2Bucket
    [K: string]: undefined | any
  }
}