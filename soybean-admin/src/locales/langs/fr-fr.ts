import enUs from './en-us';

const local: App.I18n.Schema = {
  ...enUs,
  system: { ...enUs.system, title: 'Rustdesk Api Server' },
  common: {
    ...enUs.common,
    refresh: 'Actualiser',
    noData: 'Aucune donnée',
    error: 'Erreur',
    warning: 'Avertissement'
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
      created_at: 'Créé le'
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
      created_at: 'Créé le'
    },
    mailTemplate: {
      ...enUs.dataMap.mailTemplate,
      name: 'Nom',
      type: 'Type',
      subject: 'Sujet',
      contents: 'Contenu',
      created_at: 'Créé le'
    },
    mailLog: {
      ...enUs.dataMap.mailLog,
      username: 'Utilisateur',
      from: 'De',
      to: 'À',
      subject: 'Sujet',
      status: 'Statut',
      created_at: 'Envoyé le'
    }
  },
  icon: { ...enUs.icon, lang: 'Changer de langue', reload: 'Recharger la page', fullscreen: 'Plein écran' }
};

export default local;
