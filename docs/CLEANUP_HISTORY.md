# 清理 GitHub Actions 运行记录、Release 和标签

仓库提供手动触发 workflow：

```text
.github/workflows/cleanup-history.yml
```

该流程用于清理：

- GitHub Actions workflow runs / 运行记录；
- GitHub Releases；
- Git tags / 已发布标签。

## 安全默认值

默认配置是安全模式：

| 参数 | 默认值 | 说明 |
|---|---:|---|
| `dry_run` | `true` | 只预览，不实际删除 |
| `clean_workflow_runs` | `true` | 默认预览/清理 Actions 运行记录 |
| `clean_releases` | `false` | 默认不清理 GitHub Releases |
| `clean_tags` | `false` | 默认不清理 Git tags |
| `tag_pattern` | 空 | 空表示匹配全部 tag；填写 `v1.4.8` 或 `v` 可限制范围 |
| `confirm` | 空 | 实际删除时必须填写确认词 |

> 重点：标签默认不清理。清理标签属于危险操作，需要额外确认。

## 只预览将删除什么

进入 GitHub：

```text
Actions -> 清理运行记录和发布标签 -> Run workflow
```

保持默认：

```text
dry_run = true
clean_workflow_runs = true
clean_releases = false
clean_tags = false
```

执行后只会打印将要删除的 Actions 运行记录，不会真正删除。

## 清理 Actions 运行记录

参数：

```text
dry_run = false
clean_workflow_runs = true
clean_releases = false
clean_tags = false
confirm = CLEAN
```

说明：

- 会删除历史 workflow runs；
- 当前正在执行的清理 run 会自动跳过；
- 删除 run 会同时删除对应日志记录；
- 不会删除 Release；
- 不会删除 tag。

## 清理 Releases，但保留 tags

参数：

```text
dry_run = false
clean_workflow_runs = false
clean_releases = true
clean_tags = false
confirm = CLEAN
```

说明：

- 会删除 GitHub Releases；
- 默认不会删除 Git tags；
- 如果填写 `tag_pattern=v1.4.8`，只清理以 `v1.4.8` 开头的 Release。

## 清理 Releases 和 Tags

危险操作。只有确认不再需要这些标签时才执行。

参数：

```text
dry_run = false
clean_workflow_runs = false
clean_releases = true
clean_tags = true
confirm = CLEAN_TAGS
```

说明：

- 会删除 Release；
- 会删除 Git tags；
- `clean_tags=true` 时确认词必须是 `CLEAN_TAGS`，普通 `CLEAN` 不会通过。

## 只清理某个 tag

例如只清理 `v1.4.8`：

```text
dry_run = false
clean_workflow_runs = false
clean_releases = true
clean_tags = true
tag_pattern = v1.4.8
confirm = CLEAN_TAGS
```

## 注意事项

- 删除 workflow run 后，GitHub Actions 页面中对应历史记录和日志会消失。
- 删除 Release 不一定删除 tag，除非同时开启 `clean_tags`。
- 删除 tag 可能影响后续重复发布、版本追溯和用户下载链接。
- 建议先用 `dry_run=true` 预览一次，再执行真实删除。
- 该流程需要 `actions: write` 和 `contents: write` 权限。
