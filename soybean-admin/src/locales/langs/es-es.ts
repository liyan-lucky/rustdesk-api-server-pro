import enUs from './en-us';

const local: App.I18n.Schema = {
  ...enUs,
  system: { ...enUs.system, title: 'Rustdesk Api Server' },
  common: {
    ...enUs.common,
    refresh: 'Actualizar',
    noData: 'Sin datos',
    error: 'Error',
    warning: 'Advertencia'
  },
  route: {
    ...enUs.route,
    home: 'Inicio',
    audit: 'Auditoría',
    user: 'Gestión de usuarios',
    user_list: 'Lista de usuarios',
    user_sessions: 'Sesiones',
    system: 'Gestión del sistema',
    devices: 'Dispositivos'
  },
  page: {
    ...enUs.page,
    home: {
      ...enUs.page.home,
      greeting: 'Buenos días, {userName}!',
      userCount: 'Usuarios',
      deviceCount: 'Dispositivos',
      onlineCount: 'En línea',
      visitsCount: 'Visitas',
      operatingSystem: 'Sistema operativo',
      oneWeek: 'Una semana',
      changeLogs: 'Registro de cambios',
      serverConfig: {
        ...enUs.page.home.serverConfig,
        title: 'Configuración de conexión del cliente',
        tip: 'Copia estos valores en el cliente RustDesk. Si KEY está vacío, configura `RUSTDESK_KEY`.',
        refresh: 'Actualizar',
        copyAll: 'Copiar todo',
        copyTemplate: 'Copiar plantilla RustDesk',
        connectivity: {
          ...enUs.page.home.serverConfig.connectivity,
          check: 'Comprobar conectividad',
          checkOne: 'Comprobar',
          clear: 'Limpiar resultados'
        }
      }
    }
  },
  icon: { ...enUs.icon, lang: 'Cambiar idioma', reload: 'Recargar página', fullscreen: 'Pantalla completa' }
};

export default local;
