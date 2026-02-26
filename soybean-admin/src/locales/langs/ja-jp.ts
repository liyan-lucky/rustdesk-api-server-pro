import enUs from './en-us';

const local: App.I18n.Schema = {
  ...enUs,
  system: { ...enUs.system, title: 'Rustdesk Api Server' },
  common: {
    ...enUs.common,
    action: '??',
    add: '??',
    addSuccess: '????',
    backToHome: '??????',
    batchDelete: '????',
    cancel: '?????',
    close: '???',
    check: '??',
    expandColumn: '????',
    columnSetting: '???',
    config: '??',
    confirm: '??',
    delete: '??',
    deleteSuccess: '????',
    confirmDelete: '?????????????',
    edit: '??',
    look: '??',
    warning: '??',
    error: '???',
    index: '??',
    keywordSearch: '??????????????',
    logout: '?????',
    logoutConfirm: '????????????????',
    lookForward: '????',
    modify: '??',
    modifySuccess: '????',
    noData: '?????',
    operate: '??',
    pleaseCheckValue: '?????????????',
    refresh: '??',
    reset: '????',
    search: '??',
    switch: '??',
    tip: '???',
    trigger: '??',
    update: '??',
    updateSuccess: '????',
    userCenter: '????????',
    yesOrNo: {
      yes: '??',
      no: '???'
    }
  },
  request: {
    ...enUs.request,
    logout: '???????????????????',
    logoutMsg: '????????????????????????',
    logoutWithModal: '????????????????????????',
    logoutWithModalMsg: '????????????????????????',
    refreshToken: '????????????????????',
    tokenExpired: '?????????????????????'
  },
  theme: {
    ...enUs.theme,
    themeSchema: {
      ...enUs.theme.themeSchema,
      title: '???????',
      light: '???',
      dark: '???',
      auto: '???????'
    },
    grayscale: '???????',
    colourWeakness: '??????',
    layoutMode: {
      ...enUs.theme.layoutMode,
      title: '????????'
    },
    recommendColor: '??????????????',
    themeColor: {
      ...enUs.theme.themeColor,
      title: '??????'
    },
    scrollMode: {
      ...enUs.theme.scrollMode,
      title: '????????'
    },
    page: {
      ...enUs.theme.page,
      animate: '??????????',
      mode: {
        ...enUs.theme.page.mode,
        title: '??????????'
      }
    },
    fixedHeaderAndTab: '??????????',
    header: {
      ...enUs.theme.header,
      height: '??????'
    },
    tab: {
      ...enUs.theme.tab,
      visible: '????',
      mode: {
        ...enUs.theme.tab.mode,
        title: '?????'
      }
    },
    sider: {
      ...enUs.theme.sider,
      width: '??????'
    },
    footer: {
      ...enUs.theme.footer,
      visible: '??????'
    },
    themeDrawerTitle: '?????',
    pageFunTitle: '?????',
    configOperation: {
      ...enUs.theme.configOperation,
      copyConfig: '??????',
      resetConfig: '???????',
      resetSuccessMsg: '??????'
    }
  },
  route: {
    ...enUs.route,
    login: '????',
    403: '????',
    404: '???????????',
    500: '???????',
    'iframe-page': 'Iframe',
    home: '???',
    audit: '??',
    user: '??????',
    user_list: '??????',
    user_sessions: '?????',
    system: '??????',
    system_mail_template: '?????????',
    system_mail_logs: '?????',
    system_mail: '???',
    audit_baselogs: '????',
    audit_filetransferlogs: '????????',
    devices: '????'
  },
  page: {
    ...enUs.page,
      login: {
        ...enUs.page.login,
        common: {
          ...enUs.page.login.common,
          loginOrRegister: '???? / ??',
          userNamePlaceholder: '??????????????',
          phonePlaceholder: '?????????????',
          codePlaceholder: '??????????????',
          passwordPlaceholder: '??????????????',
          confirmPasswordPlaceholder: '???????????????',
          codeLogin: '?????????',
          confirm: '??',
          back: '??',
          validateSuccess: '????',
          loginSuccess: '??????',
          welcomeBack: '????????{userName} ???'
        },
        pwdLogin: {
          ...enUs.page.login.pwdLogin,
          title: '?????????',
          rememberMe: '?????????'
        }
      },
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
        password: '?????',
      name: 'ニックネーム',
      email: 'メール',
      licensed_devices: '許可デバイス数',
      login_verify: 'ログイン認証',
      status: '状態',
      is_admin: '管理者',
        tfa_secret: '2FA??????',
        tfa_code: '2FA???',
      created_at: '作成日時',
      statusLabel: {
        ...enUs.dataMap.user.statusLabel,
        disabled: '無効',
        unverified: '未確認',
        normal: '正常'
      },
      loginVerifyLabel: {
        ...enUs.dataMap.user.loginVerifyLabel,
        none: '不要',
        emailCheck: 'メール認証',
        tfaCheck: '2FA'
      }
    },
      session: {
        ...enUs.dataMap.session,
        expired: '??????',
        created_at: '????'
      },
    device: {
      ...enUs.dataMap.device,
      username: 'ユーザー名',
      hostname: 'ホスト名',
      version: 'RustDesk バージョン',
        memory: '???',
      os: 'OS',
      rustdesk_id: 'Rustdesk ID'
    },
    audit: {
      ...enUs.dataMap.audit,
      username: 'ユーザー名',
      type: '種類',
        conn_id: '??ID',
      rustdesk_id: 'Rustdesk ID',
        peer_id: 'Peer ID',
      ip: 'IP',
        session_id: '?????ID',
        uuid: 'UUID',
      created_at: '作成日時',
        closed_at: '????',
      typeLabel: {
        ...enUs.dataMap.audit.typeLabel,
        remote_control: 'リモート操作',
        file_transfer: 'ファイル転送',
        tcp_tunnel: 'TCP トンネル'
      },
      fileTransferTypeLabel: {
        ...enUs.dataMap.audit.fileTransferTypeLabel,
        master_controlled: '操作側 -> 被操作側',
        controlled_master: '被操作側 -> 操作側'
      },
        path: '??'
    },
    mailTemplate: {
      ...enUs.dataMap.mailTemplate,
      name: '名前',
      type: '種類',
      subject: '件名',
      contents: '内容',
      created_at: '作成日時',
      typeLabel: {
        ...enUs.dataMap.mailTemplate.typeLabel,
        loginVerify: 'ログイン認証',
        registerVerify: '登録認証',
        other: 'その他'
      }
    },
    mailLog: {
      ...enUs.dataMap.mailLog,
      username: 'ユーザー名',
        uuid: 'UUID',
      from: '送信元',
      to: '宛先',
      subject: '件名',
        contents: '??',
      status: '状態',
      created_at: '送信日時',
      statusLabel: {
        ...enUs.dataMap.mailLog.statusLabel,
        ok: '成功',
        err: '失敗'
      }
    }
  },
  api: {
    ...enUs.api,
    CaptchaError: 'CAPTCHA エラー',
    UserNotExists: 'ユーザーが存在しません',
    UsernameOrPasswordError: 'アカウントまたはパスワードが正しくありません',
    UserExists: 'ユーザー名は既に使用されています',
    UsernameEmpty: 'ユーザー名を入力してください',
    PasswordEmpty: 'パスワードを入力してください',
    UserAddSuccess: 'ユーザーを作成しました',
    DataError: 'データエラー',
    UserUpdateSuccess: 'ユーザーを更新しました',
    UserDeleteSuccess: 'ユーザーを削除しました',
    SessionKillSuccess: 'セッションを終了しました',
    MailTemplateNameEmpty: 'テンプレート名を入力してください',
    MailTemplateSubjectEmpty: '件名を入力してください',
    MailTemplateContentsEmpty: '内容を入力してください',
    MailTemplateAddSuccess: 'メールテンプレートを作成しました',
    MailTemplateUpdateSuccess: 'メールテンプレートを更新しました',
    NoEmailAddress: 'メールアドレスが設定されていません',
    VerificationCodeError: '認証コードエラー',
    UUIDEmpty: 'UUID を入力してください'
  },
  icon: { ...enUs.icon, lang: '言語を切り替え', reload: 'ページ再読み込み', fullscreen: '全画面' }
};

export default local;
