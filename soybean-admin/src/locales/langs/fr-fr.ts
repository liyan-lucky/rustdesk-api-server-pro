import enUs from './en-us';

const local: App.I18n.Schema = {
  ...enUs,
  system: {
    ...enUs.system,
    title: 'Rustdesk Api Server',
    updateTitle: 'Mise à jour de la version du système',
    updateContent: 'Une nouvelle version du système est disponible. Actualiser la page maintenant ?',
    updateConfirm: 'Actualiser maintenant',
    updateCancel: 'Plus tard'
  },
  common: {
    ...enUs.common,
    action: 'Action',
    add: 'Ajouter',
    addSuccess: 'Ajout réussi',
    backToHome: 'Retour à l’accueil',
    batchDelete: 'Suppression par lot',
    cancel: 'Annuler',
    close: 'Fermer',
    check: 'Vérifier',
    expandColumn: 'Développer la colonne',
    columnSetting: 'Paramètres des colonnes',
    config: 'Configuration',
    confirm: 'Confirmer',
    delete: 'Supprimer',
    deleteSuccess: 'Suppression réussie',
    confirmDelete: 'Voulez-vous vraiment supprimer ?',
    edit: 'Modifier',
    import: 'Importer CSV',
    export: 'Exporter CSV',
    look: 'Voir',
    warning: 'Avertissement',
    error: 'Erreur',
    index: 'Index',
    keywordSearch: 'Veuillez saisir un mot-clé',
    logout: 'Se déconnecter',
    logoutConfirm: 'Voulez-vous vraiment vous déconnecter ?',
    lookForward: 'Bientôt disponible',
    modify: 'Modifier',
    modifySuccess: 'Modification réussie',
    noData: 'Aucune donnée',
    operate: 'Opération',
    pleaseCheckValue: 'Veuillez vérifier si la valeur est valide',
    refresh: 'Actualiser',
    reset: 'Réinitialiser',
    search: 'Rechercher',
    switch: 'Basculer',
    tip: 'Conseil',
    trigger: 'Déclencher',
    update: 'Mettre à jour',
    updateSuccess: 'Mise à jour réussie',
    userCenter: 'Centre utilisateur',
    yesOrNo: {
      yes: 'Oui',
      no: 'Non'
    }
  },
  request: {
    ...enUs.request,
    logout: 'Déconnecter l’utilisateur après échec de la requête',
    logoutMsg: 'État utilisateur invalide, veuillez vous reconnecter',
    logoutWithModal: 'Afficher une fenêtre après échec de la requête puis déconnecter',
    logoutWithModalMsg: 'État utilisateur invalide, veuillez vous reconnecter',
    refreshToken: 'Le jeton a expiré, actualisation du jeton',
    tokenExpired: 'Le jeton de la requête a expiré'
  },
  theme: {
    ...enUs.theme,
    themeSchema: {
      ...enUs.theme.themeSchema,
      title: 'Schéma de thème',
      light: 'Clair',
      dark: 'Sombre',
      auto: 'Suivre le système'
    },
    grayscale: 'Niveaux de gris',
    colourWeakness: 'Daltonisme',
    layoutMode: {
      ...enUs.theme.layoutMode,
      title: 'Mode de mise en page',
      vertical: 'Menu vertical',
      horizontal: 'Menu horizontal',
      'vertical-mix': 'Menu mixte vertical',
      'horizontal-mix': 'Menu mixte horizontal',
      reverseHorizontalMix: 'Inverser la position des menus de premier niveau et des sous-menus'
    },
    recommendColor: 'Appliquer l’algorithme de couleur recommandé',
    recommendColorDesc: 'L’algorithme de couleur recommandé fait référence à',
    themeColor: {
      ...enUs.theme.themeColor,
      title: 'Couleur du thème',
      primary: 'Primaire',
      info: 'Information',
      success: 'Succès',
      warning: 'Avertissement',
      error: 'Erreur',
      followPrimary: 'Suivre le primaire'
    },
    scrollMode: {
      ...enUs.theme.scrollMode,
      title: 'Mode de défilement',
      wrapper: 'Conteneur',
      content: 'Contenu'
    },
    page: {
      ...enUs.theme.page,
      animate: 'Animation de page',
      mode: {
        ...enUs.theme.page.mode,
        title: 'Mode d’animation de page',
        fade: 'Fondu',
        'fade-slide': 'Glissement',
        'fade-bottom': 'Fondu zoom',
        'fade-scale': 'Fondu échelle',
        'zoom-fade': 'Zoom fondu',
        'zoom-out': 'Zoom arrière',
        none: 'Aucune'
      }
    },
    fixedHeaderAndTab: 'En-tête et onglet fixes',
    header: {
      ...enUs.theme.header,
      height: 'Hauteur d’en-tête',
      breadcrumb: {
        visible: 'Fil d’Ariane visible',
        showIcon: 'Icône du fil d’Ariane visible'
      }
    },
    tab: {
      ...enUs.theme.tab,
      visible: 'Onglet visible',
      cache: 'Cache d’onglet',
      height: 'Hauteur d’onglet',
      mode: {
        ...enUs.theme.tab.mode,
        title: 'Mode d’onglet',
        chrome: 'Chrome',
        button: 'Bouton'
      }
    },
    sider: {
      ...enUs.theme.sider,
      inverted: 'Barre latérale sombre',
      width: 'Largeur de la barre latérale',
      collapsedWidth: 'Largeur repliée de la barre latérale',
      mixWidth: 'Largeur de la barre latérale mixte',
      mixCollapsedWidth: 'Largeur repliée de la barre latérale mixte',
      mixChildMenuWidth: 'Largeur du sous-menu mixte'
    },
    footer: {
      ...enUs.theme.footer,
      visible: 'Pied de page visible',
      fixed: 'Pied de page fixe',
      height: 'Hauteur du pied de page',
      right: 'Pied de page droit'
    },
    watermark: {
      ...enUs.theme.watermark,
      visible: 'Filigrane plein écran visible',
      text: 'Texte du filigrane'
    },
    themeDrawerTitle: 'Configuration du thème',
    pageFunTitle: 'Fonction de page',
    configOperation: {
      ...enUs.theme.configOperation,
      copyConfig: 'Copier la configuration',
      copySuccessMsg: 'Copie réussie, veuillez remplacer la variable "themeSettings" dans "src/theme/settings.ts"',
      resetConfig: 'Réinitialiser la configuration',
      resetSuccessMsg: 'Réinitialisation réussie'
    }
  },
  route: {
    ...enUs.route,
    login: 'Connexion',
    403: 'Accès interdit',
    404: 'Page non trouvée',
    500: 'Erreur serveur',
    'iframe-page': 'Iframe',
    home: 'Accueil',
    about: 'À propos',
    about_version: 'Informations de version',
    about_help: 'Aide sur les codes d’erreur',
    audit: 'Audit',
    user: 'Gestion des utilisateurs',
    user_list: 'Liste des utilisateurs',
    user_sessions: 'Sessions',
    user_profile: 'Profil',
    system: 'Gestion du système',
    system_mail_template: 'Modèle de courriel',
    system_mail_logs: 'Journaux de courriel',
    system_mail: 'Gestion des courriels',
    system_server: 'Configuration du serveur',
    system_tokens: 'Jetons utilisateur',
    system_oauth: 'Connexion tierce',
    audit_baselogs: 'Journal de base',
    audit_filetransferlogs: 'Journaux de transfert de fichiers',
    audit_loginlogs: 'Journaux de connexion',
    'audit_error-logs': "Journaux d'erreurs",
    devices: 'Appareils',
    'my-devices': 'Contacts',
    'my-devices_peers': 'Mes pairs',
    'my-devices_manage': 'Gestion du carnet d’adresses',
    'my-devices_tags': 'Gestion des étiquettes',
    workspace: 'Mon espace de travail',
    workspace_overview: 'Vue d’ensemble',
    workspace_devices: 'Mes appareils',
    workspace_sessions: 'Sessions de connexion',
    workspace_security: 'Événements de sécurité',
    workspace_profile: 'Profil'
  },
  page: {
    ...enUs.page,
    login: {
      ...enUs.page.login,
      common: {
        ...enUs.page.login.common,
        loginOrRegister: 'Connexion / Inscription',
        userNamePlaceholder: 'Veuillez saisir le nom d’utilisateur',
        phonePlaceholder: 'Veuillez saisir le numéro de téléphone',
        codePlaceholder: 'Veuillez saisir le code de vérification',
        passwordPlaceholder: 'Veuillez saisir le mot de passe',
        confirmPasswordPlaceholder: 'Veuillez saisir le mot de passe à nouveau',
        codeLogin: 'Connexion par code',
        confirm: 'Confirmer',
        back: 'Retour',
        validateSuccess: 'Vérification réussie',
        loginSuccess: 'Connexion réussie',
        welcomeBack: 'Bienvenue, {userName} !',
        thirdPartyLogin: 'Connexion tierce',
        continueWith: 'Continuer avec {provider}',
        providerUnavailable: 'La connexion {provider} n’est pas disponible',
        oauthAccountNotBound:
          'Aucun compte correspondant ne peut être lié. Ajoutez la même adresse e-mail vérifiée au compte cible ou activez la création automatique de compte.',
        oauthProviderUnreachable:
          'Le serveur ne peut pas atteindre le fournisseur de connexion. Vérifiez la connectivité HTTPS sortante et réessayez.',
        oauthStateExpired: 'La demande de connexion a expiré ou a déjà été utilisée. Relancez la connexion.',
        oauthAuthFailed: 'La connexion tierce a échoué. Vérifiez la configuration du fournisseur et le journal d’audit.'
      },
      pwdLogin: {
        ...enUs.page.login.pwdLogin,
        title: 'Connexion par mot de passe',
        rememberMe: 'Se souvenir de moi',
        switchToUser: 'Connexion utilisateur'
      },
      userLogin: {
        title: 'Connexion utilisateur',
        switchToAdmin: 'Connexion administrateur'
      }
    },
    home: {
      ...enUs.page.home,
      greeting: 'Bonjour, {userName} !',
      userCount: 'Utilisateurs',
      deviceCount: 'Appareils',
      onlineCount: 'En ligne',
      visitsCount: 'Visites',
      operatingSystem: 'Système d’exploitation',
      oneWeek: 'Une semaine',
      changeLogs: 'Journal des modifications',
      cardDetail: {
        viewHint: 'Cliquez pour voir les détails',
        recentUsers: 'Utilisateurs récents',
        recentDevices: 'Appareils récents',
        recentVisits: 'Journaux de visites récents',
        desc: {
          userCount: 'Affiche le nombre total d’utilisateurs dans le système.',
          deviceCount: 'Affiche le nombre total d’appareils dans le système.',
          onlineCount: 'Affiche le nombre d’appareils en ligne basé sur les statistiques de heartbeat.',
          visitCount: 'Affiche les statistiques de visite à partir des journaux d’audit.'
        }
      },
      serverConfig: {
        ...enUs.page.home.serverConfig,
        title: 'Configuration de connexion client',
        tip: 'Copiez les valeurs suivantes dans le client RustDesk. KEY est lu d’abord depuis la variable d’environnement RUSTDESK_KEY, puis détecté automatiquement depuis id_ed25519.pub dans les répertoires montés.',
        idServer: 'Serveur ID',
        relayServer: 'Serveur Relay',
        apiServer: 'Serveur API',
        key: 'KEY',
        idServerPlaceholder: 'ex. votre.domaine.com',
        relayServerPlaceholder: 'ex. votre.domaine.com',
        apiServerPlaceholder: 'ex. https://votre.domaine.com',
        keyPlaceholder: 'Détecté automatiquement ou via RUSTDESK_KEY / RUSTDESK_HBBS_DIR',
        copy: 'Copier',
        copyAll: 'Tout copier',
        copyTemplate: 'Copier le modèle RustDesk',
        showQr: 'Afficher le code QR',
        qrTitle: 'Code QR d’importation RustDesk',
        qrTip:
          'Scannez ce code QR dans l’application RustDesk mobile pour importer le même texte de modèle que le bouton copier.',
        qrPayload: 'Texte du modèle RustDesk',
        qrFailed: 'Échec de la génération du code QR',
        refresh: 'Actualiser',
        clearCacheReload: 'Vider le cache et recharger',
        source: 'Source',
        lastUpdated: 'Dernière mise à jour',
        show: 'Afficher',
        hide: 'Masquer',
        missingTip:
          'Les champs suivants sont vides, veuillez les configurer d’abord dans les variables d’environnement du conteneur : {fields}',
        copyEmpty: '{label} est vide et ne peut pas être copié',
        copySuccess: '{label} copié',
        copyFailed: 'Échec de la copie de {label}',
        fetchFailed: 'Échec du chargement de la configuration du serveur',
        cacheCleared: 'Cache vidé, rechargement de la configuration du serveur',
        sourceType: {
          ...enUs.page.home.serverConfig.sourceType,
          remote: 'Distant',
          'memory-cache': 'Cache mémoire',
          'session-cache': 'Cache de session',
          env: 'Environnement',
          inferred: 'Déduit',
          empty: 'Vide',
          auto: 'Détecté automatiquement'
        },
        sourceHint: {
          ...enUs.page.home.serverConfig.sourceHint,
          env: 'Cette valeur provient de la variable d’environnement du conteneur.',
          inferred: 'Cette valeur est déduite ou détectée automatiquement à partir du conteneur.',
          empty: 'Aucune valeur encore configurée ou déduite.'
        },
        connectivity: {
          ...enUs.page.home.serverConfig.connectivity,
          clear: 'Effacer les résultats',
          check: 'Vérifier la connectivité',
          checkOne: 'Vérifier',
          checked: 'Vérification de connectivité terminée',
          checkedOne: 'Connectivité de {field} vérifiée',
          checkedCached: 'Utilisation du résultat récent de la vérification (cache)',
          checkFailed: 'Échec de la vérification de connectivité',
          cleared: 'Résultats de connectivité effacés',
          source: 'Source de vérification',
          lastChecked: 'Dernière vérification',
          target: 'Cible',
          duration: 'Durée',
          notChecked: 'Pas encore vérifié',
          checkSourceType: {
            ...enUs.page.home.serverConfig.connectivity.checkSourceType,
            remote: 'Distant',
            cache: 'Cache'
          },
          status: {
            ...enUs.page.home.serverConfig.connectivity.status,
            idle: 'Non vérifié',
            ok: 'Accessible',
            error: 'Échoué',
            skip: 'Ignoré'
          }
        }
      }
    },
    user: {
      ...enUs.page.user,
      list: {
        ...enUs.page.user.list,
        addUser: 'Ajouter un utilisateur',
        editUser: 'Modifier l’utilisateur',
        inputUsername: 'Saisir le nom d’utilisateur',
        inputPassword: 'Saisir le mot de passe',
        inputNickname: 'Saisir le pseudonyme',
        emailFormatError: 'Format d’e-mail invalide',
        selectUserStatus: 'Veuillez sélectionner le statut de l’utilisateur',
        searchPlaceholder: 'Nom d’utilisateur\\Pseudonyme\\E-mail',
        tfa_secret_bind: 'Liaison de l’appareil 2FA',
        require2FASecret: 'Secret 2FA vide',
        require2FACode: 'Le code 2FA ne peut pas être vide'
      },
      sessions: {
        ...enUs.page.user.sessions,
        kill: 'Terminer',
        confirmKill: 'Confirmer la terminaison ?'
      },
      audit: {
        ...enUs.page.user.audit,
        logsSearchPlaceholder: 'Nom d’utilisateur\\Action\\RustdeskID\\IP'
      },
      devices: {
        ...enUs.page.user.devices,
        logsSearchPlaceholder: 'Nom d’utilisateur\\Nom d’hôte\\RustdeskID'
      }
    },
    system: {
      ...enUs.page.system,
      mailTemplate: {
        ...enUs.page.system.mailTemplate,
        addMailTemplate: 'Ajouter un modèle',
        editMailTemplate: 'Modifier le modèle',
        inputName: 'Saisir le nom',
        inputSubject: 'Saisir l’objet',
        inputContents: 'Saisir le contenu',
        selectType: 'Veuillez sélectionner le type'
      },
      mailLog: {
        ...enUs.page.system.mailLog,
        info: 'Détails'
      }
    },
    myDevices: {
      title: 'Contacts',
      welcome: 'Bienvenue, {userName}',
      status: 'Statut',
      online: 'En ligne',
      offline: 'Hors ligne',
      conns: 'Connexions',
      lastSync: 'Dernière synchronisation',
      logout: 'Se déconnecter'
    },
    workspace: {
      ...enUs.page.workspace,
      scopeTitle: 'Espace de travail personnel',
      scopeTip:
        'Seuls vos appareils, sessions, événements de sécurité et carnets d’adresses autorisés sont affichés ici.',
      myDevices: 'Mes appareils',
      activeSessions: 'Sessions actives',
      addressBooks: 'Carnets d’adresses',
      securityEvents: 'Événements de sécurité',
      currentSession: 'Session actuelle',
      revokeConfirm: 'Révoquer cette session de connexion ?',
      revoke: 'Révoquer',
      accountRole: 'Rôle du compte',
      adminRole: 'Administrateur',
      userRole: 'Utilisateur',
      permissionScope: 'Portée des autorisations',
      userScope: 'Ressources personnelles et carnets d’adresses explicitement partagés',
      active: 'Actif'
    },
    oauth: {
      ...enUs.page.oauth,
      configTitle: 'Configuration de la connexion tierce',
      bindingsTitle: 'Liaisons de compte',
      addProvider: 'Ajouter un fournisseur',
      editProvider: 'Modifier le fournisseur',
      providerName: 'Clé du fournisseur',
      displayName: 'Nom affiché',
      clientId: 'ID client',
      clientSecret: 'Secret client',
      secretPlaceholder: 'Laisser vide pour conserver le secret configuré',
      redirectUrl: 'URL de rappel',
      scopes: 'Portées',
      accountRole: 'Rôle du compte',
      allowedDomains: 'Domaines e-mail autorisés',
      bindByEmail: 'Lier via e-mail vérifié',
      autoCreateAdmin: 'Créer automatiquement un administrateur',
      autoCreateUser: 'Créer automatiquement un utilisateur',
      testConfig: 'Vérifier les champs obligatoires',
      testSuccess:
        'Les champs obligatoires sont complets et une URL d’autorisation peut être générée. La validité des identifiants nécessite un rappel GitHub complété.',
      copyCallback: 'Copier le rappel',
      githubOnlyTip:
        'GitHub est disponible en premier. Configurez-le ici ; server.yaml reste une option de compatibilité et de récupération.',
      adminRole: 'Administrateur',
      userRole: 'Utilisateur',
      useDefault: 'Utiliser la valeur par défaut',
      listPlaceholder: 'Séparez plusieurs valeurs par des espaces ou des virgules',
      copied: 'Copié'
    },
    about: {
      runningVersion: 'Version en cours',
      buildTime: 'Date de compilation',
      compatVersion: 'Version RustDesk compatible',
      latestVersion: 'Dernière version',
      updateAvailable: 'Mise à jour disponible',
      upToDate: 'À jour',
      updateCheck: 'Vérification des mises à jour en ligne',
      urlTip:
        'L’URL de vérification est modifiable et sauvegardée dans ce navigateur. Elle peut retourner une version sémantique en texte brut ou un JSON contenant version, latest_version, tag_name ou server.version. Le site distant doit autoriser les requêtes CORS du navigateur.',
      urlPlaceholder: 'URL de vérification des mises à jour',
      checkNow: 'Vérifier maintenant',
      restoreDefault: 'Restaurer par défaut',
      checkFailed: 'Échec de la vérification des mises à jour',
      invalidUrl: 'Seules les URL HTTP et HTTPS sont prises en charge',
      invalidResponse: 'Aucune version sémantique trouvée dans la réponse',
      updateCommand: 'Commande de mise à jour du conteneur',
      commandTip:
        'Exécutez directement la commande la plus récente. « L’image est à jour » signifie que l’opération a réussi et que l’image installée était déjà actuelle. Utilisez la commande épinglée lorsque vous devez vérifier une version exacte.',
      copyCommand: 'Copier la commande',
      latestCommand: 'Mettre à jour vers la dernière version',
      pinnedCommand: 'Mettre à jour et vérifier la version détectée',
      customCommand: 'Modèle de commande personnalisé',
      versionInfo: 'Informations de version',
      errorHelp: 'Aide sur les codes d’erreur',
      errcodeTip:
        'Tous les codes d’erreur renvoyés par le serveur sont listés ci-dessous. Lorsque vous rencontrez une erreur, recherchez le code pour trouver la cause et la solution.',
      searchPlaceholder: 'Rechercher un code, un message ou une description',
      moduleFilter: 'Filtrer par module',
      errCode: 'Code',
      errMessage: 'Message',
      errModule: 'Module',
      errDescription: 'Description',
      errSolution: 'Solution'
    }
  },
  dropdown: {
    ...enUs.dropdown,
    closeCurrent: 'Fermer l’onglet actuel',
    closeOther: 'Fermer les autres',
    closeLeft: 'Fermer à gauche',
    closeRight: 'Fermer à droite',
    closeAll: 'Fermer tout'
  },
  icon: {
    ...enUs.icon,
    themeConfig: 'Configuration du thème',
    themeSchema: 'Schéma de thème',
    lang: 'Changer de langue',
    fullscreen: 'Plein écran',
    fullscreenExit: 'Quitter le plein écran',
    reload: 'Recharger la page',
    collapse: 'Réduire le menu',
    expand: 'Développer le menu',
    pin: 'Épingler',
    unpin: 'Désépingler'
  },
  datatable: {
    ...enUs.datatable,
    itemCount: 'Total {total} éléments'
  },
  dataMap: {
    ...enUs.dataMap,
    user: {
      ...enUs.dataMap.user,
      username: 'Nom d’utilisateur',
      password: 'Mot de passe',
      name: 'Pseudonyme',
      email: 'E-mail',
      licensed_devices: 'Appareils sous licence',
      login_verify: 'Vérification de connexion',
      status: 'Statut',
      is_admin: 'Administrateur',
      tfa_secret: 'Secret 2FA',
      tfa_code: 'Code 2FA',
      created_at: 'Créé le',
      statusLabel: {
        ...enUs.dataMap.user.statusLabel,
        disabled: 'Désactivé',
        unverified: 'Non vérifié',
        normal: 'Normal'
      },
      loginVerifyLabel: {
        ...enUs.dataMap.user.loginVerifyLabel,
        none: 'Aucune',
        emailCheck: 'Vérification par e-mail',
        tfaCheck: '2FA'
      }
    },
    session: {
      ...enUs.dataMap.session,
      expired: 'Expire le',
      created_at: 'Créé le'
    },
    device: {
      ...enUs.dataMap.device,
      username: 'Nom d’utilisateur',
      hostname: 'Nom d’hôte',
      version: 'Version RustDesk',
      memory: 'Mémoire',
      os: 'Système d’exploitation',
      rustdesk_id: 'Rustdesk ID'
    },
    audit: {
      ...enUs.dataMap.audit,
      username: 'Utilisateur',
      type: 'Type',
      conn_id: 'ID de connexion',
      rustdesk_id: 'Rustdesk ID',
      ip: 'IP',
      session_id: 'ID de session',
      uuid: 'UUID',
      created_at: 'Créé le',
      closed_at: 'Fermé le',
      typeLabel: {
        ...enUs.dataMap.audit.typeLabel,
        remote_control: 'Contrôle à distance',
        file_transfer: 'Transfert de fichiers',
        tcp_tunnel: 'Tunnel TCP'
      },
      fileTransferTypeLabel: {
        ...enUs.dataMap.audit.fileTransferTypeLabel,
        master_controlled: 'Contrôlant -> Contrôlé',
        controlled_master: 'Contrôlé -> Contrôlant'
      },
      peer_id: 'ID pair',
      path: 'Chemin'
    },
    mailTemplate: {
      ...enUs.dataMap.mailTemplate,
      name: 'Nom',
      type: 'Type',
      subject: 'Objet',
      contents: 'Contenu',
      created_at: 'Créé le',
      typeLabel: {
        ...enUs.dataMap.mailTemplate.typeLabel,
        loginVerify: 'Vérification de connexion',
        registerVerify: 'Vérification d’inscription',
        other: 'Autre'
      }
    },
    mailLog: {
      ...enUs.dataMap.mailLog,
      username: 'Utilisateur',
      uuid: 'UUID',
      from: 'De',
      to: 'À',
      subject: 'Objet',
      contents: 'Contenu',
      status: 'Statut',
      created_at: 'Envoyé le',
      statusLabel: {
        ...enUs.dataMap.mailLog.statusLabel,
        ok: 'Succès',
        err: 'Erreur'
      }
    },
    ab: {
      ...enUs.dataMap.ab,
      rustdesk_id: 'Rustdesk ID',
      username: 'Nom d’utilisateur',
      hostname: 'Nom d’hôte',
      tags: 'Étiquettes',
      alias: 'Alias',
      hash: 'Hash',
      owner: 'Propriétaire',
      name: 'Nom du carnet d’adresses',
      user_id: 'ID utilisateur',
      guid: 'GUID',
      rule: 'Règle',
      max_peer: 'Pairs maximum',
      shared: 'Partagé',
      ab_id: 'ID du carnet d’adresses',
      tagName: 'Nom',
      tagColor: 'Couleur',
      updated_at: 'Mis à jour le',
      personal: 'Mon carnet d’adresses',
      legacy: 'Carnet d’adresses hérité',
      note: 'Note',
      platform: 'Plateforme',
      personalReadOnly: 'Personnel (lecture seule)',
      nameRequired: 'Le nom est obligatoire',
      deviceIdRequired: 'L’ID de l’appareil est obligatoire',
      tagsHint: 'Séparez plusieurs étiquettes par des virgules',
      read: 'Lecture',
      readWrite: 'Lecture et écriture',
      fullControl: 'Contrôle total'
    },
    token: {
      device_os: 'SE de l’appareil',
      device_name: 'Nom de l’appareil',
      token_hash: 'Hash du jeton',
      is_admin: 'Administrateur',
      status: 'Actif'
    },
    oauth: {
      provider: 'Fournisseur',
      subject: 'Sujet',
      email: 'E-mail',
      name: 'Nom',
      last_login_at: 'Dernière connexion'
    },
    loginLog: {
      allEvents: 'Tous les événements',
      event: 'Événement',
      userAgent: 'Agent utilisateur',
      success: 'Succès',
      reason: 'Raison'
    }
  },
  api: {
    ...enUs.api,
    CaptchaError: 'Erreur CAPTCHA',
    UserNotExists: 'L’utilisateur n’existe pas',
    UsernameOrPasswordError: 'Compte ou mot de passe incorrect',
    UserExists: 'Le nom d’utilisateur est déjà utilisé',
    UsernameEmpty: 'Le nom d’utilisateur ne peut pas être vide',
    PasswordEmpty: 'Le mot de passe ne peut pas être vide',
    UserAddSuccess: 'Utilisateur créé avec succès',
    DataError: 'Erreur de données',
    RequestError: 'Requête échouée',
    UserUpdateSuccess: 'Utilisateur mis à jour avec succès',
    UserDeleteSuccess: 'Utilisateur supprimé avec succès',
    SessionKillSuccess: 'Session terminée avec succès',
    MailTemplateNameEmpty: 'Le nom ne peut pas être vide',
    MailTemplateSubjectEmpty: 'L’objet ne peut pas être vide',
    MailTemplateContentsEmpty: 'Le contenu ne peut pas être vide',
    MailTemplateAddSuccess: 'Modèle de courriel créé avec succès',
    MailTemplateUpdateSuccess: 'Modèle de courriel mis à jour avec succès',
    NoEmailAddress: 'Aucune adresse e-mail définie',
    VerificationCodeError: 'Erreur du code de vérification',
    UUIDEmpty: 'L’UUID ne peut pas être vide'
  }
};

export default local;
