/* eslint-disable no-underscore-dangle, n/prefer-global/process */
import fs from 'node:fs';
import path from 'node:path';
import url from 'node:url';

type LocaleSummary = {
  locale: string;
  translated: number;
  translatedRatio: number;
  fallbackRatio: number;
  missing: number;
  extra: number;
  suspect: number;
  pass: boolean;
};

type QualitySummary = {
  strictMode: boolean;
  locales: LocaleSummary[];
};

const __dirname = path.dirname(url.fileURLToPath(import.meta.url));
const summaryPath = path.join(path.resolve(__dirname, '..'), 'docs', 'i18n-quality-summary.json');

function main() {
  if (!fs.existsSync(summaryPath)) {
    console.error(`Missing summary file: ${summaryPath}`);
    process.exit(1);
  }

  const data = JSON.parse(fs.readFileSync(summaryPath, 'utf8')) as QualitySummary;
  const mode = data.strictMode ? 'strict' : 'standard';
  console.log(`i18n quality summary (${mode})`);
  for (const item of data.locales) {
    console.log(
      `${item.pass ? 'OK  ' : 'FAIL'} ${item.locale} | translated=${(item.translatedRatio * 100).toFixed(2)}% | fallback=${(item.fallbackRatio * 100).toFixed(2)}% | missing=${item.missing} | extra=${item.extra} | suspect=${item.suspect}`
    );
  }
}

main();
