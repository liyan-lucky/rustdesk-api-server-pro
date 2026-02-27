/* eslint-disable no-underscore-dangle, no-continue, n/prefer-global/process */
import fs from 'node:fs';
import path from 'node:path';
import url from 'node:url';

import locales from '../src/locales/locale';
import { isPlainObject, isSuspect } from './i18n-utils';
import type { JsonValue } from './i18n-utils';

const __dirname = path.dirname(url.fileURLToPath(import.meta.url));
const rootDir = path.resolve(__dirname, '..');
const reportPath = path.join(rootDir, 'docs', 'i18n-sanitize-report.json');

const localeFileMap: Record<string, string> = {
  'ja-JP': 'src/locales/langs/ja-jp.ts',
  'ko-KR': 'src/locales/langs/ko-kr.ts',
  'ru-RU': 'src/locales/langs/ru-ru.ts',
  'fr-FR': 'src/locales/langs/fr-fr.ts',
  'de-DE': 'src/locales/langs/de-de.ts',
  'es-ES': 'src/locales/langs/es-es.ts',
  'zh-CN': 'src/locales/langs/zh-cn.ts'
};

function parseArgs() {
  const args = process.argv.slice(2);
  const dryRun = args.includes('--dry-run');
  const failOnDetect = args.includes('--fail-on-detect');
  const json = args.includes('--json');
  const localeArg = args.find(arg => arg.startsWith('--locale='));
  const localesArg = localeArg
    ? localeArg
        .replace('--locale=', '')
        .split(',')
        .map(v => v.trim())
        .filter(Boolean)
    : [];
  return {
    dryRun,
    failOnDetect,
    json,
    targetLocales: localesArg.length ? localesArg : ['ja-JP', 'ko-KR', 'ru-RU']
  };
}

function sanitizeAgainstBase(target: unknown, base: unknown): [unknown, number] {
  if (typeof target === 'string') {
    if (typeof base === 'string' && isSuspect(base, target)) {
      return [base, 1];
    }
    return [target, 0];
  }

  if (Array.isArray(target)) {
    let count = 0;
    const baseArr = Array.isArray(base) ? base : [];
    const next = target.map((item, idx) => {
      const [sanitized, changed] = sanitizeAgainstBase(item, baseArr[idx]);
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
      const [sanitized, changed] = sanitizeAgainstBase(value, baseObj[key]);
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

function ensureDir(filePath: string) {
  fs.mkdirSync(path.dirname(filePath), { recursive: true });
}

function main() {
  const { dryRun, failOnDetect, json, targetLocales } = parseArgs();
  const baseLocale = locales['en-US'];
  const report: Array<{ locale: string; file: string; replaced: number; applied: boolean }> = [];
  let totalReplaced = 0;

  for (const locale of targetLocales) {
    const file = localeFileMap[locale];
    if (!file) {
      console.warn(`skip unsupported locale: ${locale}`);
      continue;
    }
    const source = locales[locale as keyof typeof locales];
    if (!source) {
      console.warn(`skip missing locale in runtime map: ${locale}`);
      continue;
    }

    const [sanitized, changed] = sanitizeAgainstBase(source, baseLocale);
    const filePath = path.join(rootDir, file);
    const nextText = toTsModule(sanitized);
    const prev = fs.existsSync(filePath) ? fs.readFileSync(filePath, 'utf8') : '';
    const changedText = prev !== nextText;

    if (!dryRun && changedText) {
      fs.writeFileSync(filePath, nextText, 'utf8');
    }

    report.push({
      locale,
      file,
      replaced: changed,
      applied: !dryRun && changedText
    });
    totalReplaced += changed;
    if (!json) {
      console.log(
        `${locale}: replaced=${changed}, fileChanged=${changedText ? 'yes' : 'no'}, mode=${dryRun ? 'dry-run' : 'write'}`
      );
    }
  }

  ensureDir(reportPath);
  const result = { dryRun, failOnDetect, targetLocales, totalReplaced, report };
  fs.writeFileSync(reportPath, JSON.stringify(result, null, 2), 'utf8');

  if (json) {
    console.log(JSON.stringify(result, null, 2));
  }

  if (failOnDetect && totalReplaced > 0) {
    console.error(`sanitize guard failed: detected ${totalReplaced} suspicious translation(s)`);
    process.exit(1);
  }
}

main();
