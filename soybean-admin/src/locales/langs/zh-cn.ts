import enUs from './en-us';

const local: App.I18n.Schema = {
  ...enUs,
  system: {
    ...enUs.system,
    title: 'Rustdesk Api Server',
    updateTitle: '系统版本更新通知',
    updateContent: '检测到系统有新版本发布，是否立即刷新页面？',
    updateConfirm: '立即刷新',
    updateCancel: '稍后再说'
  },
  common: {
    ...enUs.common,
    action: '??',
    add: '??',
    addSuccess: '????',
    backToHome: '????',
    batchDelete: '????',
    cancel: '??',
    close: '??',
    check: '??',
    expandColumn: '???',
    columnSetting: '???',
    config: '??',
    confirm: '??',
    delete: '??',
    deleteSuccess: '????',
    confirmDelete: '??????',
    edit: '??',
    look: '??',
    warning: '??',
    error: '??',
    index: '??',
    keywordSearch: '??????',
    logout: '????',
    logoutConfirm: '????????',
    lookForward: '????',
    modify: '??',
    modifySuccess: '????',
    noData: '????',
    operate: '??',
    pleaseCheckValue: '??????????',
    refresh: '??',
    reset: '??',
    search: '??',
    switch: '??',
    tip: '??',
    trigger: '??',
    update: '??',
    updateSuccess: '????',
    userCenter: '????',
    yesOrNo: {
      yes: '?',
      no: '?'
    }
  },
  request: {
    ...enUs.request,
    logout: '?????????',
    logoutMsg: '????????????',
    logoutWithModal: '??????????????',
    logoutWithModalMsg: '????????????',
    refreshToken: '??????????????',
    tokenExpired: '???????'
  },
  theme: {
    ...enUs.theme,
    themeSchema: {
      ...enUs.theme.themeSchema,
      title: '????',
      light: '??',
      dark: '??',
      auto: '????'
    },
    grayscale: '????',
    colourWeakness: '????',
    layoutMode: {
      ...enUs.theme.layoutMode,
      title: '????',
      vertical: '??????',
      horizontal: '??????',
      'vertical-mix': '????????',
      'horizontal-mix': '????????',
      reverseHorizontalMix: '????????????'
    },
    recommendColor: '????????',
    recommendColorDesc: '????????',
    themeColor: {
      ...enUs.theme.themeColor,
      title: '????',
      primary: '??',
      info: '??',
      success: '??',
      warning: '??',
      error: '??',
      followPrimary: '????'
    },
    scrollMode: {
      ...enUs.theme.scrollMode,
      title: '????',
      wrapper: '??',
      content: '??'
    },
    page: {
      ...enUs.theme.page,
      animate: '????',
      mode: {
        ...enUs.theme.page.mode,
        title: '??????',
        fade: '????',
        'fade-slide': '??',
        'fade-bottom': '????',
        'fade-scale': '????(??)',
        'zoom-fade': '????',
        'zoom-out': '????',
        none: '?'
      }
    },
    fixedHeaderAndTab: '????????',
    header: {
      ...enUs.theme.header,
      height: '????',
      breadcrumb: {
        ...enUs.theme.header.breadcrumb,
        visible: '?????',
        showIcon: '???????'
      }
    },
    tab: {
      ...enUs.theme.tab,
      visible: '?????',
      cache: '?????',
      height: '?????',
      mode: {
        ...enUs.theme.tab.mode,
        title: '?????',
        chrome: 'Chrome',
        button: '??'
      }
    },
    sider: {
      ...enUs.theme.sider,
      inverted: '?????',
      width: '?????',
      collapsedWidth: '???????',
      mixWidth: '???????',
      mixCollapsedWidth: '?????????',
      mixChildMenuWidth: '???????'
    },
    footer: {
      ...enUs.theme.footer,
      visible: '????',
      fixed: '????',
      height: '????',
      right: '????'
    },
    watermark: {
      ...enUs.theme.watermark,
      visible: '??????',
      text: '????'
    },
    themeDrawerTitle: '????',
    pageFunTitle: '????',
    configOperation: {
      ...enUs.theme.configOperation,
      copyConfig: '????',
      copySuccessMsg: '???????? "src/theme/settings.ts" ???? "themeSettings"',
      resetConfig: '????',
      resetSuccessMsg: '????'
    }
  },
  route: {
    ...enUs.route,
    login: '??',
    403: '???',
    404: '?????',
    500: '?????',
    'iframe-page': '????',
    home: '??',
    audit: '????',
    user: '????',
    user_list: '????',
    user_sessions: '????',
    system: '????',
    system_mail_template: '????',
    system_mail_logs: '????',
    system_mail: '????',
    audit_baselogs: '????',
    audit_filetransferlogs: '??????',
    devices: '????'
  },
  page: {
    ...enUs.page,
      login: {
        ...enUs.page.login,
        common: {
          ...enUs.page.login.common,
          loginOrRegister: '?? / ??',
          userNamePlaceholder: '??????',
          phonePlaceholder: '??????',
          codePlaceholder: '??????',
          passwordPlaceholder: '?????',
          confirmPasswordPlaceholder: '???????',
          codeLogin: '?????',
          confirm: '??',
          back: '??',
          validateSuccess: '????',
          loginSuccess: '????',
          welcomeBack: '?????{userName}?'
        },
        pwdLogin: {
          ...enUs.page.login.pwdLogin,
          title: '????',
          rememberMe: '???'
        }
      },
    home: {
      ...enUs.page.home,
      greeting: '你好，{userName}，今天又是充满活力的一天！',
      userCount: '用户数量',
      deviceCount: '设备数量',
      onlineCount: '在线数量',
      visitsCount: '访问次数',
      operatingSystem: '操作系统',
      oneWeek: '一周内',
      changeLogs: '更新日志',
      serverConfig: {
        title: '客户端连接配置',
        tip: '复制以下配置后可直接填写到 RustDesk 客户端（如缺少 KEY，请在容器环境变量中设置 `RUSTDESK_KEY`）。',
        idServer: 'ID服务器',
        relayServer: '中继服务器',
        apiServer: 'API服务器',
        key: 'KEY',
        idServerPlaceholder: '例如 your.domain.com',
        relayServerPlaceholder: '例如 your.domain.com',
        apiServerPlaceholder: '例如 https://your.domain.com',
        keyPlaceholder: '可通过环境变量 RUSTDESK_KEY 提供',
        copy: '复制',
        copyAll: '复制全部',
        copyTemplate: '复制RustDesk模板',
        refresh: '刷新配置',
        clearCacheReload: '清缓存并重载',
        source: '来源',
        lastUpdated: '最后更新',
        show: '显示',
        hide: '隐藏',
        missingTip: '以下字段为空，请先在容器环境变量中配置：{fields}',
        copyEmpty: '{label}为空，无法复制',
        copySuccess: '{label}已复制',
        copyFailed: '{label}复制失败',
        fetchFailed: '获取服务器配置失败，请稍后重试',
        cacheCleared: '已清除缓存，正在重新加载服务器配置',
        sourceType: {
          remote: '远端接口',
          'memory-cache': '内存缓存',
          'session-cache': '会话缓存',
          env: '环境变量',
          inferred: '自动推断',
          empty: '未配置'
        },
        sourceHint: {
          env: '该值来自容器环境变量配置。',
          inferred: '该值根据当前访问地址自动推断生成。',
          empty: '当前既没有配置，也无法自动推断。'
        },
        connectivity: {
          clear: '清除检测结果',
          check: '检测连通性',
          checkOne: '检测',
          checked: '连通性检测完成',
          checkedOne: '{field} 连通性检测完成',
          checkedCached: '使用最近一次检测结果（缓存）',
          checkFailed: '连通性检测失败',
          cleared: '已清除连通性检测结果',
          source: '检测来源',
          lastChecked: '最后检测',
          target: '目标',
          duration: '耗时',
          notChecked: '尚未检测',
          checkSourceType: {
            remote: '远端检测',
            cache: '缓存结果'
          },
          status: {
            idle: '未检测',
            ok: '可连接',
            error: '失败',
            skip: '跳过'
          }
        }
      }
    },
    user: {
      ...enUs.page.user,
      list: {
        ...enUs.page.user.list,
        addUser: '添加用户',
        editUser: '编辑用户',
        inputUsername: '请输入用户名',
        inputPassword: '请输入密码',
        inputNickname: '请输入昵称',
        emailFormatError: '邮箱格式错误',
        selectUserStatus: '请选择用户状态',
        searchPlaceholder: '用户名、昵称、邮箱'
      },
      sessions: {
        ...enUs.page.user.sessions,
        kill: '终结',
        confirmKill: '是否终结该会话？'
      },
      audit: {
        ...enUs.page.user.audit,
        logsSearchPlaceholder: '用户名、行为、RustdeskID、IP'
      },
      devices: {
        ...enUs.page.user.devices,
        logsSearchPlaceholder: '用户名、主机名、RustdeskID'
      }
    },
    system: {
      ...enUs.page.system,
      mailTemplate: {
        ...enUs.page.system.mailTemplate,
        addMailTemplate: '添加模板',
        editMailTemplate: '编辑模板'
      },
      mailLog: {
        ...enUs.page.system.mailLog,
        info: '详情'
      }
    }
  },
  dropdown: {
    ...enUs.dropdown,
    closeCurrent: '????',
    closeOther: '????',
    closeLeft: '????',
    closeRight: '????',
    closeAll: '????'
  },
  icon: {
    ...enUs.icon,
    themeConfig: '????',
    themeSchema: '????',
    lang: '????',
    fullscreen: '??',
    fullscreenExit: '????',
    reload: '????',
    collapse: '????',
    expand: '????',
    pin: '??',
    unpin: '????'
  },
  datatable: {
    ...enUs.datatable,
    itemCount: '? {total} ?'
  },
  dataMap: {
    ...enUs.dataMap,
    user: {
      ...enUs.dataMap.user,
      username: '用户名',
      password: '密码',
      name: '昵称',
      email: '邮箱',
      licensed_devices: '授权设备数',
      login_verify: '登录验证',
      status: '状态',
      is_admin: '管理员',
      tfa_secret: '2FA密钥',
      tfa_code: '2FA验证码',
      created_at: '创建时间',
      statusLabel: {
        ...enUs.dataMap.user.statusLabel,
        disabled: '禁用',
        unverified: '未验证',
        normal: '正常'
      },
      loginVerifyLabel: {
        ...enUs.dataMap.user.loginVerifyLabel,
        none: '无需验证',
        emailCheck: '邮箱验证',
        tfaCheck: '双重认证'
      }
    },
    session: {
      ...enUs.dataMap.session,
      expired: '过期时间',
      created_at: '创建时间'
    },
    device: {
      ...enUs.dataMap.device,
      username: '用户名',
      hostname: '主机名',
      version: 'RustDesk版本',
      memory: '内存',
      os: '操作系统',
      rustdesk_id: 'RustdeskID'
    },
    audit: {
      ...enUs.dataMap.audit,
      username: '用户名',
      type: '类型',
      conn_id: '连接ID',
      rustdesk_id: 'RustdeskID',
      session_id: '会话ID',
      created_at: '创建时间',
      closed_at: '结束时间',
      peer_id: '对端ID',
      path: '路径'
    }
  }
};

export default local;
