# i18n Coverage Report

Base locale: `en-US`

Metrics:
- `Translated` = leaf value differs from `en-US` and does not match suspicious placeholder patterns
- `Suspect` = leaf differs from `en-US` but looks corrupted (e.g. many `?` placeholders)
- `Fallback-identical` = leaf exists but equals `en-US` (usually untranslated fallback)
- `Missing` = leaf key not present in locale object

| Locale | Base Keys | Translated | Suspect | Fallback | Missing | Extra | Translated/Base | Translated/(Base-Missing) |
| --- | ---: | ---: | ---: | ---: | ---: | ---: | ---: | ---: |
| de-DE | 392 | 344 | 0 | 48 | 0 | 0 | 87.76% | 87.76% |
| es-ES | 392 | 356 | 0 | 36 | 0 | 0 | 90.82% | 90.82% |
| fr-FR | 392 | 350 | 0 | 42 | 0 | 0 | 89.29% | 89.29% |
| ja-JP | 392 | 176 | 0 | 216 | 0 | 0 | 44.90% | 44.90% |
| ko-KR | 392 | 176 | 0 | 216 | 0 | 0 | 44.90% | 44.90% |
| ru-RU | 392 | 224 | 0 | 168 | 0 | 0 | 57.14% | 57.14% |
| zh-CN | 392 | 385 | 0 | 7 | 0 | 0 | 98.21% | 98.21% |

## de-DE

- Base leaf keys: 392
- Translated leaves: 344 (87.76%)
- Suspect translated leaves: 0
- Fallback-identical leaves: 48 (12.24%)
- Missing leaves: 0
- Extra leaves: 0

**Sample Fallback Keys**
  - `common.index`
  - `dataMap.ab.alias`
  - `dataMap.ab.guid`
  - `dataMap.ab.hash`
  - `dataMap.ab.hostname`
  - `dataMap.ab.rustdesk_id`
  - `dataMap.ab.tagName`
  - `dataMap.audit.ip`
  - `dataMap.audit.rustdesk_id`
  - `dataMap.audit.uuid`
  - `dataMap.device.os`
  - `dataMap.device.rustdesk_id`
  - `dataMap.mailLog.status`
  - `dataMap.mailLog.uuid`
  - `dataMap.mailTemplate.name`
  - `dataMap.oauth.name`
  - `dataMap.token.is_admin`
  - `dataMap.user.loginVerifyLabel.tfaCheck`
  - `dataMap.user.status`
  - `dataMap.user.statusLabel.normal`
  - `page.home.serverConfig.ageSeconds`
  - `page.home.serverConfig.cacheTtlHint`
  - `page.home.serverConfig.connectivity.checkSourceType.cache`
  - `page.home.serverConfig.connectivity.checkSourceType.remote`
  - `page.home.serverConfig.key`
  - `page.home.serverConfig.sourceType.remote`
  - `page.myDevices.offline`
  - `page.myDevices.online`
  - `page.myDevices.status`
  - `page.user.list.emailFormatError`
  - `page.user.list.inputNickname`
  - `page.user.list.inputPassword`
  - `page.user.list.inputUsername`
  - `page.user.list.require2FACode`
  - `page.user.list.require2FASecret`
  - `page.user.list.selectUserStatus`
  - `page.user.list.tfa_secret_bind`
  - `route.audit`
  - `route.iframe-page`
  - `system.title`
  - ... and 8 more

**Suspect Keys**
-

**Missing Keys**
-

**Extra Keys**
-

## es-ES

- Base leaf keys: 392
- Translated leaves: 356 (90.82%)
- Suspect translated leaves: 0
- Fallback-identical leaves: 36 (9.18%)
- Missing leaves: 0
- Extra leaves: 0

**Sample Fallback Keys**
  - `common.error`
  - `common.yesOrNo.no`
  - `dataMap.ab.alias`
  - `dataMap.ab.guid`
  - `dataMap.ab.hash`
  - `dataMap.ab.rustdesk_id`
  - `dataMap.ab.tagColor`
  - `dataMap.audit.ip`
  - `dataMap.audit.rustdesk_id`
  - `dataMap.audit.uuid`
  - `dataMap.device.rustdesk_id`
  - `dataMap.mailLog.statusLabel.err`
  - `dataMap.mailLog.uuid`
  - `dataMap.token.is_admin`
  - `dataMap.user.loginVerifyLabel.tfaCheck`
  - `dataMap.user.statusLabel.normal`
  - `page.home.serverConfig.ageSeconds`
  - `page.home.serverConfig.cacheTtlHint`
  - `page.home.serverConfig.key`
  - `page.user.list.emailFormatError`
  - `page.user.list.inputNickname`
  - `page.user.list.inputPassword`
  - `page.user.list.inputUsername`
  - `page.user.list.require2FACode`
  - `page.user.list.require2FASecret`
  - `page.user.list.selectUserStatus`
  - `page.user.list.tfa_secret_bind`
  - `route.iframe-page`
  - `system.title`
  - `system.updateCancel`
  - `system.updateConfirm`
  - `system.updateContent`
  - `system.updateTitle`
  - `theme.tab.mode.chrome`
  - `theme.themeColor.error`
  - `theme.themeColor.info`

**Suspect Keys**
-

**Missing Keys**
-

**Extra Keys**
-

## fr-FR

- Base leaf keys: 392
- Translated leaves: 350 (89.29%)
- Suspect translated leaves: 0
- Fallback-identical leaves: 42 (10.71%)
- Missing leaves: 0
- Extra leaves: 0

**Sample Fallback Keys**
  - `common.action`
  - `common.index`
  - `dataMap.ab.alias`
  - `dataMap.ab.guid`
  - `dataMap.ab.hash`
  - `dataMap.audit.ip`
  - `dataMap.audit.rustdesk_id`
  - `dataMap.audit.type`
  - `dataMap.audit.uuid`
  - `dataMap.device.os`
  - `dataMap.device.rustdesk_id`
  - `dataMap.mailLog.uuid`
  - `dataMap.mailTemplate.type`
  - `dataMap.token.is_admin`
  - `dataMap.user.loginVerifyLabel.tfaCheck`
  - `dataMap.user.statusLabel.normal`
  - `page.home.serverConfig.ageSeconds`
  - `page.home.serverConfig.cacheTtlHint`
  - `page.home.serverConfig.connectivity.checkSourceType.cache`
  - `page.home.serverConfig.key`
  - `page.home.serverConfig.source`
  - `page.home.serverConfig.sourceType.env`
  - `page.myDevices.title`
  - `page.user.list.emailFormatError`
  - `page.user.list.inputNickname`
  - `page.user.list.inputPassword`
  - `page.user.list.inputUsername`
  - `page.user.list.require2FACode`
  - `page.user.list.require2FASecret`
  - `page.user.list.selectUserStatus`
  - `page.user.list.tfa_secret_bind`
  - `route.audit`
  - `route.iframe-page`
  - `route.my-devices`
  - `route.user_sessions`
  - `system.title`
  - `system.updateCancel`
  - `system.updateConfirm`
  - `system.updateContent`
  - `system.updateTitle`
  - ... and 2 more

**Suspect Keys**
-

**Missing Keys**
-

**Extra Keys**
-

## ja-JP

- Base leaf keys: 392
- Translated leaves: 176 (44.90%)
- Suspect translated leaves: 0
- Fallback-identical leaves: 216 (55.10%)
- Missing leaves: 0
- Extra leaves: 0

**Sample Fallback Keys**
  - `common.action`
  - `common.add`
  - `common.addSuccess`
  - `common.backToHome`
  - `common.batchDelete`
  - `common.cancel`
  - `common.check`
  - `common.close`
  - `common.columnSetting`
  - `common.config`
  - `common.confirm`
  - `common.confirmDelete`
  - `common.delete`
  - `common.deleteSuccess`
  - `common.edit`
  - `common.error`
  - `common.expandColumn`
  - `common.index`
  - `common.keywordSearch`
  - `common.logout`
  - `common.logoutConfirm`
  - `common.look`
  - `common.lookForward`
  - `common.modify`
  - `common.modifySuccess`
  - `common.noData`
  - `common.operate`
  - `common.pleaseCheckValue`
  - `common.refresh`
  - `common.reset`
  - `common.search`
  - `common.switch`
  - `common.tip`
  - `common.trigger`
  - `common.userCenter`
  - `common.warning`
  - `common.yesOrNo.no`
  - `common.yesOrNo.yes`
  - `dataMap.ab.ab_id`
  - `dataMap.ab.alias`
  - ... and 176 more

**Suspect Keys**
-

**Missing Keys**
-

**Extra Keys**
-

## ko-KR

- Base leaf keys: 392
- Translated leaves: 176 (44.90%)
- Suspect translated leaves: 0
- Fallback-identical leaves: 216 (55.10%)
- Missing leaves: 0
- Extra leaves: 0

**Sample Fallback Keys**
  - `common.action`
  - `common.add`
  - `common.addSuccess`
  - `common.backToHome`
  - `common.batchDelete`
  - `common.cancel`
  - `common.check`
  - `common.close`
  - `common.columnSetting`
  - `common.config`
  - `common.confirm`
  - `common.confirmDelete`
  - `common.delete`
  - `common.deleteSuccess`
  - `common.edit`
  - `common.error`
  - `common.expandColumn`
  - `common.index`
  - `common.keywordSearch`
  - `common.logout`
  - `common.logoutConfirm`
  - `common.look`
  - `common.lookForward`
  - `common.modify`
  - `common.modifySuccess`
  - `common.noData`
  - `common.operate`
  - `common.pleaseCheckValue`
  - `common.refresh`
  - `common.reset`
  - `common.search`
  - `common.switch`
  - `common.tip`
  - `common.trigger`
  - `common.userCenter`
  - `common.warning`
  - `common.yesOrNo.no`
  - `common.yesOrNo.yes`
  - `dataMap.ab.ab_id`
  - `dataMap.ab.alias`
  - ... and 176 more

**Suspect Keys**
-

**Missing Keys**
-

**Extra Keys**
-

## ru-RU

- Base leaf keys: 392
- Translated leaves: 224 (57.14%)
- Suspect translated leaves: 0
- Fallback-identical leaves: 168 (42.86%)
- Missing leaves: 0
- Extra leaves: 0

**Sample Fallback Keys**
  - `common.action`
  - `common.add`
  - `common.addSuccess`
  - `common.backToHome`
  - `common.batchDelete`
  - `common.cancel`
  - `common.check`
  - `common.close`
  - `common.columnSetting`
  - `common.config`
  - `common.confirm`
  - `common.confirmDelete`
  - `common.delete`
  - `common.deleteSuccess`
  - `common.edit`
  - `common.error`
  - `common.expandColumn`
  - `common.index`
  - `common.keywordSearch`
  - `common.logout`
  - `common.logoutConfirm`
  - `common.look`
  - `common.lookForward`
  - `common.modify`
  - `common.modifySuccess`
  - `common.noData`
  - `common.operate`
  - `common.pleaseCheckValue`
  - `common.refresh`
  - `common.reset`
  - `common.search`
  - `common.switch`
  - `common.tip`
  - `common.trigger`
  - `common.userCenter`
  - `common.warning`
  - `common.yesOrNo.no`
  - `common.yesOrNo.yes`
  - `dataMap.ab.ab_id`
  - `dataMap.ab.alias`
  - ... and 128 more

**Suspect Keys**
-

**Missing Keys**
-

**Extra Keys**
-

## zh-CN

- Base leaf keys: 392
- Translated leaves: 385 (98.21%)
- Suspect translated leaves: 0
- Fallback-identical leaves: 7 (1.79%)
- Missing leaves: 0
- Extra leaves: 0

**Sample Fallback Keys**
  - `dataMap.ab.guid`
  - `dataMap.audit.ip`
  - `dataMap.audit.uuid`
  - `dataMap.mailLog.uuid`
  - `page.home.serverConfig.key`
  - `system.title`
  - `theme.tab.mode.chrome`

**Suspect Keys**
-

**Missing Keys**
-

**Extra Keys**
-
