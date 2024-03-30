export function debug(_?: boolean) {
  // @ts-expect-error
  globalThis['__DEBUG'] = !!_
}

export function isDebug() {
  // @ts-expect-error
  return !!globalThis['__DEBUG']
}

const noop = () => { }
export const logger = new Proxy(console, {
  get(target, p, receiver) {
    // @ts-expect-error
    if (typeof target[p] == 'function' && !isDebug()) return noop
    // @ts-expect-error
    return target[p]
  },
})