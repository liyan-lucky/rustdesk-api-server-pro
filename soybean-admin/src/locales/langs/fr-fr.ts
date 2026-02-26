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
    }
  },
  icon: { ...enUs.icon, lang: 'Changer de langue', reload: 'Recharger la page', fullscreen: 'Plein écran' }
};

export default local;
