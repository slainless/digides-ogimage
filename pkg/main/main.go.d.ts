import { Payload } from '../../src/core/payload'

export interface Parameters extends Pick<Payload, "title" | "subtitle"> { }

declare const __default: {
  draw: (params: Payload) => Promise<any>
}

export default __default