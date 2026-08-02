import enUs from './en-us';

const local: App.I18n.Schema = {
  ...enUs,
  system: { ...enUs.system, title: 'Rustdesk Api Server' },
  common: {
    ...enUs.common,
    action: 'Aktion',
    add: 'Hinzufügen',
    addSuccess: 'Erfolgreich hinzugefügt',
    backToHome: 'Zur Startseite',
    batchDelete: 'Stapel löschen',
    cancel: 'Abbrechen',
    close: 'Schließen',
    check: 'Prüfen',
    expandColumn: 'Spalte erweitern',
    columnSetting: 'Spalteneinstellungen',
    config: 'Konfiguration',
    confirm: 'Bestätigen',
    delete: 'löschen',
    deleteSuccess: 'Erfolgreich gelöscht',
    confirmDelete: 'Möchten Sie wirklich löschen?',
    edit: 'Bearbeiten',
    import: 'CSV importieren',
    export: 'CSV exportieren',
    look: 'Anzeigen',
    warning: 'Warnung',
    error: 'Fehler',
    index: 'Index',
    keywordSearch: 'Bitte Schlüsselwort eingeben',
    logout: 'Abmelden',
    logoutConfirm: 'Möchten Sie sich wirklich Abmelden?',
    lookForward: 'Demnächst verfügbar',
    modify: 'Ändern',
    modifySuccess: 'Erfolgreich geändert',
    noData: 'Keine Daten',
    operate: 'Vorgang',
    pleaseCheckValue: 'Bitte Prüfen Sie, ob der Wert gültig ist',
    refresh: 'Aktualisieren',
    reset: 'Zurücksetzen',
    search: 'Suchen',
    switch: 'Umschalten',
    tip: 'Hinweis',
    trigger: 'Auslösen',
    update: 'Aktualisieren',
    updateSuccess: 'Aktualisierung erfolgreich',
    userCenter: 'Benutzerzentrum',
    yesOrNo: {
      yes: 'Ja',
      no: 'Nein'
    }
  },
  request: {
    ...enUs.request,
    logout: 'Benutzer nach fehlgeschlagener Anfrage Abmelden',
    logoutMsg: 'Benutzerstatus ungültig, bitte erneut anmelden',
    logoutWithModal: 'Nach fehlgeschlagener Anfrage Dialog anzeigen und dann Abmelden',
    logoutWithModalMsg: 'Benutzerstatus ungültig, bitte erneut anmelden',
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
    colourWeakness: 'Farbschwäche',
    layoutMode: {
      ...enUs.theme.layoutMode,
      title: 'Layoutmodus',
      vertical: 'Vertikales Menü',
      horizontal: 'Horizontales Menü',
      'vertical-mix': 'Vertikaler Mix-Modus',
      'horizontal-mix': 'Horizontaler Mix-Modus',
      reverseHorizontalMix: 'Position von Haupt- und UnterMenüs umkehren'
    },
    recommendColor: 'Empfohlenen Farbalgorithmus anwenden',
    recommendColorDesc: 'Der empfohlene Farbalgorithmus bezieht sich auf',
    themeColor: {
      ...enUs.theme.themeColor,
      title: 'Theme-Farbe',
      primary: 'Primär',
      info: 'Info',
      success: 'Erfolg',
      warning: 'Warnung',
      error: 'Fehler',
      followPrimary: 'Primärfarbe folgen'
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
      height: 'Headerhöhe',
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
      height: 'Tab-Höhe',
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
      mixChildMenuWidth: 'Mix-UnterMenübreite'
    },
    footer: {
      ...enUs.theme.footer,
      visible: 'Footer sichtbar',
      fixed: 'Footer fixieren',
      height: 'Footerhöhe',
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
      resetConfig: 'Konfiguration Zurücksetzen',
      resetSuccessMsg: 'Zurücksetzen erfolgreich'
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
    user_profile: 'Profil',
    system: 'Systemverwaltung',
    system_mail_template: 'Mail-Vorlagen',
    system_mail_logs: 'Mail-Protokolle',
    system_mail: 'Mail',
    audit_baselogs: 'Basisprotokolle',
    audit_filetransferlogs: 'Dateiuebertragungsprotokolle',
    audit_loginlogs: 'Anmeldeprotokolle',
    devices: 'Geraete',
    'my-devices': 'Kontakte',
    'my-devices_peers': 'Meine Kontakte',
    'my-devices_manage': 'Adressbuchverwaltung',
    'my-devices_tags': 'Tag-Verwaltung',
    workspace: 'Mein Arbeitsbereich', workspace_overview: 'Übersicht', workspace_devices: 'Meine Geräte', workspace_sessions: 'Anmeldesitzungen', workspace_security: 'Sicherheitsereignisse', workspace_profile: 'Profil', about: 'Über & Updates',
    system_server: 'Serverkonfiguration',
    system_tokens: 'Benutzer-Token',
    system_oauth: 'OAuth-Verwaltung'
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
          codePlaceholder: 'Bestätigungscode eingeben',
          passwordPlaceholder: 'Passwort eingeben',
          confirmPasswordPlaceholder: 'Passwort erneut eingeben',
          codeLogin: 'Code-Anmeldung',
          confirm: 'Bestätigen',
          back: 'Zurück',
          validateSuccess: 'Prüfung erfolgreich',
          loginSuccess: 'Anmeldung erfolgreich',
          welcomeBack: 'Willkommen Zurück, {userName} !',
          thirdPartyLogin: 'Drittanbieter-Anmeldung',
          continueWith: 'Mit {provider} fortfahren',
          providerUnavailable: '{provider}-Anmeldung ist derzeit nicht verfügbar'
        },
        pwdLogin: {
          ...enUs.page.login.pwdLogin,
          title: 'Passwort-Anmeldung',
          rememberMe: 'Angemeldet bleiben',
          switchToUser: 'Benutzer-Anmeldung'
         },
         userLogin: {
           title: 'Benutzer-Anmeldung',
           switchToAdmin: 'Admin-Anmeldung'
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
      cardDetail: {
        viewHint: 'Klicken, um Details anzuzeigen',
        recentUsers: 'Neueste Benutzer',
        recentDevices: 'Neueste Geraete',
        recentVisits: 'Neueste Zugriffsprotokolle',
        desc: {
          userCount: 'Zeigt die Gesamtzahl der Benutzer im System.',
          deviceCount: 'Zeigt die Gesamtzahl der Geraete im System.',
          onlineCount: 'Zeigt die Anzahl online Geraete basierend auf Heartbeat-Statistiken.',
          visitCount: 'Zeigt Besuchsstatistiken aus Audit-Logs.'
        }
      },
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
        keyPlaceholder: 'Über Umgebungsvariable RUSTDESK_KEY bereitstellen',
        copy: 'Kopieren',
        copyAll: 'Alles kopieren',
        copyTemplate: 'RustDesk-Vorlage kopieren',
        showQr: 'QR-Code anzeigen',
        qrTitle: 'RustDesk-Import-QR-Code',
        qrTip: 'Scannen Sie diesen QR-Code in der RustDesk-Mobile-App, um die Konfiguration zu importieren.',
        qrPayload: 'RustDesk-Vorlagentext',
        qrFailed: 'QR-Code konnte nicht generiert werden',
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
          empty: 'Leer',
           auto: 'Automatisch erkannt'
        },
        sourceHint: {
          ...enUs.page.home.serverConfig.sourceHint,
          env: 'Dieser Wert stammt aus einer Container-Umgebungsvariable.',
          inferred: 'Dieser Wert wurde aus der aktuellen Zugriffsadresse automatisch abgeleitet.',
          empty: 'Noch kein Wert konfiguriert oder ableitbar.'
        },
        connectivity: {
          ...enUs.page.home.serverConfig.connectivity,
          clear: 'Ergebnisse löschen',
          check: 'Konnektivität Prüfen',
          checkOne: 'Prüfen',
          checked: 'KonnektivitätsPrüfung abgeschlossen',
          checkedOne: 'Konnektivität von {field} geprüft',
          checkedCached: 'Letztes Prüfergebnis aus Cache verwendet',
          checkFailed: 'KonnektivitätsPrüfung fehlgeschlagen',
          cleared: 'Konnektivitätsergebnisse gelöscht',
          source: 'Prüfquelle',
          lastChecked: 'Zuletzt geprüft',
          target: 'Ziel',
          duration: 'Dauer',
          notChecked: 'Noch nicht geprüft',
          checkSourceType: {
            ...enUs.page.home.serverConfig.connectivity.checkSourceType,
            remote: 'Remote',
            cache: 'Cache'
          },
          status: {
            ...enUs.page.home.serverConfig.connectivity.status,
            idle: 'Ungeprüft',
            ok: 'Erreichbar',
            error: 'Fehlgeschlagen',
            skip: 'Übersprungen'
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
    },
    myDevices: {
      title: 'Kontakte',
      welcome: 'Willkommen, {userName}',
      status: 'Status',
      online: 'Online',
      offline: 'Offline',
      conns: 'Verbindungen',
      lastSync: 'Letzte Synchronisierung',
      logout: 'Abmelden'
    },
    workspace: {
      ...enUs.page.workspace,
      scopeTitle: 'Persönlicher Arbeitsbereich', scopeTip: 'Hier werden nur Ihre Geräte, Sitzungen, Sicherheitsereignisse und freigegebenen Adressbücher angezeigt.', myDevices: 'Meine Geräte', activeSessions: 'Aktive Sitzungen', addressBooks: 'Adressbücher', securityEvents: 'Sicherheitsereignisse', currentSession: 'Aktuelle Sitzung', revokeConfirm: 'Diese Anmeldesitzung widerrufen?', revoke: 'Widerrufen', accountRole: 'Kontorolle', adminRole: 'Administrator', userRole: 'Benutzer', permissionScope: 'Berechtigungsumfang', userScope: 'Persönliche Ressourcen und ausdrücklich freigegebene Adressbücher', active: 'Aktiv'
    },
    about: {
      runningVersion: 'Laufende Version', buildTime: 'Build-Zeit', compatVersion: 'Kompatible RustDesk-Version', latestVersion: 'Neueste Version', updateAvailable: 'Update verfügbar', upToDate: 'Aktuell', updateCheck: 'Online-Updateprüfung', urlTip: 'Die Prüfadresse kann geändert und in diesem Browser gespeichert werden. Sie darf eine semantische Version als Text oder JSON mit version, latest_version, tag_name oder server.version liefern. Die Zielseite muss Browser-CORS-Anfragen erlauben.', urlPlaceholder: 'Adresse für die Updateprüfung', checkNow: 'Jetzt prüfen', restoreDefault: 'Standard wiederherstellen', checkFailed: 'Updateprüfung fehlgeschlagen', invalidUrl: 'Nur HTTP- und HTTPS-Adressen werden unterstützt', invalidResponse: 'In der Antwort wurde keine semantische Version gefunden', updateCommand: 'Container-Aktualisierungsbefehl', commandTip: 'Die Befehlsvorlage ist editierbar; {version} wird durch die neueste Version ersetzt.', copyCommand: 'Befehl kopieren'
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
        expired: 'Läuft ab am',
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
    ab: {
      ...enUs.dataMap.ab,
      rustdesk_id: 'Rustdesk ID',
      username: 'Benutzername',
      hostname: 'Hostname',
      tags: 'Schlagwörter',
      alias: 'Alias',
      hash: 'Hash',
      owner: 'Eigentümer',
      name: 'Adressbuchname',
      user_id: 'Benutzer-ID',
      guid: 'GUID',
      rule: 'Berechtigung',
      max_peer: 'Max. Kontakte',
      shared: 'Geteilt',
      ab_id: 'Adressbuch-ID',
      tagName: 'Name',
      tagColor: 'Farbe',
      updated_at: 'Aktualisiert am',
      personal: 'Mein Adressbuch', note: 'Notiz', platform: 'Plattform', personalReadOnly: 'Persönlich (schreibgeschützt)', nameRequired: 'Name ist erforderlich', deviceIdRequired: 'Geräte-ID ist erforderlich', tagsHint: 'Mehrere Tags durch Kommas trennen', read: 'Lesen', readWrite: 'Lesen und schreiben', fullControl: 'Vollzugriff'
    },
    token: {
      device_os: 'Gerätebetriebssystem',
      device_name: 'Gerätename',
      token_hash: 'Token-Hash',
      is_admin: 'Admin',
      status: 'Aktiv'
    },
    oauth: {
      provider: 'Anbieter',
      subject: 'Subjekt',
      email: 'E-Mail',
      name: 'Name',
      last_login_at: 'Letzte Anmeldung'
    },
    loginLog: {
      allEvents: 'Alle Ereignisse',
      event: 'Ereignis',
      userAgent: 'Benutzer-Agent',
      success: 'Erfolg',
      reason: 'Grund'
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
    RequestError: 'Anfrage fehlgeschlagen',
    UserUpdateSuccess: 'Benutzer erfolgreich aktualisiert',
    UserDeleteSuccess: 'Benutzer erfolgreich gelöscht',
    SessionKillSuccess: 'Sitzung erfolgreich beendet',
    MailTemplateNameEmpty: 'Name darf nicht leer sein',
    MailTemplateSubjectEmpty: 'Betreff darf nicht leer sein',
    MailTemplateContentsEmpty: 'Inhalt darf nicht leer sein',
    MailTemplateAddSuccess: 'Mail-Vorlage erfolgreich erstellt',
    MailTemplateUpdateSuccess: 'E-Mail-Vorlage erfolgreich aktualisiert',
    NoEmailAddress: 'Keine E-Mail-Adresse gesetzt',
    VerificationCodeError: 'Fehler beim Verifizierungscode',
    UUIDEmpty: 'UUID darf nicht leer sein'
  },
  dropdown: {
    ...enUs.dropdown,
    closeCurrent: 'Aktuelles Schließen',
    closeOther: 'Andere Schließen',
    closeLeft: 'Links Schließen',
    closeRight: 'Rechts Schließen',
    closeAll: 'Alle Schließen'
  },
  icon: {
    ...enUs.icon,
    themeConfig: 'Theme-Konfiguration',
    themeSchema: 'Theme-Schema',
    lang: 'Sprache wechseln',
    fullscreen: 'Vollbild',
    fullscreenExit: 'Vollbild verlassen',
    reload: 'Seite neu laden',
    collapse: 'Menü einklappen',
    expand: 'Menü ausklappen',
    pin: 'Anheften',
    unpin: 'Lösen'
  },
  datatable: {
    ...enUs.datatable,
    itemCount: 'Insgesamt {total} Einträge'
  }
};

export default local;
