# 線路相關接口

## 1. 獲取線路列表

### 接口描述

獲取域名支持的線路列表。

### 輸入參數

| 參數名稱 | 必選 | 類型 | 描述 |
|---------|------|------|------|
| Action | 是 | String | 公共參數，本接口取值：DescribeRecordLineList |
| Domain | 是 | String | 域名 |
| DomainId | 否 | Integer | 域名 ID |

### 輸出參數

| 參數名稱 | 類型 | 描述 |
|---------|------|------|
| LineList | Array of LineInfo | 線路列表 |
| LineGroupList | Array of LineGroupInfo | 線路分組列表 |
| RequestId | String | 唯一請求 ID |

---

## 2. 獲取記錄類型

### 接口描述

獲取域名支持的記錄類型。

### 輸入參數

| 參數名稱 | 必選 | 類型 | 描述 |
|---------|------|------|------|
| Action | 是 | String | 公共參數，本接口取值：DescribeRecordType |
| Domain | 否 | String | 域名 |
| DomainId | 否 | Integer | 域名 ID |

### 輸出參數

| 參數名稱 | 類型 | 描述 |
|---------|------|------|
| TypeList | Array of String | 支持的記錄類型列表 |
| RequestId | String | 唯一請求 ID |

---

## 3. 獲取域名自定義線路列表

### 接口描述

獲取域名自定義線路列表。

### 輸入參數

| 參數名稱 | 必選 | 類型 | 描述 |
|---------|------|------|------|
| Action | 是 | String | 公共參數，本接口取值：DescribeDomainCustomLineList |
| Domain | 是 | String | 域名 |
| DomainId | 否 | Integer | 域名 ID |

### 輸出參數

| 參數名稱 | 類型 | 描述 |
|---------|------|------|
| LineList | Array of CustomLineInfo | 自定義線路列表 |
| RequestId | String | 唯一請求 ID |

---

## 4. 創建域名的自定義線路

### 接口描述

創建域名的自定義線路。

### 輸入參數

| 參數名稱 | 必選 | 類型 | 描述 |
|---------|------|------|------|
| Action | 是 | String | 公共參數，本接口取值：CreateDomainCustomLine |
| Domain | 是 | String | 域名 |
| Name | 是 | String | 自定義線路名稱 |
| Area | 是 | String | 自定義線路 IP 段，用 "-" 連接 |
| DomainId | 否 | Integer | 域名 ID |

### 輸出參數

| 參數名稱 | 類型 | 描述 |
|---------|------|------|
| RequestId | String | 唯一請求 ID |

---

## 5. 刪除域名的自定義線路

### 接口描述

刪除域名的自定義線路。

### 輸入參數

| 參數名稱 | 必選 | 類型 | 描述 |
|---------|------|------|------|
| Action | 是 | String | 公共參數，本接口取值：DeleteDomainCustomLine |
| Domain | 是 | String | 域名 |
| Name | 是 | String | 自定義線路名稱 |
| DomainId | 否 | Integer | 域名 ID |

### 輸出參數

| 參數名稱 | 類型 | 描述 |
|---------|------|------|
| RequestId | String | 唯一請求 ID |

---

## 6. 修改域名的自定義線路

### 接口描述

修改域名的自定義線路。

### 輸入參數

| 參數名稱 | 必選 | 類型 | 描述 |
|---------|------|------|------|
| Action | 是 | String | 公共參數，本接口取值：ModifyDomainCustomLine |
| Domain | 是 | String | 域名 |
| Name | 是 | String | 自定義線路名稱 |
| Area | 是 | String | 自定義線路 IP 段 |
| DomainId | 否 | Integer | 域名 ID |

### 輸出參數

| 參數名稱 | 類型 | 描述 |
|---------|------|------|
| RequestId | String | 唯一請求 ID |

---

## 相關數據結構

### LineInfo

線路信息

| 名稱 | 類型 | 描述 |
|-----|------|------|
| Name | String | 線路名稱 |
| LineId | String | 線路 ID |

### LineGroupInfo

線路分組信息

| 名稱 | 類型 | 描述 |
|-----|------|------|
| GroupId | Integer | 線路分組 ID |
| GroupName | String | 線路分組名稱 |
| LineList | Array of LineInfo | 線路列表 |

### CustomLineInfo

自定義線路詳情

| 名稱 | 類型 | 描述 |
|-----|------|------|
| DomainId | Integer | 域名 ID |
| Name | String | 自定義線路名稱 |
| Area | String | 自定義線路 IP 段 |
| UseCount | Integer | 已使用 IP 段個數 |
| MaxCount | Integer | 允許使用 IP 段最大個數 |
