const local: App.I18n.Schema = {
  "system": {
    "title": "Rustdesk Api Server",
    "updateTitle": "系统版本更新通知",
    "updateContent": "检测到系统有新版本发布，是否立即刷新页面？",
    "updateConfirm": "立即刷新",
    "updateCancel": "稍后再说"
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
    "update": "更新",
    "updateSuccess": "更新成功",
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
    "home": "홈",
    "audit": "감사",
    "user": "사용자 관리",
    "user_list": "사용자 목록",
    "user_sessions": "세션",
    "system": "시스템 관리",
    "system_mail_template": "메일 템플릿",
    "system_mail_logs": "메일 로그",
    "system_mail": "메일 관리",
    "system_server": "서버 설정",
    "audit_baselogs": "기본 로그",
    "audit_filetransferlogs": "파일 전송 로그",
    "devices": "장치"
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
        "providerUnavailable": "{provider} 로그인은 현재 사용할 수 없습니다"
      },
      "pwdLogin": {
        "title": "비밀번호 로그인",
        "rememberMe": "로그인 상태 유지"
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
        "refresh": "Refresh",
        "clearCacheReload": "Clear Cache & Reload",
        "cacheTtlHint": "Cache TTL: config {configSeconds}s, connectivity {connectivitySeconds}s",
        "source": "Source",
        "lastUpdated": "最后更新",
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
          "empty": "Empty"
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
    "UserUpdateSuccess": "用户修改成功",
    "UserDeleteSuccess": "사용자가 삭제되었습니다",
    "SessionKillSuccess": "세션이 종료되었습니다",
    "MailTemplateNameEmpty": "템플릿 이름을 입력하세요",
    "MailTemplateSubjectEmpty": "제목을 입력하세요",
    "MailTemplateContentsEmpty": "내용을 입력하세요",
    "MailTemplateAddSuccess": "메일 템플릿이 생성되었습니다",
    "MailTemplateUpdateSuccess": "邮件模板修改成功",
    "NoEmailAddress": "이메일 주소가 설정되지 않았습니다",
    "VerificationCodeError": "인증 코드 오류",
    "UUIDEmpty": "UUID를 입력하세요"
  }
};

export default local;
