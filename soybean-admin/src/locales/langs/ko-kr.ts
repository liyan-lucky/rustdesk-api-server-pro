import enUs from './en-us';

const local: App.I18n.Schema = {
  ...enUs,
  system: { ...enUs.system, title: 'Rustdesk Api Server' },
  common: {
    ...enUs.common,
    refresh: '새로고침',
    noData: '데이터 없음',
    error: '오류',
    warning: '경고'
  },
  route: {
    ...enUs.route,
    home: '홈',
    audit: '감사',
    user: '사용자 관리',
    user_list: '사용자 목록',
    user_sessions: '세션',
    system: '시스템 관리',
    system_mail_template: '메일 템플릿',
    system_mail_logs: '메일 로그',
    system_mail: '메일',
    devices: '장치'
  },
  page: {
    ...enUs.page,
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
      name: '닉네임',
      email: '이메일',
      licensed_devices: '허용 장치 수',
      login_verify: '로그인 인증',
      status: '상태',
      is_admin: '관리자',
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
    device: {
      ...enUs.dataMap.device,
      username: '사용자명',
      hostname: '호스트명',
      version: 'RustDesk 버전',
      os: '운영체제',
      rustdesk_id: 'Rustdesk ID'
    },
    audit: {
      ...enUs.dataMap.audit,
      username: '사용자명',
      type: '유형',
      rustdesk_id: 'Rustdesk ID',
      ip: 'IP',
      created_at: '생성일',
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
      }
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
      from: '발신자',
      to: '수신자',
      subject: '제목',
      status: '상태',
      created_at: '전송 시간',
      statusLabel: {
        ...enUs.dataMap.mailLog.statusLabel,
        ok: '성공',
        err: '실패'
      }
    }
  },
  icon: { ...enUs.icon, lang: '언어 전환', reload: '페이지 새로고침', fullscreen: '전체화면' }
};

export default local;
