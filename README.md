# DNSPod.cn Go SDK

[![Build Status](https://img.shields.io/travis/com/MMHK/dnspodcn?style=for-the-badge)](https://travis-ci.com/MMHK/dnspodcn)
[![Go.Dev reference](https://img.shields.io/badge/go.dev-reference-blue?logo=go&logoColor=white&style=for-the-badge)](https://pkg.go.dev/github.com/MMHK/dnspodcn)
[![Go Report Card](https://goreportcard.com/badge/github.com/MMHK/dnspodcn?style=for-the-badge)](https://goreportcard.com/report/github.com/MMHK/dnspodcn)
[![Release](https://img.shields.io/github/release/MMHK/dnspodcn.svg?style=for-the-badge)](https://github.com/MMHK/dnspodcn/releases)

`dnspodcn` 是 [DNSPod](https://cloud.tencent.com/document/api/1427/56194) API 3.0 的 Go SDK。其中 `Provider` 实现了 [libdns](https://github.com/libdns/libdns/) 接口。

## 安裝

```shell
go get -u github.com/MMHK/dnspodcn
```

## 使用

```go
import "github.com/MMHK/dnspodcn"

client := dnspodcn.NewClient("SECRET_ID", "SECRET_KEY")
```

> `SECRET_ID` 和 `SECRET_KEY` 是從騰訊雲控制台獲取的訪問密鑰（CAM 密鑰）。

### 配置區域（可選）

```go
client.Region = "ap-guangzhou" // 可選：ap-guangzhou, ap-shanghai, ap-beijing
```

## 認證方式

本 SDK 使用騰訊雲 API 3.0 的 **TC3-HMAC-SHA256** 簽名方法進行身份驗證。詳情請參考 [騰訊雲 API 文檔](https://cloud.tencent.com/document/api/1427/56189)。

## 接口

### DNS 記錄相關

| 功能 | 方法 | 對應 API |
|------|------|----------|
| 獲取解析記錄列表 | `client.DescribeRecordList` | DescribeRecordList |
| 添加解析記錄 | `client.CreateRecord` | CreateRecord |
| 修改解析記錄 | `client.UpdateRecord` | ModifyRecord |
| 刪除解析記錄 | `client.DeleteRecord` | DeleteRecord |

## 數據結構

### RecordListItem - 解析記錄列表項

| 欄位 | 類型 | 描述 |
|------|------|------|
| RecordID | string | 記錄 ID |
| Value | string | 記錄值 |
| Status | string | 記錄狀態：ENABLE、DISABLE |
| UpdatedOn | string | 更新時間 |
| Name | string | 主機記錄 |
| Line | string | 線路 |
| LineID | string | 線路 ID |
| Type | string | 記錄類型 |
| Weight | int | 權重 |
| MonitorStatus | string | 監控狀態 |
| Remark | string | 備註 |
| TTL | int | TTL 值 |
| MX | int | MX 優先級 |
| DefaultNS | bool | 是否為默認 NS 記錄 |
| GroupID | int | 分組 ID |
| Enabled | int | 是否啟用 |

### RecordCountInfo - 記錄數量統計

| 欄位 | 類型 | 描述 |
|------|------|------|
| TotalCount | int | 符合條件的記錄總數 |
| ListCount | int | 本次返回的記錄數量 |
| SubdomainCount | int | 符合條件的子域名數量 |

## 貢獻

如果你有其他接口需求或發現問題，可以通過 [Issue](https://github.com/MMHK/dnspodcn/issues/new) 反饋，也歡迎提交 Pull Request。

## 許可證

[MIT](LICENSE)
