/* eslint-disable no-underscore-dangle, no-continue, no-lonely-if, n/prefer-global/process */
import fs from 'node:fs';
import path from 'node:path';
import url from 'node:url';

import locales from '../src/locales/locale';
import { flattenLeaves, isSuspect, normalize } from './i18n-utils';
import type { FlatMap } from './i18n-utils';

type LocaleReport = {
  locale: string;
  totalBaseLeafKeys: number;
  missingKeys: string[];
  extraKeys: string[];
  fallbackKeys: string[];
  suspectKeys: string[];
  translatedKeys: string[];
};

const __dirname = path.dirname(url.fileURLToPath(import.meta.url));
const rootDir = path.resolve(__dirname, '..');
const reportPath = path.join(rootDir, 'docs', 'i18n-coverage-report.md');
const baseLocaleKey = 'en-US' as const;
const sampleLimit = 40;

function buildReport(localeName: string, baseFlat: FlatMap, targetFlat: FlatMap): LocaleReport {
  const missingKeys: string[] = [];
  const fallbackKeys: string[] = [];
  const suspectKeys: string[] = [];
  const translatedKeys: string[] = [];

  for (const [key, baseValue] of Object.entries(baseFlat)) {
    if (!(key in targetFlat)) {
      missingKeys.push(key);
      continue;
    }

    const targetValue = targetFlat[key];
    const sameAsBase = normalize(targetValue) === normalize(baseValue);

    if (sameAsBase) {
      fallbackKeys.push(key);
    } else {
      if (isSuspect(baseValue, targetValue)) {
        suspectKeys.push(key);
      } else {
        translatedKeys.push(key);
      }
    }
  }

  const extraKeys = Object.keys(targetFlat).filter(key => !(key in baseFlat)).sort();

  return {
    locale: localeName,
    totalBaseLeafKeys: Object.keys(baseFlat).length,
    missingKeys: missingKeys.sort(),
    extraKeys,
    fallbackKeys: fallbackKeys.sort(),
    suspectKeys: suspectKeys.sort(),
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

      return `| ${report.locale} | ${report.totalBaseLeafKeys} | ${report.translatedKeys.length} | ${report.suspectKeys.length} | ${report.fallbackKeys.length} | ${report.missingKeys.length} | ${report.extraKeys.length} | ${translatedRatio} | ${nonFallbackRatio} |`;
    })
    .join('\n');

  const details = reports
    .map(report => {
      return [
        `## ${report.locale}`,
        '',
        `- Base leaf keys: ${report.totalBaseLeafKeys}`,
        `- Translated leaves: ${report.translatedKeys.length} (${pct(report.translatedKeys.length, report.totalBaseLeafKeys)})`,
        `- Suspect translated leaves: ${report.suspectKeys.length}`,
        `- Fallback-identical leaves: ${report.fallbackKeys.length} (${pct(report.fallbackKeys.length, report.totalBaseLeafKeys)})`,
        `- Missing leaves: ${report.missingKeys.length}`,
        `- Extra leaves: ${report.extraKeys.length}`,
        '',
        '**Sample Fallback Keys**',
        formatSample(report.fallbackKeys),
        '',
        '**Suspect Keys**',
        formatSample(report.suspectKeys),
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
    `Base locale: \`${baseLocaleKey}\``,
    '',
    'Metrics:',
    '- `Translated` = leaf value differs from `en-US` and does not match suspicious placeholder patterns',
    '- `Suspect` = leaf differs from `en-US` but looks corrupted (e.g. many `?` placeholders)',
    '- `Fallback-identical` = leaf exists but equals `en-US` (usually untranslated fallback)',
    '- `Missing` = leaf key not present in locale object',
    '',
    '| Locale | Base Keys | Translated | Suspect | Fallback | Missing | Extra | Translated/Base | Translated/(Base-Missing) |',
    '| --- | ---: | ---: | ---: | ---: | ---: | ---: | ---: | ---: |',
    summaryRows,
    '',
    details
  ].join('\n');

  ensureDir(reportPath);
  const prev = fs.existsSync(reportPath) ? fs.readFileSync(reportPath, 'utf8') : '';
  const changed = prev !== markdown;
  if (changed) {
    fs.writeFileSync(reportPath, markdown, 'utf8');
  }

  console.log(`${changed ? 'Wrote' : 'Unchanged'} report: ${path.relative(process.cwd(), reportPath)}`);
  for (const report of reports) {
    console.log(
      `${report.locale}: translated=${report.translatedKeys.length}, suspect=${report.suspectKeys.length}, fallback=${report.fallbackKeys.length}, missing=${report.missingKeys.length}, extra=${report.extraKeys.length}`
    );
  }
}

main();
