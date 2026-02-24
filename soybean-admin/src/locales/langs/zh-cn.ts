import enUs from './en-us';

const local: App.I18n.Schema = {
  ...enUs,
  system: {
    ...enUs.system,
    updateTitle: '版本更新提示（ChatGPT生成）',
    updateContent: '此更新内容由 ChatGPT 生成/整理，可能存在错误或遗漏。请先自行审查并测试，尤其在生产环境中务必谨慎使用。是否现在刷新？'
  }
};

export default local;
