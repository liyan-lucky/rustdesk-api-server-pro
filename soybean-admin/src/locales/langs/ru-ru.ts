const local: App.I18n.Schema = {
  system: {
    title: 'Rustdesk Api Server',
    updateTitle: '╨₧╨▒╨╜╨╛╨▓╨╗╨╡╨╜╨╕╨╡ ╨▓╨╡╤Ç╤ü╨╕╨╕ ╤ü╨╕╤ü╤é╨╡╨╝╤ï',
    updateContent:
      '╨ö╨╛╤ü╤é╤â╨┐╨╜╨░ ╨╜╨╛╨▓╨░╤Å ╨▓╨╡╤Ç╤ü╨╕╤Å ╤ü╨╕╤ü╤é╨╡╨╝╤ï. ╨₧╨▒╨╜╨╛╨▓╨╕╤é╤î ╤ü╤é╤Ç╨░╨╜╨╕╤å╤â ╤ü╨╡╨╣╤ç╨░╤ü?',
    updateConfirm: '╨₧╨▒╨╜╨╛╨▓╨╕╤é╤î ╤ü╨╡╨╣╤ç╨░╤ü',
    updateCancel: '╨ƒ╨╛╨╖╨╢╨╡'
  },
  common: {
    action: 'Action',
    add: 'Add',
    addSuccess: 'Add Success',
    backToHome: 'Back to home',
    batchDelete: 'Batch Delete',
    cancel: 'Cancel',
    close: 'Close',
    check: 'Check',
    expandColumn: 'Expand Column',
    columnSetting: 'Column Setting',
    config: 'Config',
    confirm: 'Confirm',
    delete: 'Delete',
    deleteSuccess: 'Delete Success',
    confirmDelete: 'Are you sure you want to delete?',
    edit: 'Edit',
    import: 'CSV Import',
    export: 'CSV Export',
    look: 'Look',
    warning: 'Warning',
    error: 'Error',
    index: 'Index',
    keywordSearch: 'Please enter keyword',
    logout: 'Logout',
    logoutConfirm: 'Are you sure you want to log out?',
    lookForward: 'Coming soon',
    modify: 'Modify',
    modifySuccess: 'Modify Success',
    noData: 'No Data',
    operate: 'Operate',
    pleaseCheckValue: 'Please check whether the value is valid',
    refresh: 'Refresh',
    reset: 'Reset',
    search: 'Search',
    switch: 'Switch',
    tip: 'Tip',
    trigger: 'Trigger',
    update: '╨₧╨▒╨╜╨╛╨▓╨╕╤é╤î',
    updateSuccess: '╨₧╨▒╨╜╨╛╨▓╨╗╨╡╨╜╨╕╨╡ ╨▓╤ï╨┐╨╛╨╗╨╜╨╡╨╜╨╛',
    userCenter: 'User Center',
    yesOrNo: { yes: 'Yes', no: 'No' }
  },
  request: {
    logout: 'Logout user after request failed',
    logoutMsg: 'User status is invalid, please log in again',
    logoutWithModal: 'Pop up modal after request failed and then log out user',
    logoutWithModalMsg: 'User status is invalid, please log in again',
    refreshToken: 'The requested token has expired, refresh the token',
    tokenExpired: 'The requested token has expired'
  },
  theme: {
    themeSchema: { title: 'Theme Schema', light: 'Light', dark: 'Dark', auto: 'Follow System' },
    grayscale: 'Grayscale',
    colourWeakness: 'Colour Weakness',
    layoutMode: {
      title: 'Layout Mode',
      vertical: 'Vertical Menu Mode',
      horizontal: 'Horizontal Menu Mode',
      'vertical-mix': 'Vertical Mix Menu Mode',
      'horizontal-mix': 'Horizontal Mix menu Mode',
      reverseHorizontalMix: 'Reverse first level menus and child level menus position'
    },
    recommendColor: 'Apply Recommended Color Algorithm',
    recommendColorDesc: 'The recommended color algorithm refers to',
    themeColor: {
      title: 'Theme Color',
      primary: 'Primary',
      info: 'Info',
      success: 'Success',
      warning: 'Warning',
      error: 'Error',
      followPrimary: 'Follow Primary'
    },
    scrollMode: { title: 'Scroll Mode', wrapper: 'Wrapper', content: 'Content' },
    page: {
      animate: 'Page Animate',
      mode: {
        title: 'Page Animate Mode',
        fade: 'Fade',
        'fade-slide': 'Slide',
        'fade-bottom': 'Fade Zoom',
        'fade-scale': 'Fade Scale',
        'zoom-fade': 'Zoom Fade',
        'zoom-out': 'Zoom Out',
        none: 'None'
      }
    },
    fixedHeaderAndTab: 'Fixed Header And Tab',
    header: {
      height: 'Header Height',
      breadcrumb: { visible: 'Breadcrumb Visible', showIcon: 'Breadcrumb Icon Visible' }
    },
    tab: {
      visible: 'Tab Visible',
      cache: 'Tab Cache',
      height: 'Tab Height',
      mode: { title: 'Tab Mode', chrome: 'Chrome', button: 'Button' }
    },
    sider: {
      inverted: 'Dark Sider',
      width: 'Sider Width',
      collapsedWidth: 'Sider Collapsed Width',
      mixWidth: 'Mix Sider Width',
      mixCollapsedWidth: 'Mix Sider Collapse Width',
      mixChildMenuWidth: 'Mix Child Menu Width'
    },
    footer: { visible: 'Footer Visible', fixed: 'Fixed Footer', height: 'Footer Height', right: 'Right Footer' },
    watermark: { visible: 'Watermark Full Screen Visible', text: 'Watermark Text' },
    themeDrawerTitle: 'Theme Configuration',
    pageFunTitle: 'Page Function',
    configOperation: {
      copyConfig: 'Copy Config',
      copySuccessMsg: 'Copy Success, Please replace the variable "themeSettings" in "src/theme/settings.ts"',
      resetConfig: 'Reset Config',
      resetSuccessMsg: 'Reset Success'
    }
  },
  route: {
    '403': 'No Permission',
    '404': 'Page Not Found',
    '500': 'Server Error',
    login: '╨Æ╤à╨╛╨┤',
    'iframe-page': '╨Æ╤ü╤é╤Ç╨╛╨╡╨╜╨╜╨░╤Å ╤ü╤é╤Ç╨░╨╜╨╕╤å╨░',
    home: '╨ô╨╗╨░╨▓╨╜╨░╤Å',
    about: 'О программе',
    about_version: 'Информация о версии',
    about_help: 'Справка по кодам ошибок',
    audit: '╨É╤â╨┤╨╕╤é',
    user: '╨ú╨┐╤Ç╨░╨▓╨╗╨╡╨╜╨╕╨╡ ╨┐╨╛╨╗╤î╨╖╨╛╨▓╨░╤é╨╡╨╗╤Å╨╝╨╕',
    user_list: '╨í╨┐╨╕╤ü╨╛╨║ ╨┐╨╛╨╗╤î╨╖╨╛╨▓╨░╤é╨╡╨╗╨╡╨╣',
    user_sessions: '╨í╨╡╤ü╤ü╨╕╨╕',
    user_profile: '╨ƒ╤Ç╨╛╤ä╨╕╨╗╤î',
    system: '╨ú╨┐╤Ç╨░╨▓╨╗╨╡╨╜╨╕╨╡ ╤ü╨╕╤ü╤é╨╡╨╝╨╛╨╣',
    system_mail_template: '╨¿╨░╨▒╨╗╨╛╨╜╤ï ╨┐╨╕╤ü╨╡╨╝',
    system_mail_logs: '╨¢╨╛╨│╨╕ ╨┐╨╛╤ç╤é╤ï',
    system_mail: '╨ú╨┐╤Ç╨░╨▓╨╗╨╡╨╜╨╕╨╡ ╨┐╨╛╤ç╤é╨╛╨╣',
    system_server: '╨Ü╨╛╨╜╤ä╨╕╨│╤â╤Ç╨░╤å╨╕╤Å ╤ü╨╡╤Ç╨▓╨╡╤Ç╨░',
    system_tokens: '╨ó╨╛╨║╨╡╨╜╤ï ╨┐╨╛╨╗╤î╨╖╨╛╨▓╨░╤é╨╡╨╗╨╡╨╣',
    system_oauth: '╨í╤é╨╛╤Ç╨╛╨╜╨╜╨╕╨╣ ╨▓╤à╨╛╨┤',
    audit_baselogs: '╨æ╨░╨╖╨╛╨▓╤ï╨╡ ╨╗╨╛╨│╨╕',
    audit_filetransferlogs: '╨¢╨╛╨│╨╕ ╨┐╨╡╤Ç╨╡╨┤╨░╤ç╨╕ ╤ä╨░╨╣╨╗╨╛╨▓',
    audit_loginlogs: '╨¢╨╛╨│╨╕ ╨▓╤à╨╛╨┤╨░',
    'audit_error-logs': 'Журналы ошибок',
    devices: '╨ú╤ü╤é╤Ç╨╛╨╣╤ü╤é╨▓╨░',
    'my-devices': '╨Ü╨╛╨╜╤é╨░╨║╤é╤ï',
    'my-devices_peers': '╨£╨╛╨╕ ╨┐╨╕╤Ç╤ï',
    'my-devices_manage': '╨ú╨┐╤Ç╨░╨▓╨╗╨╡╨╜╨╕╨╡ ╨░╨┤╤Ç╨╡╤ü╨╜╨╛╨╣ ╨║╨╜╨╕╨│╨╛╨╣',
    'my-devices_tags': '╨ú╨┐╤Ç╨░╨▓╨╗╨╡╨╜╨╕╨╡ ╤é╨╡╨│╨░╨╝╨╕',
    workspace: '╨£╨╛╤æ ╨┐╤Ç╨╛╤ü╤é╤Ç╨░╨╜╤ü╤é╨▓╨╛',
    workspace_overview: '╨₧╨▒╨╖╨╛╤Ç',
    workspace_devices: '╨£╨╛╨╕ ╤â╤ü╤é╤Ç╨╛╨╣╤ü╤é╨▓╨░',
    workspace_sessions: '╨í╨╡╨░╨╜╤ü╤ï ╨▓╤à╨╛╨┤╨░',
    workspace_security: '╨í╨╛╨▒╤ï╤é╨╕╤Å ╨▒╨╡╨╖╨╛╨┐╨░╤ü╨╜╨╛╤ü╤é╨╕',
    workspace_profile: '╨ƒ╤Ç╨╛╤ä╨╕╨╗╤î'
  },
  page: {
    login: {
      common: {
        loginOrRegister: '╨Æ╤à╨╛╨┤ / ╨á╨╡╨│╨╕╤ü╤é╤Ç╨░╤å╨╕╤Å',
        userNamePlaceholder: '╨Æ╨▓╨╡╨┤╨╕╤é╨╡ ╨╕╨╝╤Å ╨┐╨╛╨╗╤î╨╖╨╛╨▓╨░╤é╨╡╨╗╤Å',
        phonePlaceholder: 'Please enter phone number',
        codePlaceholder: '╨Æ╨▓╨╡╨┤╨╕╤é╨╡ ╨║╨╛╨┤ ╨┐╨╛╨┤╤é╨▓╨╡╤Ç╨╢╨┤╨╡╨╜╨╕╤Å',
        passwordPlaceholder: '╨Æ╨▓╨╡╨┤╨╕╤é╨╡ ╨┐╨░╤Ç╨╛╨╗╤î',
        confirmPasswordPlaceholder: '╨Æ╨▓╨╡╨┤╨╕╤é╨╡ ╨┐╨░╤Ç╨╛╨╗╤î ╨╡╤ë╨╡ ╤Ç╨░╨╖',
        codeLogin: '╨Æ╤à╨╛╨┤ ╨┐╨╛ ╨║╨╛╨┤╤â ╨┐╨╛╨┤╤é╨▓╨╡╤Ç╨╢╨┤╨╡╨╜╨╕╤Å',
        confirm: '╨ƒ╨╛╨┤╤é╨▓╨╡╤Ç╨┤╨╕╤é╤î',
        back: '╨¥╨░╨╖╨░╨┤',
        validateSuccess: '╨ƒ╤Ç╨╛╨▓╨╡╤Ç╨║╨░ ╨┐╤Ç╨╛╨╣╨┤╨╡╨╜╨░',
        loginSuccess: '╨Æ╤à╨╛╨┤ ╨▓╤ï╨┐╨╛╨╗╨╜╨╡╨╜ ╤â╤ü╨┐╨╡╤ê╨╜╨╛',
        welcomeBack: '╨í ╨▓╨╛╨╖╨▓╤Ç╨░╤ë╨╡╨╜╨╕╨╡╨╝, {userName}!',
        thirdPartyLogin: '╨Æ╤à╨╛╨┤ ╤ç╨╡╤Ç╨╡╨╖ ╤ü╤é╨╛╤Ç╨╛╨╜╨╜╨╕╨╣ ╤ü╨╡╤Ç╨▓╨╕╤ü',
        continueWith: '╨ƒ╤Ç╨╛╨┤╨╛╨╗╨╢╨╕╤é╤î ╤ç╨╡╤Ç╨╡╨╖ {provider}',
        providerUnavailable: '╨Æ╤à╨╛╨┤ ╤ç╨╡╤Ç╨╡╨╖ {provider} ╤ü╨╡╨╣╤ç╨░╤ü ╨╜╨╡╨┤╨╛╤ü╤é╤â╨┐╨╡╨╜',
        oauthAccountNotBound:
          '╨ƒ╨╛╨┤╤à╨╛╨┤╤Å╤ë╤â╤Ä ╤â╤ç╨╡╤é╨╜╤â╤Ä ╨╖╨░╨┐╨╕╤ü╤î ╨┐╤Ç╨╕╨▓╤Å╨╖╨░╤é╤î ╨╜╨╡╨╗╤î╨╖╤Å. ╨ú╨║╨░╨╢╨╕╤é╨╡ ╤é╤â ╨╢╨╡ ╨┐╨╛╨┤╤é╨▓╨╡╤Ç╨╢╨┤╨╡╨╜╨╜╤â╤Ä ╨┐╨╛╤ç╤é╤â ╨╕╨╗╨╕ ╨▓╨║╨╗╤Ä╤ç╨╕╤é╨╡ ╨░╨▓╤é╨╛╨╝╨░╤é╨╕╤ç╨╡╤ü╨║╨╛╨╡ ╤ü╨╛╨╖╨┤╨░╨╜╨╕╨╡.',
        oauthProviderUnreachable:
          '╨í╨╡╤Ç╨▓╨╡╤Ç ╨╜╨╡ ╨╝╨╛╨╢╨╡╤é ╨┐╨╛╨┤╨║╨╗╤Ä╤ç╨╕╤é╤î╤ü╤Å ╨║ ╨┐╤Ç╨╛╨▓╨░╨╣╨┤╨╡╤Ç╤â. ╨ƒ╤Ç╨╛╨▓╨╡╤Ç╤î╤é╨╡ ╨╕╤ü╤à╨╛╨┤╤Å╤ë╨╡╨╡ HTTPS-╤ü╨╛╨╡╨┤╨╕╨╜╨╡╨╜╨╕╨╡.',
        oauthStateExpired:
          '╨ù╨░╨┐╤Ç╨╛╤ü ╨▓╤à╨╛╨┤╨░ ╨╕╤ü╤é╨╡╨║ ╨╕╨╗╨╕ ╤â╨╢╨╡ ╨╕╤ü╨┐╨╛╨╗╤î╨╖╨╛╨▓╨░╨╜. ╨¥╨░╤ç╨╜╨╕╤é╨╡ ╨▓╤à╨╛╨┤ ╨╖╨░╨╜╨╛╨▓╨╛.',
        oauthAuthFailed:
          '╨₧╤ê╨╕╨▒╨║╨░ ╤ü╤é╨╛╤Ç╨╛╨╜╨╜╨╡╨│╨╛ ╨▓╤à╨╛╨┤╨░. ╨ƒ╤Ç╨╛╨▓╨╡╤Ç╤î╤é╨╡ ╨╜╨░╤ü╤é╤Ç╨╛╨╣╨║╨╕ ╨╕ ╨╢╤â╤Ç╨╜╨░╨╗ ╨▒╨╡╨╖╨╛╨┐╨░╤ü╨╜╨╛╤ü╤é╨╕.'
      },
      pwdLogin: {
        title: '╨Æ╤à╨╛╨┤ ╨┐╨╛ ╨┐╨░╤Ç╨╛╨╗╤Ä',
        rememberMe: '╨ù╨░╨┐╨╛╨╝╨╜╨╕╤é╤î ╨╝╨╡╨╜╤Å',
        switchToUser: '╨Æ╤à╨╛╨┤ ╨┐╨╛╨╗╤î╨╖╨╛╨▓╨░╤é╨╡╨╗╤Å'
      },
      userLogin: { title: '╨Æ╤à╨╛╨┤ ╨┐╨╛╨╗╤î╨╖╨╛╨▓╨░╤é╨╡╨╗╤Å', switchToAdmin: '╨Æ╤à╨╛╨┤ ╨░╨┤╨╝╨╕╨╜╨╕╤ü╤é╤Ç╨░╤é╨╛╤Ç╨░' }
    },
    home: {
      greeting: '╨ö╨╛╨▒╤Ç╨╛╨╡ ╤â╤é╤Ç╨╛, {userName}, ╤ü╨╡╨│╨╛╨┤╨╜╤Å ╨╛╤é╨╗╨╕╤ç╨╜╤ï╨╣ ╨┤╨╡╨╜╤î!',
      userCount: '╨ƒ╨╛╨╗╤î╨╖╨╛╨▓╨░╤é╨╡╨╗╨╕',
      deviceCount: '╨ú╤ü╤é╤Ç╨╛╨╣╤ü╤é╨▓╨░',
      onlineCount: '╨₧╨╜╨╗╨░╨╣╨╜',
      visitsCount: '╨ƒ╨╛╤ü╨╡╤ë╨╡╨╜╨╕╤Å',
      operatingSystem: '╨₧╨┐╨╡╤Ç╨░╤å╨╕╨╛╨╜╨╜╤ï╨╡ ╤ü╨╕╤ü╤é╨╡╨╝╤ï',
      oneWeek: '╨ù╨░ ╨╜╨╡╨┤╨╡╨╗╤Ä',
      changeLogs: '╨û╤â╤Ç╨╜╨░╨╗ ╨╕╨╖╨╝╨╡╨╜╨╡╨╜╨╕╨╣',
      cardDetail: {
        viewHint: '╨¥╨░╨╢╨╝╨╕╤é╨╡, ╤ç╤é╨╛╨▒╤ï ╨┐╨╛╤ü╨╝╨╛╤é╤Ç╨╡╤é╤î ╨┤╨╡╤é╨░╨╗╨╕',
        recentUsers: '╨¥╨╡╨┤╨░╨▓╨╜╨╕╨╡ ╨┐╨╛╨╗╤î╨╖╨╛╨▓╨░╤é╨╡╨╗╨╕',
        recentDevices: '╨¥╨╡╨┤╨░╨▓╨╜╨╕╨╡ ╤â╤ü╤é╤Ç╨╛╨╣╤ü╤é╨▓╨░',
        recentVisits: '╨¥╨╡╨┤╨░╨▓╨╜╨╕╨╡ ╨╖╨░╨┐╨╕╤ü╨╕ ╨┐╨╛╤ü╨╡╤ë╨╡╨╜╨╕╨╣',
        desc: {
          userCount:
            '╨ƒ╨╛╨║╨░╨╖╤ï╨▓╨░╨╡╤é ╨╛╨▒╤ë╨╡╨╡ ╨║╨╛╨╗╨╕╤ç╨╡╤ü╤é╨▓╨╛ ╨┐╨╛╨╗╤î╨╖╨╛╨▓╨░╤é╨╡╨╗╨╡╨╣ ╨▓ ╤ü╨╕╤ü╤é╨╡╨╝╨╡.',
          deviceCount: '╨ƒ╨╛╨║╨░╨╖╤ï╨▓╨░╨╡╤é ╨╛╨▒╤ë╨╡╨╡ ╨║╨╛╨╗╨╕╤ç╨╡╤ü╤é╨▓╨╛ ╤â╤ü╤é╤Ç╨╛╨╣╤ü╤é╨▓ ╨▓ ╤ü╨╕╤ü╤é╨╡╨╝╨╡.',
          onlineCount:
            '╨ƒ╨╛╨║╨░╨╖╤ï╨▓╨░╨╡╤é ╤ç╨╕╤ü╨╗╨╛ ╨╛╨╜╨╗╨░╨╣╨╜-╤â╤ü╤é╤Ç╨╛╨╣╤ü╤é╨▓ ╨┐╨╛ ╤ü╤é╨░╤é╨╕╤ü╤é╨╕╨║╨╡ heartbeat.',
          visitCount: '╨ƒ╨╛╨║╨░╨╖╤ï╨▓╨░╨╡╤é ╤ü╤é╨░╤é╨╕╤ü╤é╨╕╨║╤â ╨┐╨╛╤ü╨╡╤ë╨╡╨╜╨╕╨╣ ╨╕╨╖ ╨╢╤â╤Ç╨╜╨░╨╗╨╛╨▓ ╨░╤â╨┤╨╕╤é╨░.'
        }
      },
      serverConfig: {
        title: '╨Ü╨╛╨╜╤ä╨╕╨│╤â╤Ç╨░╤å╨╕╤Å ╨┐╨╛╨┤╨║╨╗╤Ä╤ç╨╡╨╜╨╕╤Å ╨║╨╗╨╕╨╡╨╜╤é╨░',
        tip: '╨í╨║╨╛╨┐╨╕╤Ç╤â╨╣╤é╨╡ ╨╖╨╜╨░╤ç╨╡╨╜╨╕╤Å ╨╜╨╕╨╢╨╡ ╨▓ ╨║╨╗╨╕╨╡╨╜╤é RustDesk. ╨ò╤ü╨╗╨╕ KEY ╨┐╤â╤ü╤é╨╛╨╣, ╨╖╨░╨┤╨░╨╣╤é╨╡ ╨┐╨╡╤Ç╨╡╨╝╨╡╨╜╨╜╤â╤Ä ╨╛╨║╤Ç╤â╨╢╨╡╨╜╨╕╤Å `RUSTDESK_KEY`.',
        idServer: 'ID ╤ü╨╡╤Ç╨▓╨╡╤Ç',
        relayServer: '╨á╨╡╨╗╨╡╨╣╨╜╤ï╨╣ ╤ü╨╡╤Ç╨▓╨╡╤Ç',
        apiServer: 'API ╤ü╨╡╤Ç╨▓╨╡╤Ç',
        key: 'KEY',
        idServerPlaceholder: '╨╜╨░╨┐╤Ç╨╕╨╝╨╡╤Ç your.domain.com',
        relayServerPlaceholder: '╨╜╨░╨┐╤Ç╨╕╨╝╨╡╤Ç your.domain.com',
        apiServerPlaceholder: '╨╜╨░╨┐╤Ç╨╕╨╝╨╡╤Ç https://your.domain.com',
        keyPlaceholder: '╨ú╨║╨░╨╢╨╕╤é╨╡ ╤ç╨╡╤Ç╨╡╨╖ ╨┐╨╡╤Ç╨╡╨╝╨╡╨╜╨╜╤â╤Ä ╨╛╨║╤Ç╤â╨╢╨╡╨╜╨╕╤Å RUSTDESK_KEY',
        copy: 'Copy',
        copyAll: '╨Ü╨╛╨┐╨╕╤Ç╨╛╨▓╨░╤é╤î ╨▓╤ü╨╡',
        copyTemplate: '╨Ü╨╛╨┐╨╕╤Ç╨╛╨▓╨░╤é╤î ╤ê╨░╨▒╨╗╨╛╨╜ RustDesk',
        showQr: '╨ƒ╨╛╨║╨░╨╖╨░╤é╤î QR-╨║╨╛╨┤',
        qrTitle: 'QR-╨║╨╛╨┤ ╨╕╨╝╨┐╨╛╤Ç╤é╨░ RustDesk',
        qrTip:
          '╨₧╤é╤ü╨║╨░╨╜╨╕╤Ç╤â╨╣╤é╨╡ ╤ì╤é╨╛╤é QR-╨║╨╛╨┤ ╨▓ ╨╝╨╛╨▒╨╕╨╗╤î╨╜╨╛╨╝ ╨┐╤Ç╨╕╨╗╨╛╨╢╨╡╨╜╨╕╨╕ RustDesk ╨┤╨╗╤Å ╨╕╨╝╨┐╨╛╤Ç╤é╨░ ╨║╨╛╨╜╤ä╨╕╨│╤â╤Ç╨░╤å╨╕╨╕.',
        qrPayload: '╨ó╨╡╨║╤ü╤é ╤ê╨░╨▒╨╗╨╛╨╜╨░ RustDesk',
        qrFailed: '╨₧╤ê╨╕╨▒╨║╨░ ╨│╨╡╨╜╨╡╤Ç╨░╤å╨╕╨╕ QR-╨║╨╛╨┤╨░',
        refresh: '╨₧╨▒╨╜╨╛╨▓╨╕╤é╤î',
        clearCacheReload: '╨₧╤ç╨╕╤ü╤é╨╕╤é╤î ╨║╤ì╤ê ╨╕ ╨┐╨╡╤Ç╨╡╨╖╨░╨│╤Ç╤â╨╖╨╕╤é╤î',
        source: '╨ÿ╤ü╤é╨╛╤ç╨╜╨╕╨║',
        lastUpdated: '╨ƒ╨╛╤ü╨╗╨╡╨┤╨╜╨╡╨╡ ╨╛╨▒╨╜╨╛╨▓╨╗╨╡╨╜╨╕╨╡',
        show: '╨ƒ╨╛╨║╨░╨╖╨░╤é╤î',
        hide: '╨í╨║╤Ç╤ï╤é╤î',
        missingTip:
          '╨í╨╗╨╡╨┤╤â╤Ä╤ë╨╕╨╡ ╨┐╨╛╨╗╤Å ╨┐╤â╤ü╤é╤ï╨╡, ╤ü╨╜╨░╤ç╨░╨╗╨░ ╨╜╨░╤ü╤é╤Ç╨╛╨╣╤é╨╡ ╨╕╤à ╨▓ ╨┐╨╡╤Ç╨╡╨╝╨╡╨╜╨╜╤ï╤à ╨╛╨║╤Ç╤â╨╢╨╡╨╜╨╕╤Å ╨║╨╛╨╜╤é╨╡╨╣╨╜╨╡╤Ç╨░: {fields}',
        copyEmpty: '{label} ╨┐╤â╤ü╤é╨╛, ╨║╨╛╨┐╨╕╤Ç╨╛╨▓╨░╨╜╨╕╨╡ ╨╜╨╡╨▓╨╛╨╖╨╝╨╛╨╢╨╜╨╛',
        copySuccess: '{label} ╤ü╨║╨╛╨┐╨╕╤Ç╨╛╨▓╨░╨╜╨╛',
        copyFailed: '╨¥╨╡ ╤â╨┤╨░╨╗╨╛╤ü╤î ╤ü╨║╨╛╨┐╨╕╤Ç╨╛╨▓╨░╤é╤î {label}',
        fetchFailed: '╨¥╨╡ ╤â╨┤╨░╨╗╨╛╤ü╤î ╨╖╨░╨│╤Ç╤â╨╖╨╕╤é╤î ╨║╨╛╨╜╤ä╨╕╨│╤â╤Ç╨░╤å╨╕╤Ä ╤ü╨╡╤Ç╨▓╨╡╤Ç╨░',
        cacheCleared:
          '╨Ü╤ì╤ê ╨╛╤ç╨╕╤ë╨╡╨╜, ╨┐╨╛╨▓╤é╨╛╤Ç╨╜╨░╤Å ╨╖╨░╨│╤Ç╤â╨╖╨║╨░ ╨║╨╛╨╜╤ä╨╕╨│╤â╤Ç╨░╤å╨╕╨╕ ╤ü╨╡╤Ç╨▓╨╡╤Ç╨░',
        cacheTtlHint: 'TTL cache: config {configSeconds}s, connectivity {connectivitySeconds}s',
        ageSeconds: '{seconds}s ago',
        sourceType: {
          remote: '╨ú╨┤╨░╨╗╤æ╨╜╨╜╤ï╨╣ ╨╕╤ü╤é╨╛╤ç╨╜╨╕╨║',
          'memory-cache': '╨Ü╤ì╤ê ╨┐╨░╨╝╤Å╤é╨╕',
          'session-cache': '╨Ü╤ì╤ê ╤ü╨╡╤ü╤ü╨╕╨╕',
          env: '╨ƒ╨╡╤Ç╨╡╨╝╨╡╨╜╨╜╨░╤Å ╨╛╨║╤Ç╤â╨╢╨╡╨╜╨╕╤Å',
          inferred: '╨É╨▓╤é╨╛╨╛╨┐╤Ç╨╡╨┤╨╡╨╗╨╡╨╜╨╕╨╡',
          empty: '╨ƒ╤â╤ü╤é╨╛',
          auto: '╨É╨▓╤é╨╛╨╛╨▒╨╜╨░╤Ç╤â╨╢╨╡╨╜╨╕╨╡'
        },
        sourceHint: {
          env: '╨¡╤é╨╛ ╨╖╨╜╨░╤ç╨╡╨╜╨╕╨╡ ╨┐╨╛╨╗╤â╤ç╨╡╨╜╨╛ ╨╕╨╖ ╨┐╨╡╤Ç╨╡╨╝╨╡╨╜╨╜╨╛╨╣ ╨╛╨║╤Ç╤â╨╢╨╡╨╜╨╕╤Å ╨║╨╛╨╜╤é╨╡╨╣╨╜╨╡╤Ç╨░.',
          inferred:
            '╨¡╤é╨╛ ╨╖╨╜╨░╤ç╨╡╨╜╨╕╨╡ ╨░╨▓╤é╨╛╨╝╨░╤é╨╕╤ç╨╡╤ü╨║╨╕ ╨╛╨┐╤Ç╨╡╨┤╨╡╨╗╨╡╨╜╨╛ ╨┐╨╛ ╤é╨╡╨║╤â╤ë╨╡╨╝╤â ╨░╨┤╤Ç╨╡╤ü╤â ╨┤╨╛╤ü╤é╤â╨┐╨░.',
          empty: '╨ù╨╜╨░╤ç╨╡╨╜╨╕╨╡ ╨╜╨╡ ╨╜╨░╤ü╤é╤Ç╨╛╨╡╨╜╨╛ ╨╕ ╨╜╨╡ ╨╛╨┐╤Ç╨╡╨┤╨╡╨╗╨╡╨╜╨╛ ╨░╨▓╤é╨╛╨╝╨░╤é╨╕╤ç╨╡╤ü╨║╨╕.'
        },
        connectivity: {
          clear: '╨₧╤ç╨╕╤ü╤é╨╕╤é╤î ╤Ç╨╡╨╖╤â╨╗╤î╤é╨░╤é╤ï',
          check: '╨ƒ╤Ç╨╛╨▓╨╡╤Ç╨╕╤é╤î ╨┤╨╛╤ü╤é╤â╨┐╨╜╨╛╤ü╤é╤î',
          checkOne: '╨ƒ╤Ç╨╛╨▓╨╡╤Ç╨╕╤é╤î',
          checked: '╨ƒ╤Ç╨╛╨▓╨╡╤Ç╨║╨░ ╨┤╨╛╤ü╤é╤â╨┐╨╜╨╛╤ü╤é╨╕ ╨╖╨░╨▓╨╡╤Ç╤ê╨╡╨╜╨░',
          checkedOne: '╨ƒ╤Ç╨╛╨▓╨╡╤Ç╨║╨░ {field} ╨╖╨░╨▓╨╡╤Ç╤ê╨╡╨╜╨░',
          checkedCached: '╨ÿ╤ü╨┐╨╛╨╗╤î╨╖╨╛╨▓╨░╨╜ ╨╜╨╡╨┤╨░╨▓╨╜╨╕╨╣ ╤Ç╨╡╨╖╤â╨╗╤î╤é╨░╤é ╨┐╤Ç╨╛╨▓╨╡╤Ç╨║╨╕ (╨║╤ì╤ê)',
          checkFailed: '╨ƒ╤Ç╨╛╨▓╨╡╤Ç╨║╨░ ╨┤╨╛╤ü╤é╤â╨┐╨╜╨╛╤ü╤é╨╕ ╨╜╨╡ ╤â╨┤╨░╨╗╨░╤ü╤î',
          cleared: '╨á╨╡╨╖╤â╨╗╤î╤é╨░╤é╤ï ╨┐╤Ç╨╛╨▓╨╡╤Ç╨║╨╕ ╨╛╤ç╨╕╤ë╨╡╨╜╤ï',
          source: '╨ÿ╤ü╤é╨╛╤ç╨╜╨╕╨║ ╨┐╤Ç╨╛╨▓╨╡╤Ç╨║╨╕',
          lastChecked: '╨ƒ╨╛╤ü╨╗╨╡╨┤╨╜╤Å╤Å ╨┐╤Ç╨╛╨▓╨╡╤Ç╨║╨░',
          target: '╨ª╨╡╨╗╤î',
          duration: '╨Æ╤Ç╨╡╨╝╤Å',
          notChecked: '╨ò╤ë╤æ ╨╜╨╡ ╨┐╤Ç╨╛╨▓╨╡╤Ç╤Å╨╗╨╛╤ü╤î',
          checkSourceType: { remote: '╨ú╨┤╨░╨╗╤æ╨╜╨╜╨░╤Å ╨┐╤Ç╨╛╨▓╨╡╤Ç╨║╨░', cache: '╨Ü╤ì╤ê' },
          status: {
            idle: '╨¥╨╡ ╨┐╤Ç╨╛╨▓╨╡╤Ç╨╡╨╜╨╛',
            ok: '╨ö╨╛╤ü╤é╤â╨┐╨╜╨╛',
            error: '╨₧╤ê╨╕╨▒╨║╨░',
            skip: '╨ƒ╤Ç╨╛╨┐╤â╤ë╨╡╨╜╨╛'
          }
        }
      }
    },
    user: {
      list: {
        addUser: '╨ö╨╛╨▒╨░╨▓╨╕╤é╤î ╨┐╨╛╨╗╤î╨╖╨╛╨▓╨░╤é╨╡╨╗╤Å',
        editUser: '╨á╨╡╨┤╨░╨║╤é╨╕╤Ç╨╛╨▓╨░╤é╤î ╨┐╨╛╨╗╤î╨╖╨╛╨▓╨░╤é╨╡╨╗╤Å',
        inputUsername: '╨Æ╨▓╨╡╨┤╨╕╤é╨╡ ╨╕╨╝╤Å ╨┐╨╛╨╗╤î╨╖╨╛╨▓╨░╤é╨╡╨╗╤Å',
        inputPassword: '╨Æ╨▓╨╡╨┤╨╕╤é╨╡ ╨┐╨░╤Ç╨╛╨╗╤î',
        inputNickname: 'Input Nickname',
        emailFormatError: 'Email format error',
        selectUserStatus: 'Please select user status',
        searchPlaceholder: '╨ÿ╨╝╤Å ╨┐╨╛╨╗╤î╨╖╨╛╨▓╨░╤é╨╡╨╗╤Å\\╨¥╨╕╨║\\Email',
        tfa_secret_bind: '2FA Device Bind',
        require2FASecret: '2FA Secret Empty',
        require2FACode: "2FA Code Can't Empty"
      },
      sessions: { kill: '╨ù╨░╨▓╨╡╤Ç╤ê╨╕╤é╤î', confirmKill: 'Confirm Kill?' },
      audit: { logsSearchPlaceholder: '╨ƒ╨╛╨╗╤î╨╖╨╛╨▓╨░╤é╨╡╨╗╤î\\╨ö╨╡╨╣╤ü╤é╨▓╨╕╨╡\\RustdeskID\\IP' },
      devices: { logsSearchPlaceholder: '╨ƒ╨╛╨╗╤î╨╖╨╛╨▓╨░╤é╨╡╨╗╤î\\╨Ñ╨╛╤ü╤é\\RustdeskID' }
    },
    system: {
      mailTemplate: {
        addMailTemplate: '╨ö╨╛╨▒╨░╨▓╨╕╤é╤î ╤ê╨░╨▒╨╗╨╛╨╜',
        editMailTemplate: '╨á╨╡╨┤╨░╨║╤é╨╕╤Ç╨╛╨▓╨░╤é╤î ╤ê╨░╨▒╨╗╨╛╨╜',
        inputName: '╨Æ╨▓╨╡╨┤╨╕╤é╨╡ ╨╕╨╝╤Å',
        inputSubject: '╨Æ╨▓╨╡╨┤╨╕╤é╨╡ ╤é╨╡╨╝╤â',
        inputContents: '╨Æ╨▓╨╡╨┤╨╕╤é╨╡ ╤ü╨╛╨┤╨╡╤Ç╨╢╨╕╨╝╨╛╨╡',
        selectType: '╨Æ╤ï╨▒╨╡╤Ç╨╕╤é╨╡ ╤é╨╕╨┐'
      },
      mailLog: { info: '╨ƒ╨╛╨┤╤Ç╨╛╨▒╨╜╨╛╤ü╤é╨╕' }
    },
    myDevices: {
      title: '╨Ü╨╛╨╜╤é╨░╨║╤é╤ï',
      welcome: '╨ö╨╛╨▒╤Ç╨╛ ╨┐╨╛╨╢╨░╨╗╨╛╨▓╨░╤é╤î, {userName}',
      status: '╨í╤é╨░╤é╤â╤ü',
      online: '╨Æ ╤ü╨╡╤é╨╕',
      offline: '╨¥╨╡ ╨▓ ╤ü╨╡╤é╨╕',
      conns: '╨í╨╛╨╡╨┤╨╕╨╜╨╡╨╜╨╕╤Å',
      lastSync: '╨ƒ╨╛╤ü╨╗╨╡╨┤╨╜╤Å╤Å ╤ü╨╕╨╜╤à╤Ç╨╛╨╜╨╕╨╖╨░╤å╨╕╤Å',
      logout: '╨Æ╤ï╤à╨╛╨┤'
    },
    workspace: {
      scopeTitle: '╨¢╨╕╤ç╨╜╨╛╨╡ ╨┐╤Ç╨╛╤ü╤é╤Ç╨░╨╜╤ü╤é╨▓╨╛',
      scopeTip:
        '╨ù╨┤╨╡╤ü╤î ╨┐╨╛╨║╨░╨╖╨░╨╜╤ï ╤é╨╛╨╗╤î╨║╨╛ ╨▓╨░╤ê╨╕ ╤â╤ü╤é╤Ç╨╛╨╣╤ü╤é╨▓╨░, ╤ü╨╡╨░╨╜╤ü╤ï, ╤ü╨╛╨▒╤ï╤é╨╕╤Å ╨▒╨╡╨╖╨╛╨┐╨░╤ü╨╜╨╛╤ü╤é╨╕ ╨╕ ╤Ç╨░╨╖╤Ç╨╡╤ê╤æ╨╜╨╜╤ï╨╡ ╨░╨┤╤Ç╨╡╤ü╨╜╤ï╨╡ ╨║╨╜╨╕╨│╨╕.',
      myDevices: '╨£╨╛╨╕ ╤â╤ü╤é╤Ç╨╛╨╣╤ü╤é╨▓╨░',
      activeSessions: '╨É╨║╤é╨╕╨▓╨╜╤ï╨╡ ╤ü╨╡╨░╨╜╤ü╤ï',
      addressBooks: '╨É╨┤╤Ç╨╡╤ü╨╜╤ï╨╡ ╨║╨╜╨╕╨│╨╕',
      securityEvents: '╨í╨╛╨▒╤ï╤é╨╕╤Å ╨▒╨╡╨╖╨╛╨┐╨░╤ü╨╜╨╛╤ü╤é╨╕',
      currentSession: '╨ó╨╡╨║╤â╤ë╨╕╨╣ ╤ü╨╡╨░╨╜╤ü',
      revokeConfirm: '╨₧╤é╨╛╨╖╨▓╨░╤é╤î ╤ì╤é╨╛╤é ╤ü╨╡╨░╨╜╤ü ╨▓╤à╨╛╨┤╨░?',
      revoke: '╨₧╤é╨╛╨╖╨▓╨░╤é╤î',
      accountRole: '╨á╨╛╨╗╤î ╤â╤ç╤æ╤é╨╜╨╛╨╣ ╨╖╨░╨┐╨╕╤ü╨╕',
      adminRole: '╨É╨┤╨╝╨╕╨╜╨╕╤ü╤é╤Ç╨░╤é╨╛╤Ç',
      userRole: '╨ƒ╨╛╨╗╤î╨╖╨╛╨▓╨░╤é╨╡╨╗╤î',
      permissionScope: '╨₧╨▒╨╗╨░╤ü╤é╤î ╨┤╨╛╤ü╤é╤â╨┐╨░',
      userScope: '╨¢╨╕╤ç╨╜╤ï╨╡ ╤Ç╨╡╤ü╤â╤Ç╤ü╤ï ╨╕ ╤Å╨▓╨╜╨╛ ╨┐╤Ç╨╡╨┤╨╛╤ü╤é╨░╨▓╨╗╨╡╨╜╨╜╤ï╨╡ ╨░╨┤╤Ç╨╡╤ü╨╜╤ï╨╡ ╨║╨╜╨╕╨│╨╕',
      active: '╨É╨║╤é╨╕╨▓╨╡╨╜'
    },
    oauth: {
      configTitle: '╨¥╨░╤ü╤é╤Ç╨╛╨╣╨║╨░ ╤ü╤é╨╛╤Ç╨╛╨╜╨╜╨╡╨│╨╛ ╨▓╤à╨╛╨┤╨░',
      bindingsTitle: '╨ƒ╤Ç╨╕╨▓╤Å╨╖╨║╨╕ ╤â╤ç╨╡╤é╨╜╤ï╤à ╨╖╨░╨┐╨╕╤ü╨╡╨╣',
      addProvider: '╨ö╨╛╨▒╨░╨▓╨╕╤é╤î ╨┐╤Ç╨╛╨▓╨░╨╣╨┤╨╡╤Ç╨░',
      editProvider: '╨ÿ╨╖╨╝╨╡╨╜╨╕╤é╤î ╨┐╤Ç╨╛╨▓╨░╨╣╨┤╨╡╤Ç╨░',
      providerName: '╨Ü╨╗╤Ä╤ç ╨┐╤Ç╨╛╨▓╨░╨╣╨┤╨╡╤Ç╨░',
      displayName: '╨₧╤é╨╛╨▒╤Ç╨░╨╢╨░╨╡╨╝╨╛╨╡ ╨╕╨╝╤Å',
      clientId: '╨ÿ╨┤╨╡╨╜╤é╨╕╤ä╨╕╨║╨░╤é╨╛╤Ç ╨║╨╗╨╕╨╡╨╜╤é╨░',
      clientSecret: '╨í╨╡╨║╤Ç╨╡╤é ╨║╨╗╨╕╨╡╨╜╤é╨░',
      secretPlaceholder:
        '╨₧╤ü╤é╨░╨▓╤î╤é╨╡ ╨┐╤â╤ü╤é╤ï╨╝, ╤ç╤é╨╛╨▒╤ï ╤ü╨╛╤à╤Ç╨░╨╜╨╕╤é╤î ╨╜╨░╤ü╤é╤Ç╨╛╨╡╨╜╨╜╤ï╨╣ ╤ü╨╡╨║╤Ç╨╡╤é',
      redirectUrl: '╨É╨┤╤Ç╨╡╤ü ╨╛╨▒╤Ç╨░╤é╨╜╨╛╨│╨╛ ╨▓╤ï╨╖╨╛╨▓╨░',
      scopes: '╨₧╨▒╨╗╨░╤ü╤é╨╕ ╨┤╨╛╤ü╤é╤â╨┐╨░',
      accountRole: '╨á╨╛╨╗╤î ╤â╤ç╨╡╤é╨╜╨╛╨╣ ╨╖╨░╨┐╨╕╤ü╨╕',
      allowedDomains: '╨á╨░╨╖╤Ç╨╡╤ê╨╡╨╜╨╜╤ï╨╡ ╨┐╨╛╤ç╤é╨╛╨▓╤ï╨╡ ╨┤╨╛╨╝╨╡╨╜╤ï',
      bindByEmail: '╨ƒ╤Ç╨╕╨▓╤Å╨╖╨░╤é╤î ╨┐╨╛ ╨┐╨╛╨┤╤é╨▓╨╡╤Ç╨╢╨┤╨╡╨╜╨╜╨╛╨╣ ╨┐╨╛╤ç╤é╨╡',
      autoCreateAdmin: '╨É╨▓╤é╨╛╨╝╨░╤é╨╕╤ç╨╡╤ü╨║╨╕ ╤ü╨╛╨╖╨┤╨░╨▓╨░╤é╤î ╨░╨┤╨╝╨╕╨╜╨╕╤ü╤é╤Ç╨░╤é╨╛╤Ç╨░',
      autoCreateUser: '╨É╨▓╤é╨╛╨╝╨░╤é╨╕╤ç╨╡╤ü╨║╨╕ ╤ü╨╛╨╖╨┤╨░╨▓╨░╤é╤î ╨┐╨╛╨╗╤î╨╖╨╛╨▓╨░╤é╨╡╨╗╤Å',
      testConfig: '╨ƒ╤Ç╨╛╨▓╨╡╤Ç╨╕╤é╤î ╨╜╨░╤ü╤é╤Ç╨╛╨╣╨║╤â',
      testSuccess: '╨¥╨░╤ü╤é╤Ç╨╛╨╣╨║╨░ ╨╖╨░╨┐╨╛╨╗╨╜╨╡╨╜╨░, ╨░╨┤╤Ç╨╡╤ü ╨░╨▓╤é╨╛╤Ç╨╕╨╖╨░╤å╨╕╨╕ ╤ü╨╛╨╖╨┤╨░╨╜',
      copyCallback: '╨Ü╨╛╨┐╨╕╤Ç╨╛╨▓╨░╤é╤î ╨░╨┤╤Ç╨╡╤ü ╨▓╨╛╨╖╨▓╤Ç╨░╤é╨░',
      githubOnlyTip:
        '╨í╨╜╨░╤ç╨░╨╗╨░ ╨┤╨╛╤ü╤é╤â╨┐╨╡╨╜ GitHub. ╨¥╨░╤ü╤é╤Ç╨╛╨╣╤é╨╡ ╨╡╨│╨╛ ╨╖╨┤╨╡╤ü╤î; server.yaml ╨╛╤ü╤é╨░╨╡╤é╤ü╤Å ╨┤╨╗╤Å ╤ü╨╛╨▓╨╝╨╡╤ü╤é╨╕╨╝╨╛╤ü╤é╨╕ ╨╕ ╨▓╨╛╤ü╤ü╤é╨░╨╜╨╛╨▓╨╗╨╡╨╜╨╕╤Å.',
      adminRole: '╨É╨┤╨╝╨╕╨╜╨╕╤ü╤é╤Ç╨░╤é╨╛╤Ç',
      userRole: '╨ƒ╨╛╨╗╤î╨╖╨╛╨▓╨░╤é╨╡╨╗╤î',
      useDefault: '╨ÿ╤ü╨┐╨╛╨╗╤î╨╖╨╛╨▓╨░╤é╤î ╨╖╨╜╨░╤ç╨╡╨╜╨╕╨╡ ╨┐╨╛ ╤â╨╝╨╛╨╗╤ç╨░╨╜╨╕╤Ä',
      listPlaceholder: '╨á╨░╨╖╨┤╨╡╨╗╤Å╨╣╤é╨╡ ╨╖╨╜╨░╤ç╨╡╨╜╨╕╤Å ╨┐╤Ç╨╛╨▒╨╡╨╗╨░╨╝╨╕ ╨╕╨╗╨╕ ╨╖╨░╨┐╤Å╤é╤ï╨╝╨╕',
      copied: '╨í╨║╨╛╨┐╨╕╤Ç╨╛╨▓╨░╨╜╨╛'
    },
    about: {
      latestCommand: '╨₧╨▒╨╜╨╛╨▓╨╕╤é╤î ╨┤╨╛ latest',
      pinnedCommand: '╨₧╨▒╨╜╨╛╨▓╨╕╤é╤î ╨╕ ╨┐╤Ç╨╛╨▓╨╡╤Ç╨╕╤é╤î ╨╜╨░╨╣╨┤╨╡╨╜╨╜╤â╤Ä ╨▓╨╡╤Ç╤ü╨╕╤Ä',
      customCommand: '╨í╨▓╨╛╨╣ ╤ê╨░╨▒╨╗╨╛╨╜ ╨║╨╛╨╝╨░╨╜╨┤╤ï',
      runningVersion: '╨ó╨╡╨║╤â╤ë╨░╤Å ╨▓╨╡╤Ç╤ü╨╕╤Å',
      buildTime: '╨Æ╤Ç╨╡╨╝╤Å ╤ü╨▒╨╛╤Ç╨║╨╕',
      compatVersion: '╨í╨╛╨▓╨╝╨╡╤ü╤é╨╕╨╝╨░╤Å ╨▓╨╡╤Ç╤ü╨╕╤Å RustDesk',
      latestVersion: '╨ƒ╨╛╤ü╨╗╨╡╨┤╨╜╤Å╤Å ╨▓╨╡╤Ç╤ü╨╕╤Å',
      updateAvailable: '╨ö╨╛╤ü╤é╤â╨┐╨╜╨╛ ╨╛╨▒╨╜╨╛╨▓╨╗╨╡╨╜╨╕╨╡',
      upToDate: '╨ú╤ü╤é╨░╨╜╨╛╨▓╨╗╨╡╨╜╨░ ╨┐╨╛╤ü╨╗╨╡╨┤╨╜╤Å╤Å ╨▓╨╡╤Ç╤ü╨╕╤Å',
      updateCheck: '╨ƒ╤Ç╨╛╨▓╨╡╤Ç╨║╨░ ╨╛╨▒╨╜╨╛╨▓╨╗╨╡╨╜╨╕╨╣',
      urlTip:
        '╨É╨┤╤Ç╨╡╤ü ╨┐╤Ç╨╛╨▓╨╡╤Ç╨║╨╕ ╨╝╨╛╨╢╨╜╨╛ ╨╕╨╖╨╝╨╡╨╜╨╕╤é╤î; ╨╛╨╜ ╤ü╨╛╤à╤Ç╨░╨╜╤Å╨╡╤é╤ü╤Å ╨▓ ╤ì╤é╨╛╨╝ ╨▒╤Ç╨░╤â╨╖╨╡╤Ç╨╡. ╨í╨░╨╣╤é ╨┤╨╛╨╗╨╢╨╡╨╜ ╤Ç╨░╨╖╤Ç╨╡╤ê╨░╤é╤î CORS.',
      urlPlaceholder: '╨É╨┤╤Ç╨╡╤ü ╨┐╤Ç╨╛╨▓╨╡╤Ç╨║╨╕ ╨╛╨▒╨╜╨╛╨▓╨╗╨╡╨╜╨╕╨╣',
      checkNow: '╨ƒ╤Ç╨╛╨▓╨╡╤Ç╨╕╤é╤î',
      restoreDefault: '╨Æ╨╡╤Ç╨╜╤â╤é╤î ╨░╨┤╤Ç╨╡╤ü ╨┐╨╛ ╤â╨╝╨╛╨╗╤ç╨░╨╜╨╕╤Ä',
      checkFailed: '╨₧╤ê╨╕╨▒╨║╨░ ╨┐╤Ç╨╛╨▓╨╡╤Ç╨║╨╕ ╨╛╨▒╨╜╨╛╨▓╨╗╨╡╨╜╨╕╨╣',
      invalidUrl: '╨ƒ╨╛╨┤╨┤╨╡╤Ç╨╢╨╕╨▓╨░╤Ä╤é╤ü╤Å ╤é╨╛╨╗╤î╨║╨╛ HTTP ╨╕ HTTPS',
      invalidResponse: '╨Æ╨╡╤Ç╤ü╨╕╤Å ╨╜╨╡ ╨╜╨░╨╣╨┤╨╡╨╜╨░ ╨▓ ╨╛╤é╨▓╨╡╤é╨╡',
      updateCommand: '╨Ü╨╛╨╝╨░╨╜╨┤╨░ ╨╛╨▒╨╜╨╛╨▓╨╗╨╡╨╜╨╕╤Å ╨║╨╛╨╜╤é╨╡╨╣╨╜╨╡╤Ç╨░',
      commandTip:
        '╨¿╨░╨▒╨╗╨╛╨╜ ╨╝╨╛╨╢╨╜╨╛ ╨╕╨╖╨╝╨╡╨╜╨╕╤é╤î; {version} ╨╖╨░╨╝╨╡╨╜╤Å╨╡╤é╤ü╤Å ╨┐╨╛╤ü╨╗╨╡╨┤╨╜╨╡╨╣ ╨▓╨╡╤Ç╤ü╨╕╨╡╨╣.',
      copyCommand: 'Копировать команду',
      versionInfo: 'Информация о версии',
      errorHelp: 'Справка по кодам ошибок',
      errcodeTip: 'Ниже перечислены все коды ошибок сервера.',
      searchPlaceholder: 'Поиск по коду, сообщению или описанию',
      moduleFilter: 'Фильтр по модулю',
      errCode: 'Код',
      errMessage: 'Сообщение',
      errModule: 'Модуль',
      errDescription: 'Описание',
      errSolution: 'Решение'
    }
  },
  dropdown: {
    closeCurrent: 'Close Current',
    closeOther: 'Close Other',
    closeLeft: 'Close Left',
    closeRight: 'Close Right',
    closeAll: 'Close All'
  },
  icon: {
    themeConfig: 'Theme Configuration',
    themeSchema: 'Theme Schema',
    lang: '╨í╨╝╨╡╨╜╨╕╤é╤î ╤Å╨╖╤ï╨║',
    fullscreen: 'Fullscreen',
    fullscreenExit: 'Exit Fullscreen',
    reload: 'Reload Page',
    collapse: 'Collapse Menu',
    expand: 'Expand Menu',
    pin: 'Pin',
    unpin: 'Unpin'
  },
  datatable: { itemCount: 'Total {total} items' },
  dataMap: {
    user: {
      username: '╨ÿ╨╝╤Å ╨┐╨╛╨╗╤î╨╖╨╛╨▓╨░╤é╨╡╨╗╤Å',
      password: 'Password',
      name: '╨¥╨╕╨║╨╜╨╡╨╣╨╝',
      email: 'Email',
      licensed_devices: '╨¢╨╕╤å╨╡╨╜╨╖╨╕╤Ç╨╛╨▓╨░╨╜╨╜╤ï╨╡ ╤â╤ü╤é╤Ç╨╛╨╣╤ü╤é╨▓╨░',
      login_verify: '╨ƒ╤Ç╨╛╨▓╨╡╤Ç╨║╨░ ╨▓╤à╨╛╨┤╨░',
      status: '╨í╤é╨░╤é╤â╤ü',
      is_admin: '╨É╨┤╨╝╨╕╨╜╨╕╤ü╤é╤Ç╨░╤é╨╛╤Ç',
      tfa_secret: '2FA Secret',
      tfa_code: '2FA Code',
      created_at: '╨í╨╛╨╖╨┤╨░╨╜╨╛',
      statusLabel: {
        disabled: '╨₧╤é╨║╨╗╤Ä╤ç╨╡╨╜',
        unverified: '╨¥╨╡ ╨┐╨╛╨┤╤é╨▓╨╡╤Ç╨╢╨┤╨╡╨╜',
        normal: '╨¥╨╛╤Ç╨╝╨░╨╗╤î╨╜╨╛'
      },
      loginVerifyLabel: { none: '╨¥╨╡╤é', emailCheck: '╨ƒ╤Ç╨╛╨▓╨╡╤Ç╨║╨░ email', tfaCheck: '2FA' }
    },
    session: { expired: 'Expired At', created_at: 'Created At' },
    device: {
      username: '╨ÿ╨╝╤Å ╨┐╨╛╨╗╤î╨╖╨╛╨▓╨░╤é╨╡╨╗╤Å',
      hostname: '╨ÿ╨╝╤Å ╨║╨╛╨╝╨┐╤î╤Ä╤é╨╡╤Ç╨░',
      version: '╨Æ╨╡╤Ç╤ü╨╕╤Å RustDesk',
      memory: '╨ƒ╨░╨╝╤Å╤é╤î',
      os: '╨₧╨í',
      rustdesk_id: 'Rustdesk ID'
    },
    audit: {
      username: '╨ƒ╨╛╨╗╤î╨╖╨╛╨▓╨░╤é╨╡╨╗╤î',
      type: '╨ó╨╕╨┐',
      conn_id: 'Connect Id',
      rustdesk_id: 'Rustdesk ID',
      ip: 'IP',
      session_id: 'Session Id',
      uuid: 'UUID',
      created_at: '╨í╨╛╨╖╨┤╨░╨╜╨╛',
      closed_at: 'Closed At',
      typeLabel: {
        remote_control: '╨ú╨┤╨░╨╗╤æ╨╜╨╜╨╛╨╡ ╤â╨┐╤Ç╨░╨▓╨╗╨╡╨╜╨╕╨╡',
        file_transfer: '╨ƒ╨╡╤Ç╨╡╨┤╨░╤ç╨░ ╤ä╨░╨╣╨╗╨╛╨▓',
        tcp_tunnel: 'TCP ╤é╤â╨╜╨╜╨╡╨╗╤î'
      },
      fileTransferTypeLabel: {
        master_controlled: '╨ú╨┐╤Ç╨░╨▓╨╗╤Å╤Ä╤ë╨╕╨╣ -> ╨ú╨┐╤Ç╨░╨▓╨╗╤Å╨╡╨╝╤ï╨╣',
        controlled_master: '╨ú╨┐╤Ç╨░╨▓╨╗╤Å╨╡╨╝╤ï╨╣ -> ╨ú╨┐╤Ç╨░╨▓╨╗╤Å╤Ä╤ë╨╕╨╣'
      },
      peer_id: 'Peer ID',
      path: 'Path'
    },
    mailTemplate: {
      name: '╨ÿ╨╝╤Å',
      type: '╨ó╨╕╨┐',
      subject: '╨ó╨╡╨╝╨░',
      contents: '╨í╨╛╨┤╨╡╤Ç╨╢╨╕╨╝╨╛╨╡',
      created_at: '╨í╨╛╨╖╨┤╨░╨╜╨╛',
      typeLabel: {
        loginVerify: '╨ƒ╤Ç╨╛╨▓╨╡╤Ç╨║╨░ ╨▓╤à╨╛╨┤╨░',
        registerVerify: '╨ƒ╤Ç╨╛╨▓╨╡╤Ç╨║╨░ ╤Ç╨╡╨│╨╕╤ü╤é╤Ç╨░╤å╨╕╨╕',
        other: '╨ö╤Ç╤â╨│╨╛╨╡'
      }
    },
    mailLog: {
      username: '╨ƒ╨╛╨╗╤î╨╖╨╛╨▓╨░╤é╨╡╨╗╤î',
      uuid: 'UUID',
      from: '╨₧╤é',
      to: '╨Ü╨╛╨╝╤â',
      subject: '╨ó╨╡╨╝╨░',
      contents: 'Content',
      status: '╨í╤é╨░╤é╤â╤ü',
      created_at: '╨Æ╤Ç╨╡╨╝╤Å ╨╛╤é╨┐╤Ç╨░╨▓╨║╨╕',
      statusLabel: { ok: '╨ú╤ü╨┐╨╡╤ê╨╜╨╛', err: '╨₧╤ê╨╕╨▒╨║╨░' }
    },
    ab: {
      rustdesk_id: 'Rustdesk ID',
      username: 'Username',
      hostname: 'Hostname',
      tags: 'Tags',
      alias: 'Alias',
      hash: 'Hash',
      owner: 'Owner',
      name: 'Address Book Name',
      user_id: 'User ID',
      guid: 'GUID',
      rule: 'Rule',
      max_peer: 'Max Peers',
      shared: 'Shared',
      ab_id: 'Address Book ID',
      tagName: 'Name',
      tagColor: 'Color',
      updated_at: 'Updated At',
      personal: '╨£╨╛╤Å ╨░╨┤╤Ç╨╡╤ü╨╜╨░╤Å ╨║╨╜╨╕╨│╨░',
      legacy: '╨ú╤ü╤é╨░╤Ç╨╡╨▓╤ê╨░╤Å ╨░╨┤╤Ç╨╡╤ü╨╜╨░╤Å ╨║╨╜╨╕╨│╨░',
      note: '╨ƒ╤Ç╨╕╨╝╨╡╤ç╨░╨╜╨╕╨╡',
      platform: '╨ƒ╨╗╨░╤é╤ä╨╛╤Ç╨╝╨░',
      personalReadOnly: '╨¢╨╕╤ç╨╜╨░╤Å (╤é╨╛╨╗╤î╨║╨╛ ╤ç╤é╨╡╨╜╨╕╨╡)',
      nameRequired: '╨ÿ╨╝╤Å ╨╛╨▒╤Å╨╖╨░╤é╨╡╨╗╤î╨╜╨╛',
      deviceIdRequired: 'ID ╤â╤ü╤é╤Ç╨╛╨╣╤ü╤é╨▓╨░ ╨╛╨▒╤Å╨╖╨░╤é╨╡╨╗╨╡╨╜',
      tagsHint: '╨á╨░╨╖╨┤╨╡╨╗╤Å╨╣╤é╨╡ ╨╝╨╡╤é╨║╨╕ ╨╖╨░╨┐╤Å╤é╤ï╨╝╨╕',
      read: '╨º╤é╨╡╨╜╨╕╨╡',
      readWrite: '╨º╤é╨╡╨╜╨╕╨╡ ╨╕ ╨╖╨░╨┐╨╕╤ü╤î',
      fullControl: '╨ƒ╨╛╨╗╨╜╤ï╨╣ ╨┤╨╛╤ü╤é╤â╨┐'
    },
    token: {
      device_os: '╨₧╨í ╤â╤ü╤é╤Ç╨╛╨╣╤ü╤é╨▓╨░',
      device_name: '╨ÿ╨╝╤Å ╤â╤ü╤é╤Ç╨╛╨╣╤ü╤é╨▓╨░',
      token_hash: '╨Ñ╨╡╤ê ╤é╨╛╨║╨╡╨╜╨░',
      is_admin: '╨É╨┤╨╝╨╕╨╜',
      status: '╨É╨║╤é╨╕╨▓╨╡╨╜'
    },
    oauth: {
      provider: '╨ƒ╤Ç╨╛╨▓╨░╨╣╨┤╨╡╤Ç',
      subject: '╨í╤â╨▒╤è╨╡╨║╤é',
      email: '╨ƒ╨╛╤ç╤é╨░',
      name: '╨ÿ╨╝╤Å',
      last_login_at: '╨ƒ╨╛╤ü╨╗╨╡╨┤╨╜╨╕╨╣ ╨▓╤à╨╛╨┤'
    },
    errorLog: {
      code: 'Error Code',
      message: 'Message',
      module: 'Module',
      path: 'Path',
      method: 'Method',
      user_name: 'User',
      client_ip: 'Client IP',
      user_agent: 'User Agent',
      created_at: 'Time'
    },
    loginLog: {
      allEvents: '╨Æ╤ü╨╡ ╤ü╨╛╨▒╤ï╤é╨╕╤Å',
      event: '╨í╨╛╨▒╤ï╤é╨╕╨╡',
      userAgent: '╨É╨│╨╡╨╜╤é ╨┐╨╛╨╗╤î╨╖╨╛╨▓╨░╤é╨╡╨╗╤Å',
      success: '╨ú╤ü╨┐╨╡╤à',
      reason: '╨ƒ╤Ç╨╕╤ç╨╕╨╜╨░'
    }
  },
  api: {
    CaptchaError: '╨₧╤ê╨╕╨▒╨║╨░ CAPTCHA',
    UserNotExists: '╨ƒ╨╛╨╗╤î╨╖╨╛╨▓╨░╤é╨╡╨╗╤î ╨╜╨╡ ╤ü╤â╤ë╨╡╤ü╤é╨▓╤â╨╡╤é',
    UsernameOrPasswordError: '╨¥╨╡╨▓╨╡╤Ç╨╜╤ï╨╣ ╨╗╨╛╨│╨╕╨╜ ╨╕╨╗╨╕ ╨┐╨░╤Ç╨╛╨╗╤î',
    UserExists: '╨ÿ╨╝╤Å ╨┐╨╛╨╗╤î╨╖╨╛╨▓╨░╤é╨╡╨╗╤Å ╤â╨╢╨╡ ╨╕╤ü╨┐╨╛╨╗╤î╨╖╤â╨╡╤é╤ü╤Å',
    UsernameEmpty: '╨ÿ╨╝╤Å ╨┐╨╛╨╗╤î╨╖╨╛╨▓╨░╤é╨╡╨╗╤Å ╨╜╨╡ ╨╝╨╛╨╢╨╡╤é ╨▒╤ï╤é╤î ╨┐╤â╤ü╤é╤ï╨╝',
    PasswordEmpty: '╨ƒ╨░╤Ç╨╛╨╗╤î ╨╜╨╡ ╨╝╨╛╨╢╨╡╤é ╨▒╤ï╤é╤î ╨┐╤â╤ü╤é╤ï╨╝',
    UserAddSuccess: '╨ƒ╨╛╨╗╤î╨╖╨╛╨▓╨░╤é╨╡╨╗╤î ╤â╤ü╨┐╨╡╤ê╨╜╨╛ ╤ü╨╛╨╖╨┤╨░╨╜',
    DataError: '╨₧╤ê╨╕╨▒╨║╨░ ╨┤╨░╨╜╨╜╤ï╤à',
    RequestError: '╨₧╤ê╨╕╨▒╨║╨░ ╨╖╨░╨┐╤Ç╨╛╤ü╨░',
    UserUpdateSuccess: '╨ƒ╨╛╨╗╤î╨╖╨╛╨▓╨░╤é╨╡╨╗╤î ╤â╤ü╨┐╨╡╤ê╨╜╨╛ ╨╛╨▒╨╜╨╛╨▓╨╗╤æ╨╜',
    UserDeleteSuccess: '╨ƒ╨╛╨╗╤î╨╖╨╛╨▓╨░╤é╨╡╨╗╤î ╤â╤ü╨┐╨╡╤ê╨╜╨╛ ╤â╨┤╨░╨╗╤æ╨╜',
    SessionKillSuccess: '╨í╨╡╤ü╤ü╨╕╤Å ╤â╤ü╨┐╨╡╤ê╨╜╨╛ ╨╖╨░╨▓╨╡╤Ç╤ê╨╡╨╜╨░',
    MailTemplateNameEmpty: '╨ÿ╨╝╤Å ╨╜╨╡ ╨╝╨╛╨╢╨╡╤é ╨▒╤ï╤é╤î ╨┐╤â╤ü╤é╤ï╨╝',
    MailTemplateSubjectEmpty: '╨ó╨╡╨╝╨░ ╨╜╨╡ ╨╝╨╛╨╢╨╡╤é ╨▒╤ï╤é╤î ╨┐╤â╤ü╤é╨╛╨╣',
    MailTemplateContentsEmpty: '╨í╨╛╨┤╨╡╤Ç╨╢╨╕╨╝╨╛╨╡ ╨╜╨╡ ╨╝╨╛╨╢╨╡╤é ╨▒╤ï╤é╤î ╨┐╤â╤ü╤é╤ï╨╝',
    MailTemplateAddSuccess: '╨¿╨░╨▒╨╗╨╛╨╜ ╨┐╨╕╤ü╤î╨╝╨░ ╤â╤ü╨┐╨╡╤ê╨╜╨╛ ╤ü╨╛╨╖╨┤╨░╨╜',
    MailTemplateUpdateSuccess: '╨ƒ╨╛╤ç╤é╨╛╨▓╤ï╨╣ ╤ê╨░╨▒╨╗╨╛╨╜ ╤â╤ü╨┐╨╡╤ê╨╜╨╛ ╨╛╨▒╨╜╨╛╨▓╨╗╤æ╨╜',
    NoEmailAddress: '╨É╨┤╤Ç╨╡╤ü ╤ì╨╗╨╡╨║╤é╤Ç╨╛╨╜╨╜╨╛╨╣ ╨┐╨╛╤ç╤é╤ï ╨╜╨╡ ╨╖╨░╨┤╨░╨╜',
    VerificationCodeError: '╨₧╤ê╨╕╨▒╨║╨░ ╨║╨╛╨┤╨░ ╨┐╨╛╨┤╤é╨▓╨╡╤Ç╨╢╨┤╨╡╨╜╨╕╤Å',
    UUIDEmpty: 'UUID ╨╜╨╡ ╨╝╨╛╨╢╨╡╤é ╨▒╤ï╤é╤î ╨┐╤â╤ü╤é╤ï╨╝'
  }
};
export default local;
