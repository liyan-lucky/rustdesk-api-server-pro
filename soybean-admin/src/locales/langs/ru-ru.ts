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
      changeLogs: 'Журнал изменений'
    },
    user: {
      ...enUs.page.user,
      devices: {
        ...enUs.page.user.devices,
        logsSearchPlaceholder: 'Username\\Hostname\\RustdeskID'
      }
    }
  },
  icon: {
    ...enUs.icon,
    lang: 'Сменить язык'
  },
  dataMap: {
    ...enUs.dataMap,
    device: {
      ...enUs.dataMap.device,
      username: 'Имя пользователя',
      hostname: 'Имя компьютера',
      version: 'Версия RustDesk',
      memory: 'Память',
      os: 'ОС',
      rustdesk_id: 'Rustdesk ID'
    }
  }
};

export default local;
