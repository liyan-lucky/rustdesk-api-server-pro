export function getVersionTag() {
  return import.meta.env.VITE_APP_VERSION || 'latest';
}

export function appendVersion(content: string) {
  return `${content} (Version: ${getVersionTag()})`;
}
