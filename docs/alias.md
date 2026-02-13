# 別名相關接口

## 1. 獲取域名別名列表

### 接口描述

獲取域名別名列表。

### 輸入參數

| 參數名稱 | 必選 | 類型 | 描述 |
|---------|------|------|------|
| Action | 是 | String | 公共參數，本接口取值：DescribeDomainAliasList |
| Domain | 是 | String | 域名 |
| DomainId | 否 | Integer | 域名 ID |

### 輸出參數

| 參數名稱 | 類型 | 描述 |
|---------|------|------|
| AliasList | Array of DomainAliasInfo | 域名別名列表 |
| RequestId | String | 唯一請求 ID |

---

## 2. 創建域名別名

### 接口描述

創建域名別名。

### 輸入參數

| 參數名稱 | 必選 | 類型 | 描述 |
|---------|------|------|------|
| Action | 是 | String | 公共參數，本接口取值：CreateDomainAlias |
| DomainAlias | 是 | String | 域名別名 |
| Domain | 否 | String | 域名 |
| DomainId | 否 | Integer | 域名 ID |

### 輸出參數

| 參數名稱 | 類型 | 描述 |
|---------|------|------|
| RequestId | String | 唯一請求 ID |

---

## 3. 刪除域名別名

### 接口描述

刪除域名別名。

### 輸入參數

| 參數名稱 | 必選 | 類型 | 描述 |
|---------|------|------|------|
| Action | 是 | String | 公共參數，本接口取值：DeleteDomainAlias |
| DomainAliasId | 是 | Integer | 域名別名 ID |

### 輸出參數

| 參數名稱 | 類型 | 描述 |
|---------|------|------|
| RequestId | String | 唯一請求 ID |

---

## 相關數據結構

### DomainAliasInfo

域名別名信息

| 名稱 | 類型 | 描述 |
|-----|------|------|
| Id | Integer | 域名別名 ID |
| DomainAlias | String | 域名別名 |
| Status | Integer | 別名狀態：1-DNS 不正確；2-正常；3-封禁 |
