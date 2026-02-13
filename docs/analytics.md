# 解析量相關接口

## 1. 域名解析量統計

### 接口描述

查詢域名的解析量統計數據，支持查詢近 3 個月的數據。

**默認接口請求頻率限制：** 20次/秒

### 輸入參數

| 參數名稱 | 必選 | 類型 | 描述 |
|---------|------|------|------|
| Action | 是 | String | 公共參數，本接口取值：DescribeDomainAnalytics |
| Domain | 是 | String | 要查詢解析量的域名 |
| StartDate | 是 | String | 查詢開始時間，格式：YYYY-MM-DD |
| EndDate | 是 | String | 查詢結束時間，格式：YYYY-MM-DD |
| DnsFormat | 否 | String | DATE：按天維度統計；HOUR：按小時維度統計 |
| DomainId | 否 | Integer | 域名 ID |

### 輸出參數

| 參數名稱 | 類型 | 描述 |
|---------|------|------|
| Data | Array of DomainAnalyticsDetail | 當前統計維度解析量小計 |
| Info | DomainAnalyticsInfo | 域名解析量統計查詢信息 |
| AliasData | Array of DomainAliasAnalyticsItem | 域名別名解析量統計信息 |
| RequestId | String | 唯一請求 ID |

### 示例

**輸出示例：**

```json
{
    "Response": {
        "Data": [
            {
                "DateKey": "20221201",
                "HourKey": 0,
                "Num": 100
            }
        ],
        "Info": {
            "DnsFormat": "HOUR",
            "DnsTotal": 8358,
            "Domain": "example.com",
            "StartDate": "2022-12-01",
            "EndDate": "2022-12-01"
        },
        "AliasData": [],
        "RequestId": "e0815831-4978-445e-b524-746df0069d39"
    }
}
```

---

## 2. 子域名解析量統計

### 接口描述

查詢子域名的解析量統計數據。

### 輸入參數

| 參數名稱 | 必選 | 類型 | 描述 |
|---------|------|------|------|
| Action | 是 | String | 公共參數，本接口取值：DescribeSubdomainAnalytics |
| Domain | 是 | String | 域名 |
| Subdomain | 是 | String | 子域名 |
| StartDate | 是 | String | 查詢開始時間，格式：YYYY-MM-DD |
| EndDate | 是 | String | 查詢結束時間，格式：YYYY-MM-DD |
| DnsFormat | 否 | String | DATE：按天維度統計；HOUR：按小時維度統計 |
| DomainId | 否 | Integer | 域名 ID |

### 輸出參數

| 參數名稱 | 類型 | 描述 |
|---------|------|------|
| Data | Array of DomainAnalyticsDetail | 當前統計維度解析量小計 |
| Info | SubdomainAnalyticsInfo | 子域名解析量統計查詢信息 |
| RequestId | String | 唯一請求 ID |

---

## 相關數據結構

### DomainAnalyticsDetail

當前統計維度解析量小計

| 名稱 | 類型 | 描述 |
|-----|------|------|
| Num | Integer | 當前統計維度解析量小計 |
| DateKey | String | 按天統計時，為統計日期 |
| HourKey | Integer | 按小時統計時，為統計的當前時間的小時數 (0-23) |

### DomainAnalyticsInfo

域名解析量統計查詢信息

| 名稱 | 類型 | 描述 |
|-----|------|------|
| DnsFormat | String | DATE：按天維度統計；HOUR：按小時維度統計 |
| DnsTotal | Integer | 當前統計周期解析量總計 |
| Domain | String | 當前查詢的域名 |
| StartDate | String | 當前統計周期開始時間 |
| EndDate | String | 當前統計周期結束時間 |

### SubdomainAnalyticsInfo

子域名解析量統計查詢信息

| 名稱 | 類型 | 描述 |
|-----|------|------|
| DnsFormat | String | DATE：按天維度統計；HOUR：按小時維度統計 |
| DnsTotal | Integer | 當前統計周期解析量總計 |
| Domain | String | 當前查詢的域名 |
| Subdomain | String | 當前查詢的子域名 |
| StartDate | String | 當前統計周期開始時間 |
| EndDate | String | 當前統計周期結束時間 |

### DomainAliasAnalyticsItem

域名別名解析量統計信息

| 名稱 | 類型 | 描述 |
|-----|------|------|
| Info | DomainAnalyticsInfo | 域名解析量統計查詢信息 |
| Data | Array of DomainAnalyticsDetail | 當前統計維度解析量小計 |
