const local: App.I18n.Schema = {
  "system": {
    "title": "Rustdesk Api Server",
    "updateTitle": "システムバージョン更新",
    "updateContent": "新しいシステムバージョンが利用可能です。今すぐページを更新しますか？",
    "updateConfirm": "今すぐ更新",
    "updateCancel": "後で"
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
    "update": "更新する",
    "updateSuccess": "更新に成功しました",
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
    "login": "ログイン",
    "iframe-page": "埋め込みページ",
    "home": "ホーム", "about": "情報と更新",
    "audit": "監査",
    "user": "ユーザー管理",
    "user_list": "ユーザー一覧",
    "user_sessions": "セッション",
    "user_profile": "プロフィール",
    "system": "システム管理",
    "system_mail_template": "メールテンプレート",
    "system_mail_logs": "メールログ",
    "system_mail": "メール管理",
    "system_server": "サーバー設定",
    "system_tokens": "ユーザートークン",
    "system_oauth": "OAuth管理",
    "audit_baselogs": "基本ログ",
    "audit_filetransferlogs": "ファイル転送ログ",
    "audit_loginlogs": "ログインログ",
    "devices": "デバイス",
    "my-devices": "連絡先",
    "my-devices_peers": "マイピア",
    "my-devices_manage": "アドレス帳管理",
    "my-devices_tags": "タグ管理", "workspace": "マイワークスペース", "workspace_overview": "概要", "workspace_devices": "マイデバイス", "workspace_sessions": "ログインセッション", "workspace_security": "セキュリティ履歴", "workspace_profile": "プロフィール"
  },
  "page": {
    "login": {
      "common": {
        "loginOrRegister": "ログイン / 登録",
        "userNamePlaceholder": "ユーザー名を入力してください",
        "phonePlaceholder": "Please enter phone number",
        "codePlaceholder": "認証コードを入力してください",
        "passwordPlaceholder": "パスワードを入力してください",
        "confirmPasswordPlaceholder": "パスワードを再入力してください",
        "codeLogin": "認証コードでログイン",
        "confirm": "確認",
        "back": "戻る",
        "validateSuccess": "検証に成功しました",
        "loginSuccess": "ログイン成功",
        "welcomeBack": "おかえりなさい、{userName} さん！",
        "thirdPartyLogin": "サードパーティーログイン",
        "continueWith": "{provider} でログイン",
        "providerUnavailable": "{provider} ログインは現在利用できません"
      },
      "pwdLogin": {
        "title": "パスワードログイン",
        "rememberMe": "ログイン状態を保持",
        "switchToUser": "ユーザーログイン"
      },
      "userLogin": {
        "title": "ユーザーログイン",
        "switchToAdmin": "管理者ログイン"
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
      "cardDetail": {
        "viewHint": "詳細を見る",
        "recentUsers": "最近のユーザー",
        "recentDevices": "最近のデバイス",
        "recentVisits": "最近のアクセスログ",
        "desc": {
          "userCount": "システム内のユーザー総数を表示します。",
          "deviceCount": "システム内のデバイス総数を表示します。",
          "onlineCount": "ハートビート統計に基づくオンラインデバイス数を表示します。",
          "visitCount": "監査ログからのアクセス統計を表示します。"
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
        "showQr": "QRコードを表示",
        "qrTitle": "RustDeskインポートQRコード",
        "qrTip": "RustDeskモバイルアプリでこのQRコードをスキャンして設定をインポートできます。",
        "qrPayload": "RustDeskテンプレートテキスト",
        "qrFailed": "QRコードの生成に失敗しました",
        "refresh": "Refresh",
        "clearCacheReload": "Clear Cache & Reload",
        "cacheTtlHint": "Cache TTL: config {configSeconds}s, connectivity {connectivitySeconds}s",
        "source": "Source",
        "lastUpdated": "最終更新",
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
          "auto": "自動検出"
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
        "inputUsername": "ユーザー名を入力",
        "inputPassword": "パスワードを入力",
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
    },
    "myDevices": {
      "title": "連絡先",
      "welcome": "ようこそ、{userName}",
      "status": "ステータス",
      "online": "オンライン",
      "offline": "オフライン",
      "conns": "接続数",
      "lastSync": "最終同期",
      "logout": "ログアウト"
    },
    "workspace": {
      "scopeTitle": "個人ワークスペース", "scopeTip": "自分のデバイス、セッション、セキュリティ履歴、許可されたアドレス帳のみ表示されます。", "myDevices": "マイデバイス", "activeSessions": "有効なセッション", "addressBooks": "アドレス帳", "securityEvents": "セキュリティ履歴", "currentSession": "現在のセッション", "revokeConfirm": "このセッションを無効にしますか？", "revoke": "無効化", "accountRole": "アカウント権限", "adminRole": "管理者", "userRole": "一般ユーザー", "permissionScope": "権限範囲", "userScope": "個人リソースと許可された共有アドレス帳", "active": "有効"
    },
    "about": {
      "runningVersion": "実行中のバージョン", "buildTime": "ビルド日時", "compatVersion": "対応RustDeskバージョン", "latestVersion": "最新バージョン", "updateAvailable": "更新があります", "upToDate": "最新です", "updateCheck": "オンライン更新確認", "urlTip": "確認URLは変更でき、このブラウザーに保存されます。対象サイトはCORSを許可する必要があります。", "urlPlaceholder": "更新確認URL", "checkNow": "今すぐ確認", "restoreDefault": "既定に戻す", "checkFailed": "更新確認に失敗しました", "invalidUrl": "HTTPまたはHTTPSのみ対応しています", "invalidResponse": "有効なバージョンが見つかりません", "updateCommand": "コンテナ更新コマンド", "commandTip": "テンプレートを編集できます。{version} は最新バージョンに置換されます。", "copyCommand": "コマンドをコピー"
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
    "lang": "言語を切り替え",
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
      "personal": "マイアドレス帳", "note": "メモ", "platform": "プラットフォーム", "personalReadOnly": "個人用（読み取り専用）", "nameRequired": "名前は必須です", "deviceIdRequired": "デバイスIDは必須です", "tagsHint": "複数のタグはカンマで区切ってください", "read": "読み取り", "readWrite": "読み取りと書き込み", "fullControl": "フルコントロール"
    },
    "token": {
      "device_os": "デバイスOS",
      "device_name": "デバイス名",
      "token_hash": "トークンハッシュ",
      "is_admin": "管理者",
      "status": "有効"
    },
    "oauth": {
      "provider": "プロバイダー",
      "subject": "サブジェクト",
      "email": "メール",
      "name": "名前",
      "last_login_at": "最終ログイン"
    },
    "loginLog": {
      "allEvents": "全イベント",
      "event": "イベント",
      "userAgent": "ユーザーエージェント",
      "success": "成功",
      "reason": "理由"
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
    "RequestError": "リクエスト失敗",
    "UserUpdateSuccess": "ユーザー更新成功",
    "UserDeleteSuccess": "ユーザーを削除しました",
    "SessionKillSuccess": "セッションを終了しました",
    "MailTemplateNameEmpty": "テンプレート名を入力してください",
    "MailTemplateSubjectEmpty": "件名を入力してください",
    "MailTemplateContentsEmpty": "内容を入力してください",
    "MailTemplateAddSuccess": "メールテンプレートを作成しました",
    "MailTemplateUpdateSuccess": "メールテンプレート更新成功",
    "NoEmailAddress": "メールアドレスが設定されていません",
    "VerificationCodeError": "認証コードエラー",
    "UUIDEmpty": "UUID を入力してください"
  }
};

export default local;
