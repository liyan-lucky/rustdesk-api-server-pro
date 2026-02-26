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
    check: '勾选',
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
  route: {
    ...enUs.route,
    login: '登录',
    403: '无权限',
    404: '页面不存在',
    500: '服务器错误',
    'iframe-page': '外链页面',
    home: '首页',
    audit: '日志审计',
    user: '用户管理',
    user_list: '用户列表',
    user_sessions: '会话管理',
    system: '系统管理',
    system_mail_template: '邮件模板',
    system_mail_logs: '邮件日志',
    system_mail: '邮件管理',
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
        welcomeBack: '欢迎回来，{userName}！'
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
          check: '检测连通性',
          checked: '连通性检测完成',
          checkedCached: '使用最近一次检测结果（缓存）',
          checkFailed: '连通性检测失败',
          target: '目标',
          notChecked: '尚未检测',
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
  icon: {
    ...enUs.icon,
    themeConfig: '主题配置',
    themeSchema: '主题模式',
    lang: '切换语言',
    fullscreen: '全屏',
    fullscreenExit: '退出全屏',
    reload: '刷新页面',
    collapse: '折叠菜单',
    expand: '展开菜单',
    pin: '固定',
    unpin: '取消固定'
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
