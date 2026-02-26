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
    update: 'Actualizar',
    updateSuccess: 'Actualizado con ?xito',
    userCenter: 'Centro de usuario',
    yesOrNo: {
      yes: 'S?',
      no: 'No'
    }
  },
  route: {
    ...enUs.route,
    home: 'Inicio',
    audit: 'Auditoría',
    user: 'Gestión de usuarios',
    user_list: 'Lista de usuarios',
    user_sessions: 'Sesiones',
    system: 'Gestión del sistema',
    system_mail_template: 'Plantillas de correo',
    system_mail_logs: 'Registros de correo',
    system_mail: 'Correo',
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
    UserUpdateSuccess: 'Usuario modificado correctamente',
    UserDeleteSuccess: 'Usuario eliminado correctamente',
    SessionKillSuccess: 'Sesión finalizada correctamente',
    MailTemplateNameEmpty: 'El nombre no puede estar vacío',
    MailTemplateSubjectEmpty: 'El asunto no puede estar vacío',
    MailTemplateContentsEmpty: 'El contenido no puede estar vacío',
    MailTemplateAddSuccess: 'Plantilla de correo creada correctamente',
    MailTemplateUpdateSuccess: 'Plantilla de correo modificada correctamente',
    NoEmailAddress: 'No hay dirección de correo configurada',
    VerificationCodeError: 'Error en el código de verificación',
    UUIDEmpty: 'UUID no puede estar vacío'
  },
  icon: { ...enUs.icon, lang: 'Cambiar idioma', reload: 'Recargar página', fullscreen: 'Pantalla completa' }
};

export default local;
