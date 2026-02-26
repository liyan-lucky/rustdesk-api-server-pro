const local: App.I18n.Schema = {
  "system": {
    "title": "Rustdesk Api Server",
    "updateTitle": "Уведомление об обновлении версии",
    "updateContent": "This update content was generated/assembled by ChatGPT and may contain mistakes. Review and test carefully before using it, especially in production. Refresh now?",
    "updateConfirm": "Обновить",
    "updateCancel": "Позже"
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
      "greeting": "Доброе утро, {userName}, сегодня отличный день!",
      "userCount": "Пользователи",
      "deviceCount": "Устройства",
      "onlineCount": "Онлайн",
      "visitsCount": "Посещения",
      "operatingSystem": "Операционные системы",
      "oneWeek": "За неделю",
      "changeLogs": "Журнал изменений",
      "serverConfig": {
        "title": "Конфигурация подключения клиента",
        "tip": "Скопируйте значения ниже в клиент RustDesk. Если KEY пустой, задайте переменную окружения `RUSTDESK_KEY`.",
        "idServer": "ID сервер",
        "relayServer": "Релейный сервер",
        "apiServer": "API сервер",
        "key": "KEY",
        "idServerPlaceholder": "например your.domain.com",
        "relayServerPlaceholder": "например your.domain.com",
        "apiServerPlaceholder": "например https://your.domain.com",
        "keyPlaceholder": "Укажите через переменную окружения RUSTDESK_KEY",
        "copy": "Copy",
        "copyAll": "Копировать все",
        "copyTemplate": "Копировать шаблон RustDesk",
        "refresh": "Обновить",
        "clearCacheReload": "Очистить кэш и перезагрузить",
        "source": "Источник",
        "lastUpdated": "Последнее обновление",
        "show": "Показать",
        "hide": "Скрыть",
        "missingTip": "Следующие поля пустые, сначала настройте их в переменных окружения контейнера: {fields}",
        "copyEmpty": "{label} пусто, копирование невозможно",
        "copySuccess": "{label} скопировано",
        "copyFailed": "Не удалось скопировать {label}",
        "fetchFailed": "Не удалось загрузить конфигурацию сервера",
        "cacheCleared": "Кэш очищен, повторная загрузка конфигурации сервера",
        "sourceType": {
          "remote": "Удалённый источник",
          "memory-cache": "Кэш памяти",
          "session-cache": "Кэш сессии",
          "env": "Переменная окружения",
          "inferred": "Автоопределение",
          "empty": "Пусто"
        },
        "sourceHint": {
          "env": "Это значение получено из переменной окружения контейнера.",
          "inferred": "Это значение автоматически определено по текущему адресу доступа.",
          "empty": "Значение не настроено и не определено автоматически."
        },
        "connectivity": {
          "clear": "Очистить результаты",
          "check": "Проверить доступность",
          "checkOne": "Проверить",
          "checked": "Проверка доступности завершена",
          "checkedOne": "Проверка {field} завершена",
          "checkedCached": "Использован недавний результат проверки (кэш)",
          "checkFailed": "Проверка доступности не удалась",
          "cleared": "Результаты проверки очищены",
          "source": "Источник проверки",
          "lastChecked": "Последняя проверка",
          "target": "Цель",
          "duration": "Время",
          "notChecked": "Ещё не проверялось",
          "checkSourceType": {
            "remote": "Удалённая проверка",
            "cache": "Кэш"
          },
          "status": {
            "idle": "Не проверено",
            "ok": "Доступно",
            "error": "Ошибка",
            "skip": "Пропущено"
          }
        }
      }
    },
    "user": {
      "list": {
        "addUser": "Добавить пользователя",
        "editUser": "Редактировать пользователя",
        "inputUsername": "Input Username",
        "inputPassword": "Input Password",
        "inputNickname": "Input Nickname",
        "emailFormatError": "Email format error",
        "selectUserStatus": "Please select user status",
        "searchPlaceholder": "Имя пользователя\\Ник\\Email",
        "tfa_secret_bind": "2FA Device Bind",
        "require2FASecret": "2FA Secret Empty",
        "require2FACode": "2FA Code Can't Empty"
      },
      "sessions": {
        "kill": "Завершить",
        "confirmKill": "Confirm Kill?"
      },
      "audit": {
        "logsSearchPlaceholder": "Пользователь\\Действие\\RustdeskID\\IP"
      },
      "devices": {
        "logsSearchPlaceholder": "Пользователь\\Хост\\RustdeskID"
      }
    },
    "system": {
      "mailTemplate": {
        "addMailTemplate": "Добавить шаблон",
        "editMailTemplate": "Редактировать шаблон",
        "inputName": "Введите имя",
        "inputSubject": "Введите тему",
        "inputContents": "Введите содержимое",
        "selectType": "Выберите тип"
      },
      "mailLog": {
        "info": "Подробности"
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
      "username": "Имя пользователя",
      "password": "Password",
      "name": "Никнейм",
      "email": "Email",
      "licensed_devices": "Лицензированные устройства",
      "login_verify": "Проверка входа",
      "status": "Статус",
      "is_admin": "Администратор",
      "tfa_secret": "2FA Secret",
      "tfa_code": "2FA Code",
      "created_at": "Создано",
      "statusLabel": {
        "disabled": "Отключен",
        "unverified": "Не подтвержден",
        "normal": "Нормально"
      },
      "loginVerifyLabel": {
        "none": "Нет",
        "emailCheck": "Проверка email",
        "tfaCheck": "2FA"
      }
    },
    "session": {
      "expired": "Expired At",
      "created_at": "Created At"
    },
    "device": {
      "username": "Имя пользователя",
      "hostname": "Имя компьютера",
      "version": "Версия RustDesk",
      "memory": "Память",
      "os": "ОС",
      "rustdesk_id": "Rustdesk ID"
    },
    "audit": {
      "username": "Пользователь",
      "type": "Тип",
      "conn_id": "Connect Id",
      "rustdesk_id": "Rustdesk ID",
      "ip": "IP",
      "session_id": "Session Id",
      "uuid": "UUID",
      "created_at": "Создано",
      "closed_at": "Closed At",
      "typeLabel": {
        "remote_control": "Удалённое управление",
        "file_transfer": "Передача файлов",
        "tcp_tunnel": "TCP туннель"
      },
      "fileTransferTypeLabel": {
        "master_controlled": "Управляющий -> Управляемый",
        "controlled_master": "Управляемый -> Управляющий"
      },
      "peer_id": "Peer ID",
      "path": "Path"
    },
    "mailTemplate": {
      "name": "Имя",
      "type": "Тип",
      "subject": "Тема",
      "contents": "Содержимое",
      "created_at": "Создано",
      "typeLabel": {
        "loginVerify": "Проверка входа",
        "registerVerify": "Проверка регистрации",
        "other": "Другое"
      }
    },
    "mailLog": {
      "username": "Пользователь",
      "uuid": "UUID",
      "from": "От",
      "to": "Кому",
      "subject": "Тема",
      "contents": "Content",
      "status": "Статус",
      "created_at": "Время отправки",
      "statusLabel": {
        "ok": "Успешно",
        "err": "Ошибка"
      }
    }
  },
  "api": {
    "CaptchaError": "Ошибка CAPTCHA",
    "UserNotExists": "Пользователь не существует",
    "UsernameOrPasswordError": "Неверный логин или пароль",
    "UserExists": "Имя пользователя уже используется",
    "UsernameEmpty": "Имя пользователя не может быть пустым",
    "PasswordEmpty": "Пароль не может быть пустым",
    "UserAddSuccess": "Пользователь успешно создан",
    "DataError": "Ошибка данных",
    "UserUpdateSuccess": "Пользователь успешно изменён",
    "UserDeleteSuccess": "Пользователь успешно удалён",
    "SessionKillSuccess": "Сессия успешно завершена",
    "MailTemplateNameEmpty": "Имя не может быть пустым",
    "MailTemplateSubjectEmpty": "Тема не может быть пустой",
    "MailTemplateContentsEmpty": "Содержимое не может быть пустым",
    "MailTemplateAddSuccess": "Шаблон письма успешно создан",
    "MailTemplateUpdateSuccess": "Шаблон письма успешно изменён",
    "NoEmailAddress": "Адрес электронной почты не задан",
    "VerificationCodeError": "Ошибка кода подтверждения",
    "UUIDEmpty": "UUID не может быть пустым"
  }
};

export default local;
