export function getAppVersion() {
  return BUILD_TIME || 'unknown';
}

export function getVersionTag() {
  return `v${getAppVersion()}`;
}

export function appendVersion(content: string) {
  return `${content} (Version: ${getVersionTag()})`;
}
