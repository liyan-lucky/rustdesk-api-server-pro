const local: App.I18n.Schema = {
  "system": {
    "title": "Rustdesk Api Server",
    "updateTitle": "시스템 버전 업데이트",
    "updateContent": "새 시스템 버전을 사용할 수 있습니다. 지금 페이지를 새로고침하시겠습니까?",
    "updateConfirm": "지금 새로고침",
    "updateCancel": "나중에"
  },
  "common": {
    "action": "Action",
    "add": "Add",
    "addSuccess": "Add Success",
    "backToHome": "Back to home",
    "batchDelete": "Batch Delete",
    "cancel": "Cancel",
    "close": "Close",
    "check": "Check",
    "expandColumn": "Expand Column",
    "columnSetting": "Column Setting",
    "config": "Config",
    "confirm": "Confirm",
    "delete": "Delete",
    "deleteSuccess": "Delete Success",
    "confirmDelete": "Are you sure you want to delete?",
    "edit": "Edit",
    "import": "CSV Import",
    "export": "CSV Export",
    "look": "Look",
    "warning": "Warning",
    "error": "Error",
    "index": "Index",
    "keywordSearch": "Please enter keyword",
    "logout": "Logout",
    "logoutConfirm": "Are you sure you want to log out?",
    "lookForward": "Coming soon",
    "modify": "Modify",
    "modifySuccess": "Modify Success",
    "noData": "No Data",
    "operate": "Operate",
    "pleaseCheckValue": "Please check whether the value is valid",
    "refresh": "Refresh",
    "reset": "Reset",
    "search": "Search",
    "switch": "Switch",
    "tip": "Tip",
    "trigger": "Trigger",
    "update": "업데이트",
    "updateSuccess": "업데이트 성공",
    "userCenter": "User Center",
    "yesOrNo": {
      "yes": "Yes",
      "no": "No"
    }
  },
  "request": {
    "logout": "Logout user after request failed",
    "logoutMsg": "User status is invalid, please log in again",
    "logoutWithModal": "Pop up modal after request failed and then log out user",
    "logoutWithModalMsg": "User status is invalid, please log in again",
    "refreshToken": "The requested token has expired, refresh the token",
    "tokenExpired": "The requested token has expired"
  },
  "theme": {
    "themeSchema": {
      "title": "Theme Schema",
      "light": "Light",
      "dark": "Dark",
      "auto": "Follow System"
    },
    "grayscale": "Grayscale",
    "colourWeakness": "Colour Weakness",
    "layoutMode": {
      "title": "Layout Mode",
      "vertical": "Vertical Menu Mode",
      "horizontal": "Horizontal Menu Mode",
      "vertical-mix": "Vertical Mix Menu Mode",
      "horizontal-mix": "Horizontal Mix menu Mode",
      "reverseHorizontalMix": "Reverse first level menus and child level menus position"
    },
    "recommendColor": "Apply Recommended Color Algorithm",
    "recommendColorDesc": "The recommended color algorithm refers to",
    "themeColor": {
      "title": "Theme Color",
      "primary": "Primary",
      "info": "Info",
      "success": "Success",
      "warning": "Warning",
      "error": "Error",
      "followPrimary": "Follow Primary"
    },
    "scrollMode": {
      "title": "Scroll Mode",
      "wrapper": "Wrapper",
      "content": "Content"
    },
    "page": {
      "animate": "Page Animate",
      "mode": {
        "title": "Page Animate Mode",
        "fade": "Fade",
        "fade-slide": "Slide",
        "fade-bottom": "Fade Zoom",
        "fade-scale": "Fade Scale",
        "zoom-fade": "Zoom Fade",
        "zoom-out": "Zoom Out",
        "none": "None"
      }
    },
    "fixedHeaderAndTab": "Fixed Header And Tab",
    "header": {
      "height": "Header Height",
      "breadcrumb": {
        "visible": "Breadcrumb Visible",
        "showIcon": "Breadcrumb Icon Visible"
      }
    },
    "tab": {
      "visible": "Tab Visible",
      "cache": "Tab Cache",
      "height": "Tab Height",
      "mode": {
        "title": "Tab Mode",
        "chrome": "Chrome",
        "button": "Button"
      }
    },
    "sider": {
      "inverted": "Dark Sider",
      "width": "Sider Width",
      "collapsedWidth": "Sider Collapsed Width",
      "mixWidth": "Mix Sider Width",
      "mixCollapsedWidth": "Mix Sider Collapse Width",
      "mixChildMenuWidth": "Mix Child Menu Width"
    },
    "footer": {
      "visible": "Footer Visible",
      "fixed": "Fixed Footer",
      "height": "Footer Height",
      "right": "Right Footer"
    },
    "watermark": {
      "visible": "Watermark Full Screen Visible",
      "text": "Watermark Text"
    },
    "themeDrawerTitle": "Theme Configuration",
    "pageFunTitle": "Page Function",
    "configOperation": {
      "copyConfig": "Copy Config",
      "copySuccessMsg": "Copy Success, Please replace the variable \"themeSettings\" in \"src/theme/settings.ts\"",
      "resetConfig": "Reset Config",
      "resetSuccessMsg": "Reset Success"
    }
  },
  "route": {
    "403": "No Permission",
    "404": "Page Not Found",
    "500": "Server Error",
    "login": "로그인",
    "iframe-page": "임베드 페이지",
    "home": "홈", "about": "정보 및 업데이트",
    "audit": "감사",
    "user": "사용자 관리",
    "user_list": "사용자 목록",
    "user_sessions": "세션",
    "user_profile": "프로필",
    "system": "시스템 관리",
    "system_mail_template": "메일 템플릿",
    "system_mail_logs": "메일 로그",
    "system_mail": "메일 관리",
    "system_server": "서버 설정",
    "system_tokens": "사용자 토큰",
    "system_oauth": "타사 로그인",
    "audit_baselogs": "기본 로그",
    "audit_filetransferlogs": "파일 전송 로그",
    "audit_loginlogs": "로그인 로그",
    "devices": "장치",
    "my-devices": "연락처",
    "my-devices_peers": "내 피어",
    "my-devices_manage": "주소록 관리",
    "my-devices_tags": "태그 관리", "workspace": "내 작업 공간", "workspace_overview": "개요", "workspace_devices": "내 장치", "workspace_sessions": "로그인 세션", "workspace_security": "보안 기록", "workspace_profile": "프로필"
  },
  "page": {
    "login": {
      "common": {
        "loginOrRegister": "로그인 / 가입",
        "userNamePlaceholder": "사용자 이름을 입력하세요",
        "phonePlaceholder": "Please enter phone number",
        "codePlaceholder": "인증 코드를 입력하세요",
        "passwordPlaceholder": "비밀번호를 입력하세요",
        "confirmPasswordPlaceholder": "비밀번호를 다시 입력하세요",
        "codeLogin": "인증 코드 로그인",
        "confirm": "확인",
        "back": "뒤로",
        "validateSuccess": "검증 성공",
        "loginSuccess": "로그인 성공",
        "welcomeBack": "환영합니다, {userName} 님!",
        "thirdPartyLogin": "서드파티 로그인",
        "continueWith": "{provider}로 로그인",
        "providerUnavailable": "{provider} 로그인은 현재 사용할 수 없습니다", "oauthAccountNotBound": "일치하는 계정을 연결할 수 없습니다. 동일한 확인된 이메일을 설정하거나 자동 생성을 활성화하세요.", "oauthProviderUnreachable": "서버가 로그인 공급자에 연결할 수 없습니다. 아웃바운드 HTTPS 연결을 확인하세요.", "oauthStateExpired": "로그인 요청이 만료되었거나 이미 사용되었습니다. 다시 시작하세요.", "oauthAuthFailed": "타사 로그인에 실패했습니다. 설정과 보안 감사 로그를 확인하세요."
      },
      "pwdLogin": {
        "title": "비밀번호 로그인",
        "rememberMe": "로그인 상태 유지",
        "switchToUser": "사용자 로그인"
      },
      "userLogin": {
        "title": "사용자 로그인",
        "switchToAdmin": "관리자 로그인"
      }
    },
    "home": {
      "greeting": "좋은 아침입니다, {userName}님!",
      "userCount": "사용자 수",
      "deviceCount": "장치 수",
      "onlineCount": "온라인 수",
      "visitsCount": "방문 수",
      "operatingSystem": "운영체제",
      "oneWeek": "최근 1주",
      "changeLogs": "업데이트 로그",
      "cardDetail": {
        "viewHint": "클릭하여 상세 보기",
        "recentUsers": "최근 사용자",
        "recentDevices": "최근 장치",
        "recentVisits": "최근 방문 로그",
        "desc": {
          "userCount": "시스템 내 전체 사용자 수를 표시합니다.",
          "deviceCount": "시스템 내 전체 장치 수를 표시합니다.",
          "onlineCount": "하트비트 통계를 기반으로 온라인 장치 수를 표시합니다.",
          "visitCount": "감사 로그 기반 방문 통계를 표시합니다."
        }
      },
      "serverConfig": {
        "title": "Client Connection Config",
        "tip": "Copy the following values into the RustDesk client. If KEY is empty, set the `RUSTDESK_KEY` container environment variable.",
        "idServer": "ID Server",
        "relayServer": "Relay Server",
        "apiServer": "API Server",
        "key": "KEY",
        "idServerPlaceholder": "e.g. your.domain.com",
        "relayServerPlaceholder": "e.g. your.domain.com",
        "apiServerPlaceholder": "e.g. https://your.domain.com",
        "keyPlaceholder": "Provide via RUSTDESK_KEY environment variable",
        "copy": "Copy",
        "copyAll": "Copy All",
        "copyTemplate": "Copy RustDesk Template",
        "showQr": "QR 코드 표시",
        "qrTitle": "RustDesk 가져오기 QR 코드",
        "qrTip": "RustDesk 모바일 앱에서 이 QR 코드를 스캔하여 설정을 가져올 수 있습니다.",
        "qrPayload": "RustDesk 템플릿 텍스트",
        "qrFailed": "QR 코드 생성 실패",
        "refresh": "Refresh",
        "clearCacheReload": "Clear Cache & Reload",
        "cacheTtlHint": "Cache TTL: config {configSeconds}s, connectivity {connectivitySeconds}s",
        "source": "Source",
        "lastUpdated": "마지막 업데이트",
        "ageSeconds": "{seconds}s ago",
        "show": "Show",
        "hide": "Hide",
        "missingTip": "The following fields are empty, please configure them in container environment variables first: {fields}",
        "copyEmpty": "{label} is empty and cannot be copied",
        "copySuccess": "{label} copied",
        "copyFailed": "{label} copy failed",
        "fetchFailed": "Failed to load server configuration",
        "cacheCleared": "Cache cleared, reloading server configuration",
        "sourceType": {
          "remote": "Remote",
          "memory-cache": "Memory Cache",
          "session-cache": "Session Cache",
          "env": "Env",
          "inferred": "Inferred",
          "empty": "Empty",
          "auto": "자동 감지"
        },
        "sourceHint": {
          "env": "This value comes from the container environment variable.",
          "inferred": "This value is auto-inferred from the current access address.",
          "empty": "No value is configured or inferred yet."
        },
        "connectivity": {
          "clear": "Clear Results",
          "check": "Check Connectivity",
          "checkOne": "Check",
          "checked": "Connectivity check completed",
          "checkedOne": "{field} connectivity checked",
          "checkedCached": "Using recent connectivity check result (cache)",
          "checkFailed": "Connectivity check failed",
          "cleared": "Connectivity results cleared",
          "source": "Check Source",
          "lastChecked": "Last Checked",
          "target": "Target",
          "duration": "Duration",
          "notChecked": "Not checked yet",
          "checkSourceType": {
            "remote": "Remote",
            "cache": "Cache"
          },
          "status": {
            "idle": "Unchecked",
            "ok": "Reachable",
            "error": "Failed",
            "skip": "Skipped"
          }
        }
      }
    },
    "user": {
      "list": {
        "addUser": "사용자 추가",
        "editUser": "사용자 수정",
        "inputUsername": "사용자 이름 입력",
        "inputPassword": "비밀번호 입력",
        "inputNickname": "Input Nickname",
        "emailFormatError": "Email format error",
        "selectUserStatus": "Please select user status",
        "searchPlaceholder": "사용자명/닉네임/이메일",
        "tfa_secret_bind": "2FA Device Bind",
        "require2FASecret": "2FA Secret Empty",
        "require2FACode": "2FA Code Can't Empty"
      },
      "sessions": {
        "kill": "종료",
        "confirmKill": "Confirm Kill?"
      },
      "audit": {
        "logsSearchPlaceholder": "사용자명/작업/RustdeskID/IP"
      },
      "devices": {
        "logsSearchPlaceholder": "사용자명/호스트명/RustdeskID"
      }
    },
    "system": {
      "mailTemplate": {
        "addMailTemplate": "템플릿 추가",
        "editMailTemplate": "템플릿 수정",
        "inputName": "이름 입력",
        "inputSubject": "제목 입력",
        "inputContents": "내용 입력",
        "selectType": "유형 선택"
      },
      "mailLog": {
        "info": "상세"
      }
    },
    "myDevices": {
      "title": "연락처",
      "welcome": "환영합니다, {userName}",
      "status": "상태",
      "online": "온라인",
      "offline": "오프라인",
      "conns": "연결 수",
      "lastSync": "마지막 동기화",
      "logout": "로그아웃"
    },
    "workspace": {
      "scopeTitle": "개인 작업 공간", "scopeTip": "본인의 장치, 세션, 보안 기록 및 권한이 있는 주소록만 표시됩니다.", "myDevices": "내 장치", "activeSessions": "활성 세션", "addressBooks": "주소록", "securityEvents": "보안 기록", "currentSession": "현재 세션", "revokeConfirm": "이 로그인 세션을 해제하시겠습니까?", "revoke": "해제", "accountRole": "계정 역할", "adminRole": "관리자", "userRole": "일반 사용자", "permissionScope": "권한 범위", "userScope": "개인 리소스 및 명시적으로 공유된 주소록", "active": "활성"
    },
    "oauth": {
      "configTitle": "타사 로그인 설정", "bindingsTitle": "계정 연결", "addProvider": "공급자 추가", "editProvider": "공급자 편집", "providerName": "공급자 키", "displayName": "표시 이름", "clientId": "클라이언트 ID", "clientSecret": "클라이언트 비밀", "secretPlaceholder": "설정된 비밀을 유지하려면 비워 두세요", "redirectUrl": "콜백 URL", "scopes": "권한 범위", "accountRole": "계정 역할", "allowedDomains": "허용 이메일 도메인", "bindByEmail": "확인된 이메일로 연결", "autoCreateAdmin": "관리자 자동 생성", "autoCreateUser": "사용자 자동 생성", "testConfig": "설정 테스트", "testSuccess": "설정이 완료되어 인증 URL을 생성했습니다", "copyCallback": "콜백 복사", "githubOnlyTip": "우선 GitHub를 지원합니다. 여기에서 설정하며 server.yaml은 호환성과 복구용입니다.", "adminRole": "관리자", "userRole": "사용자", "useDefault": "기본값 사용", "listPlaceholder": "여러 값은 공백이나 쉼표로 구분하세요", "copied": "복사됨"
    },
    "about": {
      "latestCommand": "latest로 업데이트", "pinnedCommand": "감지된 버전으로 업데이트 및 확인", "customCommand": "사용자 지정 명령 템플릿",
      "runningVersion": "실행 버전", "buildTime": "빌드 시간", "compatVersion": "호환 RustDesk 버전", "latestVersion": "최신 버전", "updateAvailable": "업데이트 있음", "upToDate": "최신 상태", "updateCheck": "온라인 업데이트 확인", "urlTip": "확인 URL은 변경할 수 있으며 이 브라우저에 저장됩니다. 대상 사이트에서 CORS를 허용해야 합니다.", "urlPlaceholder": "업데이트 확인 URL", "checkNow": "지금 확인", "restoreDefault": "기본값 복원", "checkFailed": "업데이트 확인 실패", "invalidUrl": "HTTP 또는 HTTPS URL만 지원합니다", "invalidResponse": "유효한 버전을 찾을 수 없습니다", "updateCommand": "컨테이너 업데이트 명령", "commandTip": "명령 템플릿을 편집할 수 있으며 {version}은 최신 버전으로 대체됩니다.", "copyCommand": "명령 복사"
    }
  },
  "dropdown": {
    "closeCurrent": "Close Current",
    "closeOther": "Close Other",
    "closeLeft": "Close Left",
    "closeRight": "Close Right",
    "closeAll": "Close All"
  },
  "icon": {
    "themeConfig": "Theme Configuration",
    "themeSchema": "Theme Schema",
    "lang": "언어 전환",
    "fullscreen": "Fullscreen",
    "fullscreenExit": "Exit Fullscreen",
    "reload": "Reload Page",
    "collapse": "Collapse Menu",
    "expand": "Expand Menu",
    "pin": "Pin",
    "unpin": "Unpin"
  },
  "datatable": {
    "itemCount": "Total {total} items"
  },
  "dataMap": {
    "user": {
      "username": "사용자명",
      "password": "Password",
      "name": "닉네임",
      "email": "이메일",
      "licensed_devices": "허용 장치 수",
      "login_verify": "로그인 인증",
      "status": "상태",
      "is_admin": "관리자",
      "tfa_secret": "2FA Secret",
      "tfa_code": "2FA Code",
      "created_at": "생성일",
      "statusLabel": {
        "disabled": "비활성화",
        "unverified": "미인증",
        "normal": "정상"
      },
      "loginVerifyLabel": {
        "none": "없음",
        "emailCheck": "이메일 인증",
        "tfaCheck": "2FA"
      }
    },
    "session": {
      "expired": "Expired At",
      "created_at": "Created At"
    },
    "device": {
      "username": "사용자명",
      "hostname": "호스트명",
      "version": "RustDesk 버전",
      "memory": "Memory",
      "os": "운영체제",
      "rustdesk_id": "Rustdesk ID"
    },
    "audit": {
      "username": "사용자명",
      "type": "유형",
      "conn_id": "Connect Id",
      "rustdesk_id": "Rustdesk ID",
      "ip": "IP",
      "session_id": "Session Id",
      "uuid": "UUID",
      "created_at": "생성일",
      "closed_at": "Closed At",
      "typeLabel": {
        "remote_control": "원격 제어",
        "file_transfer": "파일 전송",
        "tcp_tunnel": "TCP 터널"
      },
      "fileTransferTypeLabel": {
        "master_controlled": "제어자 -> 피제어자",
        "controlled_master": "피제어자 -> 제어자"
      },
      "peer_id": "Peer ID",
      "path": "Path"
    },
    "mailTemplate": {
      "name": "이름",
      "type": "유형",
      "subject": "제목",
      "contents": "내용",
      "created_at": "생성일",
      "typeLabel": {
        "loginVerify": "로그인 인증",
        "registerVerify": "회원가입 인증",
        "other": "기타"
      }
    },
    "mailLog": {
      "username": "사용자명",
      "uuid": "UUID",
      "from": "발신자",
      "to": "수신자",
      "subject": "제목",
      "contents": "Content",
      "status": "상태",
      "created_at": "전송 시간",
      "statusLabel": {
        "ok": "성공",
        "err": "실패"
      }
    },
    "ab": {
      "rustdesk_id": "Rustdesk ID",
      "username": "Username",
      "hostname": "Hostname",
      "tags": "Tags",
      "alias": "Alias",
      "hash": "Hash",
      "owner": "Owner",
      "name": "Address Book Name",
      "user_id": "User ID",
      "guid": "GUID",
      "rule": "Rule",
      "max_peer": "Max Peers",
      "shared": "Shared",
      "ab_id": "Address Book ID",
      "tagName": "Name",
      "tagColor": "Color",
      "updated_at": "Updated At",
      "personal": "내 주소록", "legacy": "레거시 주소록", "note": "메모", "platform": "플랫폼", "personalReadOnly": "개인용(읽기 전용)", "nameRequired": "이름은 필수입니다", "deviceIdRequired": "장치 ID는 필수입니다", "tagsHint": "여러 태그는 쉼표로 구분하세요", "read": "읽기", "readWrite": "읽기 및 쓰기", "fullControl": "전체 제어"
    },
    "token": {
      "device_os": "디바이스 OS",
      "device_name": "디바이스 이름",
      "token_hash": "토큰 해시",
      "is_admin": "관리자",
      "status": "활성"
    },
    "oauth": {
      "provider": "제공자",
      "subject": "주체",
      "email": "이메일",
      "name": "이름",
      "last_login_at": "마지막 로그인"
    },
    "loginLog": {
      "allEvents": "모든 이벤트",
      "event": "이벤트",
      "userAgent": "사용자 에이전트",
      "success": "성공",
      "reason": "사유"
    }
  },
  "api": {
    "CaptchaError": "CAPTCHA 오류",
    "UserNotExists": "사용자가 존재하지 않습니다",
    "UsernameOrPasswordError": "계정 또는 비밀번호가 올바르지 않습니다",
    "UserExists": "이미 사용 중인 사용자명입니다",
    "UsernameEmpty": "사용자명을 입력하세요",
    "PasswordEmpty": "비밀번호를 입력하세요",
    "UserAddSuccess": "사용자가 생성되었습니다",
    "DataError": "데이터 오류",
    "RequestError": "요청 실패",
    "UserUpdateSuccess": "사용자 수정 성공",
    "UserDeleteSuccess": "사용자가 삭제되었습니다",
    "SessionKillSuccess": "세션이 종료되었습니다",
    "MailTemplateNameEmpty": "템플릿 이름을 입력하세요",
    "MailTemplateSubjectEmpty": "제목을 입력하세요",
    "MailTemplateContentsEmpty": "내용을 입력하세요",
    "MailTemplateAddSuccess": "메일 템플릿이 생성되었습니다",
    "MailTemplateUpdateSuccess": "메일 템플릿 수정 성공",
    "NoEmailAddress": "이메일 주소가 설정되지 않았습니다",
    "VerificationCodeError": "인증 코드 오류",
    "UUIDEmpty": "UUID를 입력하세요"
  }
};

export default local;
