const local: App.I18n.Schema = {
  "system": {
    "title": "Rustdesk Api Server",
    "updateTitle": "Version Update Notice (ChatGPT-Generated)",
    "updateContent": "This update content was generated/assembled by ChatGPT and may contain mistakes. Review and test carefully before using it, especially in production. Refresh now?",
    "updateConfirm": "Refresh immediately",
    "updateCancel": "Later"
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
    "update": "Update",
    "updateSuccess": "Update Success",
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
    "login": "Login",
    "iframe-page": "Iframe",
    "home": "Home",
    "audit": "Audit",
    "user": "User Management",
    "user_list": "User List",
    "user_sessions": "Sessions",
    "system": "System Management",
    "system_mail_template": "Mail Template",
    "system_mail_logs": "Mail Logs",
    "system_mail": "Mail Managment",
    "audit_baselogs": "Base Logs",
    "audit_filetransferlogs": "File Transfer Logs",
    "devices": "Devices"
  },
  "page": {
    "login": {
      "common": {
        "loginOrRegister": "Login / Register",
        "userNamePlaceholder": "Please enter user name",
        "phonePlaceholder": "Please enter phone number",
        "codePlaceholder": "Please enter verification code",
        "passwordPlaceholder": "Please enter password",
        "confirmPasswordPlaceholder": "Please enter password again",
        "codeLogin": "Verification code login",
        "confirm": "Confirm",
        "back": "Back",
        "validateSuccess": "Verification passed",
        "loginSuccess": "Login successfully",
        "welcomeBack": "Welcome back, {userName} !"
      },
      "pwdLogin": {
        "title": "Password Login",
        "rememberMe": "Remember me"
      }
    },
    "home": {
      "greeting": "おはようございます、{userName}さん！",
      "userCount": "ユーザー数",
      "deviceCount": "デバイス数",
      "onlineCount": "オンライン数",
      "visitsCount": "訪問数",
      "operatingSystem": "OS",
      "oneWeek": "1週間",
      "changeLogs": "更新履歴",
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
        "source": "Source",
        "lastUpdated": "Last Updated",
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
        "addUser": "ユーザー追加",
        "editUser": "ユーザー編集",
        "inputUsername": "Input Username",
        "inputPassword": "Input Password",
        "inputNickname": "Input Nickname",
        "emailFormatError": "Email format error",
        "selectUserStatus": "Please select user status",
        "searchPlaceholder": "ユーザー名/ニックネーム/メール",
        "tfa_secret_bind": "2FA Device Bind",
        "require2FASecret": "2FA Secret Empty",
        "require2FACode": "2FA Code Can't Empty"
      },
      "sessions": {
        "kill": "切断",
        "confirmKill": "このセッションを終了しますか？"
      },
      "audit": {
        "logsSearchPlaceholder": "ユーザー名/操作/RustdeskID/IP"
      },
      "devices": {
        "logsSearchPlaceholder": "ユーザー名/ホスト名/RustdeskID"
      }
    },
    "system": {
      "mailTemplate": {
        "addMailTemplate": "テンプレート追加",
        "editMailTemplate": "テンプレート編集",
        "inputName": "名前を入力",
        "inputSubject": "件名を入力",
        "inputContents": "内容を入力",
        "selectType": "種類を選択"
      },
      "mailLog": {
        "info": "詳細"
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
    "lang": "Switch Language",
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
      "username": "ユーザー名",
      "password": "Password",
      "name": "ニックネーム",
      "email": "メール",
      "licensed_devices": "許可デバイス数",
      "login_verify": "ログイン認証",
      "status": "状態",
      "is_admin": "管理者",
      "tfa_secret": "2FA Secret",
      "tfa_code": "2FA Code",
      "created_at": "作成日時",
      "statusLabel": {
        "disabled": "無効",
        "unverified": "未確認",
        "normal": "正常"
      },
      "loginVerifyLabel": {
        "none": "不要",
        "emailCheck": "メール認証",
        "tfaCheck": "2FA"
      }
    },
    "session": {
      "expired": "Expired At",
      "created_at": "Created At"
    },
    "device": {
      "username": "ユーザー名",
      "hostname": "ホスト名",
      "version": "RustDesk バージョン",
      "memory": "Memory",
      "os": "OS",
      "rustdesk_id": "Rustdesk ID"
    },
    "audit": {
      "username": "ユーザー名",
      "type": "種類",
      "conn_id": "Connect Id",
      "rustdesk_id": "Rustdesk ID",
      "ip": "IP",
      "session_id": "Session Id",
      "uuid": "UUID",
      "created_at": "作成日時",
      "closed_at": "Closed At",
      "typeLabel": {
        "remote_control": "リモート操作",
        "file_transfer": "ファイル転送",
        "tcp_tunnel": "TCP トンネル"
      },
      "fileTransferTypeLabel": {
        "master_controlled": "操作側 -> 被操作側",
        "controlled_master": "被操作側 -> 操作側"
      },
      "peer_id": "Peer ID",
      "path": "Path"
    },
    "mailTemplate": {
      "name": "名前",
      "type": "種類",
      "subject": "件名",
      "contents": "内容",
      "created_at": "作成日時",
      "typeLabel": {
        "loginVerify": "ログイン認証",
        "registerVerify": "登録認証",
        "other": "その他"
      }
    },
    "mailLog": {
      "username": "ユーザー名",
      "uuid": "UUID",
      "from": "送信元",
      "to": "宛先",
      "subject": "件名",
      "contents": "Content",
      "status": "状態",
      "created_at": "送信日時",
      "statusLabel": {
        "ok": "成功",
        "err": "失敗"
      }
    }
  },
  "api": {
    "CaptchaError": "CAPTCHA エラー",
    "UserNotExists": "ユーザーが存在しません",
    "UsernameOrPasswordError": "アカウントまたはパスワードが正しくありません",
    "UserExists": "ユーザー名は既に使用されています",
    "UsernameEmpty": "ユーザー名を入力してください",
    "PasswordEmpty": "パスワードを入力してください",
    "UserAddSuccess": "ユーザーを作成しました",
    "DataError": "データエラー",
    "UserUpdateSuccess": "ユーザーを更新しました",
    "UserDeleteSuccess": "ユーザーを削除しました",
    "SessionKillSuccess": "セッションを終了しました",
    "MailTemplateNameEmpty": "テンプレート名を入力してください",
    "MailTemplateSubjectEmpty": "件名を入力してください",
    "MailTemplateContentsEmpty": "内容を入力してください",
    "MailTemplateAddSuccess": "メールテンプレートを作成しました",
    "MailTemplateUpdateSuccess": "メールテンプレートを更新しました",
    "NoEmailAddress": "メールアドレスが設定されていません",
    "VerificationCodeError": "認証コードエラー",
    "UUIDEmpty": "UUID を入力してください"
  }
};

export default local;
