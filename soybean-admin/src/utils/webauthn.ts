/**
 * WebAuthn 浏览器端工具函数
 *
 * 处理 base64url 编解码、ArrayBuffer 转换、以及 WebAuthn 凭据序列化
 */

/** 将 ArrayBuffer 转为 base64url 字符串 */
export function bufferToBase64url(buffer: ArrayBuffer | ArrayBufferView): string {
  const bytes = buffer instanceof ArrayBuffer ? new Uint8Array(buffer) : new Uint8Array(buffer.buffer, buffer.byteOffset, buffer.byteLength);
  let str = '';
  for (let i = 0; i < bytes.length; i++) {
    str += String.fromCharCode(bytes[i]);
  }
  const base64 = btoa(str);
  return base64.replace(/\+/g, '-').replace(/\//g, '_').replace(/=+$/, '');
}

/** 将 base64url 字符串转为 ArrayBuffer */
export function base64urlToBuffer(base64url: string): ArrayBuffer {
  const base64 = base64url.replace(/-/g, '+').replace(/_/g, '/');
  const pad = base64.length % 4 === 0 ? '' : '='.repeat(4 - (base64.length % 4));
  const binary = atob(base64 + pad);
  const bytes = new Uint8Array(binary.length);
  for (let i = 0; i < binary.length; i++) {
    bytes[i] = binary.charCodeAt(i);
  }
  return bytes.buffer;
}

/** 检测当前浏览器是否支持 WebAuthn / Passkey */
export function isWebAuthnSupported(): boolean {
  return typeof window !== 'undefined' && typeof window.PublicKeyCredential !== 'undefined';
}

/** 将后端返回的 PublicKeyCredentialRequestOptions 中的 base64url 字段转为 ArrayBuffer */
export function prepareAssertionOptions(options: Record<string, any>): PublicKeyCredentialRequestOptions {
  const prepared: PublicKeyCredentialRequestOptions = {
    challenge: base64urlToBuffer(options.challenge),
    rpId: options.rpId,
    timeout: options.timeout,
    userVerification: options.userVerification
  };

  if (options.allowCredentials && Array.isArray(options.allowCredentials)) {
    prepared.allowCredentials = options.allowCredentials.map((cred: any) => ({
      id: base64urlToBuffer(cred.id),
      type: cred.type,
      transports: cred.transports
    }));
  }

  return prepared;
}

/** 将后端返回的 PublicKeyCredentialCreationOptions 中的 base64url 字段转为 ArrayBuffer */
export function prepareCreationOptions(options: Record<string, any>): PublicKeyCredentialCreationOptions {
  const prepared: PublicKeyCredentialCreationOptions = {
    challenge: base64urlToBuffer(options.challenge),
    rp: options.rp,
    user: {
      ...options.user,
      id: base64urlToBuffer(options.user.id)
    },
    pubKeyCredParams: options.pubKeyCredParams,
    timeout: options.timeout,
    excludeCredentials: options.excludeCredentials?.map((cred: any) => ({
      ...cred,
      id: base64urlToBuffer(cred.id)
    })),
    authenticatorSelection: options.authenticatorSelection,
    attestation: options.attestation
  };

  return prepared;
}

/** 将 navigator.credentials.get() 的结果序列化为可发送给后端的 JSON */
export function serializeAssertion(credential: PublicKeyCredential): Record<string, any> {
  const assertion = credential as PublicKeyCredential & {
    response: AuthenticatorAssertionResponse;
  };

  return {
    id: credential.id,
    rawId: bufferToBase64url(credential.rawId),
    type: credential.type,
    response: {
      authenticatorData: bufferToBase64url(assertion.response.authenticatorData),
      clientDataJSON: bufferToBase64url(assertion.response.clientDataJSON),
      signature: bufferToBase64url(assertion.response.signature),
      userHandle: assertion.response.userHandle ? bufferToBase64url(assertion.response.userHandle) : null
    }
  };
}

/** 将 navigator.credentials.create() 的结果序列化为可发送给后端的 JSON */
export function serializeCreation(credential: PublicKeyCredential): Record<string, any> {
  const creation = credential as PublicKeyCredential & {
    response: AuthenticatorAttestationResponse;
  };

  return {
    id: credential.id,
    rawId: bufferToBase64url(credential.rawId),
    type: credential.type,
    response: {
      attestationObject: bufferToBase64url(creation.response.attestationObject),
      clientDataJSON: bufferToBase64url(creation.response.clientDataJSON)
    }
  };
}
