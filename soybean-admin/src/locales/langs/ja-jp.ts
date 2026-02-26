import enUs from './en-us';

const local: App.I18n.Schema = {
  ...enUs,
  system: { ...enUs.system, title: 'Rustdesk Api Server' },
  common: {
    ...enUs.common,
    refresh: '更新',
    noData: 'データなし',
    error: 'エラー',
    warning: '警告'
  },
  route: {
    ...enUs.route,
    home: 'ホーム',
    audit: '監査',
    user: 'ユーザー管理',
    user_list: 'ユーザー一覧',
    user_sessions: 'セッション',
    system: 'システム管理',
    devices: 'デバイス'
  },
  page: {
    ...enUs.page,
    home: {
      ...enUs.page.home,
      greeting: 'おはようございます、{userName}さん！',
      userCount: 'ユーザー数',
      deviceCount: 'デバイス数',
      onlineCount: 'オンライン数',
      visitsCount: '訪問数',
      operatingSystem: 'OS',
      oneWeek: '1週間',
      changeLogs: '更新履歴',
      serverConfig: {
        ...enUs.page.home.serverConfig,
        title: 'クライアント接続設定',
        tip: '以下の設定を RustDesk クライアントに貼り付けてください。KEY が空の場合は `RUSTDESK_KEY` を設定してください。',
        refresh: '更新',
        copyAll: 'すべてコピー',
        copyTemplate: 'RustDeskテンプレートをコピー',
        connectivity: {
          ...enUs.page.home.serverConfig.connectivity,
          check: '接続確認',
          checkOne: '確認',
          clear: '結果クリア'
        }
      }
    }
  },
  icon: { ...enUs.icon, lang: '言語を切り替え', reload: 'ページ再読み込み', fullscreen: '全画面' }
};

export default local;
