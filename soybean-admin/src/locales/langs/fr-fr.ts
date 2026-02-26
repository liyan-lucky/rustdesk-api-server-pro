import enUs from './en-us';

const local: App.I18n.Schema = {
  ...enUs,
  system: { ...enUs.system, title: 'Rustdesk Api Server' },
  icon: { ...enUs.icon, lang: 'Changer de langue' }
};

export default local;
