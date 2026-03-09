/* eslint-disable no-underscore-dangle, no-continue, n/prefer-global/process */
import fs from 'node:fs';
import path from 'node:path';
import url from 'node:url';

type LocaleSummary = {
  locale: string;
  translatedRatio: number;
  fallbackRatio: number;
};

type QualitySummary = {
  locales: LocaleSummary[];
};

type BaselineConfig = {
  locales: Record<string, { minTranslatedRatio: number; maxFallbackRatio: number }>;
};

const __dirname = path.dirname(url.fileURLToPath(import.meta.url));
const rootDir = path.resolve(__dirname, '..');
const summaryPath = path.join(rootDir, 'docs', 'i18n-quality-summary.json');
const baselinePath = path.join(__dirname, 'i18n-quality-baseline.json');

function readJsonFile<T>(filePath: string): T {
  return JSON.parse(fs.readFileSync(filePath, 'utf8')) as T;
}

function main() {
  if (!fs.existsSync(summaryPath)) {
    console.error(`missing summary: ${summaryPath}`);
    process.exit(1);
  }
  if (!fs.existsSync(baselinePath)) {
    console.error(`missing baseline: ${baselinePath}`);
    process.exit(1);
  }

  const summary = readJsonFile<QualitySummary>(summaryPath);
  const baseline = readJsonFile<BaselineConfig>(baselinePath);
  let failed = false;

  for (const item of summary.locales) {
    const target = baseline.locales[item.locale];
    if (!target) {
      console.log(`SKIP ${item.locale}: no baseline`);
      continue;
    }

    const translatedOk = item.translatedRatio >= target.minTranslatedRatio;
    const fallbackOk = item.fallbackRatio <= target.maxFallbackRatio;
    const ok = translatedOk && fallbackOk;
    const mark = ok ? 'OK' : 'FAIL';
    console.log(
      `${mark} ${item.locale}: translated ${(item.translatedRatio * 100).toFixed(2)}% >= ${(target.minTranslatedRatio * 100).toFixed(2)}%, fallback ${(item.fallbackRatio * 100).toFixed(2)}% <= ${(target.maxFallbackRatio * 100).toFixed(2)}%`
    );
    if (!ok) failed = true;
  }

  if (failed) {
    process.exit(1);
  }
}

main();
