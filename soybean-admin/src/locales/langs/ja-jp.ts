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
    system_mail_template: 'メールテンプレート',
    system_mail_logs: 'メールログ',
    system_mail: 'メール',
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
    },
    user: {
      ...enUs.page.user,
      list: {
        ...enUs.page.user.list,
        addUser: 'ユーザー追加',
        editUser: 'ユーザー編集',
        searchPlaceholder: 'ユーザー名/ニックネーム/メール'
      },
      sessions: {
        ...enUs.page.user.sessions,
        kill: '切断',
        confirmKill: 'このセッションを終了しますか？'
      },
      audit: {
        ...enUs.page.user.audit,
        logsSearchPlaceholder: 'ユーザー名/操作/RustdeskID/IP'
      },
      devices: {
        ...enUs.page.user.devices,
        logsSearchPlaceholder: 'ユーザー名/ホスト名/RustdeskID'
      }
    },
    system: {
      ...enUs.page.system,
      mailTemplate: {
        ...enUs.page.system.mailTemplate,
        addMailTemplate: 'テンプレート追加',
        editMailTemplate: 'テンプレート編集',
        inputName: '名前を入力',
        inputSubject: '件名を入力',
        inputContents: '内容を入力',
        selectType: '種類を選択'
      },
      mailLog: {
        ...enUs.page.system.mailLog,
        info: '詳細'
      }
    }
  },
  dataMap: {
    ...enUs.dataMap,
    user: {
      ...enUs.dataMap.user,
      username: 'ユーザー名',
      name: 'ニックネーム',
      email: 'メール',
      licensed_devices: '許可デバイス数',
      login_verify: 'ログイン認証',
      status: '状態',
      is_admin: '管理者',
      created_at: '作成日時'
    },
    device: {
      ...enUs.dataMap.device,
      username: 'ユーザー名',
      hostname: 'ホスト名',
      version: 'RustDesk バージョン',
      os: 'OS',
      rustdesk_id: 'Rustdesk ID'
    },
    audit: {
      ...enUs.dataMap.audit,
      username: 'ユーザー名',
      type: '種類',
      rustdesk_id: 'Rustdesk ID',
      ip: 'IP',
      created_at: '作成日時'
    },
    mailTemplate: {
      ...enUs.dataMap.mailTemplate,
      name: '名前',
      type: '種類',
      subject: '件名',
      contents: '内容',
      created_at: '作成日時'
    },
    mailLog: {
      ...enUs.dataMap.mailLog,
      username: 'ユーザー名',
      from: '送信元',
      to: '宛先',
      subject: '件名',
      status: '状態',
      created_at: '送信日時'
    }
  },
  icon: { ...enUs.icon, lang: '言語を切り替え', reload: 'ページ再読み込み', fullscreen: '全画面' }
};

export default local;
