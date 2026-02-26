import enUs from './en-us';

const local: App.I18n.Schema = {
  ...enUs,
  system: { ...enUs.system, title: 'Rustdesk Api Server' },
  common: {
    ...enUs.common,
    action: 'Aktion',
    add: 'Hinzuf?gen',
    addSuccess: 'Erfolgreich hinzugef?gt',
    backToHome: 'Zur Startseite',
    batchDelete: 'Stapel l?schen',
    cancel: 'Abbrechen',
    close: 'Schlie?en',
    check: 'Pr?fen',
    expandColumn: 'Spalte erweitern',
    columnSetting: 'Spalteneinstellungen',
    config: 'Konfiguration',
    confirm: 'Best?tigen',
    delete: 'L?schen',
    deleteSuccess: 'Erfolgreich gel?scht',
    confirmDelete: 'M?chten Sie wirklich l?schen?',
    edit: 'Bearbeiten',
    look: 'Anzeigen',
    warning: 'Warnung',
    error: 'Fehler',
    index: 'Index',
    keywordSearch: 'Bitte Schl?sselwort eingeben',
    logout: 'Abmelden',
    logoutConfirm: 'M?chten Sie sich wirklich abmelden?',
    lookForward: 'Demn?chst verf?gbar',
    modify: '?ndern',
    modifySuccess: 'Erfolgreich ge?ndert',
    noData: 'Keine Daten',
    operate: 'Vorgang',
    pleaseCheckValue: 'Bitte pr?fen Sie, ob der Wert g?ltig ist',
    refresh: 'Aktualisieren',
    reset: 'Zur?cksetzen',
    search: 'Suchen',
    switch: 'Umschalten',
    tip: 'Hinweis',
    trigger: 'Ausl?sen',
    update: 'Update',
    updateSuccess: 'Erfolgreich aktualisiert',
    userCenter: 'Benutzerzentrum',
    yesOrNo: {
      yes: 'Ja',
      no: 'Nein'
    }
  },
  route: {
    ...enUs.route,
    home: 'Startseite',
    audit: 'Audit',
    user: 'Benutzerverwaltung',
    user_list: 'Benutzerliste',
    user_sessions: 'Sitzungen',
    system: 'Systemverwaltung',
    system_mail_template: 'Mail-Vorlagen',
    system_mail_logs: 'Mail-Protokolle',
    system_mail: 'Mail',
    devices: 'Geräte'
  },
  page: {
    ...enUs.page,
    home: {
      ...enUs.page.home,
      greeting: 'Guten Morgen, {userName}!',
      userCount: 'Benutzer',
      deviceCount: 'Geräte',
      onlineCount: 'Online',
      visitsCount: 'Besuche',
      operatingSystem: 'Betriebssystem',
      oneWeek: 'Eine Woche',
      changeLogs: 'Änderungsprotokoll',
      serverConfig: {
        ...enUs.page.home.serverConfig,
        title: 'Client-Verbindungskonfiguration',
        tip: 'Kopieren Sie die folgenden Werte in den RustDesk-Client. Wenn KEY leer ist, setzen Sie `RUSTDESK_KEY`.',
        refresh: 'Aktualisieren',
        copyAll: 'Alles kopieren',
        copyTemplate: 'RustDesk-Vorlage kopieren',
        connectivity: {
          ...enUs.page.home.serverConfig.connectivity,
          check: 'Konnektivität prüfen',
          checkOne: 'Prüfen',
          clear: 'Ergebnisse löschen'
        }
      }
    },
    user: {
      ...enUs.page.user,
      list: {
        ...enUs.page.user.list,
        addUser: 'Benutzer hinzufügen',
        editUser: 'Benutzer bearbeiten',
        searchPlaceholder: 'Benutzername/Spitzname/E-Mail'
      },
      sessions: {
        ...enUs.page.user.sessions,
        kill: 'Beenden',
        confirmKill: 'Diese Sitzung beenden?'
      },
      audit: {
        ...enUs.page.user.audit,
        logsSearchPlaceholder: 'Benutzer/Aktion/RustdeskID/IP'
      },
      devices: {
        ...enUs.page.user.devices,
        logsSearchPlaceholder: 'Benutzer/Hostname/RustdeskID'
      }
    },
    system: {
      ...enUs.page.system,
      mailTemplate: {
        ...enUs.page.system.mailTemplate,
        addMailTemplate: 'Vorlage hinzufügen',
        editMailTemplate: 'Vorlage bearbeiten',
        inputName: 'Name eingeben',
        inputSubject: 'Betreff eingeben',
        inputContents: 'Inhalt eingeben',
        selectType: 'Typ auswählen'
      },
      mailLog: {
        ...enUs.page.system.mailLog,
        info: 'Details'
      }
    }
  },
  dataMap: {
    ...enUs.dataMap,
    user: {
      ...enUs.dataMap.user,
      username: 'Benutzername',
      name: 'Spitzname',
      email: 'E-Mail',
      licensed_devices: 'Lizenzierte Geräte',
      login_verify: 'Login-Prüfung',
      status: 'Status',
      is_admin: 'Admin',
      created_at: 'Erstellt am',
      statusLabel: {
        ...enUs.dataMap.user.statusLabel,
        disabled: 'Deaktiviert',
        unverified: 'Unbestätigt',
        normal: 'Normal'
      },
      loginVerifyLabel: {
        ...enUs.dataMap.user.loginVerifyLabel,
        none: 'Keine',
        emailCheck: 'E-Mail-Prüfung',
        tfaCheck: '2FA'
      }
    },
    device: {
      ...enUs.dataMap.device,
      username: 'Benutzername',
      hostname: 'Hostname',
      version: 'RustDesk-Version',
      os: 'OS',
      rustdesk_id: 'Rustdesk ID'
    },
    audit: {
      ...enUs.dataMap.audit,
      username: 'Benutzer',
      type: 'Typ',
      rustdesk_id: 'Rustdesk ID',
      ip: 'IP',
      created_at: 'Erstellt am',
      typeLabel: {
        ...enUs.dataMap.audit.typeLabel,
        remote_control: 'Fernsteuerung',
        file_transfer: 'Dateiübertragung',
        tcp_tunnel: 'TCP-Tunnel'
      },
      fileTransferTypeLabel: {
        ...enUs.dataMap.audit.fileTransferTypeLabel,
        master_controlled: 'Steuernd -> Gesteuert',
        controlled_master: 'Gesteuert -> Steuernd'
      }
    },
    mailTemplate: {
      ...enUs.dataMap.mailTemplate,
      name: 'Name',
      type: 'Typ',
      subject: 'Betreff',
      contents: 'Inhalt',
      created_at: 'Erstellt am',
      typeLabel: {
        ...enUs.dataMap.mailTemplate.typeLabel,
        loginVerify: 'Login-Verifizierung',
        registerVerify: 'Registrierungs-Verifizierung',
        other: 'Sonstiges'
      }
    },
    mailLog: {
      ...enUs.dataMap.mailLog,
      username: 'Benutzer',
      from: 'Von',
      to: 'An',
      subject: 'Betreff',
      status: 'Status',
      created_at: 'Gesendet am',
      statusLabel: {
        ...enUs.dataMap.mailLog.statusLabel,
        ok: 'Erfolg',
        err: 'Fehler'
      }
    }
  },
  api: {
    ...enUs.api,
    CaptchaError: 'CAPTCHA-Fehler',
    UserNotExists: 'Benutzer existiert nicht',
    UsernameOrPasswordError: 'Konto oder Passwort ist falsch',
    UserExists: 'Der Benutzername wird bereits verwendet',
    UsernameEmpty: 'Benutzername darf nicht leer sein',
    PasswordEmpty: 'Passwort darf nicht leer sein',
    UserAddSuccess: 'Benutzer erfolgreich erstellt',
    DataError: 'Datenfehler',
    UserUpdateSuccess: 'Benutzer erfolgreich geändert',
    UserDeleteSuccess: 'Benutzer erfolgreich gelöscht',
    SessionKillSuccess: 'Sitzung erfolgreich beendet',
    MailTemplateNameEmpty: 'Name darf nicht leer sein',
    MailTemplateSubjectEmpty: 'Betreff darf nicht leer sein',
    MailTemplateContentsEmpty: 'Inhalt darf nicht leer sein',
    MailTemplateAddSuccess: 'Mail-Vorlage erfolgreich erstellt',
    MailTemplateUpdateSuccess: 'Mail-Vorlage erfolgreich geändert',
    NoEmailAddress: 'Keine E-Mail-Adresse gesetzt',
    VerificationCodeError: 'Fehler beim Verifizierungscode',
    UUIDEmpty: 'UUID darf nicht leer sein'
  },
  icon: { ...enUs.icon, lang: 'Sprache wechseln', reload: 'Seite neu laden', fullscreen: 'Vollbild' }
};

export default local;
