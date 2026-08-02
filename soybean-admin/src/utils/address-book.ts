import { $t } from '@/locales';

/** Translate only server-owned built-in names; user-defined address-book names remain unchanged. */
export function localizeAddressBookName(name?: string) {
  const value = (name || '').trim();
  switch (value.toLowerCase()) {
    case 'my address book':
    case 'personal address book':
      return $t('dataMap.ab.personal');
    case 'legacy address book':
      return $t('dataMap.ab.legacy');
    default:
      return value;
  }
}

export function normalizeAddressBookFilter(value?: string) {
  const keyword = (value || '').trim();
  if (!keyword) return '';
  if ($t('dataMap.ab.personal').toLowerCase().includes(keyword.toLowerCase())) return 'My address book';
  if ($t('dataMap.ab.legacy').toLowerCase().includes(keyword.toLowerCase())) return 'Legacy address book';
  return keyword;
}
