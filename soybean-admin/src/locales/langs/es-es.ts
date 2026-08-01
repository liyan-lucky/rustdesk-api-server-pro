import enUs from './en-us';

const local: App.I18n.Schema = {
  ...enUs,
  system: { ...enUs.system, title: 'Rustdesk Api Server' },
  common: {
    ...enUs.common,
    action: 'Acción',
    add: 'Agregar',
    addSuccess: 'Agregado con Éxito',
    backToHome: 'Volver al inicio',
    batchDelete: 'Eliminar por lote',
    cancel: 'Cancelar',
    close: 'Cerrar',
    check: 'Comprobar',
    expandColumn: 'Expandir columna',
    columnSetting: 'configuración de columnas',
    config: 'configuración',
    confirm: 'Confirmar',
    delete: 'Eliminar',
    deleteSuccess: 'Eliminado con Éxito',
    confirmDelete: '?Seguro que deseas eliminar?',
    edit: 'Editar',
    look: 'Ver',
    warning: 'Advertencia',
    error: 'Error',
    index: 'Índice',
    keywordSearch: 'Introduce una palabra clave',
    logout: 'Cerrar sesión',
    logoutConfirm: '?Seguro que deseas cerrar sesión?',
    lookForward: 'Próximamente',
    modify: 'Modificar',
    modifySuccess: 'Modificado con Éxito',
    noData: 'Sin datos',
    operate: 'operación',
    pleaseCheckValue: 'Comprueba si el valor es válido',
    refresh: 'Actualizar',
    reset: 'Restablecer',
    search: 'Buscar',
    switch: 'Cambiar',
    tip: 'Consejo',
    trigger: 'Activar',
    update: 'Actualizar',
    updateSuccess: 'Actualización exitosa',
    userCenter: 'Centro de usuario',
    yesOrNo: {
      yes: 'S?',
      no: 'No'
    }
  },
  request: {
    ...enUs.request,
    logout: 'Cerrar sesión del usuario tras error de solicitud',
    logoutMsg: 'Estado de usuario inválido, inicia sesión de nuevo',
    logoutWithModal: 'Mostrar diálogo tras error de solicitud y luego cerrar sesión',
    logoutWithModalMsg: 'Estado de usuario inválido, inicia sesión de nuevo',
    refreshToken: 'El token expir?, se actualizar?',
    tokenExpired: 'El token de la solicitud ha expirado'
  },
  theme: {
    ...enUs.theme,
    themeSchema: {
      ...enUs.theme.themeSchema,
      title: 'Esquema de tema',
      light: 'Claro',
      dark: 'Oscuro',
      auto: 'Seguir sistema'
    },
    grayscale: 'Escala de grises',
    colourWeakness: 'Deficiencia de color',
    layoutMode: {
      ...enUs.theme.layoutMode,
      title: 'Modo de diseño',
      vertical: 'Menú vertical',
      horizontal: 'Menú horizontal',
      'vertical-mix': 'Modo mixto vertical',
      'horizontal-mix': 'Modo mixto horizontal',
      reverseHorizontalMix: 'Invertir posición de Menús de primer y segundo nivel'
    },
    recommendColor: 'Aplicar algoritmo de color recomendado',
    recommendColorDesc: 'El algoritmo de color recomendado se refiere a',
    themeColor: {
      ...enUs.theme.themeColor,
      title: 'Color del tema',
      primary: 'Primario',
      info: 'Info',
      success: 'Éxito',
      warning: 'Advertencia',
      error: 'Error',
      followPrimary: 'Seguir color primario'
    },
    scrollMode: {
      ...enUs.theme.scrollMode,
      title: 'Modo de desplazamiento',
      wrapper: 'Contenedor',
      content: 'Contenido'
    },
    page: {
      ...enUs.theme.page,
      animate: 'Animación de página',
      mode: {
        ...enUs.theme.page.mode,
        title: 'Modo de animación',
        fade: 'Desvanecer',
        'fade-slide': 'Deslizar',
        'fade-bottom': 'Zoom desvanecido',
        'fade-scale': 'Escala desvanecida',
        'zoom-fade': 'Zoom con desvanecido',
        'zoom-out': 'Alejar',
        none: 'Ninguno'
      }
    },
    fixedHeaderAndTab: 'Fijar cabecera y pestañas',
    header: {
      ...enUs.theme.header,
      height: 'Altura de cabecera',
      breadcrumb: {
        ...enUs.theme.header.breadcrumb,
        visible: 'Breadcrumb visible',
        showIcon: 'Icono de breadcrumb visible'
      }
    },
    tab: {
      ...enUs.theme.tab,
      visible: 'pestañas visibles',
      cache: 'Caché de pestañas',
      height: 'Altura de pestañas',
      mode: {
        ...enUs.theme.tab.mode,
        title: 'Modo de pestañas',
        chrome: 'Chrome',
        button: 'Botón'
      }
    },
    sider: {
      ...enUs.theme.sider,
      inverted: 'Barra lateral oscura',
      width: 'Ancho de barra lateral',
      collapsedWidth: 'Ancho colapsado',
      mixWidth: 'Ancho modo mixto',
      mixCollapsedWidth: 'Ancho colapsado mixto',
      mixChildMenuWidth: 'Ancho subMenú mixto'
    },
    footer: {
      ...enUs.theme.footer,
      visible: 'Pie visible',
      fixed: 'Pie fijo',
      height: 'Altura del pie',
      right: 'Pie derecho'
    },
    watermark: {
      ...enUs.theme.watermark,
      visible: 'Marca de agua visible en pantalla completa',
      text: 'Texto de marca de agua'
    },
    themeDrawerTitle: 'configuración de tema',
    pageFunTitle: 'Funciones de página',
    configOperation: {
      ...enUs.theme.configOperation,
      copyConfig: 'Copiar configuración',
      copySuccessMsg: 'Copia correcta, sustituye la variable "themeSettings" en "src/theme/settings.ts"',
      resetConfig: 'Restablecer configuración',
      resetSuccessMsg: 'Restablecido con Éxito'
    }
  },
  route: {
    ...enUs.route,
    login: 'Iniciar sesión',
    403: 'Sin permiso',
    404: 'página no encontrada',
    500: 'Error del servidor',
    'iframe-page': 'Iframe',
    home: 'Inicio',
    audit: 'Auditoría',
    user: 'Gestión de usuarios',
    user_list: 'Lista de usuarios',
    user_sessions: 'Sesiones',
    system: 'Gestion del sistema',
    system_mail_template: 'Plantillas de correo',
    system_mail_logs: 'registros de correo',
    system_mail: 'Correo',
    audit_baselogs: 'registros base',
    audit_filetransferlogs: 'registros de transferencia',
    audit_loginlogs: 'registros de inicio de sesión',
    devices: 'Dispositivos',
    'my-devices': 'Contactos',
    'my-devices_peers': 'Mis contactos',
    'my-devices_manage': 'Gestión de libretas',
    'my-devices_tags': 'Gestión de etiquetas',
    system_server: 'Configuración del servidor',
    system_tokens: 'Token de usuario',
    system_oauth: 'Gestión OAuth'
  },
  page: {
    ...enUs.page,
      login: {
        ...enUs.page.login,
        common: {
          ...enUs.page.login.common,
          loginOrRegister: 'Iniciar sesión / Registrarse',
          userNamePlaceholder: 'Introduce el nombre de usuario',
          phonePlaceholder: 'Introduce el número de teléfono',
          codePlaceholder: 'Introduce el código de verificación',
          passwordPlaceholder: 'Introduce la contraseña',
          confirmPasswordPlaceholder: 'Introduce la contraseña de nuevo',
          codeLogin: 'Inicio con código',
          confirm: 'Confirmar',
          back: 'Volver',
          validateSuccess: 'verificación correcta',
          loginSuccess: 'Inicio de sesión correcto',
          welcomeBack: 'Bienvenido de nuevo, {userName} !',
          thirdPartyLogin: 'Inicio de sesión de terceros',
          continueWith: 'Continuar con {provider}',
          providerUnavailable: 'El inicio de sesión con {provider} no está disponible'
        },
        pwdLogin: {
          ...enUs.page.login.pwdLogin,
          title: 'Inicio con contraseña',
          rememberMe: 'Recordarme',
          switchToUser: 'Inicio de sesión de usuario'
         },
         userLogin: {
           title: 'Inicio de sesión de usuario',
           switchToAdmin: 'Inicio de sesión de admin'
        }
      },
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
      cardDetail: {
        viewHint: 'Haz clic para ver detalles',
        recentUsers: 'Usuarios recientes',
        recentDevices: 'Dispositivos recientes',
        recentVisits: 'registros de acceso recientes',
        desc: {
          userCount: 'Muestra el numero total de usuarios del sistema.',
          deviceCount: 'Muestra el numero total de dispositivos del sistema.',
          onlineCount: 'Muestra los dispositivos en linea segun estadisticas de heartbeat.',
          visitCount: 'Muestra estadisticas de visitas desde los registros de auditoria.'
        }
      },
      serverConfig: {
        ...enUs.page.home.serverConfig,
        title: 'configuración de conexión del cliente',
        tip: 'Copia los siguientes valores en el cliente RustDesk. Si KEY está vacío, configura `RUSTDESK_KEY` como variable de entorno del contenedor.',
        idServer: 'Servidor ID',
        relayServer: 'Servidor relay',
        apiServer: 'Servidor API',
        key: 'KEY',
        idServerPlaceholder: 'p. ej. your.domain.com',
        relayServerPlaceholder: 'p. ej. your.domain.com',
        apiServerPlaceholder: 'p. ej. https://your.domain.com',
        keyPlaceholder: 'Proporcionar mediante la variable RUSTDESK_KEY',
        copy: 'Copiar',
        copyAll: 'Copiar todo',
        copyTemplate: 'Copiar plantilla RustDesk',
        showQr: 'Mostrar código QR',
        qrTitle: 'Código QR de importación RustDesk',
        qrTip: 'Escanee este código QR en la aplicación RustDesk móvil para importar la configuración.',
        qrPayload: 'Texto de plantilla RustDesk',
        qrFailed: 'Error al generar código QR',
        refresh: 'Actualizar configuración',
        clearCacheReload: 'Limpiar Caché y recargar',
        source: 'Origen',
        lastUpdated: 'Ultima actualizacion',
        show: 'Mostrar',
        hide: 'Ocultar',
        missingTip: 'Los siguientes campos están vacíos. Configúralos primero en las variables de entorno del contenedor: {fields}',
        copyEmpty: '{label} está vacío y no se puede copiar',
        copySuccess: '{label} copiado',
        copyFailed: 'Error al copiar {label}',
        fetchFailed: 'No se pudo cargar la configuración del servidor',
        cacheCleared: 'Caché limpiada, recargando configuración del servidor',
        sourceType: {
          ...enUs.page.home.serverConfig.sourceType,
          remote: 'Remoto',
          'memory-cache': 'Caché en memoria',
          'session-cache': 'Caché de sesión',
          env: 'Entorno',
          inferred: 'Inferido',
          empty: 'Vacío',
           auto: 'Auto-detectado'
        },
        sourceHint: {
          ...enUs.page.home.serverConfig.sourceHint,
          env: 'Este valor proviene de una variable de entorno del contenedor.',
          inferred: 'Este valor se infiere automáticamente de la dirección de acceso actual.',
          empty: 'Aún no hay valor configurado ni inferido.'
        },
        connectivity: {
          ...enUs.page.home.serverConfig.connectivity,
          clear: 'Limpiar resultados',
          check: 'Comprobar conectividad',
          checkOne: 'Comprobar',
          checked: 'Comprobación de conectividad completada',
          checkedOne: 'Conectividad de {field} comprobada',
          checkedCached: 'Usando resultado reciente de conectividad (Caché)',
          checkFailed: 'Error en la comprobación de conectividad',
          cleared: 'Resultados de conectividad limpiados',
          source: 'Origen de comprobación',
          lastChecked: 'Última comprobación',
          target: 'Destino',
          duration: 'Duración',
          notChecked: 'Aún no comprobado',
          checkSourceType: {
            ...enUs.page.home.serverConfig.connectivity.checkSourceType,
            remote: 'Remoto',
            cache: 'Caché'
          },
          status: {
            ...enUs.page.home.serverConfig.connectivity.status,
            idle: 'Sin comprobar',
            ok: 'Accesible',
            error: 'Fallido',
            skip: 'Omitido'
          }
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
    },
    system: {
      ...enUs.page.system,
      mailTemplate: {
        ...enUs.page.system.mailTemplate,
        addMailTemplate: 'Agregar plantilla',
        editMailTemplate: 'Editar plantilla',
        inputName: 'Ingresar nombre',
        inputSubject: 'Ingresar asunto',
        inputContents: 'Ingresar contenido',
        selectType: 'Seleccionar tipo'
      },
      mailLog: {
        ...enUs.page.system.mailLog,
        info: 'Detalle'
      }
    },
    myDevices: {
      title: 'Contactos',
      welcome: 'Bienvenido, {userName}',
      status: 'Estado',
      online: 'En línea',
      offline: 'Desconectado',
      conns: 'Conexiones',
      lastSync: 'Última sincronización',
      logout: 'Cerrar sesión'
    }
  },
  dataMap: {
    ...enUs.dataMap,
    user: {
      ...enUs.dataMap.user,
      username: 'Usuario',
        password: 'Contraseña',
      name: 'Apodo',
      email: 'Correo',
      licensed_devices: 'Dispositivos licenciados',
      login_verify: 'Verificación de acceso',
      status: 'Estado',
      is_admin: 'Admin',
        tfa_secret: 'Secreto 2FA',
        tfa_code: 'Código 2FA',
      created_at: 'Creado el',
      statusLabel: {
        ...enUs.dataMap.user.statusLabel,
        disabled: 'Deshabilitado',
        unverified: 'No verificado',
        normal: 'Normal'
      },
      loginVerifyLabel: {
        ...enUs.dataMap.user.loginVerifyLabel,
        none: 'Ninguna',
        emailCheck: 'Verificación por correo',
        tfaCheck: '2FA'
      }
    },
      session: {
        ...enUs.dataMap.session,
        expired: 'Expira el',
        created_at: 'Creado el'
      },
    device: {
      ...enUs.dataMap.device,
      username: 'Usuario',
      hostname: 'Nombre del host',
      version: 'Versión de RustDesk',
        memory: 'Memoria',
      os: 'SO',
      rustdesk_id: 'Rustdesk ID'
    },
    ab: {
      rustdesk_id: 'Rustdesk ID',
      username: 'Nombre de usuario',
      hostname: 'Nombre de host',
      tags: 'Etiquetas',
      alias: 'Alias',
      hash: 'Hash',
      owner: 'Propietario',
      name: 'Nombre de libreta',
      user_id: 'ID de usuario',
      guid: 'GUID',
      rule: 'Regla',
      max_peer: 'Contactos máx.',
      shared: 'Compartido',
      ab_id: 'ID de libreta',
      tagName: 'Nombre',
      tagColor: 'Color',
      updated_at: 'Actualizado el'
    },
    token: {
      device_os: 'SO del dispositivo',
      device_name: 'Nombre del dispositivo',
      token_hash: 'Hash de Token',
      is_admin: 'Admin',
      status: 'Activo'
    },
    oauth: {
      provider: 'Proveedor',
      subject: 'Sujeto',
      email: 'Correo',
      name: 'Nombre',
      last_login_at: 'Último inicio'
    },
    loginLog: {
      allEvents: 'Todos los eventos',
      event: 'Evento',
      userAgent: 'Agente de usuario',
      success: 'Éxito',
      reason: 'Razón'
    },
    audit: {
      ...enUs.dataMap.audit,
      username: 'Usuario',
      type: 'Tipo',
        conn_id: 'ID de conexión',
      rustdesk_id: 'Rustdesk ID',
        peer_id: 'ID de peer',
      ip: 'IP',
        session_id: 'ID de sesión',
        uuid: 'UUID',
      created_at: 'Creado el',
        closed_at: 'Cerrado el',
      typeLabel: {
        ...enUs.dataMap.audit.typeLabel,
        remote_control: 'Control remoto',
        file_transfer: 'Transferencia de archivos',
        tcp_tunnel: 'Túnel TCP'
      },
      fileTransferTypeLabel: {
        ...enUs.dataMap.audit.fileTransferTypeLabel,
        master_controlled: 'Controlador -> Controlado',
        controlled_master: 'Controlado -> Controlador'
      },
        path: 'Ruta'
    },
    mailTemplate: {
      ...enUs.dataMap.mailTemplate,
      name: 'Nombre',
      type: 'Tipo',
      subject: 'Asunto',
      contents: 'Contenido',
      created_at: 'Creado el',
      typeLabel: {
        ...enUs.dataMap.mailTemplate.typeLabel,
        loginVerify: 'Verificación de inicio de sesión',
        registerVerify: 'Verificación de registro',
        other: 'Otro'
      }
    },
    mailLog: {
      ...enUs.dataMap.mailLog,
      username: 'Usuario',
        uuid: 'UUID',
      from: 'De',
      to: 'Para',
      subject: 'Asunto',
        contents: 'Contenido',
      status: 'Estado',
      created_at: 'Enviado el',
      statusLabel: {
        ...enUs.dataMap.mailLog.statusLabel,
        ok: 'Éxito',
        err: 'Error'
      }
    }
  },
  api: {
    ...enUs.api,
    CaptchaError: 'Error de CAPTCHA',
    UserNotExists: 'El usuario no existe',
    UsernameOrPasswordError: 'Cuenta o contraseña incorrecta',
    UserExists: 'El nombre de usuario ya está en uso',
    UsernameEmpty: 'El nombre de usuario no puede estar vacío',
    PasswordEmpty: 'La contraseña no puede estar vacía',
    UserAddSuccess: 'Usuario creado correctamente',
    DataError: 'Error de datos',
    RequestError: 'Solicitud fallida',
    UserUpdateSuccess: 'Usuario actualizado correctamente',
    UserDeleteSuccess: 'Usuario eliminado correctamente',
    SessionKillSuccess: 'Sesión finalizada correctamente',
    MailTemplateNameEmpty: 'El nombre no puede estar vacío',
    MailTemplateSubjectEmpty: 'El asunto no puede estar vacío',
    MailTemplateContentsEmpty: 'El contenido no puede estar vacío',
    MailTemplateAddSuccess: 'Plantilla de correo creada correctamente',
    MailTemplateUpdateSuccess: 'Plantilla de correo actualizada correctamente',
    NoEmailAddress: 'No hay dirección de correo configurada',
    VerificationCodeError: 'Error en el código de verificación',
    UUIDEmpty: 'UUID no puede estar vacío'
  },
  dropdown: {
    ...enUs.dropdown,
    closeCurrent: 'Cerrar actual',
    closeOther: 'Cerrar otros',
    closeLeft: 'Cerrar izquierda',
    closeRight: 'Cerrar derecha',
    closeAll: 'Cerrar todo'
  },
  icon: {
    ...enUs.icon,
    themeConfig: 'configuración de tema',
    themeSchema: 'Esquema de tema',
    lang: 'Cambiar idioma',
    fullscreen: 'Pantalla completa',
    fullscreenExit: 'Salir de pantalla completa',
    reload: 'Recargar página',
    collapse: 'Colapsar Menú',
    expand: 'Expandir Menú',
    pin: 'Fijar',
    unpin: 'Desfijar'
  },
  datatable: {
    ...enUs.datatable,
    itemCount: 'Total {total} elementos'
  }
};

export default local;
