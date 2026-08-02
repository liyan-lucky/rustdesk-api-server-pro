import { h } from 'vue';
import { NButton } from 'naive-ui';
import { $t } from '../locales';
import { appendVersion, getVersionTag } from '@/utils/version';

export function setupAppVersionNotification() {
  let isShow = false;

  const checkVersion = async () => {
    if (isShow || import.meta.env.DEV) return;
    const [buildTime, serverVersion] = await Promise.all([getHtmlBuildTime(), getServerVersion()]);
    const currentVersion = getVersionTag().replace(/^v/, '');
    if (buildTime === BUILD_TIME && (!serverVersion || serverVersion === currentVersion)) return;
    isShow = true;
    const target = serverVersion || 'new build';
    const n = window.$notification?.create({
      title: `${$t('system.updateTitle')} (${currentVersion} → ${target})`,
      content: `${appendVersion($t('system.updateContent'))} Server: ${target}`,
      action() {
        return h('div', { style: { display: 'flex', justifyContent: 'end', gap: '12px', width: '325px' } }, [
          h(NButton, { onClick() { n?.destroy(); } }, () => $t('system.updateCancel')),
          h(NButton, { type: 'primary', onClick() { location.reload(); } }, () => $t('system.updateConfirm'))
        ]);
      },
      onClose() { isShow = false; }
    });
  };

  document.addEventListener('visibilitychange', () => {
    if (document.visibilityState === 'visible') void checkVersion();
  });
	window.setTimeout(() => void checkVersion(), 10_000);
  window.setInterval(() => void checkVersion(), 5 * 60 * 1000);
}

async function getServerVersion() {
  try {
    const baseURL = import.meta.env.VITE_BASE_URL;
    const res = await fetch(`${baseURL}api/version`, { cache: 'no-store' });
    if (!res.ok) return '';
    const data = await res.json();
    return String(data?.compat_target?.server?.version || '').replace(/^v/, '');
  } catch { return ''; }
}

async function getHtmlBuildTime() {
  const baseURL = import.meta.env.VITE_BASE_URL;

  const res = await fetch(`${baseURL}index.html`);

  const html = await res.text();

  const match = html.match(/<meta name="buildTime" content="(.*)">/);

  const buildTime = match?.[1] || '';

  return buildTime;
}
