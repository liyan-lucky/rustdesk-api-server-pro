export function getVersionTag() {
  return import.meta.env.VITE_APP_VERSION || 'latest';
}

export function getBuildTime() {
  return import.meta.env.VITE_BUILD_TIME || '';
}

export function appendVersion(content: string) {
  return `${content} (Version: ${getVersionTag()})`;
}
