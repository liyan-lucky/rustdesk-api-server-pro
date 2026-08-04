import {
  dateDeDE,
  dateEnUS,
  dateEsAR,
  dateFrFR,
  dateItIT,
  dateJaJP,
  dateKoKR,
  dateRuRU,
  dateZhCN,
  deDE,
  enUS,
  esAR,
  frFR,
  itIT,
  jaJP,
  koKR,
  ruRU,
  zhCN
} from 'naive-ui';
import type { NDateLocale, NLocale } from 'naive-ui';

export const naiveLocales: Record<App.I18n.LangType, NLocale> = {
  'zh-CN': zhCN,
  'en-US': enUS,
  'ja-JP': jaJP,
  'ko-KR': koKR,
  'fr-FR': frFR,
  'de-DE': deDE,
  'es-ES': esAR,
  'ru-RU': ruRU,
  'it-IT': itIT
};

export const naiveDateLocales: Record<App.I18n.LangType, NDateLocale> = {
  'zh-CN': dateZhCN,
  'en-US': dateEnUS,
  'ja-JP': dateJaJP,
  'ko-KR': dateKoKR,
  'fr-FR': dateFrFR,
  'de-DE': dateDeDE,
  'es-ES': dateEsAR,
  'ru-RU': dateRuRU,
  'it-IT': dateItIT
};
