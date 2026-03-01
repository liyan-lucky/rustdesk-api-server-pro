import enUs from './en-us';

const local: App.I18n.Schema = {
  ...enUs,
  system: { ...enUs.system, title: 'Rustdesk Api Server' },
  common: {
    ...enUs.common,
    action: 'Acci?n',
    add: 'Agregar',
    addSuccess: 'Agregado con ?xito',
    backToHome: 'Volver al inicio',
    batchDelete: 'Eliminar por lote',
    cancel: 'Cancelar',
    close: 'Cerrar',
    check: 'Comprobar',
    expandColumn: 'Expandir columna',
    columnSetting: 'Configuraci?n de columnas',
    config: 'Configuraci?n',
    confirm: 'Confirmar',
    delete: 'Eliminar',
    deleteSuccess: 'Eliminado con ?xito',
    confirmDelete: '?Seguro que deseas eliminar?',
    edit: 'Editar',
    look: 'Ver',
    warning: 'Advertencia',
    error: 'Error',
    index: '?ndice',
    keywordSearch: 'Introduce una palabra clave',
    logout: 'Cerrar sesi?n',
    logoutConfirm: '?Seguro que deseas cerrar sesi?n?',
    lookForward: 'Pr?ximamente',
    modify: 'Modificar',
    modifySuccess: 'Modificado con ?xito',
    noData: 'Sin datos',
    operate: 'Operaci?n',
    pleaseCheckValue: 'Comprueba si el valor es v?lido',
    refresh: 'Actualizar',
    reset: 'Restablecer',
    search: 'Buscar',
    switch: 'Cambiar',
    tip: 'Consejo',
    trigger: 'Activar',
    update: '更新',
    updateSuccess: '更新成功',
    userCenter: 'Centro de usuario',
    yesOrNo: {
      yes: 'S?',
      no: 'No'
    }
  },
  request: {
    ...enUs.request,
    logout: 'Cerrar sesi?n del usuario tras error de solicitud',
    logoutMsg: 'Estado de usuario inv?lido, inicia sesi?n de nuevo',
    logoutWithModal: 'Mostrar di?logo tras error de solicitud y luego cerrar sesi?n',
    logoutWithModalMsg: 'Estado de usuario inv?lido, inicia sesi?n de nuevo',
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
      title: 'Modo de dise?o',
      vertical: 'Men? vertical',
      horizontal: 'Men? horizontal',
      'vertical-mix': 'Modo mixto vertical',
      'horizontal-mix': 'Modo mixto horizontal',
      reverseHorizontalMix: 'Invertir posici?n de men?s de primer y segundo nivel'
    },
    recommendColor: 'Aplicar algoritmo de color recomendado',
    recommendColorDesc: 'El algoritmo de color recomendado se refiere a',
    themeColor: {
      ...enUs.theme.themeColor,
      title: 'Color del tema',
      primary: 'Primario',
      info: 'Info',
      success: '?xito',
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
      animate: 'Animaci?n de p?gina',
      mode: {
        ...enUs.theme.page.mode,
        title: 'Modo de animaci?n',
        fade: 'Desvanecer',
        'fade-slide': 'Deslizar',
        'fade-bottom': 'Zoom desvanecido',
        'fade-scale': 'Escala desvanecida',
        'zoom-fade': 'Zoom con desvanecido',
        'zoom-out': 'Alejar',
        none: 'Ninguno'
      }
    },
    fixedHeaderAndTab: 'Fijar cabecera y pesta?as',
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
      visible: 'Pesta?as visibles',
      cache: 'Cach? de pesta?as',
      height: 'Altura de pesta?as',
      mode: {
        ...enUs.theme.tab.mode,
        title: 'Modo de pesta?as',
        chrome: 'Chrome',
        button: 'Bot?n'
      }
    },
    sider: {
      ...enUs.theme.sider,
      inverted: 'Barra lateral oscura',
      width: 'Ancho de barra lateral',
      collapsedWidth: 'Ancho colapsado',
      mixWidth: 'Ancho modo mixto',
      mixCollapsedWidth: 'Ancho colapsado mixto',
      mixChildMenuWidth: 'Ancho submen? mixto'
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
    themeDrawerTitle: 'Configuraci?n de tema',
    pageFunTitle: 'Funciones de p?gina',
    configOperation: {
      ...enUs.theme.configOperation,
      copyConfig: 'Copiar configuraci?n',
      copySuccessMsg: 'Copia correcta, sustituye la variable "themeSettings" en "src/theme/settings.ts"',
      resetConfig: 'Restablecer configuraci?n',
      resetSuccessMsg: 'Restablecido con ?xito'
    }
  },
  route: {
    ...enUs.route,
    login: 'Iniciar sesi?n',
    403: 'Sin permiso',
    404: 'P?gina no encontrada',
    500: 'Error del servidor',
    'iframe-page': 'Iframe',
    home: 'Inicio',
    audit: 'Auditor?a',
    user: 'Gesti?n de usuarios',
    user_list: 'Lista de usuarios',
    user_sessions: 'Sesiones',
    system: 'Gesti?n del sistema',
    system_mail_template: 'Plantillas de correo',
    system_mail_logs: 'Registros de correo',
    system_mail: 'Correo',
    audit_baselogs: 'Registros base',
    audit_filetransferlogs: 'Registros de transferencia',
    devices: 'Dispositivos'
  },
  page: {
    ...enUs.page,
      login: {
        ...enUs.page.login,
        common: {
          ...enUs.page.login.common,
          loginOrRegister: 'Iniciar sesi?n / Registrarse',
          userNamePlaceholder: 'Introduce el nombre de usuario',
          phonePlaceholder: 'Introduce el n?mero de tel?fono',
          codePlaceholder: 'Introduce el c?digo de verificaci?n',
          passwordPlaceholder: 'Introduce la contrase?a',
          confirmPasswordPlaceholder: 'Introduce la contrase?a de nuevo',
          codeLogin: 'Inicio con c?digo',
          confirm: 'Confirmar',
          back: 'Volver',
          validateSuccess: 'Verificaci?n correcta',
          loginSuccess: 'Inicio de sesi?n correcto',
          welcomeBack: 'Bienvenido de nuevo, {userName} !'
        },
        pwdLogin: {
          ...enUs.page.login.pwdLogin,
          title: 'Inicio con contrase?a',
          rememberMe: 'Recordarme'
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
      serverConfig: {
        ...enUs.page.home.serverConfig,
        title: 'Configuraci?n de conexi?n del cliente',
        tip: 'Copia los siguientes valores en el cliente RustDesk. Si KEY est? vac?o, configura `RUSTDESK_KEY` como variable de entorno del contenedor.',
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
        refresh: 'Actualizar configuraci?n',
        clearCacheReload: 'Limpiar cach? y recargar',
        source: 'Origen',
        lastUpdated: '最后更新',
        show: 'Mostrar',
        hide: 'Ocultar',
        missingTip: 'Los siguientes campos est?n vac?os. Config?ralos primero en las variables de entorno del contenedor: {fields}',
        copyEmpty: '{label} est? vac?o y no se puede copiar',
        copySuccess: '{label} copiado',
        copyFailed: 'Error al copiar {label}',
        fetchFailed: 'No se pudo cargar la configuraci?n del servidor',
        cacheCleared: 'Cach? limpiada, recargando configuraci?n del servidor',
        sourceType: {
          ...enUs.page.home.serverConfig.sourceType,
          remote: 'Remoto',
          'memory-cache': 'Cach? en memoria',
          'session-cache': 'Cach? de sesi?n',
          env: 'Entorno',
          inferred: 'Inferido',
          empty: 'Vac?o'
        },
        sourceHint: {
          ...enUs.page.home.serverConfig.sourceHint,
          env: 'Este valor proviene de una variable de entorno del contenedor.',
          inferred: 'Este valor se infiere autom?ticamente de la direcci?n de acceso actual.',
          empty: 'A?n no hay valor configurado ni inferido.'
        },
        connectivity: {
          ...enUs.page.home.serverConfig.connectivity,
          clear: 'Limpiar resultados',
          check: 'Comprobar conectividad',
          checkOne: 'Comprobar',
          checked: 'Comprobaci?n de conectividad completada',
          checkedOne: 'Conectividad de {field} comprobada',
          checkedCached: 'Usando resultado reciente de conectividad (cach?)',
          checkFailed: 'Error en la comprobaci?n de conectividad',
          cleared: 'Resultados de conectividad limpiados',
          source: 'Origen de comprobaci?n',
          lastChecked: '?ltima comprobaci?n',
          target: 'Destino',
          duration: 'Duraci?n',
          notChecked: 'A?n no comprobado',
          checkSourceType: {
            ...enUs.page.home.serverConfig.connectivity.checkSourceType,
            remote: 'Remoto',
            cache: 'Cach?'
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
    }
  },
  dataMap: {
    ...enUs.dataMap,
    user: {
      ...enUs.dataMap.user,
      username: 'Usuario',
        password: 'Contrase?a',
      name: 'Apodo',
      email: 'Correo',
      licensed_devices: 'Dispositivos licenciados',
      login_verify: 'Verificación de acceso',
      status: 'Estado',
      is_admin: 'Admin',
        tfa_secret: 'Secreto 2FA',
        tfa_code: 'C?digo 2FA',
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
    audit: {
      ...enUs.dataMap.audit,
      username: 'Usuario',
      type: 'Tipo',
        conn_id: 'ID de conexi?n',
      rustdesk_id: 'Rustdesk ID',
        peer_id: 'ID de peer',
      ip: 'IP',
        session_id: 'ID de sesi?n',
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
    UserUpdateSuccess: '用户修改成功',
    UserDeleteSuccess: 'Usuario eliminado correctamente',
    SessionKillSuccess: 'Sesión finalizada correctamente',
    MailTemplateNameEmpty: 'El nombre no puede estar vacío',
    MailTemplateSubjectEmpty: 'El asunto no puede estar vacío',
    MailTemplateContentsEmpty: 'El contenido no puede estar vacío',
    MailTemplateAddSuccess: 'Plantilla de correo creada correctamente',
    MailTemplateUpdateSuccess: '邮件模板修改成功',
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
    themeConfig: 'Configuraci?n de tema',
    themeSchema: 'Esquema de tema',
    lang: 'Cambiar idioma',
    fullscreen: 'Pantalla completa',
    fullscreenExit: 'Salir de pantalla completa',
    reload: 'Recargar p?gina',
    collapse: 'Colapsar men?',
    expand: 'Expandir men?',
    pin: 'Fijar',
    unpin: 'Desfijar'
  },
  datatable: {
    ...enUs.datatable,
    itemCount: 'Total {total} elementos'
  }
};

export default local;
