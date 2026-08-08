import { h } from 'vue';
import { NButton } from 'naive-ui';
import { appendVersion, getVersionTag } from '@/utils/version';
import { $t } from '../locales';

const DISMISS_KEY = 'rustdesk-api-version-dismissed';

function getDismissedVersion(): string {
  try {
    return localStorage.getItem(DISMISS_KEY) || '';
  } catch {
    return '';
  }
}

function setDismissedVersion(value: string) {
  try {
    localStorage.setItem(DISMISS_KEY, value);
  } catch {
    // ignore storage errors
  }
}

function clearDismissedVersion() {
  try {
    localStorage.removeItem(DISMISS_KEY);
  } catch {
    // ignore storage errors
  }
}

export function setupAppVersionNotification() {
  let isShow = false;

  const checkVersion = async () => {
    if (isShow || import.meta.env.DEV) return;
    const [buildTime, serverVersion] = await Promise.all([getHtmlBuildTime(), getServerVersion()]);
    const currentVersion = getVersionTag().replace(/^v/, '');
    if (buildTime === BUILD_TIME && (!serverVersion || serverVersion === currentVersion)) return;

    const target = serverVersion || 'new build';
    const dismissKey = `${currentVersion}|${target}|${buildTime}`;
    if (getDismissedVersion() === dismissKey) return;

    isShow = true;
    const n = window.$notification?.create({
      title: `${$t('system.updateTitle')} (${currentVersion} → ${target})`,
      content: `${appendVersion($t('system.updateContent'))} Server: ${target}`,
      action() {
        return h('div', { style: { display: 'flex', justifyContent: 'end', gap: '12px', width: '325px' } }, [
          h(
            NButton,
            {
              onClick() {
                setDismissedVersion(dismissKey);
                n?.destroy();
              }
            },
            () => $t('system.updateCancel')
          ),
          h(
            NButton,
            {
              type: 'primary',
              onClick() {
                clearDismissedVersion();
                location.reload();
              }
            },
            () => $t('system.updateConfirm')
          )
        ]);
      },
      onClose() {
        setDismissedVersion(dismissKey);
        isShow = false;
      }
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
  } catch {
    return '';
  }
}

async function getHtmlBuildTime() {
  const baseURL = import.meta.env.VITE_BASE_URL;

  const res = await fetch(`${baseURL}index.html`);

  const html = await res.text();

  const match = html.match(/<meta name="buildTime" content="(.*)">/);

  const buildTime = match?.[1] || '';

  return buildTime;
}
