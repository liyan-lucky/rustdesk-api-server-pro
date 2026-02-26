import fs from 'node:fs';
import path from 'node:path';
import url from 'node:url';

import locales from '../src/locales/locale';

type JsonValue = null | boolean | number | string | JsonValue[] | { [key: string]: JsonValue };

type FlatMap = Record<string, JsonValue>;

type LocaleReport = {
  locale: string;
  totalBaseLeafKeys: number;
  missingKeys: string[];
  extraKeys: string[];
  fallbackKeys: string[];
  translatedKeys: string[];
};

const __dirname = path.dirname(url.fileURLToPath(import.meta.url));
const rootDir = path.resolve(__dirname, '..');
const reportPath = path.join(rootDir, 'docs', 'i18n-coverage-report.md');
const baseLocaleKey = 'en-US' as const;
const sampleLimit = 40;

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
    if (Array.isArray(value)) {
      out[nextPrefix] = value as JsonValue;
      continue;
    }
    out[nextPrefix] = (value ?? null) as JsonValue;
  }

  return out;
}

function normalizeForCompare(value: JsonValue): string {
  if (typeof value === 'string') {
    return value.trim();
  }
  return JSON.stringify(value);
}

function buildReport(localeName: string, baseFlat: FlatMap, targetFlat: FlatMap): LocaleReport {
  const missingKeys: string[] = [];
  const fallbackKeys: string[] = [];
  const translatedKeys: string[] = [];

  for (const [key, baseValue] of Object.entries(baseFlat)) {
    if (!(key in targetFlat)) {
      missingKeys.push(key);
      continue;
    }

    const targetValue = targetFlat[key];
    const sameAsBase = normalizeForCompare(targetValue) === normalizeForCompare(baseValue);

    if (sameAsBase) {
      fallbackKeys.push(key);
    } else {
      translatedKeys.push(key);
    }
  }

  const extraKeys = Object.keys(targetFlat).filter(key => !(key in baseFlat)).sort();

  return {
    locale: localeName,
    totalBaseLeafKeys: Object.keys(baseFlat).length,
    missingKeys: missingKeys.sort(),
    extraKeys,
    fallbackKeys: fallbackKeys.sort(),
    translatedKeys: translatedKeys.sort()
  };
}

function pct(numerator: number, denominator: number): string {
  if (!denominator) return '0.00%';
  return `${((numerator / denominator) * 100).toFixed(2)}%`;
}

function formatSample(keys: string[]): string {
  if (keys.length === 0) return '-';
  const sample = keys.slice(0, sampleLimit).map(key => `  - \`${key}\``).join('\n');
  const remaining = keys.length - Math.min(keys.length, sampleLimit);
  return remaining > 0 ? `${sample}\n  - ... and ${remaining} more` : sample;
}

function ensureDir(filePath: string) {
  fs.mkdirSync(path.dirname(filePath), { recursive: true });
}

function main() {
  const baseLocale = locales[baseLocaleKey];
  const baseFlat = flattenLeaves(baseLocale);

  const reports = Object.entries(locales)
    .filter(([key]) => key !== baseLocaleKey)
    .map(([key, localeValue]) => buildReport(key, baseFlat, flattenLeaves(localeValue)))
    .sort((a, b) => a.locale.localeCompare(b.locale));

  const summaryRows = reports
    .map(report => {
      const effectiveTotal = report.totalBaseLeafKeys - report.missingKeys.length;
      const translatedRatio = pct(report.translatedKeys.length, report.totalBaseLeafKeys);
      const nonFallbackRatio = pct(report.translatedKeys.length, Math.max(1, effectiveTotal));

      return `| ${report.locale} | ${report.totalBaseLeafKeys} | ${report.translatedKeys.length} | ${report.fallbackKeys.length} | ${report.missingKeys.length} | ${report.extraKeys.length} | ${translatedRatio} | ${nonFallbackRatio} |`;
    })
    .join('\n');

  const details = reports
    .map(report => {
      return [
        `## ${report.locale}`,
        '',
        `- Base leaf keys: ${report.totalBaseLeafKeys}`,
        `- Translated leaves: ${report.translatedKeys.length} (${pct(report.translatedKeys.length, report.totalBaseLeafKeys)})`,
        `- Fallback-identical leaves: ${report.fallbackKeys.length} (${pct(report.fallbackKeys.length, report.totalBaseLeafKeys)})`,
        `- Missing leaves: ${report.missingKeys.length}`,
        `- Extra leaves: ${report.extraKeys.length}`,
        '',
        '**Sample Fallback Keys**',
        formatSample(report.fallbackKeys),
        '',
        '**Missing Keys**',
        formatSample(report.missingKeys),
        '',
        '**Extra Keys**',
        formatSample(report.extraKeys),
        ''
      ].join('\n');
    })
    .join('\n');

  const markdown = [
    '# i18n Coverage Report',
    '',
    `Generated: ${new Date().toISOString()}`,
    `Base locale: \`${baseLocaleKey}\``,
    '',
    'Metrics:',
    '- `Translated` = leaf value differs from `en-US`',
    '- `Fallback-identical` = leaf exists but equals `en-US` (usually untranslated fallback)',
    '- `Missing` = leaf key not present in locale object',
    '',
    '| Locale | Base Keys | Translated | Fallback | Missing | Extra | Translated/Base | Translated/(Base-Missing) |',
    '| --- | ---: | ---: | ---: | ---: | ---: | ---: | ---: |',
    summaryRows,
    '',
    details
  ].join('\n');

  ensureDir(reportPath);
  fs.writeFileSync(reportPath, markdown, 'utf8');

  console.log(`Wrote report: ${path.relative(process.cwd(), reportPath)}`);
  for (const report of reports) {
    console.log(
      `${report.locale}: translated=${report.translatedKeys.length}, fallback=${report.fallbackKeys.length}, missing=${report.missingKeys.length}, extra=${report.extraKeys.length}`
    );
  }
}

main();
