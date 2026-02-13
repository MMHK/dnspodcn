# 賬戶相關接口

## 1. 獲取用戶信息

### 接口描述

獲取當前登錄用戶的信息。

### 輸入參數

| 參數名稱 | 必選 | 類型 | 描述 |
|---------|------|------|------|
| Action | 是 | String | 公共參數，本接口取值：DescribeUserDetail |

### 輸出參數

| 參數名稱 | 類型 | 描述 |
|---------|------|------|
| UserInfo | UserInfo | 用戶信息 |
| RequestId | String | 唯一請求 ID |

---

## 2. 修改用戶信息

### 接口描述

修改當前登錄用戶的信息。

### 輸入參數

| 參數名稱 | 必選 | 類型 | 描述 |
|---------|------|------|------|
| Action | 是 | String | 公共參數，本接口取值：ModifyUserInfo |
| Nickname | 否 | String | 用戶昵稱 |
| Phone | 否 | String | 手機號碼 |
| Email | 否 | String | 郵箱地址 |

### 輸出參數

| 參數名稱 | 類型 | 描述 |
|---------|------|------|
| RequestId | String | 唯一請求 ID |

---

## 相關數據結構

### UserInfo

用戶信息

| 名稱 | 類型 | 描述 |
|-----|------|------|
| Nickname | String | 用戶昵稱 |
| Email | String | 用戶郵箱 |
| Phone | String | 用戶手機號碼 |
| EmailVerified | String | 郵箱是否驗證：yes、no |
| PhoneVerified | String | 手機是否驗證：yes、no |
| AccountId | String | 賬戶 ID |
| RealName | String | 是否實名認證 |
