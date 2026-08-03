const local: App.I18n.Schema = {
  "system": {
    "title": "Rustdesk Api Server",
    "updateTitle": "Обновление версии системы",
    "updateContent": "Доступна новая версия системы. Обновить страницу сейчас?",
    "updateConfirm": "Обновить сейчас",
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
    "update": "Обновить",
    "updateSuccess": "Обновление выполнено",
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
    "login": "Вход",
    "iframe-page": "Встроенная страница",
    "home": "Главная", "about": "О программе и обновления",
    "audit": "Аудит",
    "user": "Управление пользователями",
    "user_list": "Список пользователей",
    "user_sessions": "Сессии",
    "user_profile": "Профиль",
    "system": "Управление системой",
    "system_mail_template": "Шаблоны писем",
    "system_mail_logs": "Логи почты",
    "system_mail": "Управление почтой",
    "system_server": "Конфигурация сервера",
    "system_tokens": "Токены пользователей",
    "system_oauth": "Сторонний вход",
    "audit_baselogs": "Базовые логи",
    "audit_filetransferlogs": "Логи передачи файлов",
    "audit_loginlogs": "Логи входа",
    "devices": "Устройства",
    "my-devices": "Контакты",
    "my-devices_peers": "Мои пиры",
    "my-devices_manage": "Управление адресной книгой",
    "my-devices_tags": "Управление тегами", "workspace": "Моё пространство", "workspace_overview": "Обзор", "workspace_devices": "Мои устройства", "workspace_sessions": "Сеансы входа", "workspace_security": "События безопасности", "workspace_profile": "Профиль"
  },
  "page": {
    "login": {
      "common": {
        "loginOrRegister": "Вход / Регистрация",
        "userNamePlaceholder": "Введите имя пользователя",
        "phonePlaceholder": "Please enter phone number",
        "codePlaceholder": "Введите код подтверждения",
        "passwordPlaceholder": "Введите пароль",
        "confirmPasswordPlaceholder": "Введите пароль еще раз",
        "codeLogin": "Вход по коду подтверждения",
        "confirm": "Подтвердить",
        "back": "Назад",
        "validateSuccess": "Проверка пройдена",
        "loginSuccess": "Вход выполнен успешно",
        "welcomeBack": "С возвращением, {userName}!",
        "thirdPartyLogin": "Вход через сторонний сервис",
        "continueWith": "Продолжить через {provider}",
        "providerUnavailable": "Вход через {provider} сейчас недоступен"
      },
      "pwdLogin": {
        "title": "Вход по паролю",
        "rememberMe": "Запомнить меня",
        "switchToUser": "Вход пользователя"
      },
      "userLogin": {
        "title": "Вход пользователя",
        "switchToAdmin": "Вход администратора"
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
      "cardDetail": {
        "viewHint": "Нажмите, чтобы посмотреть детали",
        "recentUsers": "Недавние пользователи",
        "recentDevices": "Недавние устройства",
        "recentVisits": "Недавние записи посещений",
        "desc": {
          "userCount": "Показывает общее количество пользователей в системе.",
          "deviceCount": "Показывает общее количество устройств в системе.",
          "onlineCount": "Показывает число онлайн-устройств по статистике heartbeat.",
          "visitCount": "Показывает статистику посещений из журналов аудита."
        }
      },
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
        "showQr": "Показать QR-код",
        "qrTitle": "QR-код импорта RustDesk",
        "qrTip": "Отсканируйте этот QR-код в мобильном приложении RustDesk для импорта конфигурации.",
        "qrPayload": "Текст шаблона RustDesk",
        "qrFailed": "Ошибка генерации QR-кода",
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
        "cacheTtlHint": "TTL cache: config {configSeconds}s, connectivity {connectivitySeconds}s",
        "ageSeconds": "{seconds}s ago",
        "sourceType": {
          "remote": "Удалённый источник",
          "memory-cache": "Кэш памяти",
          "session-cache": "Кэш сессии",
          "env": "Переменная окружения",
          "inferred": "Автоопределение",
          "empty": "Пусто",
          "auto": "Автообнаружение"
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
        "inputUsername": "Введите имя пользователя",
        "inputPassword": "Введите пароль",
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
    },
    "myDevices": {
      "title": "Контакты",
      "welcome": "Добро пожаловать, {userName}",
      "status": "Статус",
      "online": "В сети",
      "offline": "Не в сети",
      "conns": "Соединения",
      "lastSync": "Последняя синхронизация",
      "logout": "Выход"
    },
    "workspace": {
      "scopeTitle": "Личное пространство", "scopeTip": "Здесь показаны только ваши устройства, сеансы, события безопасности и разрешённые адресные книги.", "myDevices": "Мои устройства", "activeSessions": "Активные сеансы", "addressBooks": "Адресные книги", "securityEvents": "События безопасности", "currentSession": "Текущий сеанс", "revokeConfirm": "Отозвать этот сеанс входа?", "revoke": "Отозвать", "accountRole": "Роль учётной записи", "adminRole": "Администратор", "userRole": "Пользователь", "permissionScope": "Область доступа", "userScope": "Личные ресурсы и явно предоставленные адресные книги", "active": "Активен"
    },
    "oauth": {
      "configTitle": "Настройка стороннего входа", "bindingsTitle": "Привязки учетных записей", "addProvider": "Добавить провайдера", "editProvider": "Изменить провайдера", "providerName": "Ключ провайдера", "displayName": "Отображаемое имя", "clientId": "Идентификатор клиента", "clientSecret": "Секрет клиента", "secretPlaceholder": "Оставьте пустым, чтобы сохранить настроенный секрет", "redirectUrl": "Адрес обратного вызова", "scopes": "Области доступа", "accountRole": "Роль учетной записи", "allowedDomains": "Разрешенные почтовые домены", "bindByEmail": "Привязать по подтвержденной почте", "autoCreateAdmin": "Автоматически создавать администратора", "autoCreateUser": "Автоматически создавать пользователя", "testConfig": "Проверить настройку", "testSuccess": "Настройка заполнена, адрес авторизации создан", "copyCallback": "Копировать адрес возврата", "githubOnlyTip": "Сначала доступен GitHub. Настройте его здесь; server.yaml остается для совместимости и восстановления.", "adminRole": "Администратор", "userRole": "Пользователь", "useDefault": "Использовать значение по умолчанию", "listPlaceholder": "Разделяйте значения пробелами или запятыми", "copied": "Скопировано"
    },
    "about": {
      "latestCommand": "Обновить до latest", "pinnedCommand": "Обновить и проверить найденную версию", "customCommand": "Свой шаблон команды",
      "runningVersion": "Текущая версия", "buildTime": "Время сборки", "compatVersion": "Совместимая версия RustDesk", "latestVersion": "Последняя версия", "updateAvailable": "Доступно обновление", "upToDate": "Установлена последняя версия", "updateCheck": "Проверка обновлений", "urlTip": "Адрес проверки можно изменить; он сохраняется в этом браузере. Сайт должен разрешать CORS.", "urlPlaceholder": "Адрес проверки обновлений", "checkNow": "Проверить", "restoreDefault": "Вернуть адрес по умолчанию", "checkFailed": "Ошибка проверки обновлений", "invalidUrl": "Поддерживаются только HTTP и HTTPS", "invalidResponse": "Версия не найдена в ответе", "updateCommand": "Команда обновления контейнера", "commandTip": "Шаблон можно изменить; {version} заменяется последней версией.", "copyCommand": "Копировать команду"
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
    "lang": "Сменить язык",
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
      "personal": "Моя адресная книга", "legacy": "Устаревшая адресная книга", "note": "Примечание", "platform": "Платформа", "personalReadOnly": "Личная (только чтение)", "nameRequired": "Имя обязательно", "deviceIdRequired": "ID устройства обязателен", "tagsHint": "Разделяйте метки запятыми", "read": "Чтение", "readWrite": "Чтение и запись", "fullControl": "Полный доступ"
    },
    "token": {
      "device_os": "ОС устройства",
      "device_name": "Имя устройства",
      "token_hash": "Хеш токена",
      "is_admin": "Админ",
      "status": "Активен"
    },
    "oauth": {
      "provider": "Провайдер",
      "subject": "Субъект",
      "email": "Почта",
      "name": "Имя",
      "last_login_at": "Последний вход"
    },
    "loginLog": {
      "allEvents": "Все события",
      "event": "Событие",
      "userAgent": "Агент пользователя",
      "success": "Успех",
      "reason": "Причина"
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
    "RequestError": "Ошибка запроса",
    "UserUpdateSuccess": "Пользователь успешно обновлён",
    "UserDeleteSuccess": "Пользователь успешно удалён",
    "SessionKillSuccess": "Сессия успешно завершена",
    "MailTemplateNameEmpty": "Имя не может быть пустым",
    "MailTemplateSubjectEmpty": "Тема не может быть пустой",
    "MailTemplateContentsEmpty": "Содержимое не может быть пустым",
    "MailTemplateAddSuccess": "Шаблон письма успешно создан",
    "MailTemplateUpdateSuccess": "Почтовый шаблон успешно обновлён",
    "NoEmailAddress": "Адрес электронной почты не задан",
    "VerificationCodeError": "Ошибка кода подтверждения",
    "UUIDEmpty": "UUID не может быть пустым"
  }
};

export default local;
