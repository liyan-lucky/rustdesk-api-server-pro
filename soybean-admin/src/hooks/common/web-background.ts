import { computed, ref } from 'vue';

export type WebBackgroundMode = 'fixed' | 'bing' | 'upload';

const MODE_KEY = 'rustdesk-background-mode';
const UPLOAD_KEY = 'rustdesk-background-upload';
const GLOBAL_KEY = 'rustdesk-background-global';

function readMode(): WebBackgroundMode {
  const value = localStorage.getItem(MODE_KEY);
  return value === 'bing' || value === 'upload' ? value : 'fixed';
}

const backgroundMode = ref<WebBackgroundMode>(readMode());
const uploadedBackground = ref(localStorage.getItem(UPLOAD_KEY) || '');
const globalBackgroundEnabled = ref(localStorage.getItem(GLOBAL_KEY) !== 'false');
const backgroundUrl = computed(() => {
  if (backgroundMode.value === 'bing') return '/api/background/bing';
  if (backgroundMode.value === 'upload' && uploadedBackground.value) return uploadedBackground.value;
  return '/login-background.jpg';
});
const backgroundStyle = computed(() => ({
  backgroundImage: `linear-gradient(rgba(5, 14, 31, 0.34), rgba(5, 14, 31, 0.48)), url("${backgroundUrl.value}")`,
  backgroundPosition: 'center',
  backgroundRepeat: 'no-repeat',
  backgroundSize: 'cover'
}));

function setBackgroundMode(mode: WebBackgroundMode) {
  backgroundMode.value = mode;
  localStorage.setItem(MODE_KEY, mode);
}

function setGlobalBackgroundEnabled(enabled: boolean) {
  globalBackgroundEnabled.value = enabled;
  localStorage.setItem(GLOBAL_KEY, String(enabled));
}

function loadImage(file: File) {
  return new Promise<HTMLImageElement>((resolve, reject) => {
    const image = new Image();
    const objectUrl = URL.createObjectURL(file);
    image.onload = () => {
      URL.revokeObjectURL(objectUrl);
      resolve(image);
    };
    image.onerror = () => {
      URL.revokeObjectURL(objectUrl);
      reject(new Error('ERR-1019: 无法读取背景图片'));
    };
    image.src = objectUrl;
  });
}

export async function saveUploadedBackground(file: File) {
  if (!file.type.startsWith('image/')) throw new Error('ERR-1019: 请选择图片文件');
  const image = await loadImage(file);
  const scale = Math.min(1, 2560 / Math.max(image.naturalWidth, image.naturalHeight));
  const canvas = document.createElement('canvas');
  canvas.width = Math.max(1, Math.round(image.naturalWidth * scale));
  canvas.height = Math.max(1, Math.round(image.naturalHeight * scale));
  canvas.getContext('2d')?.drawImage(image, 0, 0, canvas.width, canvas.height);
  const dataUrl = canvas.toDataURL('image/jpeg', 0.84);
  if (dataUrl.length > 3_800_000) throw new Error('ERR-1020: 图片过大，请选择较小的图片');
  localStorage.setItem(UPLOAD_KEY, dataUrl);
  uploadedBackground.value = dataUrl;
  setBackgroundMode('upload');
}

export function clearUploadedBackground() {
  localStorage.removeItem(UPLOAD_KEY);
  uploadedBackground.value = '';
  if (backgroundMode.value === 'upload') setBackgroundMode('fixed');
}

export function useWebBackground() {
  return { backgroundMode, uploadedBackground, globalBackgroundEnabled, backgroundStyle, setBackgroundMode, setGlobalBackgroundEnabled };
}
