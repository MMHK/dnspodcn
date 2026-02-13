# DNSPod API 文檔

本文檔整理了 DNSPod 的 API 接口文檔，基於 [騰訊雲 DNSPod API 3.0](https://cloud.tencent.com/document/api/1427/56194)。

## API 概覽

DNSPod API 分為以下幾個類別：

| 類別 | 說明 |
|------|------|
| [域名相關接口](./domain.md) | 域名管理、域名分組、域名鎖定等 |
| [記錄相關接口](./record.md) | 解析記錄的增刪改查 |
| [批量操作相關接口](./batch.md) | 批量域名和記錄操作 |
| [線路相關接口](./line.md) | 線路查詢和自定義線路 |
| [快照相關接口](./snapshot.md) | 域名和記錄快照管理 |
| [解析量相關接口](./analytics.md) | 域名解析量統計 |
| [別名相關接口](./alias.md) | 域名別名管理 |
| [套餐及增值服務相關接口](./vip.md) | 套餐購買和升級 |
| [賬戶相關接口](./user.md) | 用戶信息獲取 |
| [數據結構](./datastructures.md) | 所有 API 使用的數據結構 |
| [錯誤碼](./errorcodes.md) | API 錯誤碼列表 |

## 通用信息

### 服務地址

| 接入地域 | 域名 |
|---------|------|
| 就近地域接入（推薦）| dnspod.tencentcloudapi.com |
| 華南地區（廣州）| dnspod.ap-guangzhou.tencentcloudapi.com |
| 華東地區（上海）| dnspod.ap-shanghai.tencentcloudapi.com |
| 華北地區（北京）| dnspod.ap-beijing.tencentcloudapi.com |

### 通信協議

所有接口均通過 HTTPS 進行通信。

### 請求方法

- POST（推薦）
- GET

### 字符編碼

均使用 `UTF-8` 編碼。

### 公共參數

| 參數名稱 | 類型 | 描述 |
|---------|------|------|
| Action | String | 接口名稱，如 `DescribeDomainList` |
| Version | String | 接口版本，固定為 `2021-03-23` |
| Region | String | 地域參數，可選 |
| Timestamp | Integer | 當前 UNIX 時間戳 |
| Nonce | Integer | 隨機正整數 |
| SecretId | String | 密鑰 ID |
| Signature | String | 請求簽名 |
| SignatureMethod | String | 簽名方法：HmacSHA256 或 TC3-HMAC-SHA256 |

## API Explorer

騰訊雲 API 平台 是綜合 API 文檔、錯誤碼、API Explorer 及 SDK 等資源的統一查詢平台，方便您從同一入口查詢及使用騰訊雲提供的所有 API 服務。

API Explorer 提供了在線調用、簽名驗證、SDK 代碼生成和快速檢索接口等能力。您可查看每次調用的請求內容和返回結果以及自動生成 SDK 調用示例。

## 頻率限制

不同接口有不同的頻率限制，具體請參考各接口文檔。一般限制維度為 `API + 接入地域 + 子賬號`。
