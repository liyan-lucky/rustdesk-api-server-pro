import enUs from './en-us';

const local: App.I18n.Schema = {
  ...enUs,
  system: { ...enUs.system, title: 'Rustdesk Api Server' },
  common: {
    ...enUs.common,
    action: 'Action',
    add: 'Ajouter',
    addSuccess: 'Ajout réussi',
    backToHome: 'Retour ? l\'accueil',
    batchDelete: 'Suppression par lot',
    cancel: 'Annuler',
    close: 'Fermer',
    check: 'vérifier',
    expandColumn: 'Développer la colonne',
    columnSetting: 'Paramètres des colonnes',
    config: 'Configuration',
    confirm: 'Confirmer',
    delete: 'Supprimer',
    deleteSuccess: 'Suppression réussie',
    confirmDelete: 'Voulez-vous vraiment supprimer ?',
    edit: 'Modifier',
    import: 'Importer CSV',
    export: 'Exporter CSV',
    look: 'Voir',
    warning: 'Avertissement',
    error: 'Erreur',
    index: 'Index',
    keywordSearch: 'Veuillez saisir un mot-clé',
    logout: 'Se déconnecter',
    logoutConfirm: 'Voulez-vous vraiment vous déconnecter ?',
    lookForward: 'Bientôt disponible',
    modify: 'Modifier',
    modifySuccess: 'Modification réussie',
    noData: 'Aucune donnée',
    operate: 'opération',
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
    logout: 'déconnecter l’utilisateur après échec de requête',
    logoutMsg: 'État utilisateur invalide, veuillez vous reconnecter',
    logoutWithModal: 'Afficher une fenêtre après échec de requête puis déconnecter',
    logoutWithModalMsg: 'État utilisateur invalide, veuillez vous reconnecter',
    refreshToken: 'Le jeton a expir?, actualisation du jeton',
    tokenExpired: 'Le jeton de la requête a expir?'
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
    colourWeakness: 'Déficience des couleurs',
    layoutMode: {
      ...enUs.theme.layoutMode,
      title: 'Mode de mise en page',
      vertical: 'Menu vertical',
      horizontal: 'Menu horizontal',
      'vertical-mix': 'Mode mixte vertical',
      'horizontal-mix': 'Mode mixte horizontal',
      reverseHorizontalMix: 'Inverser la position des menus de niveau 1 et enfants'
    },
    recommendColor: 'Appliquer l’algorithme de couleur recommandé',
    recommendColorDesc: "L'algorithme de couleur recommandé fait référence à",
    themeColor: {
      ...enUs.theme.themeColor,
      title: 'Couleur du thème',
      primary: 'Primaire',
      info: 'Info',
      success: 'Succès',
      warning: 'Avertissement',
      error: 'Erreur',
      followPrimary: 'Suivre la couleur primaire'
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
        title: 'Mode d’animation',
        fade: 'Fondu',
        'fade-slide': 'Glisser',
        'fade-bottom': 'Fondu zoom',
        'fade-scale': 'Fondu échelle',
        'zoom-fade': 'Zoom fondu',
        'zoom-out': 'Zoom arrière',
        none: 'Aucun'
      }
    },
    fixedHeaderAndTab: 'En-tête et onglets fixes',
    header: {
      ...enUs.theme.header,
      height: 'Hauteur de l?En-tête',
      breadcrumb: {
        ...enUs.theme.header.breadcrumb,
        visible: 'Fil d’Ariane visible',
        showIcon: 'Icône du fil d’Ariane visible'
      }
    },
    tab: {
      ...enUs.theme.tab,
      visible: 'Onglets visibles',
      cache: 'Cache des onglets',
      height: 'Hauteur des onglets',
      mode: {
        ...enUs.theme.tab.mode,
        title: 'Mode des onglets',
        chrome: 'Chrome',
        button: 'Bouton'
      }
    },
    sider: {
      ...enUs.theme.sider,
      inverted: 'Barre latérale sombre',
      width: 'Largeur de la barre latérale',
      collapsedWidth: 'Largeur repliée',
      mixWidth: 'Largeur mixte',
      mixCollapsedWidth: 'Largeur mixte repliée',
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
      visible: 'Filigrane visible en plein écran',
      text: 'Texte du filigrane'
    },
    themeDrawerTitle: 'Configuration du thème',
    pageFunTitle: 'Fonctions de page',
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
    login: 'connexion',
    403: 'Accès refusé',
    404: 'Page introuvable',
    500: 'Erreur serveur',
    'iframe-page': 'Iframe',
    home: 'Accueil',
    audit: 'Audit',
    user: 'Gestion des utilisateurs',
    user_list: 'Liste des utilisateurs',
    user_sessions: 'Sessions',
    user_profile: 'Profil',
    system: 'Gestion systeme',
    system_mail_template: 'Modèles e-mail',
    system_mail_logs: 'Logs e-mail',
    system_mail: 'E-mail',
    audit_baselogs: 'Journaux de base',
    audit_filetransferlogs: 'Journaux de transfert de fichiers',
    audit_loginlogs: 'Journaux de connexion',
    devices: 'Appareils',
    'my-devices': 'Contacts',
    'my-devices_peers': 'Mes contacts',
    'my-devices_manage': 'Gestion des carnets',
    'my-devices_tags': 'Gestion des tags',
    workspace: 'Mon espace', workspace_overview: 'Vue d’ensemble', workspace_devices: 'Mes appareils', workspace_sessions: 'Sessions de connexion', workspace_security: 'Événements de sécurité', workspace_profile: 'Profil', about: 'À propos et mises à jour',
    system_server: 'Configuration serveur',
    system_tokens: 'Jeton utilisateur',
    system_oauth: 'Connexion tierce'
  },
  page: {
    ...enUs.page,
      login: {
        ...enUs.page.login,
        common: {
          ...enUs.page.login.common,
          loginOrRegister: 'connexion / Inscription',
          userNamePlaceholder: 'Veuillez saisir le nom d’utilisateur',
          phonePlaceholder: 'Veuillez saisir le numéro de téléphone',
          codePlaceholder: 'Veuillez saisir le code de vérification',
          passwordPlaceholder: 'Veuillez saisir le mot de passe',
          confirmPasswordPlaceholder: 'Veuillez saisir ? nouveau le mot de passe',
          codeLogin: 'connexion par code',
          confirm: 'Confirmer',
          back: 'Retour',
          validateSuccess: 'Vérification réussie',
          loginSuccess: 'connexion réussie',
          welcomeBack: 'Bon retour, {userName} !',
          thirdPartyLogin: 'connexion tierce',
          continueWith: 'Continuer avec {provider}',
          providerUnavailable: 'La connexion {provider} est indisponible', oauthAccountNotBound: 'Aucun compte correspondant ne peut être lié. Utilisez le même courriel vérifié ou activez la création automatique.', oauthProviderUnreachable: 'Le serveur ne peut pas joindre le fournisseur. Vérifiez la connexion HTTPS sortante.', oauthStateExpired: 'La demande de connexion a expiré ou a déjà été utilisée. Recommencez.', oauthAuthFailed: 'Échec de la connexion tierce. Vérifiez la configuration et le journal de sécurité.'
        },
        pwdLogin: {
          ...enUs.page.login.pwdLogin,
          title: 'connexion par mot de passe',
          rememberMe: 'Se souvenir de moi',
          switchToUser: 'connexion utilisateur'
         },
         userLogin: {
           title: 'connexion utilisateur',
           switchToAdmin: 'connexion admin'
        }
      },
    home: {
      ...enUs.page.home,
      greeting: 'Bonjour {userName}, excellente journée !',
      userCount: 'Utilisateurs',
      deviceCount: 'Appareils',
      onlineCount: 'En ligne',
      visitsCount: 'Visites',
      operatingSystem: 'Systeme d exploitation',
      oneWeek: 'Une semaine',
      changeLogs: 'Journal des modifications',
      cardDetail: {
        viewHint: 'Cliquez pour voir les details',
        recentUsers: 'Utilisateurs recents',
        recentDevices: 'Appareils recents',
        recentVisits: 'Journaux de visite recents',
        desc: {
          userCount: "Affiche le nombre total d'utilisateurs dans le systeme.",
          deviceCount: "Affiche le nombre total d'appareils dans le systeme.",
          onlineCount: "Affiche le nombre d'appareils en ligne base sur les heartbeats.",
          visitCount: 'Affiche les statistiques de visites a partir des journaux audit.'
        }
      },
      serverConfig: {
        ...enUs.page.home.serverConfig,
        title: 'Configuration de connexion client',
        tip: 'Copiez les valeurs suivantes dans le client RustDesk. Si KEY est vide, définissez `RUSTDESK_KEY` dans les variables d’environnement du conteneur.',
        idServer: 'Serveur ID',
        relayServer: 'Serveur relais',
        apiServer: 'Serveur API',
        key: 'KEY',
        idServerPlaceholder: 'ex. your.domain.com',
        relayServerPlaceholder: 'ex. your.domain.com',
        apiServerPlaceholder: 'ex. https://your.domain.com',
        keyPlaceholder: 'Fournir via la variable RUSTDESK_KEY',
        copy: 'Copier',
        copyAll: 'Tout copier',
        copyTemplate: 'Copier le modèle RustDesk',
        showQr: 'Afficher le code QR',
        qrTitle: 'Code QR d\'importation RustDesk',
        qrTip: 'Scannez ce code QR dans l\'application RustDesk mobile pour importer la configuration.',
        qrPayload: 'Texte du modèle RustDesk',
        qrFailed: 'Échec de la génération du code QR',
        refresh: 'Actualiser la configuration',
        clearCacheReload: 'Vider le cache et recharger',
        source: 'Source',
        lastUpdated: 'Derniere mise a jour',
        show: 'Afficher',
        hide: 'Masquer',
        missingTip: 'Les champs suivants sont vides. Veuillez d’abord les configurer dans les variables d’environnement du conteneur : {fields}',
        copyEmpty: '{label} est vide et ne peut pas être copié',
        copySuccess: '{label} copié',
        copyFailed: 'échec de la copie de {label}',
        fetchFailed: 'échec du chargement de la configuration serveur',
        cacheCleared: 'Cache vid?, rechargement de la configuration serveur',
        sourceType: {
          ...enUs.page.home.serverConfig.sourceType,
          remote: 'Distant',
          'memory-cache': 'Cache mémoire',
          'session-cache': 'Cache session',
          env: 'Env',
          inferred: 'Déduit',
          empty: 'Vide',
           auto: 'Auto-détecté'
        },
        sourceHint: {
          ...enUs.page.home.serverConfig.sourceHint,
          env: 'Cette valeur provient d’une variable d’environnement du conteneur.',
          inferred: 'Cette valeur est déduite automatiquement de l’adresse d’accès actuelle.',
          empty: 'Aucune valeur configurée ni déduite pour le moment.'
        },
        connectivity: {
          ...enUs.page.home.serverConfig.connectivity,
          clear: 'Effacer les résultats',
          check: 'Tester la connectivité',
          checkOne: 'Tester',
          checked: 'Vérification de connectivité terminée',
          checkedOne: 'Connectivit? de {field} vérifiée',
          checkedCached: 'Résultat récent de connectivité utilisé (cache)',
          checkFailed: 'échec de la vérification de connectivité',
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
            error: 'échec',
            skip: 'ignoré'
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
        searchPlaceholder: 'Nom utilisateur/Pseudo/E-mail'
      },
      sessions: {
        ...enUs.page.user.sessions,
        kill: 'Terminer',
        confirmKill: 'Terminer cette session ?'
      },
      audit: {
        ...enUs.page.user.audit,
        logsSearchPlaceholder: 'Utilisateur/Action/RustdeskID/IP'
      },
      devices: {
        ...enUs.page.user.devices,
        logsSearchPlaceholder: 'Utilisateur/Hôte/RustdeskID'
      }
    },
    system: {
      ...enUs.page.system,
      mailTemplate: {
        ...enUs.page.system.mailTemplate,
        addMailTemplate: 'Ajouter un modèle',
        editMailTemplate: 'Modifier le modèle',
        inputName: 'Saisir le nom',
        inputSubject: 'Saisir le sujet',
        inputContents: 'Saisir le contenu',
        selectType: 'Sélectionner le type'
      },
      mailLog: {
        ...enUs.page.system.mailLog,
        info: 'Détail'
      }
    },
    myDevices: {
      title: 'Contacts',
      welcome: 'Bienvenue, {userName}',
      status: 'Statut',
      online: 'En ligne',
      offline: 'Hors ligne',
      conns: 'connexions',
      lastSync: 'Dernière synchro',
      logout: 'Déconnexion'
    },
    workspace: {
      ...enUs.page.workspace,
      scopeTitle: 'Espace personnel', scopeTip: 'Seuls vos appareils, sessions, événements de sécurité et carnets autorisés sont affichés ici.', myDevices: 'Mes appareils', activeSessions: 'Sessions actives', addressBooks: 'Carnets d’adresses', securityEvents: 'Événements de sécurité', currentSession: 'Session actuelle', revokeConfirm: 'Révoquer cette session de connexion ?', revoke: 'Révoquer', accountRole: 'Rôle du compte', adminRole: 'Administrateur', userRole: 'Utilisateur', permissionScope: 'Périmètre des droits', userScope: 'Ressources personnelles et carnets explicitement partagés', active: 'Actif'
    },
    oauth: {
      configTitle: 'Configuration de la connexion tierce', bindingsTitle: 'Liaisons de comptes', addProvider: 'Ajouter un fournisseur', editProvider: 'Modifier le fournisseur', providerName: 'Clé du fournisseur', displayName: 'Nom affiché', clientId: 'Identifiant client', clientSecret: 'Secret client', secretPlaceholder: 'Laisser vide pour conserver le secret configuré', redirectUrl: 'Adresse de rappel', scopes: 'Autorisations', accountRole: 'Rôle du compte', allowedDomains: 'Domaines de messagerie autorisés', bindByEmail: 'Lier par courriel vérifié', autoCreateAdmin: 'Créer automatiquement un administrateur', autoCreateUser: 'Créer automatiquement un utilisateur', testConfig: 'Tester la configuration', testSuccess: 'Configuration complète et adresse d’autorisation générée', copyCallback: 'Copier le rappel', githubOnlyTip: 'GitHub est disponible en premier. Configurez-le ici ; server.yaml reste réservé à la compatibilité et à la récupération.', adminRole: 'Administrateur', userRole: 'Utilisateur', useDefault: 'Utiliser la valeur par défaut', listPlaceholder: 'Séparer les valeurs par des espaces ou des virgules', copied: 'Copié'
    },
    about: {
      latestCommand: 'Mettre à jour vers latest', pinnedCommand: 'Mettre à jour et vérifier la version détectée', customCommand: 'Modèle de commande personnalisé',
      runningVersion: 'Version en cours', buildTime: 'Date de compilation', compatVersion: 'Version RustDesk compatible', latestVersion: 'Dernière version', updateAvailable: 'Mise à jour disponible', upToDate: 'À jour', updateCheck: 'Recherche de mise à jour en ligne', urlTip: 'L’adresse de vérification est modifiable et enregistrée dans ce navigateur. Elle peut renvoyer une version sémantique en texte ou un JSON contenant version, latest_version, tag_name ou server.version. Le site distant doit autoriser les requêtes CORS du navigateur.', urlPlaceholder: 'Adresse de vérification des mises à jour', checkNow: 'Vérifier maintenant', restoreDefault: 'Rétablir l’adresse par défaut', checkFailed: 'Échec de la vérification', invalidUrl: 'Seules les adresses HTTP et HTTPS sont acceptées', invalidResponse: 'Aucune version sémantique trouvée dans la réponse', updateCommand: 'Commande de mise à jour', commandTip: 'Modifiez le modèle; {version} est remplacé par la dernière version.', copyCommand: 'Copier la commande'
    }
  },
  dataMap: {
    ...enUs.dataMap,
    user: {
      ...enUs.dataMap.user,
      username: 'Nom utilisateur',
        password: 'Mot de passe',
      name: 'Pseudo',
      email: 'E-mail',
      licensed_devices: 'Appareils autorisés',
      login_verify: 'Vérification connexion',
      status: 'Statut',
      is_admin: 'Admin',
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
        emailCheck: 'Vérification e-mail',
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
      username: 'Nom utilisateur',
      hostname: "Nom de l'hôte",
      version: 'Version RustDesk',
        memory: 'Mémoire',
      os: 'OS',
      rustdesk_id: 'Rustdesk ID'
    },
    ab: {
      ...enUs.dataMap.ab,
      rustdesk_id: 'Identifiant Rustdesk',
      username: 'Nom d\'utilisateur',
      hostname: "Nom d'hôte",
      tags: 'Étiquettes',
      alias: 'Alias',
      hash: 'Hash',
      owner: 'Propriétaire',
      name: 'Nom du carnet',
      user_id: 'ID utilisateur',
      guid: 'GUID',
      rule: 'Règle',
      max_peer: 'Contacts max',
      shared: 'Partagé',
      ab_id: 'ID carnet',
      tagName: 'Nom',
      tagColor: 'Couleur',
      updated_at: 'Mis à jour le',
      personal: "Mon carnet d'adresses", legacy: "Carnet d'adresses hérité", note: 'Note', platform: 'Plateforme', personalReadOnly: 'Personnel (lecture seule)', nameRequired: 'Le nom est obligatoire', deviceIdRequired: "L’identifiant de l’appareil est obligatoire", tagsHint: 'Séparez les étiquettes par des virgules', read: 'Lecture', readWrite: 'Lecture et écriture', fullControl: 'Contrôle total'
    },
    token: {
      device_os: 'SE appareil',
      device_name: 'Nom appareil',
      token_hash: 'Hachage Token',
      is_admin: 'Admin',
      status: 'Actif'
    },
    oauth: {
      provider: 'Fournisseur',
      subject: 'Sujet',
      email: 'Courriel',
      name: 'Nom',
      last_login_at: 'Dernière connexion'
    },
    loginLog: {
      allEvents: 'Tous les événements',
      event: 'Événement',
      userAgent: 'Agent utilisateur',
      success: 'Succès',
      reason: 'Raison'
    },
    audit: {
      ...enUs.dataMap.audit,
      username: 'Utilisateur',
      type: 'Type',
        conn_id: 'ID de connexion',
      rustdesk_id: 'Rustdesk ID',
        peer_id: 'ID pair',
      ip: 'IP',
        session_id: 'ID de session',
        uuid: 'UUID',
      created_at: 'Créé le',
        closed_at: 'fermé le',
      typeLabel: {
        ...enUs.dataMap.audit.typeLabel,
        remote_control: 'Contrôle à distance',
        file_transfer: 'Transfert de fichiers',
        tcp_tunnel: 'Tunnel TCP'
      },
      fileTransferTypeLabel: {
        ...enUs.dataMap.audit.fileTransferTypeLabel,
        master_controlled: 'Contrôleur -> Contrôlé',
        controlled_master: 'Contrôlé -> Contrôleur'
      },
        path: 'Chemin'
    },
    mailTemplate: {
      ...enUs.dataMap.mailTemplate,
      name: 'Nom',
      type: 'Type',
      subject: 'Sujet',
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
      subject: 'Sujet',
        contents: 'Contenu',
      status: 'Statut',
      created_at: 'Envoyé le',
      statusLabel: {
        ...enUs.dataMap.mailLog.statusLabel,
        ok: 'Succès',
        err: 'Erreur'
      }
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
    RequestError: 'Échec de la requête',
    UserUpdateSuccess: 'Utilisateur modifié avec succès',
    UserDeleteSuccess: 'Utilisateur supprimé avec succès',
    SessionKillSuccess: 'Session terminée avec succès',
    MailTemplateNameEmpty: 'Le nom ne peut pas être vide',
    MailTemplateSubjectEmpty: 'Le sujet ne peut pas être vide',
    MailTemplateContentsEmpty: 'Le contenu ne peut pas être vide',
    MailTemplateAddSuccess: 'Modèle e-mail créé avec succès',
    MailTemplateUpdateSuccess: 'Modèle e-mail modifié avec succès',
    NoEmailAddress: 'Aucune adresse e-mail définie',
    VerificationCodeError: 'Code de vérification incorrect',
    UUIDEmpty: 'UUID ne peut pas être vide'
  },
  dropdown: {
    ...enUs.dropdown,
    closeCurrent: "Fermer l'onglet courant",
    closeOther: 'Fermer les autres',
    closeLeft: 'Fermer ? gauche',
    closeRight: 'Fermer ? droite',
    closeAll: 'Tout fermer'
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
  }
};

export default local;
