import enUs from './en-us';

const local: App.I18n.Schema = {
  ...enUs,
  system: { ...enUs.system, title: 'Rustdesk Api Server' },
  icon: { ...enUs.icon, lang: '言語を切り替え' }
};

export default local;
