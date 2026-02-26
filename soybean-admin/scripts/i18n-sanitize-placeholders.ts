import fs from 'node:fs';
import path from 'node:path';
import url from 'node:url';

import locales from '../src/locales/locale';

type JsonValue = null | boolean | number | string | JsonValue[] | { [key: string]: JsonValue };

const __dirname = path.dirname(url.fileURLToPath(import.meta.url));
const rootDir = path.resolve(__dirname, '..');

const targets = [
  { locale: 'ja-JP', file: 'src/locales/langs/ja-jp.ts' },
  { locale: 'ko-KR', file: 'src/locales/langs/ko-kr.ts' },
  { locale: 'ru-RU', file: 'src/locales/langs/ru-ru.ts' }
] as const;

function isPlainObject(value: unknown): value is Record<string, unknown> {
  return Object.prototype.toString.call(value) === '[object Object]';
}

function isSuspiciousString(value: string): boolean {
  if (!value) return false;
  if (value.includes('�')) return true;
  if (value.includes('?')) return true;
  if (/\?{3,}/.test(value)) return true;

  const compact = value.replace(/\s/g, '');
  if (compact.length < 3) return false;

  const questionCount = [...compact].filter(ch => ch === '?').length;
  return questionCount >= 2 && questionCount / compact.length >= 0.3;
}

function sanitizeAgainstBase(target: unknown, base: unknown, pathParts: string[] = []): [unknown, number] {
  if (typeof target === 'string') {
    if (isSuspiciousString(target) && typeof base === 'string') {
      return [base, 1];
    }
    return [target, 0];
  }

  if (Array.isArray(target)) {
    let count = 0;
    const baseArr = Array.isArray(base) ? base : [];
    const next = target.map((item, idx) => {
      const [sanitized, changed] = sanitizeAgainstBase(item, baseArr[idx], [...pathParts, String(idx)]);
      count += changed;
      return sanitized;
    });
    return [next, count];
  }

  if (isPlainObject(target)) {
    let count = 0;
    const baseObj = isPlainObject(base) ? base : {};
    const out: Record<string, JsonValue> = {};
    for (const [key, value] of Object.entries(target)) {
      const [sanitized, changed] = sanitizeAgainstBase(value, baseObj[key], [...pathParts, key]);
      out[key] = sanitized as JsonValue;
      count += changed;
    }
    return [out, count];
  }

  return [target, 0];
}

function toTsModule(obj: unknown): string {
  return `const local: App.I18n.Schema = ${JSON.stringify(obj, null, 2)};\n\nexport default local;\n`;
}

function main() {
  const baseLocale = locales['en-US'];

  for (const target of targets) {
    const source = locales[target.locale];
    const [sanitized, changed] = sanitizeAgainstBase(source, baseLocale);
    const filePath = path.join(rootDir, target.file);
    fs.writeFileSync(filePath, toTsModule(sanitized), 'utf8');
    console.log(`${target.locale}: replaced ${changed} suspicious string(s) -> ${path.relative(rootDir, filePath)}`);
  }
}

main();
