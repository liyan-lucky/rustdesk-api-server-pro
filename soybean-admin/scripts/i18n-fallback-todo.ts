/* eslint-disable no-underscore-dangle, no-continue */
import fs from 'node:fs';
import path from 'node:path';
import url from 'node:url';

import locales from '../src/locales/locale';
import { flattenLeaves, normalize } from './i18n-utils';

const __dirname = path.dirname(url.fileURLToPath(import.meta.url));
const rootDir = path.resolve(__dirname, '..');
const markdownPath = path.join(rootDir, 'docs', 'i18n-fallback-todo.md');
const jsonPath = path.join(rootDir, 'docs', 'i18n-fallback-todo.json');
const baseLocale = 'en-US' as const;

function parseArgs() {
  const args = process.argv.slice(2);
  const localeArg = args.find(arg => arg.startsWith('--locale='));
  const topArg = args.find(arg => arg.startsWith('--top='));
  const localesArg = localeArg
    ? localeArg
        .replace('--locale=', '')
        .split(',')
        .map(v => v.trim())
        .filter(Boolean)
    : [];
  const top = topArg ? Number(topArg.replace('--top=', '')) : 80;
  return {
    localesArg,
    top: Number.isFinite(top) && top > 0 ? Math.floor(top) : 80
  };
}

function ensureDir(filePath: string) {
  fs.mkdirSync(path.dirname(filePath), { recursive: true });
}

function writeIfChanged(filePath: string, content: string) {
  ensureDir(filePath);
  const prev = fs.existsSync(filePath) ? fs.readFileSync(filePath, 'utf8') : '';
  if (prev !== content) {
    fs.writeFileSync(filePath, content, 'utf8');
  }
}

function byModule(keys: string[]) {
  const counts: Record<string, number> = {};
  for (const key of keys) {
    const moduleName = key.split('.')[0] || 'misc';
    counts[moduleName] = (counts[moduleName] || 0) + 1;
  }
  return Object.entries(counts)
    .sort((a, b) => b[1] - a[1] || a[0].localeCompare(b[0]))
    .map(([module, count]) => ({ module, count }));
}

function main() {
  const cli = parseArgs();
  const baseFlat = flattenLeaves(locales[baseLocale]);
  const baseKeys = Object.keys(baseFlat);
  const targetLocales = Object.keys(locales).filter(locale => {
    if (locale === baseLocale) return false;
    if (!cli.localesArg.length) return true;
    return cli.localesArg.includes(locale);
  });

  const output = targetLocales.map(locale => {
    const targetFlat = flattenLeaves(locales[locale as keyof typeof locales]);
    const fallbackKeys = baseKeys.filter(key => key in targetFlat && normalize(baseFlat[key]) === normalize(targetFlat[key]));
    const topKeys = fallbackKeys.slice(0, cli.top);
    return {
      locale,
      baseKeys: baseKeys.length,
      fallbackCount: fallbackKeys.length,
      translatedCount: baseKeys.length - fallbackKeys.length,
      fallbackRatio: fallbackKeys.length / Math.max(1, baseKeys.length),
      topFallbackKeys: topKeys,
      modules: byModule(fallbackKeys)
    };
  });

  const jsonText = JSON.stringify({ baseLocale, generatedAt: null, locales: output }, null, 2);
  writeIfChanged(jsonPath, jsonText);

  const markdown = [
    '# i18n Fallback TODO',
    '',
    `Base locale: \`${baseLocale}\``,
    '',
    ...output.flatMap(item => [
      `## ${item.locale}`,
      '',
      `- Base keys: ${item.baseKeys}`,
      `- Fallback keys: ${item.fallbackCount} (${(item.fallbackRatio * 100).toFixed(2)}%)`,
      `- Translated keys: ${item.translatedCount}`,
      '',
      '**Top Modules To Translate**',
      ...item.modules.slice(0, 8).map(x => `- \`${x.module}\`: ${x.count}`),
      '',
      `**Top ${Math.min(cli.top, item.topFallbackKeys.length)} Fallback Keys**`,
      ...item.topFallbackKeys.map(key => `- \`${key}\``),
      ''
    ])
  ].join('\n');
  writeIfChanged(markdownPath, markdown);

  for (const item of output) {
    console.log(
      `${item.locale}: fallback=${item.fallbackCount}/${item.baseKeys} (${(item.fallbackRatio * 100).toFixed(2)}%), translated=${item.translatedCount}`
    );
  }
}

main();
