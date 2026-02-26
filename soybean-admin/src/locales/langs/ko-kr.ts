import enUs from './en-us';

const local: App.I18n.Schema = {
  ...enUs,
  system: { ...enUs.system, title: 'Rustdesk Api Server' },
  icon: { ...enUs.icon, lang: '언어 전환' }
};

export default local;
