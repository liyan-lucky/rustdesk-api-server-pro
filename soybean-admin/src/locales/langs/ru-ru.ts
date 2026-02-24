import enUs from './en-us';

const local: App.I18n.Schema = {
  ...enUs,
  system: {
    ...enUs.system,
    updateTitle: 'Version Update Notice (ChatGPT-Generated)',
    updateContent:
      'This update content was generated/assembled by ChatGPT and may contain mistakes. Review and test carefully before using it, especially in production. Refresh now?'
  }
};

export default local;
