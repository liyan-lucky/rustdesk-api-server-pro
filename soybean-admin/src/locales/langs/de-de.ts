import enUs from './en-us';

const local: App.I18n.Schema = {
  ...enUs,
  system: { ...enUs.system, title: 'Rustdesk Api Server' },
  common: {
    ...enUs.common,
    refresh: 'Aktualisieren',
    noData: 'Keine Daten',
    error: 'Fehler',
    warning: 'Warnung'
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
      created_at: 'Erstellt am'
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
      created_at: 'Erstellt am'
    },
    mailTemplate: {
      ...enUs.dataMap.mailTemplate,
      name: 'Name',
      type: 'Typ',
      subject: 'Betreff',
      contents: 'Inhalt',
      created_at: 'Erstellt am'
    },
    mailLog: {
      ...enUs.dataMap.mailLog,
      username: 'Benutzer',
      from: 'Von',
      to: 'An',
      subject: 'Betreff',
      status: 'Status',
      created_at: 'Gesendet am'
    }
  },
  icon: { ...enUs.icon, lang: 'Sprache wechseln', reload: 'Seite neu laden', fullscreen: 'Vollbild' }
};

export default local;
