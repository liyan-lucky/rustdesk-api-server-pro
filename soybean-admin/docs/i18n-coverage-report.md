# i18n Coverage Report

Base locale: `en-US`

Metrics:
- `Translated` = leaf value differs from `en-US` and does not match suspicious placeholder patterns
- `Suspect` = leaf differs from `en-US` but looks corrupted (e.g. many `?` placeholders)
- `Fallback-identical` = leaf exists but equals `en-US` (usually untranslated fallback)
- `Missing` = leaf key not present in locale object

| Locale | Base Keys | Translated | Suspect | Fallback | Missing | Extra | Translated/Base | Translated/(Base-Missing) |
| --- | ---: | ---: | ---: | ---: | ---: | ---: | ---: | ---: |
| de-DE | 320 | 284 | 0 | 36 | 0 | 0 | 88.75% | 88.75% |
| es-ES | 320 | 292 | 0 | 28 | 0 | 0 | 91.25% | 91.25% |
| fr-FR | 320 | 285 | 1 | 34 | 0 | 0 | 89.06% | 89.06% |
| ja-JP | 320 | 81 | 0 | 239 | 0 | 0 | 25.31% | 25.31% |
| ko-KR | 320 | 81 | 0 | 239 | 0 | 0 | 25.31% | 25.31% |
| ru-RU | 320 | 135 | 0 | 185 | 0 | 0 | 42.19% | 42.19% |
| zh-CN | 320 | 314 | 0 | 6 | 0 | 0 | 98.13% | 98.13% |

## de-DE

- Base leaf keys: 320
- Translated leaves: 284 (88.75%)
- Suspect translated leaves: 0
- Fallback-identical leaves: 36 (11.25%)
- Missing leaves: 0
- Extra leaves: 0

**Sample Fallback Keys**
  - `common.index`
  - `common.update`
  - `dataMap.audit.ip`
  - `dataMap.audit.rustdesk_id`
  - `dataMap.audit.uuid`
  - `dataMap.device.os`
  - `dataMap.device.rustdesk_id`
  - `dataMap.mailLog.status`
  - `dataMap.mailLog.uuid`
  - `dataMap.mailTemplate.name`
  - `dataMap.user.loginVerifyLabel.tfaCheck`
  - `dataMap.user.status`
  - `dataMap.user.statusLabel.normal`
  - `page.home.serverConfig.connectivity.checkSourceType.cache`
  - `page.home.serverConfig.connectivity.checkSourceType.remote`
  - `page.home.serverConfig.key`
  - `page.home.serverConfig.sourceType.remote`
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
  - `system.updateCancel`
  - `system.updateConfirm`
  - `system.updateContent`
  - `system.updateTitle`
  - `theme.scrollMode.wrapper`
  - `theme.tab.mode.button`
  - `theme.tab.mode.chrome`
  - `theme.themeColor.info`

**Suspect Keys**
-

**Missing Keys**
-

**Extra Keys**
-

## es-ES

- Base leaf keys: 320
- Translated leaves: 292 (91.25%)
- Suspect translated leaves: 0
- Fallback-identical leaves: 28 (8.75%)
- Missing leaves: 0
- Extra leaves: 0

**Sample Fallback Keys**
  - `common.error`
  - `common.yesOrNo.no`
  - `dataMap.audit.ip`
  - `dataMap.audit.rustdesk_id`
  - `dataMap.audit.uuid`
  - `dataMap.device.rustdesk_id`
  - `dataMap.mailLog.statusLabel.err`
  - `dataMap.mailLog.uuid`
  - `dataMap.user.loginVerifyLabel.tfaCheck`
  - `dataMap.user.statusLabel.normal`
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

- Base leaf keys: 320
- Translated leaves: 285 (89.06%)
- Suspect translated leaves: 1
- Fallback-identical leaves: 34 (10.63%)
- Missing leaves: 0
- Extra leaves: 0

**Sample Fallback Keys**
  - `common.action`
  - `common.index`
  - `dataMap.audit.ip`
  - `dataMap.audit.rustdesk_id`
  - `dataMap.audit.type`
  - `dataMap.audit.uuid`
  - `dataMap.device.os`
  - `dataMap.device.rustdesk_id`
  - `dataMap.mailLog.uuid`
  - `dataMap.mailTemplate.type`
  - `dataMap.user.loginVerifyLabel.tfaCheck`
  - `dataMap.user.statusLabel.normal`
  - `page.home.serverConfig.connectivity.checkSourceType.cache`
  - `page.home.serverConfig.key`
  - `page.home.serverConfig.source`
  - `page.home.serverConfig.sourceType.env`
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
  - `route.user_sessions`
  - `system.title`
  - `system.updateCancel`
  - `system.updateConfirm`
  - `system.updateContent`
  - `system.updateTitle`
  - `theme.tab.mode.chrome`
  - `theme.themeColor.info`

**Suspect Keys**
  - `dataMap.session.created_at`

**Missing Keys**
-

**Extra Keys**
-

## ja-JP

- Base leaf keys: 320
- Translated leaves: 81 (25.31%)
- Suspect translated leaves: 0
- Fallback-identical leaves: 239 (74.69%)
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
  - `common.update`
  - `common.updateSuccess`
  - `common.userCenter`
  - `common.warning`
  - `common.yesOrNo.no`
  - `common.yesOrNo.yes`
  - ... and 199 more

**Suspect Keys**
-

**Missing Keys**
-

**Extra Keys**
-

## ko-KR

- Base leaf keys: 320
- Translated leaves: 81 (25.31%)
- Suspect translated leaves: 0
- Fallback-identical leaves: 239 (74.69%)
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
  - `common.update`
  - `common.updateSuccess`
  - `common.userCenter`
  - `common.warning`
  - `common.yesOrNo.no`
  - `common.yesOrNo.yes`
  - ... and 199 more

**Suspect Keys**
-

**Missing Keys**
-

**Extra Keys**
-

## ru-RU

- Base leaf keys: 320
- Translated leaves: 135 (42.19%)
- Suspect translated leaves: 0
- Fallback-identical leaves: 185 (57.81%)
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
  - `common.update`
  - `common.updateSuccess`
  - `common.userCenter`
  - `common.warning`
  - `common.yesOrNo.no`
  - `common.yesOrNo.yes`
  - ... and 145 more

**Suspect Keys**
-

**Missing Keys**
-

**Extra Keys**
-

## zh-CN

- Base leaf keys: 320
- Translated leaves: 314 (98.13%)
- Suspect translated leaves: 0
- Fallback-identical leaves: 6 (1.88%)
- Missing leaves: 0
- Extra leaves: 0

**Sample Fallback Keys**
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
