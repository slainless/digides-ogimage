import { Buffer } from 'node:buffer'
import { logger } from './debug'
import { Static, Type } from '@sinclair/typebox'
import { Value } from '@sinclair/typebox/value'
// import { TypeCompiler } from '@sinclair/typebox/compiler'

export const ErrorInvalidPayload = new Error("Invalid payload")

export const PayloadSchema = Type.Object({
  title: Type.String(),
  subtitle: Type.String(),
  icon: Type.String(),
  background: Type.String(),
  titleFont: Type.Optional(Type.String()),
  subtitleFont: Type.Optional(Type.String()),
})

// export const CompiledPayloadSchema = TypeCompiler.Compile(PayloadSchema)

export type Payload = Static<typeof PayloadSchema>

export async function decode(key: CryptoKey, data: string): Promise<Payload> {
  const arr = Buffer.from(data, "base64url")
  const iv = arr.subarray(0, 12)
  const encrypted = arr.subarray(12)

  const decrypted = await crypto.subtle.decrypt({
    name: "AES-GCM",
    iv
  }, key, encrypted)

  const parsed = JSON.parse(Buffer.from(decrypted).toString("utf8"))
  // if (Value.Check(PayloadSchema, parsed)) {
  //   throw ErrorInvalidPayload
  // }
  return parsed as Payload
}

export async function encode(key: CryptoKey, payload: Payload): Promise<string> {
  const data = Buffer.from(JSON.stringify(payload))

  const iv = Buffer.allocUnsafe(12)
  crypto.getRandomValues(iv)

  const encrypted = await crypto.subtle.encrypt({
    name: "AES-GCM",
    iv
  }, key, data)

  const result = Buffer.allocUnsafe(12 + encrypted.byteLength)
  result.set(iv, 0)
  result.set(new Uint8Array(encrypted), 12)

  return result.toString("base64url")
}