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
    update: 'Mettre ? jour',
    updateSuccess: 'Mise ? jour r?ussie',
    userCenter: 'Centre utilisateur',
    yesOrNo: {
      yes: 'Oui',
      no: 'Non'
    }
  },
  route: {
    ...enUs.route,
    home: 'Accueil',
    audit: 'Audit',
    user: 'Gestion des utilisateurs',
    user_list: 'Liste des utilisateurs',
    user_sessions: 'Sessions',
    system: 'Gestion système',
    system_mail_template: 'Modèles e-mail',
    system_mail_logs: 'Logs e-mail',
    system_mail: 'E-mail',
    devices: 'Appareils'
  },
  page: {
    ...enUs.page,
    home: {
      ...enUs.page.home,
      greeting: 'Bonjour {userName}, excellente journée !',
      userCount: 'Utilisateurs',
      deviceCount: 'Appareils',
      onlineCount: 'En ligne',
      visitsCount: 'Visites',
      operatingSystem: 'Système d’exploitation',
      oneWeek: 'Une semaine',
      changeLogs: 'Journal des modifications',
      serverConfig: {
        ...enUs.page.home.serverConfig,
        title: 'Configuration de connexion client',
        tip: 'Copiez ces valeurs dans le client RustDesk. Si KEY est vide, configurez `RUSTDESK_KEY`.',
        refresh: 'Actualiser',
        copyAll: 'Tout copier',
        copyTemplate: 'Copier le modèle RustDesk',
        connectivity: {
          ...enUs.page.home.serverConfig.connectivity,
          check: 'Tester la connectivité',
          checkOne: 'Tester',
          clear: 'Effacer les résultats'
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
      name: 'Pseudo',
      email: 'E-mail',
      licensed_devices: 'Appareils autorisés',
      login_verify: 'Vérification connexion',
      status: 'Statut',
      is_admin: 'Admin',
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
    device: {
      ...enUs.dataMap.device,
      username: 'Nom utilisateur',
      hostname: 'Nom de l’hôte',
      version: 'Version RustDesk',
      os: 'OS',
      rustdesk_id: 'Rustdesk ID'
    },
    audit: {
      ...enUs.dataMap.audit,
      username: 'Utilisateur',
      type: 'Type',
      rustdesk_id: 'Rustdesk ID',
      ip: 'IP',
      created_at: 'Créé le',
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
      }
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
      from: 'De',
      to: 'À',
      subject: 'Sujet',
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
  icon: { ...enUs.icon, lang: 'Changer de langue', reload: 'Recharger la page', fullscreen: 'Plein écran' }
};

export default local;
