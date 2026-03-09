import process from 'node:process';
import path from 'node:path';
import unocss from '@unocss/vite';
import presetIcons from '@unocss/preset-icons';
import { FileSystemIconLoader } from '@iconify/utils/lib/loader/node-loaders';
import { createRequire } from 'node:module';

const require = createRequire(import.meta.url);
const antDesign = require('@iconify-json/ant-design/icons.json');
const fluent = require('@iconify-json/fluent/icons.json');
const gridicons = require('@iconify-json/gridicons/icons.json');
const ic = require('@iconify-json/ic/icons.json');
const mdi = require('@iconify-json/mdi/icons.json');
const uil = require('@iconify-json/uil/icons.json');

export function setupUnocss(viteEnv: Env.ImportMeta) {
  const { VITE_ICON_PREFIX, VITE_ICON_LOCAL_PREFIX } = viteEnv;

  const localIconPath = path.join(process.cwd(), 'src/assets/svg-icon');

  /** The name of the local icon collection */
  const collectionName = VITE_ICON_LOCAL_PREFIX.replace(`${VITE_ICON_PREFIX}-`, '');

  return unocss({
    presets: [
      presetIcons({
        prefix: `${VITE_ICON_PREFIX}-`,
        scale: 1,
        extraProperties: {
          display: 'inline-block'
        },
        collections: {
          [collectionName]: FileSystemIconLoader(localIconPath, svg =>
            svg.replace(/^<svg\s/, '<svg width="1em" height="1em" ')
          ),
          'ant-design': antDesign,
          fluent,
          gridicons,
          ic,
          mdi,
          uil
        },
        warn: false
      })
    ]
  });
}
