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
  request: {
    ...enUs.request,
    logout: 'Benutzer nach fehlgeschlagener Anfrage abmelden',
    logoutMsg: 'Benutzerstatus ung?ltig, bitte erneut anmelden',
    logoutWithModal: 'Nach fehlgeschlagener Anfrage Dialog anzeigen und dann abmelden',
    logoutWithModalMsg: 'Benutzerstatus ung?ltig, bitte erneut anmelden',
    refreshToken: 'Token abgelaufen, Token wird aktualisiert',
    tokenExpired: 'Anfrage-Token ist abgelaufen'
  },
  theme: {
    ...enUs.theme,
    themeSchema: {
      ...enUs.theme.themeSchema,
      title: 'Designschema',
      light: 'Hell',
      dark: 'Dunkel',
      auto: 'System folgen'
    },
    grayscale: 'Graustufen',
    colourWeakness: 'Farbschw?che',
    layoutMode: {
      ...enUs.theme.layoutMode,
      title: 'Layoutmodus',
      vertical: 'Vertikales Men?',
      horizontal: 'Horizontales Men?',
      'vertical-mix': 'Vertikaler Mix-Modus',
      'horizontal-mix': 'Horizontaler Mix-Modus',
      reverseHorizontalMix: 'Position von Haupt- und Untermen?s umkehren'
    },
    recommendColor: 'Empfohlenen Farbalgorithmus anwenden',
    recommendColorDesc: 'Der empfohlene Farbalgorithmus bezieht sich auf',
    themeColor: {
      ...enUs.theme.themeColor,
      title: 'Theme-Farbe',
      primary: 'Prim?r',
      info: 'Info',
      success: 'Erfolg',
      warning: 'Warnung',
      error: 'Fehler',
      followPrimary: 'Prim?rfarbe folgen'
    },
    scrollMode: {
      ...enUs.theme.scrollMode,
      title: 'Scrollmodus',
      wrapper: 'Wrapper',
      content: 'Inhalt'
    },
    page: {
      ...enUs.theme.page,
      animate: 'Seitenanimation',
      mode: {
        ...enUs.theme.page.mode,
        title: 'Animationsmodus',
        fade: 'Einblenden',
        'fade-slide': 'Gleiten',
        'fade-bottom': 'Fade-Zoom',
        'fade-scale': 'Fade-Skalierung',
        'zoom-fade': 'Zoom-Fade',
        'zoom-out': 'Zoom-Out',
        none: 'Keine'
      }
    },
    fixedHeaderAndTab: 'Header und Tabs fixieren',
    header: {
      ...enUs.theme.header,
      height: 'Headerh?he',
      breadcrumb: {
        ...enUs.theme.header.breadcrumb,
        visible: 'Breadcrumb sichtbar',
        showIcon: 'Breadcrumb-Symbol sichtbar'
      }
    },
    tab: {
      ...enUs.theme.tab,
      visible: 'Tab sichtbar',
      cache: 'Tab-Cache',
      height: 'Tab-H?he',
      mode: {
        ...enUs.theme.tab.mode,
        title: 'Tab-Modus',
        chrome: 'Chrome',
        button: 'Button'
      }
    },
    sider: {
      ...enUs.theme.sider,
      inverted: 'Dunkle Seitenleiste',
      width: 'Seitenleistenbreite',
      collapsedWidth: 'Breite eingeklappt',
      mixWidth: 'Mix-Seitenleistenbreite',
      mixCollapsedWidth: 'Mix eingeklappt Breite',
      mixChildMenuWidth: 'Mix-Untermen?breite'
    },
    footer: {
      ...enUs.theme.footer,
      visible: 'Footer sichtbar',
      fixed: 'Footer fixieren',
      height: 'Footer-H?he',
      right: 'Rechter Footer'
    },
    watermark: {
      ...enUs.theme.watermark,
      visible: 'Wasserzeichen Vollbild sichtbar',
      text: 'Wasserzeichentext'
    },
    themeDrawerTitle: 'Theme-Konfiguration',
    pageFunTitle: 'Seitenfunktionen',
    configOperation: {
      ...enUs.theme.configOperation,
      copyConfig: 'Konfiguration kopieren',
      copySuccessMsg: 'Kopieren erfolgreich, bitte Variable "themeSettings" in "src/theme/settings.ts" ersetzen',
      resetConfig: 'Konfiguration zur?cksetzen',
      resetSuccessMsg: 'Zur?cksetzen erfolgreich'
    }
  },
  route: {
    ...enUs.route,
    login: 'Anmeldung',
    403: 'Keine Berechtigung',
    404: 'Seite nicht gefunden',
    500: 'Serverfehler',
    'iframe-page': 'Iframe',
    home: 'Startseite',
    audit: 'Audit',
    user: 'Benutzerverwaltung',
    user_list: 'Benutzerliste',
    user_sessions: 'Sitzungen',
    system: 'Systemverwaltung',
    system_mail_template: 'Mail-Vorlagen',
    system_mail_logs: 'Mail-Protokolle',
    system_mail: 'Mail',
    audit_baselogs: 'Basisprotokolle',
    audit_filetransferlogs: 'Datei?bertragungsprotokolle',
    devices: 'Ger?te'
  },
  page: {
    ...enUs.page,
      login: {
        ...enUs.page.login,
        common: {
          ...enUs.page.login.common,
          loginOrRegister: 'Anmelden / Registrieren',
          userNamePlaceholder: 'Benutzernamen eingeben',
          phonePlaceholder: 'Telefonnummer eingeben',
          codePlaceholder: 'Best?tigungscode eingeben',
          passwordPlaceholder: 'Passwort eingeben',
          confirmPasswordPlaceholder: 'Passwort erneut eingeben',
          codeLogin: 'Code-Anmeldung',
          confirm: 'Best?tigen',
          back: 'Zur?ck',
          validateSuccess: 'Pr?fung erfolgreich',
          loginSuccess: 'Anmeldung erfolgreich',
          welcomeBack: 'Willkommen zur?ck, {userName} !'
        },
        pwdLogin: {
          ...enUs.page.login.pwdLogin,
          title: 'Passwort-Anmeldung',
          rememberMe: 'Angemeldet bleiben'
        }
      },
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
        tip: 'Kopieren Sie die folgenden Werte in den RustDesk-Client. Wenn KEY leer ist, setzen Sie `RUSTDESK_KEY` als Container-Umgebungsvariable.',
        idServer: 'ID-Server',
        relayServer: 'Relay-Server',
        apiServer: 'API-Server',
        key: 'KEY',
        idServerPlaceholder: 'z. B. your.domain.com',
        relayServerPlaceholder: 'z. B. your.domain.com',
        apiServerPlaceholder: 'z. B. https://your.domain.com',
        keyPlaceholder: '?ber Umgebungsvariable RUSTDESK_KEY bereitstellen',
        copy: 'Kopieren',
        copyAll: 'Alles kopieren',
        copyTemplate: 'RustDesk-Vorlage kopieren',
        refresh: 'Konfiguration aktualisieren',
        clearCacheReload: 'Cache leeren & neu laden',
        source: 'Quelle',
        lastUpdated: 'Zuletzt aktualisiert',
        show: 'Anzeigen',
        hide: 'Verbergen',
        missingTip: 'Die folgenden Felder sind leer. Bitte zuerst in den Container-Umgebungsvariablen konfigurieren: {fields}',
        copyEmpty: '{label} ist leer und kann nicht kopiert werden',
        copySuccess: '{label} kopiert',
        copyFailed: '{label} konnte nicht kopiert werden',
        fetchFailed: 'Serverkonfiguration konnte nicht geladen werden',
        cacheCleared: 'Cache geleert, Serverkonfiguration wird neu geladen',
        sourceType: {
          ...enUs.page.home.serverConfig.sourceType,
          remote: 'Remote',
          'memory-cache': 'Speicher-Cache',
          'session-cache': 'Sitzungs-Cache',
          env: 'Umgebung',
          inferred: 'Abgeleitet',
          empty: 'Leer'
        },
        sourceHint: {
          ...enUs.page.home.serverConfig.sourceHint,
          env: 'Dieser Wert stammt aus einer Container-Umgebungsvariable.',
          inferred: 'Dieser Wert wurde aus der aktuellen Zugriffsadresse automatisch abgeleitet.',
          empty: 'Noch kein Wert konfiguriert oder ableitbar.'
        },
        connectivity: {
          ...enUs.page.home.serverConfig.connectivity,
          clear: 'Ergebnisse l?schen',
          check: 'Konnektivit?t pr?fen',
          checkOne: 'Pr?fen',
          checked: 'Konnektivit?tspr?fung abgeschlossen',
          checkedOne: 'Konnektivit?t von {field} gepr?ft',
          checkedCached: 'Letztes Pr?fergebnis aus Cache verwendet',
          checkFailed: 'Konnektivit?tspr?fung fehlgeschlagen',
          cleared: 'Konnektivit?tsergebnisse gel?scht',
          source: 'Pr?fquelle',
          lastChecked: 'Zuletzt gepr?ft',
          target: 'Ziel',
          duration: 'Dauer',
          notChecked: 'Noch nicht gepr?ft',
          checkSourceType: {
            ...enUs.page.home.serverConfig.connectivity.checkSourceType,
            remote: 'Remote',
            cache: 'Cache'
          },
          status: {
            ...enUs.page.home.serverConfig.connectivity.status,
            idle: 'Ungepr?ft',
            ok: 'Erreichbar',
            error: 'Fehlgeschlagen',
            skip: '?bersprungen'
          }
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
        password: 'Passwort',
      name: 'Spitzname',
      email: 'E-Mail',
      licensed_devices: 'Lizenzierte Geräte',
      login_verify: 'Login-Prüfung',
      status: 'Status',
      is_admin: 'Admin',
        tfa_secret: '2FA-Geheimnis',
        tfa_code: '2FA-Code',
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
      session: {
        ...enUs.dataMap.session,
        expired: 'L?uft ab am',
        created_at: 'Erstellt am'
      },
    device: {
      ...enUs.dataMap.device,
      username: 'Benutzername',
      hostname: 'Hostname',
      version: 'RustDesk-Version',
        memory: 'Speicher',
      os: 'OS',
      rustdesk_id: 'Rustdesk ID'
    },
    audit: {
      ...enUs.dataMap.audit,
      username: 'Benutzer',
      type: 'Typ',
        conn_id: 'Verbindungs-ID',
      rustdesk_id: 'Rustdesk ID',
        peer_id: 'Peer-ID',
      ip: 'IP',
        session_id: 'Sitzungs-ID',
        uuid: 'UUID',
      created_at: 'Erstellt am',
        closed_at: 'Geschlossen am',
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
      },
        path: 'Pfad'
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
        uuid: 'UUID',
      from: 'Von',
      to: 'An',
      subject: 'Betreff',
        contents: 'Inhalt',
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
  dropdown: {
    ...enUs.dropdown,
    closeCurrent: 'Aktuelles schlie?en',
    closeOther: 'Andere schlie?en',
    closeLeft: 'Links schlie?en',
    closeRight: 'Rechts schlie?en',
    closeAll: 'Alle schlie?en'
  },
  icon: {
    ...enUs.icon,
    themeConfig: 'Theme-Konfiguration',
    themeSchema: 'Theme-Schema',
    lang: 'Sprache wechseln',
    fullscreen: 'Vollbild',
    fullscreenExit: 'Vollbild verlassen',
    reload: 'Seite neu laden',
    collapse: 'Men? einklappen',
    expand: 'Men? ausklappen',
    pin: 'Anheften',
    unpin: 'L?sen'
  },
  datatable: {
    ...enUs.datatable,
    itemCount: 'Insgesamt {total} Eintr?ge'
  }
};

export default local;
