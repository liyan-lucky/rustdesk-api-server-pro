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
    cancel: '取消',
    close: '关闭',
    confirm: '确认',
    delete: '删除',
    edit: '编辑',
    error: '错误',
    refresh: '刷新',
    reset: '重置',
    search: '搜索',
    warning: '警告',
    logout: '退出登录',
    logoutConfirm: '确认退出登录吗？',
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
      deviceCount: '主机数量',
      onlineCount: '在线数量',
      visitsCount: '访问次数',
      operatingSystem: '操作系统',
      oneWeek: '一周内',
      changeLogs: '更新日志'
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
    device: {
      ...enUs.dataMap.device,
      username: '用户名',
      hostname: '主机名',
      version: 'RustDesk版本',
      memory: '内存',
      os: '操作系统',
      rustdesk_id: 'RustdeskID'
    }
  }
};

export default local;
