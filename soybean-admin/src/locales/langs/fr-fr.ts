import enUs from './en-us';

const local: App.I18n.Schema = {
  ...enUs,
  system: { ...enUs.system, title: 'Rustdesk Api Server' },
  common: {
    ...enUs.common,
    action: 'Action',
    add: 'Ajouter',
    addSuccess: 'Ajout r?ussi',
    backToHome: 'Retour ? l\'accueil',
    batchDelete: 'Suppression par lot',
    cancel: 'Annuler',
    close: 'Fermer',
    check: 'V?rifier',
    expandColumn: 'D?velopper la colonne',
    columnSetting: 'Param?tres des colonnes',
    config: 'Configuration',
    confirm: 'Confirmer',
    delete: 'Supprimer',
    deleteSuccess: 'Suppression r?ussie',
    confirmDelete: 'Voulez-vous vraiment supprimer ?',
    edit: 'Modifier',
    look: 'Voir',
    warning: 'Avertissement',
    error: 'Erreur',
    index: 'Index',
    keywordSearch: 'Veuillez saisir un mot-cl?',
    logout: 'Se d?connecter',
    logoutConfirm: 'Voulez-vous vraiment vous d?connecter ?',
    lookForward: 'Bient?t disponible',
    modify: 'Modifier',
    modifySuccess: 'Modification r?ussie',
    noData: 'Aucune donn?e',
    operate: 'Op?ration',
    pleaseCheckValue: 'Veuillez v?rifier si la valeur est valide',
    refresh: 'Actualiser',
    reset: 'R?initialiser',
    search: 'Rechercher',
    switch: 'Basculer',
    tip: 'Conseil',
    trigger: 'D?clencher',
    update: '更新',
    updateSuccess: '更新成功',
    userCenter: 'Centre utilisateur',
    yesOrNo: {
      yes: 'Oui',
      no: 'Non'
    }
  },
  request: {
    ...enUs.request,
    logout: 'D?connecter l?utilisateur apr?s ?chec de requ?te',
    logoutMsg: '?tat utilisateur invalide, veuillez vous reconnecter',
    logoutWithModal: 'Afficher une fen?tre apr?s ?chec de requ?te puis d?connecter',
    logoutWithModalMsg: '?tat utilisateur invalide, veuillez vous reconnecter',
    refreshToken: 'Le jeton a expir?, actualisation du jeton',
    tokenExpired: 'Le jeton de la requ?te a expir?'
  },
  theme: {
    ...enUs.theme,
    themeSchema: {
      ...enUs.theme.themeSchema,
      title: 'Sch?ma de th?me',
      light: 'Clair',
      dark: 'Sombre',
      auto: 'Suivre le syst?me'
    },
    grayscale: 'Niveaux de gris',
    colourWeakness: 'D?ficience des couleurs',
    layoutMode: {
      ...enUs.theme.layoutMode,
      title: 'Mode de mise en page',
      vertical: 'Menu vertical',
      horizontal: 'Menu horizontal',
      'vertical-mix': 'Mode mixte vertical',
      'horizontal-mix': 'Mode mixte horizontal',
      reverseHorizontalMix: 'Inverser la position des menus de niveau 1 et enfants'
    },
    recommendColor: 'Appliquer l?algorithme de couleur recommand?',
    recommendColorDesc: 'L?algorithme de couleur recommand? fait r?f?rence ?',
    themeColor: {
      ...enUs.theme.themeColor,
      title: 'Couleur du th?me',
      primary: 'Primaire',
      info: 'Info',
      success: 'Succ?s',
      warning: 'Avertissement',
      error: 'Erreur',
      followPrimary: 'Suivre la couleur primaire'
    },
    scrollMode: {
      ...enUs.theme.scrollMode,
      title: 'Mode de d?filement',
      wrapper: 'Conteneur',
      content: 'Contenu'
    },
    page: {
      ...enUs.theme.page,
      animate: 'Animation de page',
      mode: {
        ...enUs.theme.page.mode,
        title: 'Mode d?animation',
        fade: 'Fondu',
        'fade-slide': 'Glisser',
        'fade-bottom': 'Fondu zoom',
        'fade-scale': 'Fondu ?chelle',
        'zoom-fade': 'Zoom fondu',
        'zoom-out': 'Zoom arri?re',
        none: 'Aucun'
      }
    },
    fixedHeaderAndTab: 'En-t?te et onglets fixes',
    header: {
      ...enUs.theme.header,
      height: 'Hauteur de l?en-t?te',
      breadcrumb: {
        ...enUs.theme.header.breadcrumb,
        visible: 'Fil d?Ariane visible',
        showIcon: 'Ic?ne du fil d?Ariane visible'
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
      inverted: 'Barre lat?rale sombre',
      width: 'Largeur de la barre lat?rale',
      collapsedWidth: 'Largeur repli?e',
      mixWidth: 'Largeur mixte',
      mixCollapsedWidth: 'Largeur mixte repli?e',
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
      visible: 'Filigrane visible en plein ?cran',
      text: 'Texte du filigrane'
    },
    themeDrawerTitle: 'Configuration du th?me',
    pageFunTitle: 'Fonctions de page',
    configOperation: {
      ...enUs.theme.configOperation,
      copyConfig: 'Copier la configuration',
      copySuccessMsg: 'Copie r?ussie, veuillez remplacer la variable "themeSettings" dans "src/theme/settings.ts"',
      resetConfig: 'R?initialiser la configuration',
      resetSuccessMsg: 'R?initialisation r?ussie'
    }
  },
  route: {
    ...enUs.route,
    login: 'Connexion',
    403: 'Acc?s refus?',
    404: 'Page introuvable',
    500: 'Erreur serveur',
    'iframe-page': 'Iframe',
    home: 'Accueil',
    audit: 'Audit',
    user: 'Gestion des utilisateurs',
    user_list: 'Liste des utilisateurs',
    user_sessions: 'Sessions',
    system: 'Gestion systeme',
    system_mail_template: 'Mod?les e-mail',
    system_mail_logs: 'Logs e-mail',
    system_mail: 'E-mail',
    audit_baselogs: 'Journaux de base',
    audit_filetransferlogs: 'Journaux de transfert de fichiers',
    devices: 'Appareils'
  },
  page: {
    ...enUs.page,
      login: {
        ...enUs.page.login,
        common: {
          ...enUs.page.login.common,
          loginOrRegister: 'Connexion / Inscription',
          userNamePlaceholder: 'Veuillez saisir le nom d?utilisateur',
          phonePlaceholder: 'Veuillez saisir le num?ro de t?l?phone',
          codePlaceholder: 'Veuillez saisir le code de v?rification',
          passwordPlaceholder: 'Veuillez saisir le mot de passe',
          confirmPasswordPlaceholder: 'Veuillez saisir ? nouveau le mot de passe',
          codeLogin: 'Connexion par code',
          confirm: 'Confirmer',
          back: 'Retour',
          validateSuccess: 'V?rification r?ussie',
          loginSuccess: 'Connexion r?ussie',
          welcomeBack: 'Bon retour, {userName} !'
        },
        pwdLogin: {
          ...enUs.page.login.pwdLogin,
          title: 'Connexion par mot de passe',
          rememberMe: 'Se souvenir de moi'
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
        tip: 'Copiez les valeurs suivantes dans le client RustDesk. Si KEY est vide, d?finissez `RUSTDESK_KEY` dans les variables d?environnement du conteneur.',
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
        copyTemplate: 'Copier le mod?le RustDesk',
        refresh: 'Actualiser la configuration',
        clearCacheReload: 'Vider le cache et recharger',
        source: 'Source',
        lastUpdated: 'Derniere mise a jour',
        show: 'Afficher',
        hide: 'Masquer',
        missingTip: 'Les champs suivants sont vides. Veuillez d?abord les configurer dans les variables d?environnement du conteneur : {fields}',
        copyEmpty: '{label} est vide et ne peut pas ?tre copi?',
        copySuccess: '{label} copi?',
        copyFailed: '?chec de la copie de {label}',
        fetchFailed: '?chec du chargement de la configuration serveur',
        cacheCleared: 'Cache vid?, rechargement de la configuration serveur',
        sourceType: {
          ...enUs.page.home.serverConfig.sourceType,
          remote: 'Distant',
          'memory-cache': 'Cache m?moire',
          'session-cache': 'Cache session',
          env: 'Env',
          inferred: 'D?duit',
          empty: 'Vide'
        },
        sourceHint: {
          ...enUs.page.home.serverConfig.sourceHint,
          env: 'Cette valeur provient d?une variable d?environnement du conteneur.',
          inferred: 'Cette valeur est d?duite automatiquement de l?adresse d?acc?s actuelle.',
          empty: 'Aucune valeur configur?e ni d?duite pour le moment.'
        },
        connectivity: {
          ...enUs.page.home.serverConfig.connectivity,
          clear: 'Effacer les r?sultats',
          check: 'Tester la connectivit?',
          checkOne: 'Tester',
          checked: 'V?rification de connectivit? termin?e',
          checkedOne: 'Connectivit? de {field} v?rifi?e',
          checkedCached: 'R?sultat r?cent de connectivit? utilis? (cache)',
          checkFailed: '?chec de la v?rification de connectivit?',
          cleared: 'R?sultats de connectivit? effac?s',
          source: 'Source de v?rification',
          lastChecked: 'Derni?re v?rification',
          target: 'Cible',
          duration: 'Dur?e',
          notChecked: 'Pas encore v?rifi?',
          checkSourceType: {
            ...enUs.page.home.serverConfig.connectivity.checkSourceType,
            remote: 'Distant',
            cache: 'Cache'
          },
          status: {
            ...enUs.page.home.serverConfig.connectivity.status,
            idle: 'Non v?rifi?',
            ok: 'Accessible',
            error: '?chec',
            skip: 'Ignor?'
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
      hostname: 'Nom de l’hôte',
      version: 'Version RustDesk',
        memory: 'M?moire',
      os: 'OS',
      rustdesk_id: 'Rustdesk ID'
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
        closed_at: 'Ferm? le',
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
    closeCurrent: 'Fermer l?onglet courant',
    closeOther: 'Fermer les autres',
    closeLeft: 'Fermer ? gauche',
    closeRight: 'Fermer ? droite',
    closeAll: 'Tout fermer'
  },
  icon: {
    ...enUs.icon,
    themeConfig: 'Configuration du th?me',
    themeSchema: 'Sch?ma de th?me',
    lang: 'Changer de langue',
    fullscreen: 'Plein ?cran',
    fullscreenExit: 'Quitter le plein ?cran',
    reload: 'Recharger la page',
    collapse: 'R?duire le menu',
    expand: 'D?velopper le menu',
    pin: '?pingler',
    unpin: 'D?s?pingler'
  },
  datatable: {
    ...enUs.datatable,
    itemCount: 'Total {total} ?l?ments'
  }
};

export default local;
