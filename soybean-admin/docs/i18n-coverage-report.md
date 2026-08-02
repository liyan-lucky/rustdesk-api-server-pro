# i18n Coverage Report

Base locale: `en-US`

Metrics:
- `Translated` = leaf value differs from `en-US` and does not match suspicious placeholder patterns
- `Suspect` = leaf differs from `en-US` but looks corrupted (e.g. many `?` placeholders)
- `Fallback-identical` = leaf exists but equals `en-US` (usually untranslated fallback)
- `Missing` = leaf key not present in locale object

| Locale | Base Keys | Translated | Suspect | Fallback | Missing | Extra | Translated/Base | Translated/(Base-Missing) |
| --- | ---: | ---: | ---: | ---: | ---: | ---: | ---: | ---: |
| de-DE | 439 | 390 | 0 | 49 | 0 | 0 | 88.84% | 88.84% |
| es-ES | 439 | 403 | 0 | 36 | 0 | 0 | 91.80% | 91.80% |
| fr-FR | 439 | 396 | 0 | 43 | 0 | 0 | 90.21% | 90.21% |
| it-IT | 439 | 404 | 0 | 35 | 0 | 0 | 92.03% | 92.03% |
| ja-JP | 439 | 223 | 0 | 216 | 0 | 0 | 50.80% | 50.80% |
| ko-KR | 439 | 223 | 0 | 216 | 0 | 0 | 50.80% | 50.80% |
| ru-RU | 439 | 271 | 0 | 168 | 0 | 0 | 61.73% | 61.73% |
| zh-CN | 439 | 432 | 0 | 7 | 0 | 0 | 98.41% | 98.41% |

## de-DE

- Base leaf keys: 439
- Translated leaves: 390 (88.84%)
- Suspect translated leaves: 0
- Fallback-identical leaves: 49 (11.16%)
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
  - `page.workspace.adminRole`
  - `route.audit`
  - `route.iframe-page`
  - ... and 9 more

**Suspect Keys**
-

**Missing Keys**
-

**Extra Keys**
-

## es-ES

- Base leaf keys: 439
- Translated leaves: 403 (91.80%)
- Suspect translated leaves: 0
- Fallback-identical leaves: 36 (8.20%)
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

- Base leaf keys: 439
- Translated leaves: 396 (90.21%)
- Suspect translated leaves: 0
- Fallback-identical leaves: 43 (9.79%)
- Missing leaves: 0
- Extra leaves: 0

**Sample Fallback Keys**
  - `common.action`
  - `common.index`
  - `dataMap.ab.alias`
  - `dataMap.ab.guid`
  - `dataMap.ab.hash`
  - `dataMap.ab.note`
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
  - ... and 3 more

**Suspect Keys**
-

**Missing Keys**
-

**Extra Keys**
-

## it-IT

- Base leaf keys: 439
- Translated leaves: 404 (92.03%)
- Suspect translated leaves: 0
- Fallback-identical leaves: 35 (7.97%)
- Missing leaves: 0
- Extra leaves: 0

**Sample Fallback Keys**
  - `common.yesOrNo.no`
  - `dataMap.ab.ab_id`
  - `dataMap.ab.alias`
  - `dataMap.ab.guid`
  - `dataMap.ab.hash`
  - `dataMap.ab.hostname`
  - `dataMap.ab.max_peer`
  - `dataMap.ab.name`
  - `dataMap.ab.owner`
  - `dataMap.ab.rule`
  - `dataMap.ab.rustdesk_id`
  - `dataMap.ab.shared`
  - `dataMap.ab.tagColor`
  - `dataMap.ab.tagName`
  - `dataMap.ab.tags`
  - `dataMap.ab.updated_at`
  - `dataMap.ab.user_id`
  - `dataMap.ab.username`
  - `dataMap.audit.ip`
  - `dataMap.audit.uuid`
  - `dataMap.mailLog.uuid`
  - `dataMap.oauth.email`
  - `dataMap.oauth.provider`
  - `dataMap.token.is_admin`
  - `dataMap.user.email`
  - `dataMap.user.loginVerifyLabel.tfaCheck`
  - `dataMap.user.name`
  - `dataMap.user.password`
  - `page.home.serverConfig.connectivity.checkSourceType.cache`
  - `page.home.serverConfig.key`
  - `page.myDevices.offline`
  - `page.myDevices.online`
  - `route.audit`
  - `route.home`
  - `route.iframe-page`

**Suspect Keys**
-

**Missing Keys**
-

**Extra Keys**
-

## ja-JP

- Base leaf keys: 439
- Translated leaves: 223 (50.80%)
- Suspect translated leaves: 0
- Fallback-identical leaves: 216 (49.20%)
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

- Base leaf keys: 439
- Translated leaves: 223 (50.80%)
- Suspect translated leaves: 0
- Fallback-identical leaves: 216 (49.20%)
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

- Base leaf keys: 439
- Translated leaves: 271 (61.73%)
- Suspect translated leaves: 0
- Fallback-identical leaves: 168 (38.27%)
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

- Base leaf keys: 439
- Translated leaves: 432 (98.41%)
- Suspect translated leaves: 0
- Fallback-identical leaves: 7 (1.59%)
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
