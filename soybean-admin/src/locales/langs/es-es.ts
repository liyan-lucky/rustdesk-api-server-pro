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
    },
    user: {
      ...enUs.page.user,
      list: {
        ...enUs.page.user.list,
        addUser: 'Agregar usuario',
        editUser: 'Editar usuario',
        searchPlaceholder: 'Usuario/Apodo/Correo'
      },
      sessions: {
        ...enUs.page.user.sessions,
        kill: 'Finalizar',
        confirmKill: '¿Finalizar esta sesión?'
      },
      audit: {
        ...enUs.page.user.audit,
        logsSearchPlaceholder: 'Usuario/Acción/RustdeskID/IP'
      },
      devices: {
        ...enUs.page.user.devices,
        logsSearchPlaceholder: 'Usuario/Host/RustdeskID'
      }
    }
  },
  dataMap: {
    ...enUs.dataMap,
    user: {
      ...enUs.dataMap.user,
      username: 'Usuario',
      name: 'Apodo',
      email: 'Correo',
      licensed_devices: 'Dispositivos licenciados',
      login_verify: 'Verificación de acceso',
      status: 'Estado',
      is_admin: 'Admin',
      created_at: 'Creado el'
    },
    device: {
      ...enUs.dataMap.device,
      username: 'Usuario',
      hostname: 'Nombre del host',
      version: 'Versión de RustDesk',
      os: 'SO',
      rustdesk_id: 'Rustdesk ID'
    },
    audit: {
      ...enUs.dataMap.audit,
      username: 'Usuario',
      type: 'Tipo',
      rustdesk_id: 'Rustdesk ID',
      ip: 'IP',
      created_at: 'Creado el'
    }
  },
  icon: { ...enUs.icon, lang: 'Cambiar idioma', reload: 'Recargar página', fullscreen: 'Pantalla completa' }
};

export default local;
