type JsonValue = null | boolean | number | string | JsonValue[] | { [key: string]: JsonValue };
type FlatMap = Record<string, JsonValue>;

function isPlainObject(value: unknown): value is Record<string, unknown> {
  return Object.prototype.toString.call(value) === '[object Object]';
}

function flattenLeaves(input: unknown, prefix = '', out: FlatMap = {}): FlatMap {
  if (Array.isArray(input)) {
    out[prefix] = input as JsonValue;
    return out;
  }

  if (!isPlainObject(input)) {
    if (prefix) {
      out[prefix] = (input ?? null) as JsonValue;
    }
    return out;
  }

  for (const [key, value] of Object.entries(input)) {
    const nextPrefix = prefix ? `${prefix}.${key}` : key;
    if (isPlainObject(value)) {
      flattenLeaves(value, nextPrefix, out);
      continue;
    }
    out[nextPrefix] = (value ?? null) as JsonValue;
  }

  return out;
}

function normalize(value: JsonValue): string {
  if (typeof value === 'string') return value.trim();
  return JSON.stringify(value);
}

function isSuspect(baseValue: JsonValue, targetValue: JsonValue): boolean {
  if (typeof baseValue !== 'string' || typeof targetValue !== 'string') return false;
  if (targetValue.includes('\uFFFD')) return true;

  const compact = targetValue.replace(/\s/g, '');
  if (!compact) return false;

  const qCount = [...compact].filter(ch => ch === '?').length;
  if (qCount === 0) return false;

  if (/^\?+$/.test(compact)) return true;
  if (/\?{3,}/.test(targetValue)) return true;
  return qCount >= 2 && qCount / compact.length >= 0.3;
}

export { flattenLeaves, isPlainObject, isSuspect, normalize };
export type { FlatMap, JsonValue };
