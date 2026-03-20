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
    action: '操作',
    add: '新增',
    addSuccess: '新增成功',
    backToHome: '返回首页',
    batchDelete: '批量删除',
    cancel: '取消',
    close: '关闭',
    check: '检查',
    expandColumn: '展开列',
    columnSetting: '列设置',
    config: '配置',
    confirm: '确认',
    delete: '删除',
    deleteSuccess: '删除成功',
    confirmDelete: '确认删除吗？',
    edit: '编辑',
    look: '查看',
    warning: '警告',
    error: '错误',
    index: '序号',
    keywordSearch: '请输入关键词',
    logout: '退出登录',
    logoutConfirm: '确认退出登录吗？',
    lookForward: '敬请期待',
    modify: '修改',
    modifySuccess: '修改成功',
    noData: '暂无数据',
    operate: '操作',
    pleaseCheckValue: '请检查输入值是否有效',
    refresh: '刷新',
    reset: '重置',
    search: '搜索',
    switch: '切换',
    tip: '提示',
    trigger: '触发',
    update: '更新',
    updateSuccess: '更新成功',
    userCenter: '个人中心',
    yesOrNo: {
      yes: '是',
      no: '否'
    }
  },
  request: {
    ...enUs.request,
    logout: '请求失败后退出登录',
    logoutMsg: '用户状态无效，请重新登录',
    logoutWithModal: '请求失败后弹窗提示并退出登录',
    logoutWithModalMsg: '用户状态无效，请重新登录',
    refreshToken: '请求令牌已过期，正在刷新令牌',
    tokenExpired: '请求令牌已过期'
  },
  theme: {
    ...enUs.theme,
    themeSchema: {
      ...enUs.theme.themeSchema,
      title: '主题方案',
      light: '亮色',
      dark: '暗色',
      auto: '跟随系统'
    },
    grayscale: '灰度模式',
    colourWeakness: '色弱模式',
    layoutMode: {
      ...enUs.theme.layoutMode,
      title: '布局模式',
      vertical: '垂直菜单模式',
      horizontal: '水平菜单模式',
      'vertical-mix': '垂直混合菜单模式',
      'horizontal-mix': '水平混合菜单模式',
      reverseHorizontalMix: '反转一级菜单与子菜单位置'
    },
    recommendColor: '应用推荐配色算法',
    recommendColorDesc: '推荐配色算法参考',
    themeColor: {
      ...enUs.theme.themeColor,
      title: '主题颜色',
      primary: '主色',
      info: '信息',
      success: '成功',
      warning: '警告',
      error: '错误',
      followPrimary: '跟随主色'
    },
    scrollMode: {
      ...enUs.theme.scrollMode,
      title: '滚动模式',
      wrapper: '容器',
      content: '内容'
    },
    page: {
      ...enUs.theme.page,
      animate: '页面动画',
      mode: {
        ...enUs.theme.page.mode,
        title: '页面动画模式',
        fade: '淡入淡出',
        'fade-slide': '滑动',
        'fade-bottom': '淡入缩放',
        'fade-scale': '淡入缩放（比例）',
        'zoom-fade': '缩放淡入',
        'zoom-out': '缩小淡出',
        none: '无'
      }
    },
    fixedHeaderAndTab: '固定头部与标签页',
    header: {
      ...enUs.theme.header,
      height: '头部高度',
      breadcrumb: {
        ...enUs.theme.header.breadcrumb,
        visible: '显示面包屑',
        showIcon: '显示面包屑图标'
      }
    },
    tab: {
      ...enUs.theme.tab,
      visible: '显示标签页',
      cache: '标签页缓存',
      height: '标签页高度',
      mode: {
        ...enUs.theme.tab.mode,
        title: '标签页模式',
        chrome: 'Chrome',
        button: '按钮'
      }
    },
    sider: {
      ...enUs.theme.sider,
      inverted: '深色侧边栏',
      width: '侧边栏宽度',
      collapsedWidth: '侧边栏收起宽度',
      mixWidth: '混合侧边栏宽度',
      mixCollapsedWidth: '混合侧边栏收起宽度',
      mixChildMenuWidth: '混合子菜单宽度'
    },
    footer: {
      ...enUs.theme.footer,
      visible: '显示页脚',
      fixed: '固定页脚',
      height: '页脚高度',
      right: '右侧页脚'
    },
    watermark: {
      ...enUs.theme.watermark,
      visible: '显示全屏水印',
      text: '水印文本'
    },
    themeDrawerTitle: '主题配置',
    pageFunTitle: '页面功能',
    configOperation: {
      ...enUs.theme.configOperation,
      copyConfig: '复制配置',
      copySuccessMsg: '复制成功，请替换 "src/theme/settings.ts" 中的变量 "themeSettings"',
      resetConfig: '重置配置',
      resetSuccessMsg: '重置成功'
    }
  },
  route: {
    ...enUs.route,
    login: '登录',
    403: '无权限',
    404: '页面不存在',
    500: '服务器错误',
    'iframe-page': '内嵌页面',
    home: '首页',
    audit: '日志审计',
    user: '用户管理',
    user_list: '用户列表',
    user_sessions: '会话管理',
    system: '系统管理',
    system_mail_template: '邮件模板',
    system_mail_logs: '邮件日志',
    system_mail: '邮件管理',
    system_server: '服务器配置',
    audit_baselogs: '基础日志',
    audit_filetransferlogs: '文件传输日志',
    devices: '设备管理'
  },
  page: {
    ...enUs.page,
    login: {
      ...enUs.page.login,
      common: {
        ...enUs.page.login.common,
        loginOrRegister: '登录 / 注册',
        userNamePlaceholder: '请输入用户名',
        phonePlaceholder: '请输入手机号',
        codePlaceholder: '请输入验证码',
        passwordPlaceholder: '请输入密码',
        confirmPasswordPlaceholder: '请再次输入密码',
        codeLogin: '验证码登录',
        confirm: '确定',
        back: '返回',
        validateSuccess: '验证成功',
        loginSuccess: '登录成功',
        welcomeBack: '欢迎回来，{userName}！',
        thirdPartyLogin: '第三方登录',
        continueWith: '使用 {provider} 登录',
        providerUnavailable: '{provider} 登录暂不可用'
      },
      pwdLogin: {
        ...enUs.page.login.pwdLogin,
        title: '密码登录',
        rememberMe: '记住我'
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
      cardDetail: {
        viewHint: '点击查看详情',
        recentUsers: '最近用户',
        recentDevices: '最近设备',
        recentVisits: '最近访问记录',
        desc: {
          userCount: '展示当前系统内用户总数。',
          deviceCount: '展示当前系统内设备总数。',
          onlineCount: '展示基于心跳统计得到的在线设备数量。',
          visitCount: '展示来自审计日志的访问次数统计。'
        }
      },
      serverConfig: {
        ...enUs.page.home.serverConfig,
        title: '客户端连接配置',
        tip: '复制以下配置后可直接填写到 RustDesk 客户端。如 KEY 为空，请在容器环境变量中设置 `RUSTDESK_KEY`。',
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
        cacheTtlHint: '缓存时长：配置 {configSeconds} 秒，连通性 {connectivitySeconds} 秒',
        source: '来源',
        lastUpdated: '最后更新',
        ageSeconds: '{seconds} 秒前',
        show: '显示',
        hide: '隐藏',
        missingTip: '以下字段为空，请先在容器环境变量中配置：{fields}',
        copyEmpty: '{label} 为空，无法复制',
        copySuccess: '{label} 已复制',
        copyFailed: '{label} 复制失败',
        fetchFailed: '获取服务器配置失败，请稍后重试',
        cacheCleared: '已清除缓存，正在重新加载服务器配置',
        sourceType: {
          ...enUs.page.home.serverConfig.sourceType,
          remote: '远端接口',
          'memory-cache': '内存缓存',
          'session-cache': '会话缓存',
          env: '环境变量',
          inferred: '自动推断',
          empty: '未配置'
        },
        sourceHint: {
          ...enUs.page.home.serverConfig.sourceHint,
          env: '该值来自容器环境变量配置。',
          inferred: '该值根据当前访问地址自动推断生成。',
          empty: '当前未配置且无法自动推断。'
        },
        connectivity: {
          ...enUs.page.home.serverConfig.connectivity,
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
            ...enUs.page.home.serverConfig.connectivity.checkSourceType,
            remote: '远端检测',
            cache: '缓存结果'
          },
          status: {
            ...enUs.page.home.serverConfig.connectivity.status,
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
        searchPlaceholder: '用户名/昵称/邮箱',
        tfa_secret_bind: '绑定2FA设备',
        require2FASecret: '2FA密钥为空',
        require2FACode: '2FA验证码不能为空'
      },
      sessions: {
        ...enUs.page.user.sessions,
        kill: '终止',
        confirmKill: '是否终止该会话？'
      },
      audit: {
        ...enUs.page.user.audit,
        logsSearchPlaceholder: '用户名/操作/RustdeskID/IP'
      },
      devices: {
        ...enUs.page.user.devices,
        logsSearchPlaceholder: '用户名/主机名/RustdeskID'
      }
    },
    system: {
      ...enUs.page.system,
      mailTemplate: {
        ...enUs.page.system.mailTemplate,
        addMailTemplate: '添加模板',
        editMailTemplate: '编辑模板',
        inputName: '请输入名称',
        inputSubject: '请输入主题',
        inputContents: '请输入内容',
        selectType: '请选择类型'
      },
      mailLog: {
        ...enUs.page.system.mailLog,
        info: '详情'
      }
    }
  },
  dropdown: {
    ...enUs.dropdown,
    closeCurrent: '关闭当前',
    closeOther: '关闭其他',
    closeLeft: '关闭左侧',
    closeRight: '关闭右侧',
    closeAll: '关闭全部'
  },
  icon: {
    ...enUs.icon,
    themeConfig: '主题配置',
    themeSchema: '主题方案',
    lang: '切换语言',
    fullscreen: '全屏',
    fullscreenExit: '退出全屏',
    reload: '刷新页面',
    collapse: '收起菜单',
    expand: '展开菜单',
    pin: '固定',
    unpin: '取消固定'
  },
  datatable: {
    ...enUs.datatable,
    itemCount: '共 {total} 条'
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
      ip: 'IP',
      session_id: '会话ID',
      uuid: 'UUID',
      created_at: '创建时间',
      closed_at: '结束时间',
      peer_id: '对端ID',
      path: '路径',
      typeLabel: {
        ...enUs.dataMap.audit.typeLabel,
        remote_control: '远程控制',
        file_transfer: '文件传输',
        tcp_tunnel: 'TCP隧道'
      },
      fileTransferTypeLabel: {
        ...enUs.dataMap.audit.fileTransferTypeLabel,
        master_controlled: '控制端 -> 被控端',
        controlled_master: '被控端 -> 控制端'
      }
    },
    mailTemplate: {
      ...enUs.dataMap.mailTemplate,
      name: '名称',
      type: '类型',
      subject: '主题',
      contents: '内容',
      created_at: '创建时间',
      typeLabel: {
        ...enUs.dataMap.mailTemplate.typeLabel,
        loginVerify: '登录验证',
        registerVerify: '注册验证',
        other: '其他'
      }
    },
    mailLog: {
      ...enUs.dataMap.mailLog,
      username: '用户名',
      uuid: 'UUID',
      from: '发件人',
      to: '收件人',
      subject: '主题',
      contents: '内容',
      status: '状态',
      created_at: '发送时间',
      statusLabel: {
        ...enUs.dataMap.mailLog.statusLabel,
        ok: '成功',
        err: '失败'
      }
    }
  },
  api: {
    ...enUs.api,
    CaptchaError: '验证码错误',
    UserNotExists: '用户不存在',
    UsernameOrPasswordError: '账号或密码错误',
    UserExists: '用户名已被使用',
    UsernameEmpty: '用户名不能为空',
    PasswordEmpty: '密码不能为空',
    UserAddSuccess: '用户创建成功',
    DataError: '数据错误',
    UserUpdateSuccess: '用户修改成功',
    UserDeleteSuccess: '用户删除成功',
    SessionKillSuccess: '会话终止成功',
    MailTemplateNameEmpty: '模板名称不能为空',
    MailTemplateSubjectEmpty: '模板主题不能为空',
    MailTemplateContentsEmpty: '模板内容不能为空',
    MailTemplateAddSuccess: '邮件模板创建成功',
    MailTemplateUpdateSuccess: '邮件模板修改成功',
    NoEmailAddress: '未设置邮箱地址',
    VerificationCodeError: '验证码错误',
    UUIDEmpty: 'UUID 不能为空'
  }
};

export default local;
