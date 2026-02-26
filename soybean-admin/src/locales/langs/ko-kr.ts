import enUs from './en-us';

const local: App.I18n.Schema = {
  ...enUs,
  system: { ...enUs.system, title: 'Rustdesk Api Server' },
  common: {
    ...enUs.common,
    action: '??',
    add: '??',
    addSuccess: '?? ??',
    backToHome: '??? ????',
    batchDelete: '?? ??',
    cancel: '??',
    close: '??',
    check: '??',
    expandColumn: '? ??',
    columnSetting: '? ??',
    config: '??',
    confirm: '??',
    delete: '??',
    deleteSuccess: '?? ??',
    confirmDelete: '?? ?????????',
    edit: '??',
    look: '??',
    warning: '??',
    error: '??',
    index: '??',
    keywordSearch: '???? ?????',
    logout: '????',
    logoutConfirm: '???????????',
    lookForward: '? ?? ??',
    modify: '??',
    modifySuccess: '?? ??',
    noData: '??? ??',
    operate: '??',
    pleaseCheckValue: '?? ???? ?????',
    refresh: '????',
    reset: '???',
    search: '??',
    switch: '??',
    tip: '?',
    trigger: '??',
    update: '????',
    updateSuccess: '???? ??',
    userCenter: '??? ??',
    yesOrNo: {
      yes: '?',
      no: '???'
    }
  },
  request: {
    ...enUs.request,
    logout: '?? ?? ? ??? ????',
    logoutMsg: '??? ??? ???? ????. ?? ??????',
    logoutWithModal: '?? ?? ? ??? ???? ????',
    logoutWithModalMsg: '??? ??? ???? ????. ?? ??????',
    refreshToken: '??? ???? ?????',
    tokenExpired: '?? ??? ???????'
  },
  theme: {
    ...enUs.theme,
    themeSchema: {
      ...enUs.theme.themeSchema,
      title: '?? ???',
      light: '???',
      dark: '??',
      auto: '??? ??'
    },
    grayscale: '??????',
    colourWeakness: '?? ??',
    layoutMode: {
      ...enUs.theme.layoutMode,
      title: '???? ??'
    },
    recommendColor: '?? ?? ???? ??',
    themeColor: {
      ...enUs.theme.themeColor,
      title: '?? ??'
    },
    scrollMode: {
      ...enUs.theme.scrollMode,
      title: '??? ??'
    },
    page: {
      ...enUs.theme.page,
      animate: '??? ?????',
      mode: {
        ...enUs.theme.page.mode,
        title: '????? ??'
      }
    },
    fixedHeaderAndTab: '?? ? ? ??',
    header: {
      ...enUs.theme.header,
      height: '?? ??'
    },
    tab: {
      ...enUs.theme.tab,
      visible: '? ??',
      mode: {
        ...enUs.theme.tab.mode,
        title: '? ??'
      }
    },
    sider: {
      ...enUs.theme.sider,
      width: '???? ??'
    },
    footer: {
      ...enUs.theme.footer,
      visible: '?? ??'
    },
    themeDrawerTitle: '?? ??',
    pageFunTitle: '??? ??',
    configOperation: {
      ...enUs.theme.configOperation,
      copyConfig: '?? ??',
      resetConfig: '?? ???',
      resetSuccessMsg: '??? ??'
    }
  },
  route: {
    ...enUs.route,
    login: '???',
    403: '?? ??',
    404: '???? ?? ? ??',
    500: '?? ??',
    'iframe-page': 'Iframe',
    home: '?',
    audit: '??',
    user: '??? ??',
    user_list: '??? ??',
    user_sessions: '??',
    system: '??? ??',
    system_mail_template: '?? ???',
    system_mail_logs: '?? ??',
    system_mail: '??',
    audit_baselogs: '?? ??',
    audit_filetransferlogs: '?? ?? ??',
    devices: '??'
  },
  page: {
    ...enUs.page,
      login: {
        ...enUs.page.login,
        common: {
          ...enUs.page.login.common,
          loginOrRegister: '??? / ????',
          userNamePlaceholder: '??? ??? ?????',
          phonePlaceholder: '????? ?????',
          codePlaceholder: '?? ??? ?????',
          passwordPlaceholder: '????? ?????',
          confirmPasswordPlaceholder: '????? ?? ?????',
          codeLogin: '???? ???',
          confirm: '??',
          back: '??',
          validateSuccess: '?? ??',
          loginSuccess: '??? ??',
          welcomeBack: '?? ?? ?? ?????, {userName} !'
        },
        pwdLogin: {
          ...enUs.page.login.pwdLogin,
          title: '???? ???',
          rememberMe: '??? ?? ??'
        }
      },
    home: {
      ...enUs.page.home,
      greeting: '좋은 아침입니다, {userName}님!',
      userCount: '사용자 수',
      deviceCount: '장치 수',
      onlineCount: '온라인 수',
      visitsCount: '방문 수',
      operatingSystem: '운영체제',
      oneWeek: '최근 1주',
      changeLogs: '업데이트 로그',
      serverConfig: {
        ...enUs.page.home.serverConfig,
        title: '클라이언트 연결 설정',
        tip: '아래 설정을 RustDesk 클라이언트에 붙여넣으세요. KEY가 비어 있으면 `RUSTDESK_KEY` 환경변수를 설정하세요.',
        refresh: '새로고침',
        copyAll: '전체 복사',
        copyTemplate: 'RustDesk 템플릿 복사',
        connectivity: {
          ...enUs.page.home.serverConfig.connectivity,
          check: '연결 확인',
          checkOne: '확인',
          clear: '결과 지우기'
        }
      }
    },
    user: {
      ...enUs.page.user,
      list: {
        ...enUs.page.user.list,
        addUser: '사용자 추가',
        editUser: '사용자 수정',
        searchPlaceholder: '사용자명/닉네임/이메일'
      },
      sessions: {
        ...enUs.page.user.sessions,
        kill: '종료',
        confirmKill: '이 세션을 종료하시겠습니까?'
      },
      audit: {
        ...enUs.page.user.audit,
        logsSearchPlaceholder: '사용자명/작업/RustdeskID/IP'
      },
      devices: {
        ...enUs.page.user.devices,
        logsSearchPlaceholder: '사용자명/호스트명/RustdeskID'
      }
    },
    system: {
      ...enUs.page.system,
      mailTemplate: {
        ...enUs.page.system.mailTemplate,
        addMailTemplate: '템플릿 추가',
        editMailTemplate: '템플릿 수정',
        inputName: '이름 입력',
        inputSubject: '제목 입력',
        inputContents: '내용 입력',
        selectType: '유형 선택'
      },
      mailLog: {
        ...enUs.page.system.mailLog,
        info: '상세'
      }
    }
  },
  dataMap: {
    ...enUs.dataMap,
    user: {
      ...enUs.dataMap.user,
      username: '사용자명',
        password: '????',
      name: '닉네임',
      email: '이메일',
      licensed_devices: '허용 장치 수',
      login_verify: '로그인 인증',
      status: '상태',
      is_admin: '관리자',
        tfa_secret: '2FA ???',
        tfa_code: '2FA ??',
      created_at: '생성일',
      statusLabel: {
        ...enUs.dataMap.user.statusLabel,
        disabled: '비활성화',
        unverified: '미인증',
        normal: '정상'
      },
      loginVerifyLabel: {
        ...enUs.dataMap.user.loginVerifyLabel,
        none: '없음',
        emailCheck: '이메일 인증',
        tfaCheck: '2FA'
      }
    },
      session: {
        ...enUs.dataMap.session,
        expired: '?? ??',
        created_at: '?? ??'
      },
    device: {
      ...enUs.dataMap.device,
      username: '사용자명',
      hostname: '호스트명',
      version: 'RustDesk 버전',
        memory: '???',
      os: '운영체제',
      rustdesk_id: 'Rustdesk ID'
    },
    audit: {
      ...enUs.dataMap.audit,
      username: '사용자명',
      type: '유형',
        conn_id: '?? ID',
      rustdesk_id: 'Rustdesk ID',
        peer_id: 'Peer ID',
      ip: 'IP',
        session_id: '?? ID',
        uuid: 'UUID',
      created_at: '생성일',
        closed_at: '?? ??',
      typeLabel: {
        ...enUs.dataMap.audit.typeLabel,
        remote_control: '원격 제어',
        file_transfer: '파일 전송',
        tcp_tunnel: 'TCP 터널'
      },
      fileTransferTypeLabel: {
        ...enUs.dataMap.audit.fileTransferTypeLabel,
        master_controlled: '제어자 -> 피제어자',
        controlled_master: '피제어자 -> 제어자'
      },
        path: '??'
    },
    mailTemplate: {
      ...enUs.dataMap.mailTemplate,
      name: '이름',
      type: '유형',
      subject: '제목',
      contents: '내용',
      created_at: '생성일',
      typeLabel: {
        ...enUs.dataMap.mailTemplate.typeLabel,
        loginVerify: '로그인 인증',
        registerVerify: '회원가입 인증',
        other: '기타'
      }
    },
    mailLog: {
      ...enUs.dataMap.mailLog,
      username: '사용자명',
        uuid: 'UUID',
      from: '발신자',
      to: '수신자',
      subject: '제목',
        contents: '??',
      status: '상태',
      created_at: '전송 시간',
      statusLabel: {
        ...enUs.dataMap.mailLog.statusLabel,
        ok: '성공',
        err: '실패'
      }
    }
  },
  api: {
    ...enUs.api,
    CaptchaError: 'CAPTCHA 오류',
    UserNotExists: '사용자가 존재하지 않습니다',
    UsernameOrPasswordError: '계정 또는 비밀번호가 올바르지 않습니다',
    UserExists: '이미 사용 중인 사용자명입니다',
    UsernameEmpty: '사용자명을 입력하세요',
    PasswordEmpty: '비밀번호를 입력하세요',
    UserAddSuccess: '사용자가 생성되었습니다',
    DataError: '데이터 오류',
    UserUpdateSuccess: '사용자가 수정되었습니다',
    UserDeleteSuccess: '사용자가 삭제되었습니다',
    SessionKillSuccess: '세션이 종료되었습니다',
    MailTemplateNameEmpty: '템플릿 이름을 입력하세요',
    MailTemplateSubjectEmpty: '제목을 입력하세요',
    MailTemplateContentsEmpty: '내용을 입력하세요',
    MailTemplateAddSuccess: '메일 템플릿이 생성되었습니다',
    MailTemplateUpdateSuccess: '메일 템플릿이 수정되었습니다',
    NoEmailAddress: '이메일 주소가 설정되지 않았습니다',
    VerificationCodeError: '인증 코드 오류',
    UUIDEmpty: 'UUID를 입력하세요'
  },
  dropdown: {
    ...enUs.dropdown,
    closeCurrent: '?? ??',
    closeOther: '?? ? ??',
    closeLeft: '?? ??',
    closeRight: '??? ??',
    closeAll: '?? ??'
  },
  icon: {
    ...enUs.icon,
    themeConfig: '?? ??',
    themeSchema: '?? ???',
    lang: '?? ??',
    fullscreen: '?? ??',
    fullscreenExit: '?? ?? ??',
    reload: '??? ????',
    collapse: '?? ??',
    expand: '?? ???',
    pin: '??',
    unpin: '?? ??'
  },
  datatable: {
    ...enUs.datatable,
    itemCount: '? {total}?'
  }
};

export default local;
