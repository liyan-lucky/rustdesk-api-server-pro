# i18n Coverage Report

Base locale: `en-US`

Metrics:
- `Translated` = leaf value differs from `en-US` and does not match suspicious placeholder patterns
- `Suspect` = leaf differs from `en-US` but looks corrupted (e.g. many `?` placeholders)
- `Fallback-identical` = leaf exists but equals `en-US` (usually untranslated fallback)
- `Missing` = leaf key not present in locale object

| Locale | Base Keys | Translated | Suspect | Fallback | Missing | Extra | Translated/Base | Translated/(Base-Missing) |
| --- | ---: | ---: | ---: | ---: | ---: | ---: | ---: | ---: |
| de-DE | 320 | 280 | 0 | 40 | 0 | 0 | 87.50% | 87.50% |
| es-ES | 320 | 287 | 0 | 33 | 0 | 0 | 89.69% | 89.69% |
| fr-FR | 320 | 283 | 0 | 37 | 0 | 0 | 88.44% | 88.44% |
| ja-JP | 320 | 95 | 0 | 225 | 0 | 0 | 29.69% | 29.69% |
| ko-KR | 320 | 95 | 0 | 225 | 0 | 0 | 29.69% | 29.69% |
| ru-RU | 320 | 145 | 0 | 175 | 0 | 0 | 45.31% | 45.31% |
| zh-CN | 320 | 305 | 0 | 15 | 0 | 0 | 95.31% | 95.31% |

## de-DE

- Base leaf keys: 320
- Translated leaves: 280 (87.50%)
- Suspect translated leaves: 0
- Fallback-identical leaves: 40 (12.50%)
- Missing leaves: 0
- Extra leaves: 0

**Sample Fallback Keys**
  - `api.MailTemplateUpdateSuccess`
  - `api.UserUpdateSuccess`
  - `common.index`
  - `common.update`
  - `common.updateSuccess`
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
  - `page.home.serverConfig.lastUpdated`
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
- Translated leaves: 287 (89.69%)
- Suspect translated leaves: 0
- Fallback-identical leaves: 33 (10.31%)
- Missing leaves: 0
- Extra leaves: 0

**Sample Fallback Keys**
  - `api.MailTemplateUpdateSuccess`
  - `api.UserUpdateSuccess`
  - `common.error`
  - `common.update`
  - `common.updateSuccess`
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
  - `page.home.serverConfig.lastUpdated`
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
- Translated leaves: 283 (88.44%)
- Suspect translated leaves: 0
- Fallback-identical leaves: 37 (11.56%)
- Missing leaves: 0
- Extra leaves: 0

**Sample Fallback Keys**
  - `common.action`
  - `common.index`
  - `common.update`
  - `common.updateSuccess`
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
  - `page.home.serverConfig.lastUpdated`
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
-

**Missing Keys**
-

**Extra Keys**
-

## ja-JP

- Base leaf keys: 320
- Translated leaves: 95 (29.69%)
- Suspect translated leaves: 0
- Fallback-identical leaves: 225 (70.31%)
- Missing leaves: 0
- Extra leaves: 0

**Sample Fallback Keys**
  - `api.MailTemplateUpdateSuccess`
  - `api.UserUpdateSuccess`
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
  - ... and 185 more

**Suspect Keys**
-

**Missing Keys**
-

**Extra Keys**
-

## ko-KR

- Base leaf keys: 320
- Translated leaves: 95 (29.69%)
- Suspect translated leaves: 0
- Fallback-identical leaves: 225 (70.31%)
- Missing leaves: 0
- Extra leaves: 0

**Sample Fallback Keys**
  - `api.MailTemplateUpdateSuccess`
  - `api.UserUpdateSuccess`
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
  - ... and 185 more

**Suspect Keys**
-

**Missing Keys**
-

**Extra Keys**
-

## ru-RU

- Base leaf keys: 320
- Translated leaves: 145 (45.31%)
- Suspect translated leaves: 0
- Fallback-identical leaves: 175 (54.69%)
- Missing leaves: 0
- Extra leaves: 0

**Sample Fallback Keys**
  - `api.MailTemplateUpdateSuccess`
  - `api.UserUpdateSuccess`
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
  - ... and 135 more

**Suspect Keys**
-

**Missing Keys**
-

**Extra Keys**
-

## zh-CN

- Base leaf keys: 320
- Translated leaves: 305 (95.31%)
- Suspect translated leaves: 0
- Fallback-identical leaves: 15 (4.69%)
- Missing leaves: 0
- Extra leaves: 0

**Sample Fallback Keys**
  - `api.MailTemplateUpdateSuccess`
  - `api.UserUpdateSuccess`
  - `common.update`
  - `common.updateSuccess`
  - `dataMap.audit.ip`
  - `dataMap.audit.uuid`
  - `dataMap.mailLog.uuid`
  - `page.home.serverConfig.key`
  - `page.home.serverConfig.lastUpdated`
  - `system.title`
  - `system.updateCancel`
  - `system.updateConfirm`
  - `system.updateContent`
  - `system.updateTitle`
  - `theme.tab.mode.chrome`

**Suspect Keys**
-

**Missing Keys**
-

**Extra Keys**
-
