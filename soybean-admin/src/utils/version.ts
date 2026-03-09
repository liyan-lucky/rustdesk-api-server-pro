export function getVersionTag() {
  return 'latest';
}

export function appendVersion(content: string) {
  return `${content} (Version: ${getVersionTag()})`;
}
