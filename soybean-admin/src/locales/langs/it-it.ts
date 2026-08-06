import enUs from './en-us';

const local: App.I18n.Schema = {
  ...enUs,
  system: {
    ...enUs.system,
    title: 'Rustdesk Api Server',
    updateTitle: 'Aggiornamento versione di sistema',
    updateContent: 'Una nuova versione di sistema è disponibile. Aggiornare la pagina ora?',
    updateConfirm: 'Aggiorna ora',
    updateCancel: 'Più tardi'
  },
  common: {
    ...enUs.common,
    action: 'Azione',
    add: 'Aggiungi',
    addSuccess: 'Aggiunto con successo',
    backToHome: 'Torna alla home',
    batchDelete: 'Elimina in blocco',
    cancel: 'Annulla',
    close: 'Chiudi',
    check: 'Controlla',
    expandColumn: 'Espandi colonna',
    columnSetting: 'Impostazioni colonna',
    config: 'Configura',
    confirm: 'Conferma',
    delete: 'Elimina',
    deleteSuccess: 'Eliminato con successo',
    confirmDelete: 'Sei sicuro di voler eliminare?',
    edit: 'Modifica',
    import: 'Importa CSV',
    export: 'Esporta CSV',
    look: 'Visualizza',
    warning: 'Avviso',
    error: 'Errore',
    index: 'Indice',
    keywordSearch: 'Inserisci parola chiave',
    logout: 'Disconnetti',
    logoutConfirm: 'Sei sicuro di voler uscire?',
    lookForward: 'Prossimamente',
    modify: 'Modifica',
    modifySuccess: 'Modificato con successo',
    noData: 'Nessun dato',
    operate: 'Operazione',
    pleaseCheckValue: 'Controlla se il valore è valido',
    refresh: 'Aggiorna',
    reset: 'Reimposta',
    search: 'Cerca',
    switch: 'Cambia',
    tip: 'Suggerimento',
    trigger: 'Attiva',
    update: 'Aggiorna',
    updateSuccess: 'Aggiornamento riuscito',
    userCenter: 'Centro utente',
    yesOrNo: {
      yes: 'Sì',
      no: 'No'
    }
  },
  request: {
    ...enUs.request,
    logout: 'Disconnetti utente dopo errore richiesta',
    logoutMsg: 'Stato utente non valido, effettua nuovamente l’accesso',
    logoutWithModal: 'Mostra finestra dopo errore richiesta e poi disconnetti utente',
    logoutWithModalMsg: 'Stato utente non valido, effettua nuovamente l’accesso',
    refreshToken: 'Il token richiesto è scaduto, aggiorna il token',
    tokenExpired: 'Il token richiesto è scaduto'
  },
  theme: {
    ...enUs.theme,
    themeSchema: {
      ...enUs.theme.themeSchema,
      title: 'Schema tema',
      light: 'Chiaro',
      dark: 'Scuro',
      auto: 'Segui sistema'
    },
    grayscale: 'Scala di grigi',
    colourWeakness: 'Daltonismo',
    layoutMode: {
      ...enUs.theme.layoutMode,
      title: 'Modalità layout',
      vertical: 'Menu verticale',
      horizontal: 'Menu orizzontale',
      'vertical-mix': 'Menu misto verticale',
      'horizontal-mix': 'Menu misto orizzontale',
      reverseHorizontalMix: 'Inverti posizione menu primo livello e sottomenu'
    },
    recommendColor: 'Applica algoritmo colore consigliato',
    recommendColorDesc: 'L’algoritmo colore consigliato fa riferimento a',
    themeColor: {
      ...enUs.theme.themeColor,
      title: 'Colore tema',
      primary: 'Primario',
      info: 'Informazioni',
      success: 'Successo',
      warning: 'Avviso',
      error: 'Errore',
      followPrimary: 'Segui primario'
    },
    scrollMode: {
      ...enUs.theme.scrollMode,
      title: 'Modalità scorrimento',
      wrapper: 'Contenitore',
      content: 'Contenuto'
    },
    page: {
      ...enUs.theme.page,
      animate: 'Animazione pagina',
      mode: {
        ...enUs.theme.page.mode,
        title: 'Modalità animazione pagina',
        fade: 'Dissolvenza',
        'fade-slide': 'Scorrimento',
        'fade-bottom': 'Dissolvenza zoom',
        'fade-scale': 'Dissolvenza scala',
        'zoom-fade': 'Zoom dissolvenza',
        'zoom-out': 'Zoom indietro',
        none: 'Nessuna'
      }
    },
    fixedHeaderAndTab: 'Intestazione e scheda fisse',
    header: {
      ...enUs.theme.header,
      height: 'Altezza intestazione',
      breadcrumb: {
        visible: 'Breadcrumbs visibili',
        showIcon: 'Icona breadcrumbs visibile'
      }
    },
    tab: {
      ...enUs.theme.tab,
      visible: 'Scheda visibile',
      cache: 'Cache scheda',
      height: 'Altezza scheda',
      mode: {
        ...enUs.theme.tab.mode,
        title: 'Modalità scheda',
        chrome: 'Chrome',
        button: 'Pulsante'
      }
    },
    sider: {
      ...enUs.theme.sider,
      inverted: 'Barra laterale scura',
      width: 'Larghezza barra laterale',
      collapsedWidth: 'Larghezza barra laterale compressa',
      mixWidth: 'Larghezza barra laterale mista',
      mixCollapsedWidth: 'Larghezza barra laterale mista compressa',
      mixChildMenuWidth: 'Larghezza sottomenu misto'
    },
    footer: {
      ...enUs.theme.footer,
      visible: 'Piè di pagina visibile',
      fixed: 'Piè di pagina fisso',
      height: 'Altezza piè di pagina',
      right: 'Piè di pagina destro'
    },
    watermark: {
      ...enUs.theme.watermark,
      visible: 'Filigrana a schermo intero visibile',
      text: 'Testo filigrana'
    },
    themeDrawerTitle: 'Configurazione tema',
    pageFunTitle: 'Funzione pagina',
    configOperation: {
      ...enUs.theme.configOperation,
      copyConfig: 'Copia configurazione',
      copySuccessMsg: 'Copia riuscita, sostituisci la variabile "themeSettings" in "src/theme/settings.ts"',
      resetConfig: 'Reimposta configurazione',
      resetSuccessMsg: 'Reimpostazione riuscita'
    }
  },
  route: {
    ...enUs.route,
    login: 'Accesso',
    403: 'Nessun permesso',
    404: 'Pagina non trovata',
    500: 'Errore server',
    'iframe-page': 'Iframe',
    home: 'Home',
    about: 'Informazioni',
    about_version: 'Informazioni versione',
    about_help: 'Guida codici di errore',
    audit: 'Audit',
    user: 'Gestione utenti',
    user_list: 'Elenco utenti',
    user_sessions: 'Sessioni',
    user_profile: 'Profilo',
    system: 'Gestione sistema',
    system_mail_template: 'Modello email',
    system_mail_logs: 'Log email',
    system_mail: 'Gestione email',
    system_server: 'Configurazione server',
    system_tokens: 'Token utente',
    system_oauth: 'Accesso terze parti',
    audit_baselogs: 'Log di base',
    audit_filetransferlogs: 'Log trasferimento file',
    audit_loginlogs: 'Log di accesso',
    'audit_error-logs': 'Log degli errori',
    'audit_system-logs': 'Log di sistema',
    devices: 'Dispositivi',
    'my-devices': 'Contatti',
    'my-devices_peers': 'Miei peer',
    'my-devices_manage': 'Gestione rubrica',
    'my-devices_tags': 'Gestione tag',
    workspace: 'Area di lavoro',
    workspace_overview: 'Panoramica',
    workspace_devices: 'Miei dispositivi',
    workspace_sessions: 'Sessioni di accesso',
    workspace_security: 'Eventi di sicurezza',
    workspace_profile: 'Profilo'
  },
  page: {
    ...enUs.page,
    login: {
      ...enUs.page.login,
      common: {
        ...enUs.page.login.common,
        loginOrRegister: 'Accedi / Registrati',
        userNamePlaceholder: 'Inserisci nome utente',
        phonePlaceholder: 'Inserisci numero di telefono',
        codePlaceholder: 'Inserisci codice di verifica',
        passwordPlaceholder: 'Inserisci password',
        confirmPasswordPlaceholder: 'Inserisci nuovamente la password',
        codeLogin: 'Accesso con codice',
        confirm: 'Conferma',
        back: 'Indietro',
        validateSuccess: 'Verifica superata',
        loginSuccess: 'Accesso riuscito',
        welcomeBack: 'Bentornato, {userName}!',
        thirdPartyLogin: 'Accesso terze parti',
        continueWith: '{provider}',
        providerUnavailable: 'L’accesso {provider} non è disponibile',
        oauthAccountNotBound:
          'Nessun account corrispondente può essere associato. Aggiungi la stessa email verificata all’account di destinazione o abilita la creazione automatica degli account.',
        oauthProviderUnreachable:
          'Il server non può raggiungere il provider di accesso. Verifica la connettività HTTPS in uscita e riprova.',
        oauthStateExpired: 'La richiesta di accesso è scaduta o è già stata utilizzata. Avvia nuovamente l’accesso.',
        oauthAuthFailed: 'Accesso terze parti fallito. Verifica la configurazione del provider e il log di audit.'
      },
      pwdLogin: {
        ...enUs.page.login.pwdLogin,
        title: 'Accesso con password',
        rememberMe: 'Ricordami',
        switchToUser: 'Accesso utente'
      },
      userLogin: {
        title: 'Accesso utente',
        switchToAdmin: 'Accesso amministratore'
      }
    },
    home: {
      ...enUs.page.home,
      greeting: 'Buongiorno, {userName}!',
      userCount: 'Utenti',
      deviceCount: 'Dispositivi',
      onlineCount: 'Online',
      visitsCount: 'Visite',
      operatingSystem: 'Sistema operativo',
      oneWeek: 'Una settimana',
      changeLogs: 'Registro modifiche',
      cardDetail: {
        viewHint: 'Clicca per visualizzare i dettagli',
        recentUsers: 'Utenti recenti',
        recentDevices: 'Dispositivi recenti',
        recentVisits: 'Log di visite recenti',
        desc: {
          userCount: 'Mostra il numero totale di utenti nel sistema.',
          deviceCount: 'Mostra il numero totale di dispositivi nel sistema.',
          onlineCount: 'Mostra il numero di dispositivi online in base alle statistiche heartbeat.',
          visitCount: 'Mostra le statistiche di visita dai log di audit.'
        }
      },
      serverConfig: {
        ...enUs.page.home.serverConfig,
        title: 'Configurazione connessione client',
        tip: 'Copia i seguenti valori nel client RustDesk. KEY viene letto prima dalla variabile d’ambiente RUSTDESK_KEY, poi rilevato automaticamente da id_ed25519.pub nelle directory montate.',
        idServer: 'Server ID',
        relayServer: 'Server Relay',
        apiServer: 'Server API',
        key: 'KEY',
        idServerPlaceholder: 'es. tuo.dominio.com',
        relayServerPlaceholder: 'es. tuo.dominio.com',
        apiServerPlaceholder: 'es. https://tuo.dominio.com',
        keyPlaceholder: 'Rilevato automaticamente o tramite RUSTDESK_KEY / RUSTDESK_HBBS_DIR',
        copy: 'Copia',
        copyAll: 'Copia tutto',
        copyTemplate: 'Copia modello RustDesk',
        showQr: 'Mostra codice QR',
        qrTitle: 'Codice QR importazione RustDesk',
        qrTip:
          'Scansiona questo codice QR nell’app RustDesk mobile per importare lo stesso testo del modello del pulsante copia.',
        qrPayload: 'Testo modello RustDesk',
        qrFailed: 'Generazione codice QR fallita',
        refresh: 'Aggiorna',
        clearCacheReload: 'Cancella cache e ricarica',
        source: 'Origine',
        lastUpdated: 'Ultimo aggiornamento',
        show: 'Mostra',
        hide: 'Nascondi',
        missingTip: 'I seguenti campi sono vuoti, configurali prima nelle variabili d’ambiente del container: {fields}',
        copyEmpty: '{label} è vuoto e non può essere copiato',
        copySuccess: '{label} copiato',
        copyFailed: 'Copia di {label} fallita',
        fetchFailed: 'Caricamento configurazione server fallito',
        cacheCleared: 'Cache cancellata, ricaricamento configurazione server',
        sourceType: {
          ...enUs.page.home.serverConfig.sourceType,
          remote: 'Remoto',
          'memory-cache': 'Cache memoria',
          'session-cache': 'Cache sessione',
          env: 'Ambiente',
          inferred: 'Derivato',
          empty: 'Vuoto',
          auto: 'Rilevato automaticamente'
        },
        sourceHint: {
          ...enUs.page.home.serverConfig.sourceHint,
          env: 'Questo valore proviene dalla variabile d’ambiente del container.',
          inferred: 'Questo valore è derivato o rilevato automaticamente dal container.',
          empty: 'Nessun valore ancora configurato o derivato.'
        },
        connectivity: {
          ...enUs.page.home.serverConfig.connectivity,
          clear: 'Cancella risultati',
          check: 'Verifica connettività',
          checkOne: 'Verifica',
          checked: 'Verifica connettività completata',
          checkedOne: 'Connettività di {field} verificata',
          checkedCached: 'Utilizzo risultato recente della verifica (cache)',
          checkFailed: 'Verifica connettività fallita',
          cleared: 'Risultati connettività cancellati',
          source: 'Origine verifica',
          lastChecked: 'Ultima verifica',
          target: 'Destinazione',
          duration: 'Durata',
          notChecked: 'Non ancora verificato',
          checkSourceType: {
            ...enUs.page.home.serverConfig.connectivity.checkSourceType,
            remote: 'Remoto',
            cache: 'Cache'
          },
          status: {
            ...enUs.page.home.serverConfig.connectivity.status,
            idle: 'Non verificato',
            ok: 'Raggiungibile',
            error: 'Fallito',
            skip: 'Saltato'
          }
        }
      }
    },
    user: {
      ...enUs.page.user,
      list: {
        ...enUs.page.user.list,
        addUser: 'Aggiungi utente',
        editUser: 'Modifica utente',
        inputUsername: 'Inserisci nome utente',
        inputPassword: 'Inserisci password',
        inputNickname: 'Inserisci soprannome',
        emailFormatError: 'Formato email non valido',
        selectUserStatus: 'Seleziona stato utente',
        searchPlaceholder: 'Nome utente\\Soprannome\\Email',
        tfa_secret_bind: 'Associazione dispositivo 2FA',
        require2FASecret: 'Segreto 2FA vuoto',
        require2FACode: 'Il codice 2FA non può essere vuoto'
      },
      sessions: {
        ...enUs.page.user.sessions,
        kill: 'Termina',
        confirmKill: 'Confermi la terminazione?'
      },
      audit: {
        ...enUs.page.user.audit,
        logsSearchPlaceholder: 'Nome utente\\Azione\\RustdeskID\\IP'
      },
      devices: {
        ...enUs.page.user.devices,
        logsSearchPlaceholder: 'Nome utente\\Hostname\\RustdeskID'
      }
    },
    system: {
      ...enUs.page.system,
      mailTemplate: {
        ...enUs.page.system.mailTemplate,
        addMailTemplate: 'Aggiungi modello',
        editMailTemplate: 'Modifica modello',
        inputName: 'Inserisci nome',
        inputSubject: 'Inserisci oggetto',
        inputContents: 'Inserisci contenuto',
        selectType: 'Seleziona tipo'
      },
      mailLog: {
        ...enUs.page.system.mailLog,
        info: 'Dettagli'
      }
    },
    myDevices: {
      title: 'Contatti',
      welcome: 'Benvenuto, {userName}',
      status: 'Stato',
      online: 'Online',
      offline: 'Offline',
      conns: 'Connessioni',
      lastSync: 'Ultima sincronizzazione',
      logout: 'Disconnetti'
    },
    workspace: {
      ...enUs.page.workspace,
      scopeTitle: 'Area di lavoro personale',
      scopeTip: 'Qui vengono mostrati solo i tuoi dispositivi, sessioni, eventi di sicurezza e rubriche autorizzate.',
      myDevices: 'Miei dispositivi',
      activeSessions: 'Sessioni attive',
      addressBooks: 'Rubriche',
      securityEvents: 'Eventi di sicurezza',
      currentSession: 'Sessione corrente',
      revokeConfirm: 'Revocare questa sessione di accesso?',
      revoke: 'Revoca',
      accountRole: 'Ruolo account',
      adminRole: 'Amministratore',
      userRole: 'Utente',
      permissionScope: 'Ambito autorizzazioni',
      userScope: 'Risorse personali e rubriche condivise esplicitamente',
      active: 'Attivo'
    },
    oauth: {
      ...enUs.page.oauth,
      configTitle: 'Configurazione accesso terze parti',
      bindingsTitle: 'Associazioni account',
      addProvider: 'Aggiungi provider',
      editProvider: 'Modifica provider',
      providerName: 'Chiave provider',
      displayName: 'Nome visualizzato',
      clientId: 'Client ID',
      clientSecret: 'Client secret',
      secretPlaceholder: 'Lascia vuoto per mantenere il secret configurato',
      redirectUrl: 'URL callback',
      scopes: 'Ambiti',
      accountRole: 'Ruolo account',
      allowedDomains: 'Domini email consentiti',
      bindByEmail: 'Associa tramite email verificata',
      autoCreateAdmin: 'Crea automaticamente amministratore',
      autoCreateUser: 'Crea automaticamente utente',
      testConfig: 'Verifica campi obbligatori',
      testSuccess:
        'I campi obbligatori sono completi e può essere generato un URL di autorizzazione. La validità delle credenziali richiede un callback GitHub completato.',
      copyCallback: 'Copia callback',
      githubOnlyTip:
        'GitHub è disponibile per primo. Configuralo qui; server.yaml rimane un’opzione di compatibilità e ripristino.',
      adminRole: 'Amministratore',
      userRole: 'Utente',
      useDefault: 'Usa predefinito',
      listPlaceholder: 'Separa più valori con spazi o virgole',
      copied: 'Copiato'
    },
    about: {
      runningVersion: 'Versione in esecuzione',
      buildTime: 'Data build',
      compatVersion: 'Versione RustDesk compatibile',
      latestVersion: 'Ultima versione',
      updateAvailable: 'Aggiornamento disponibile',
      upToDate: 'Aggiornato',
      updateCheck: 'Verifica aggiornamenti online',
      urlTip:
        'L’URL di verifica è modificabile e salvato in questo browser. Può restituire una versione semantica in testo semplice o JSON contenente version, latest_version, tag_name o server.version. Il sito remoto deve consentire richieste CORS del browser.',
      urlPlaceholder: 'URL verifica aggiornamenti',
      checkNow: 'Verifica ora',
      restoreDefault: 'Ripristina predefinito',
      checkFailed: 'Verifica aggiornamenti fallita',
      invalidUrl: 'Solo URL HTTP e HTTPS sono supportati',
      invalidResponse: 'Nessuna versione semantica trovata nella risposta',
      updateCommand: 'Comando aggiornamento container',
      commandTip:
        'Esegui direttamente il comando più recente. "L’immagine è aggiornata" significa che l’operazione è riuscita e l’immagine installata era già corrente. Usa il comando fissato quando devi verificare una versione esatta.',
      copyCommand: 'Copia comando',
      latestCommand: 'Aggiorna alla versione più recente',
      pinnedCommand: 'Aggiorna e verifica la versione rilevata',
      customCommand: 'Modello comando personalizzato',
      versionInfo: 'Informazioni versione',
      errorHelp: 'Guida codici di errore',
      errcodeTip:
        'Tutti i codici di errore restituiti dal server sono elencati di seguito. Quando incontri un errore, cerca il codice per trovare la causa e la soluzione.',
      searchPlaceholder: 'Cerca codice, messaggio o descrizione',
      moduleFilter: 'Filtra per modulo',
      errCode: 'Codice',
      errMessage: 'Messaggio',
      errModule: 'Modulo',
      errDescription: 'Descrizione',
      errSolution: 'Soluzione'
    }
  },
  dropdown: {
    ...enUs.dropdown,
    closeCurrent: 'Chiudi corrente',
    closeOther: 'Chiudi altri',
    closeLeft: 'Chiudi a sinistra',
    closeRight: 'Chiudi a destra',
    closeAll: 'Chiudi tutto'
  },
  icon: {
    ...enUs.icon,
    themeConfig: 'Configurazione tema',
    themeSchema: 'Schema tema',
    lang: 'Cambia lingua',
    fullscreen: 'Schermo intero',
    fullscreenExit: 'Esci da schermo intero',
    reload: 'Ricarica pagina',
    collapse: 'Comprimi menu',
    expand: 'Espandi menu',
    pin: 'Fissa',
    unpin: 'Rilascia'
  },
  datatable: {
    ...enUs.datatable,
    itemCount: 'Totale {total} elementi'
  },
  dataMap: {
    ...enUs.dataMap,
    user: {
      ...enUs.dataMap.user,
      username: 'Nome utente',
      password: 'Password',
      name: 'Soprannome',
      email: 'Email',
      licensed_devices: 'Dispositivi con licenza',
      login_verify: 'Verifica accesso',
      status: 'Stato',
      is_admin: 'Amministratore',
      tfa_secret: 'Segreto 2FA',
      tfa_code: 'Codice 2FA',
      created_at: 'Creato il',
      statusLabel: {
        ...enUs.dataMap.user.statusLabel,
        disabled: 'Disabilitato',
        unverified: 'Non verificato',
        normal: 'Normale'
      },
      loginVerifyLabel: {
        ...enUs.dataMap.user.loginVerifyLabel,
        none: 'Nessuna',
        emailCheck: 'Verifica email',
        tfaCheck: '2FA'
      }
    },
    session: {
      ...enUs.dataMap.session,
      expired: 'Scade il',
      created_at: 'Creato il'
    },
    device: {
      ...enUs.dataMap.device,
      username: 'Nome utente',
      hostname: 'Nome computer',
      version: 'Versione RustDesk',
      memory: 'Memoria',
      os: 'Sistema operativo',
      rustdesk_id: 'Rustdesk ID'
    },
    audit: {
      ...enUs.dataMap.audit,
      username: 'Utente',
      type: 'Tipo',
      conn_id: 'ID connessione',
      rustdesk_id: 'Rustdesk ID',
      ip: 'IP',
      session_id: 'ID sessione',
      uuid: 'UUID',
      created_at: 'Creato il',
      closed_at: 'Chiuso il',
      typeLabel: {
        ...enUs.dataMap.audit.typeLabel,
        remote_control: 'Controllo remoto',
        file_transfer: 'Trasferimento file',
        tcp_tunnel: 'Tunnel TCP'
      },
      fileTransferTypeLabel: {
        ...enUs.dataMap.audit.fileTransferTypeLabel,
        master_controlled: 'Controllore -> Controllato',
        controlled_master: 'Controllato -> Controllore'
      },
      peer_id: 'Peer ID',
      path: 'Percorso'
    },
    mailTemplate: {
      ...enUs.dataMap.mailTemplate,
      name: 'Nome',
      type: 'Tipo',
      subject: 'Oggetto',
      contents: 'Contenuto',
      created_at: 'Creato il',
      typeLabel: {
        ...enUs.dataMap.mailTemplate.typeLabel,
        loginVerify: 'Verifica accesso',
        registerVerify: 'Verifica registrazione',
        other: 'Altro'
      }
    },
    mailLog: {
      ...enUs.dataMap.mailLog,
      username: 'Utente',
      uuid: 'UUID',
      from: 'Da',
      to: 'A',
      subject: 'Oggetto',
      contents: 'Contenuto',
      status: 'Stato',
      created_at: 'Inviato il',
      statusLabel: {
        ...enUs.dataMap.mailLog.statusLabel,
        ok: 'Successo',
        err: 'Errore'
      }
    },
    ab: {
      ...enUs.dataMap.ab,
      rustdesk_id: 'Rustdesk ID',
      username: 'Nome utente',
      hostname: 'Hostname',
      tags: 'Tag',
      alias: 'Alias',
      hash: 'Hash',
      owner: 'Proprietario',
      name: 'Nome rubrica',
      user_id: 'ID utente',
      guid: 'GUID',
      rule: 'Regola',
      max_peer: 'Peer massimi',
      shared: 'Condiviso',
      ab_id: 'ID rubrica',
      tagName: 'Nome',
      tagColor: 'Colore',
      updated_at: 'Aggiornato il',
      personal: 'Rubrica personale',
      legacy: 'Rubrica legacy',
      note: 'Nota',
      platform: 'Piattaforma',
      personalReadOnly: 'Personale (sola lettura)',
      nameRequired: 'Il nome è obbligatorio',
      deviceIdRequired: 'L’ID dispositivo è obbligatorio',
      tagsHint: 'Separa più tag con virgole',
      read: 'Lettura',
      readWrite: 'Lettura e scrittura',
      fullControl: 'Controllo completo'
    },
    token: {
      device_os: 'Sistema operativo dispositivo',
      device_name: 'Nome dispositivo',
      token_hash: 'Hash token',
      is_admin: 'Amministratore',
      status: 'Attivo'
    },
    oauth: {
      provider: 'Provider',
      subject: 'Soggetto',
      email: 'Email',
      name: 'Nome',
      last_login_at: 'Ultimo accesso'
    },
    loginLog: {
      allEvents: 'Tutti gli eventi',
      event: 'Evento',
      userAgent: 'User Agent',
      success: 'Successo',
      reason: 'Motivo'
    },
    containerLog: {
      timestamp: 'Ora',
      level: 'Livello',
      source: 'Origine',
      message: 'Messaggio',
      method: 'Metodo',
      path: 'Percorso',
      status_code: 'Stato',
      duration_ms: 'Durata',
      user_name: 'Utente',
      client_ip: 'IP client',
      user_agent: 'User Agent',
      request_id: 'ID richiesta',
      error_msg: 'Errore',
      created_at: 'Registrato'
    }
  },
  api: {
    ...enUs.api,
    CaptchaError: 'Errore CAPTCHA',
    UserNotExists: 'L’utente non esiste',
    UsernameOrPasswordError: 'Account o password errati',
    UserExists: 'Il nome utente è già in uso',
    UsernameEmpty: 'Il nome utente non può essere vuoto',
    PasswordEmpty: 'La password non può essere vuota',
    UserAddSuccess: 'Utente creato con successo',
    DataError: 'Errore dati',
    RequestError: 'Richiesta fallita',
    UserUpdateSuccess: 'Utente aggiornato con successo',
    UserDeleteSuccess: 'Utente eliminato con successo',
    SessionKillSuccess: 'Sessione terminata con successo',
    MailTemplateNameEmpty: 'Il nome non può essere vuoto',
    MailTemplateSubjectEmpty: 'L’oggetto non può essere vuoto',
    MailTemplateContentsEmpty: 'Il contenuto non può essere vuoto',
    MailTemplateAddSuccess: 'Modello email creato con successo',
    MailTemplateUpdateSuccess: 'Modello email aggiornato con successo',
    NoEmailAddress: 'Nessun indirizzo email impostato',
    VerificationCodeError: 'Errore codice di verifica',
    UUIDEmpty: 'L’UUID non può essere vuoto'
  }
};

export default local;
