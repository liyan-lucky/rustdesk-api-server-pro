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
    const serverVersion = await getServerVersion();
    const currentVersion = getVersionTag().replace(/^v/, '');
    if (!serverVersion || !isVersionHigher(serverVersion, currentVersion)) return;

    const dismissKey = `${currentVersion}|${serverVersion}`;
    if (getDismissedVersion() === dismissKey) return;

    isShow = true;
    const n = window.$notification?.create({
      title: `${$t('system.updateTitle')} (${currentVersion} → ${serverVersion})`,
      content: `${appendVersion($t('system.updateContent'))} Server: ${serverVersion}`,
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

function isVersionHigher(a: string, b: string): boolean {
  const pa = a.split('.').map(Number);
  const pb = b.split('.').map(Number);
  for (let i = 0; i < Math.max(pa.length, pb.length); i++) {
    const va = pa[i] || 0;
    const vb = pb[i] || 0;
    if (va > vb) return true;
    if (va < vb) return false;
  }
  return false;
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
