import enUs from './en-us';

const local: App.I18n.Schema = {
  ...enUs,
  system: {
    ...enUs.system,
    title: 'Rustdesk Api Server',
    updateTitle: 'Уведомление об обновлении версии',
    updateContent: 'Обнаружена новая версия системы. Обновить страницу сейчас?',
    updateConfirm: 'Обновить',
    updateCancel: 'Позже'
  },
  route: {
    ...enUs.route,
    login: 'Вход',
    403: 'Нет доступа',
    404: 'Страница не найдена',
    500: 'Ошибка сервера',
    home: 'Главная',
    audit: 'Аудит',
    user: 'Пользователи',
    user_list: 'Список пользователей',
    user_sessions: 'Сессии',
    system: 'Система',
    system_mail_template: 'Шаблоны почты',
    system_mail_logs: 'Логи почты',
    system_mail: 'Почта',
    audit_baselogs: 'Базовые логи',
    audit_filetransferlogs: 'Логи передачи файлов',
    devices: 'Устройства'
  },
  page: {
    ...enUs.page,
    home: {
      ...enUs.page.home,
      greeting: 'Доброе утро, {userName}, сегодня отличный день!',
      userCount: 'Пользователи',
      deviceCount: 'Устройства',
      onlineCount: 'Онлайн',
      visitsCount: 'Посещения',
      operatingSystem: 'Операционные системы',
      oneWeek: 'За неделю',
      changeLogs: 'Журнал изменений',
      serverConfig: {
        ...enUs.page.home.serverConfig,
        title: 'Конфигурация подключения клиента',
        tip: 'Скопируйте значения ниже в клиент RustDesk. Если KEY пустой, задайте переменную окружения `RUSTDESK_KEY`.',
        copyAll: 'Копировать все',
        copyTemplate: 'Копировать шаблон RustDesk',
        refresh: 'Обновить',
        clearCacheReload: 'Очистить кэш и перезагрузить',
        source: 'Источник',
        lastUpdated: 'Последнее обновление',
        show: 'Показать',
        hide: 'Скрыть',
        missingTip: 'Следующие поля пустые, сначала настройте их в переменных окружения контейнера: {fields}',
        copyEmpty: '{label} пусто, копирование невозможно',
        copySuccess: '{label} скопировано',
        copyFailed: 'Не удалось скопировать {label}',
        fetchFailed: 'Не удалось загрузить конфигурацию сервера',
        cacheCleared: 'Кэш очищен, повторная загрузка конфигурации сервера',
        sourceType: {
          ...enUs.page.home.serverConfig.sourceType,
          remote: 'Удалённый источник',
          'memory-cache': 'Кэш памяти',
          'session-cache': 'Кэш сессии',
          env: 'Переменная окружения',
          inferred: 'Автоопределение',
          empty: 'Пусто'
        },
        sourceHint: {
          ...enUs.page.home.serverConfig.sourceHint,
          env: 'Это значение получено из переменной окружения контейнера.',
          inferred: 'Это значение автоматически определено по текущему адресу доступа.',
          empty: 'Значение не настроено и не определено автоматически.'
        },
        connectivity: {
          ...enUs.page.home.serverConfig.connectivity,
          clear: 'Очистить результаты',
          check: 'Проверить доступность',
          checkOne: 'Проверить',
          checked: 'Проверка доступности завершена',
          checkedOne: 'Проверка {field} завершена',
          checkedCached: 'Использован недавний результат проверки (кэш)',
          checkFailed: 'Проверка доступности не удалась',
          cleared: 'Результаты проверки очищены',
          source: 'Источник проверки',
          lastChecked: 'Последняя проверка',
          target: 'Цель',
          duration: 'Время',
          notChecked: 'Ещё не проверялось',
          checkSourceType: {
            ...enUs.page.home.serverConfig.connectivity.checkSourceType,
            remote: 'Удалённая проверка',
            cache: 'Кэш'
          },
          status: {
            ...enUs.page.home.serverConfig.connectivity.status,
            idle: 'Не проверено',
            ok: 'Доступно',
            error: 'Ошибка',
            skip: 'Пропущено'
          }
        }
      }
    },
    user: {
      ...enUs.page.user,
      list: {
        ...enUs.page.user.list,
        addUser: 'Добавить пользователя',
        editUser: 'Редактировать пользователя',
        searchPlaceholder: 'Имя пользователя\\Ник\\Email'
      },
      sessions: {
        ...enUs.page.user.sessions,
        kill: 'Завершить',
        confirmKill: 'Завершить эту сессию?'
      },
      audit: {
        ...enUs.page.user.audit,
        logsSearchPlaceholder: 'Пользователь\\Действие\\RustdeskID\\IP'
      },
      devices: {
        ...enUs.page.user.devices,
        logsSearchPlaceholder: 'Пользователь\\Хост\\RustdeskID'
      }
    },
    system: {
      ...enUs.page.system,
      mailTemplate: {
        ...enUs.page.system.mailTemplate,
        addMailTemplate: 'Добавить шаблон',
        editMailTemplate: 'Редактировать шаблон',
        inputName: 'Введите имя',
        inputSubject: 'Введите тему',
        inputContents: 'Введите содержимое',
        selectType: 'Выберите тип'
      },
      mailLog: {
        ...enUs.page.system.mailLog,
        info: 'Подробности'
      }
    }
  },
  icon: {
    ...enUs.icon,
    lang: 'Сменить язык',
    reload: 'Перезагрузить страницу',
    fullscreen: 'Полный экран'
  },
  dataMap: {
    ...enUs.dataMap,
    user: {
      ...enUs.dataMap.user,
      username: 'Имя пользователя',
      name: 'Никнейм',
      email: 'Email',
      licensed_devices: 'Лицензированные устройства',
      login_verify: 'Проверка входа',
      status: 'Статус',
      is_admin: 'Администратор',
      created_at: 'Создано',
      statusLabel: {
        ...enUs.dataMap.user.statusLabel,
        disabled: 'Отключен',
        unverified: 'Не подтвержден',
        normal: 'Нормально'
      },
      loginVerifyLabel: {
        ...enUs.dataMap.user.loginVerifyLabel,
        none: 'Нет',
        emailCheck: 'Проверка email',
        tfaCheck: '2FA'
      }
    },
    device: {
      ...enUs.dataMap.device,
      username: 'Имя пользователя',
      hostname: 'Имя компьютера',
      version: 'Версия RustDesk',
      memory: 'Память',
      os: 'ОС',
      rustdesk_id: 'Rustdesk ID'
    },
    audit: {
      ...enUs.dataMap.audit,
      username: 'Пользователь',
      type: 'Тип',
      rustdesk_id: 'Rustdesk ID',
      ip: 'IP',
      created_at: 'Создано',
      typeLabel: {
        ...enUs.dataMap.audit.typeLabel,
        remote_control: 'Удалённое управление',
        file_transfer: 'Передача файлов',
        tcp_tunnel: 'TCP туннель'
      },
      fileTransferTypeLabel: {
        ...enUs.dataMap.audit.fileTransferTypeLabel,
        master_controlled: 'Управляющий -> Управляемый',
        controlled_master: 'Управляемый -> Управляющий'
      }
    },
    mailTemplate: {
      ...enUs.dataMap.mailTemplate,
      name: 'Имя',
      type: 'Тип',
      subject: 'Тема',
      contents: 'Содержимое',
      created_at: 'Создано',
      typeLabel: {
        ...enUs.dataMap.mailTemplate.typeLabel,
        loginVerify: 'Проверка входа',
        registerVerify: 'Проверка регистрации',
        other: 'Другое'
      }
    },
    mailLog: {
      ...enUs.dataMap.mailLog,
      username: 'Пользователь',
      from: 'От',
      to: 'Кому',
      subject: 'Тема',
      status: 'Статус',
      created_at: 'Время отправки',
      statusLabel: {
        ...enUs.dataMap.mailLog.statusLabel,
        ok: 'Успешно',
        err: 'Ошибка'
      }
    }
  },
  api: {
    ...enUs.api,
    CaptchaError: 'Ошибка CAPTCHA',
    UserNotExists: 'Пользователь не существует',
    UsernameOrPasswordError: 'Неверный логин или пароль',
    UserExists: 'Имя пользователя уже используется',
    UsernameEmpty: 'Имя пользователя не может быть пустым',
    PasswordEmpty: 'Пароль не может быть пустым',
    UserAddSuccess: 'Пользователь успешно создан',
    DataError: 'Ошибка данных',
    UserUpdateSuccess: 'Пользователь успешно изменён',
    UserDeleteSuccess: 'Пользователь успешно удалён',
    SessionKillSuccess: 'Сессия успешно завершена',
    MailTemplateNameEmpty: 'Имя не может быть пустым',
    MailTemplateSubjectEmpty: 'Тема не может быть пустой',
    MailTemplateContentsEmpty: 'Содержимое не может быть пустым',
    MailTemplateAddSuccess: 'Шаблон письма успешно создан',
    MailTemplateUpdateSuccess: 'Шаблон письма успешно изменён',
    NoEmailAddress: 'Адрес электронной почты не задан',
    VerificationCodeError: 'Ошибка кода подтверждения',
    UUIDEmpty: 'UUID не может быть пустым'
  }
};

export default local;
