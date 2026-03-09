import zhCN from './langs/zh-cn';
import enUS from './langs/en-us';
import jaJP from './langs/ja-jp';
import koKR from './langs/ko-kr';
import frFR from './langs/fr-fr';
import deDE from './langs/de-de';
import esES from './langs/es-es';
import ruRU from './langs/ru-ru';

const locales: Record<App.I18n.LangType, App.I18n.Schema> = {
  'zh-CN': zhCN,
  'en-US': enUS,
  'ja-JP': jaJP,
  'ko-KR': koKR,
  'fr-FR': frFR,
  'de-DE': deDE,
  'es-ES': esES,
  'ru-RU': ruRU,
};

export default locales;
