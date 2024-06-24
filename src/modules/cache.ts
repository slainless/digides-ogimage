
export class WarmCache {
  private weakMap = new WeakMap<any, any>;
  private map = new Map<any, any>

  private decideMap(key: any) {
    if (typeof key == "string" || typeof key == "number" || typeof key == "boolean" || typeof key == 'undefined' || key === null)
      return this.map
    return this.weakMap
  }

  get<T>(key: any, defaultValue?: () => T): T {
    const map = this.decideMap(key)
    if (defaultValue == null || map.has(key)) return map.get(key)
    const value = defaultValue()
    map.set(key, value)
    return value
  }

  set(key: any, value: any) {
    const map = this.decideMap(key)
    map.set(key, value)
    return this
  }
}
