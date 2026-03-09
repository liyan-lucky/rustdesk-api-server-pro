/* eslint-disable no-underscore-dangle, no-continue, n/prefer-global/process */
import fs from 'node:fs';
import url from 'node:url';
import path from 'node:path';

import locales from '../src/locales/locale';
import { flattenLeaves, isSuspect, normalize } from './i18n-utils';

const __dirname = path.dirname(url.fileURLToPath(import.meta.url));
const rootDir = path.resolve(__dirname, '..');
const configPath = path.join(__dirname, 'i18n-quality.config.json');
const summaryPath = path.join(rootDir, 'docs', 'i18n-quality-summary.json');
const baseLocaleKey = 'en-US' as const;

type QualityConfig = {
  mode?: 'standard' | 'strict';
  maxSuspectPerLocale: number;
  maxMissingPerLocale: number;
  maxExtraPerLocale: number;
  ignoreKeyPrefixes: string[];
  ignoreExactKeys: string[];
  maxFallbackRatio: Partial<Record<string, number>>;
  minTranslatedRatio: Partial<Record<string, number>>;
  strictOverrides?: Partial<Pick<QualityConfig, 'maxSuspectPerLocale' | 'maxMissingPerLocale' | 'maxExtraPerLocale'>>;
};

function parseArgs() {
  const args = process.argv.slice(2);
  const strict = args.includes('--strict') || process.env.I18N_QUALITY_MODE === 'strict';
  const json = args.includes('--json');
  const localeArg = args.find(arg => arg.startsWith('--locale='));
  const localeList = localeArg
    ? localeArg
        .replace('--locale=', '')
        .split(',')
        .map(v => v.trim())
        .filter(Boolean)
    : [];
  return { strict, json, locales: localeList };
}

function readQualityConfig(): QualityConfig {
  const fallback: QualityConfig = {
    mode: 'standard',
    maxSuspectPerLocale: 0,
    maxMissingPerLocale: 0,
    maxExtraPerLocale: 0,
    ignoreKeyPrefixes: [],
    ignoreExactKeys: [],
    maxFallbackRatio: {},
    minTranslatedRatio: {},
    strictOverrides: {
      maxSuspectPerLocale: 0,
      maxMissingPerLocale: 0,
      maxExtraPerLocale: 0
    }
  };
  try {
    const raw = fs.readFileSync(configPath, 'utf8');
    const parsed = JSON.parse(raw) as QualityConfig;
    return {
      mode: parsed.mode || fallback.mode,
      maxSuspectPerLocale:
        typeof parsed.maxSuspectPerLocale === 'number' && parsed.maxSuspectPerLocale >= 0
          ? parsed.maxSuspectPerLocale
          : fallback.maxSuspectPerLocale,
      maxMissingPerLocale:
        typeof parsed.maxMissingPerLocale === 'number' && parsed.maxMissingPerLocale >= 0
          ? parsed.maxMissingPerLocale
          : fallback.maxMissingPerLocale,
      maxExtraPerLocale:
        typeof parsed.maxExtraPerLocale === 'number' && parsed.maxExtraPerLocale >= 0
          ? parsed.maxExtraPerLocale
          : fallback.maxExtraPerLocale,
      ignoreKeyPrefixes: Array.isArray(parsed.ignoreKeyPrefixes) ? parsed.ignoreKeyPrefixes : fallback.ignoreKeyPrefixes,
      ignoreExactKeys: Array.isArray(parsed.ignoreExactKeys) ? parsed.ignoreExactKeys : fallback.ignoreExactKeys,
      maxFallbackRatio: parsed.maxFallbackRatio || fallback.maxFallbackRatio,
      minTranslatedRatio: parsed.minTranslatedRatio || {},
      strictOverrides: parsed.strictOverrides || fallback.strictOverrides
    };
  } catch {
    return fallback;
  }
}

function isIgnoredKey(key: string, config: QualityConfig): boolean {
  if (config.ignoreExactKeys.includes(key)) return true;
  return config.ignoreKeyPrefixes.some(prefix => key.startsWith(prefix));
}

function ensureDir(filePath: string) {
  fs.mkdirSync(path.dirname(filePath), { recursive: true });
}

// eslint-disable-next-line complexity
function main() {
  const cli = parseArgs();
  const qualityConfig = readQualityConfig();
  const baseFlat = flattenLeaves(locales[baseLocaleKey]);
  const baseKeysAll = Object.keys(baseFlat);
  const baseKeys = baseKeysAll.filter(key => !isIgnoredKey(key, qualityConfig));
  const baseKeySet = new Set(baseKeys);
  let hasError = false;
  const strictMode = cli.strict || qualityConfig.mode === 'strict';

  const maxSuspectPerLocale = strictMode
    ? qualityConfig.strictOverrides?.maxSuspectPerLocale ?? qualityConfig.maxSuspectPerLocale
    : qualityConfig.maxSuspectPerLocale;
  const maxMissingPerLocale = strictMode
    ? qualityConfig.strictOverrides?.maxMissingPerLocale ?? qualityConfig.maxMissingPerLocale
    : qualityConfig.maxMissingPerLocale;
  const maxExtraPerLocale = strictMode
    ? qualityConfig.strictOverrides?.maxExtraPerLocale ?? qualityConfig.maxExtraPerLocale
    : qualityConfig.maxExtraPerLocale;
  const localeSummaries: Array<{
    locale: string;
    translated: number;
    translatedRatio: number;
    missing: number;
    extra: number;
    suspect: number;
    minTranslatedRatio: number;
    fallbackRatio: number;
    maxFallbackRatio: number;
    pass: boolean;
  }> = [];

  for (const [locale, schema] of Object.entries(locales)) {
    if (locale === baseLocaleKey) continue;
    if (cli.locales.length && !cli.locales.includes(locale)) continue;

    const targetFlat = flattenLeaves(schema);
    const targetKeysAll = Object.keys(targetFlat);
    const missing = baseKeys.filter(key => !(key in targetFlat));
    const extra = targetKeysAll.filter(key => !baseKeySet.has(key) && !isIgnoredKey(key, qualityConfig));
    const suspect = baseKeys.filter(key => {
      if (!(key in targetFlat)) return false;
      const baseVal = baseFlat[key];
      const targetVal = targetFlat[key];
      if (normalize(baseVal) === normalize(targetVal)) return false;
      return isSuspect(baseVal, targetVal);
    });
    const translated = baseKeys.filter(key => {
      if (!(key in targetFlat)) return false;
      const baseVal = baseFlat[key];
      const targetVal = targetFlat[key];
      if (normalize(baseVal) === normalize(targetVal)) return false;
      return !isSuspect(baseVal, targetVal);
    });
    const fallback = baseKeys.filter(key => {
      if (!(key in targetFlat)) return false;
      return normalize(baseFlat[key]) === normalize(targetFlat[key]);
    });

    const minRatio = qualityConfig.minTranslatedRatio[locale] ?? 0;
    const denominator = Math.max(1, baseKeys.length);
    const translatedRatio = translated.length / denominator;
    const fallbackRatio = fallback.length / denominator;
    const maxFallbackRatio = qualityConfig.maxFallbackRatio[locale] ?? 1;

    const ok =
      missing.length <= maxMissingPerLocale &&
      extra.length <= maxExtraPerLocale &&
      suspect.length <= maxSuspectPerLocale &&
      translatedRatio >= minRatio &&
      fallbackRatio <= maxFallbackRatio;
    if (!cli.json) {
      const mark = ok ? 'OK' : 'FAIL';
      console.log(
        `${mark} ${locale}: translated=${translated.length}, ratio=${(translatedRatio * 100).toFixed(2)}%, fallback=${(fallbackRatio * 100).toFixed(2)}%, missing=${missing.length}, extra=${extra.length}, suspect=${suspect.length}`
      );
    }
    localeSummaries.push({
      locale,
      translated: translated.length,
      translatedRatio,
      fallbackRatio,
      maxFallbackRatio,
      missing: missing.length,
      extra: extra.length,
      suspect: suspect.length,
      minTranslatedRatio: minRatio,
      pass: ok
    });

    if (!ok) {
      hasError = true;
      const sample = [...missing, ...extra, ...suspect].slice(0, 10);
      if (!cli.json && sample.length) {
        console.log(`  sample: ${sample.join(', ')}`);
      }
      if (!cli.json && translatedRatio < minRatio) {
        console.log(`  ratio threshold failed: got ${(translatedRatio * 100).toFixed(2)}% < required ${(minRatio * 100).toFixed(2)}%`);
      }
      if (!cli.json && fallbackRatio > maxFallbackRatio) {
        console.log(
          `  fallback threshold failed: got ${(fallbackRatio * 100).toFixed(2)}% > max ${(maxFallbackRatio * 100).toFixed(2)}%`
        );
      }
    }
  }

  const summary = {
    baseLocale: baseLocaleKey,
    strictMode,
    thresholds: {
      maxSuspectPerLocale,
      maxMissingPerLocale,
      maxExtraPerLocale
    },
    locales: localeSummaries
  };
  const summaryText = JSON.stringify(summary, null, 2);
  ensureDir(summaryPath);
  const prev = fs.existsSync(summaryPath) ? fs.readFileSync(summaryPath, 'utf8') : '';
  if (prev !== summaryText) {
    fs.writeFileSync(summaryPath, summaryText, 'utf8');
  }

  if (cli.json) {
    console.log(summaryText);
  }

  if (hasError) {
    console.error(`i18n quality check failed (${path.relative(process.cwd(), rootDir)})`);
    process.exit(1);
  }
}

main();
