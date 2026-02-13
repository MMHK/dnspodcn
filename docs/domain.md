# 域名相關接口

## 1. 獲取域名列表

### 接口描述

獲取域名列表。

**接口請求域名：** dnspod.tencentcloudapi.com

**默認接口請求頻率限制：** 20次/秒

### 輸入參數

| 參數名稱 | 必選 | 類型 | 描述 |
|---------|------|------|------|
| Action | 是 | String | 公共參數，本接口取值：DescribeDomainList |
| Version | 是 | String | 公共參數，本接口取值：2021-03-23 |
| Type | 否 | String | 域名分組類型，默認為 ALL。可取值為 ALL、MINE、SHARE、ISMARK、PAUSE、VIP、RECENT、SHARE_OUT、FREE |
| Offset | 否 | Integer | 記錄開始的偏移，第一條記錄為 0，默認值為 0 |
| Limit | 否 | Integer | 要獲取的域名數量，默認值為 3000 |
| GroupId | 否 | Integer | 分組 ID |
| Keyword | 否 | String | 根據關鍵字搜索域名 |

### 輸出參數

| 參數名稱 | 類型 | 描述 |
|---------|------|------|
| DomainCountInfo | DomainCountInfo | 列表頁統計信息 |
| DomainList | Array of DomainListItem | 域名列表 |
| RequestId | String | 唯一請求 ID |

### 示例

**輸入示例：**

```json
{
    "Type": "ALL",
    "Keyword": "dnspod",
    "Limit": 1,
    "Offset": 0,
    "GroupId": 1
}
```

**輸出示例：**

```json
{
    "Response": {
        "DomainCountInfo": {
            "AllTotal": 35,
            "DomainTotal": 1,
            "MineTotal": 28,
            "ShareTotal": 7,
            "VipTotal": 4
        },
        "DomainList": [
            {
                "DomainId": 12614766,
                "Name": "dnspod.com",
                "Status": "ENABLE",
                "TTL": 600,
                "CNAMESpeedup": "DISABLE",
                "Grade": "DP_ULTRA",
                "GroupId": 1,
                "IsVip": "YES",
                "RecordCount": 0
            }
        ],
        "RequestId": "bfb3f27e-4dba-4a5c-9aff-08d1c27d1c61"
    }
}
```

---

## 2. 刪除域名

### 接口描述

刪除域名。

**默認接口請求頻率限制：** 20次/秒

### 輸入參數

| 參數名稱 | 必選 | 類型 | 描述 |
|---------|------|------|------|
| Action | 是 | String | 公共參數，本接口取值：DeleteDomain |
| Domain | 是 | String | 域名 |
| DomainId | 否 | Integer | 域名 ID，優先級比 Domain 高 |

### 輸出參數

| 參數名稱 | 類型 | 描述 |
|---------|------|------|
| RequestId | String | 唯一請求 ID |

---

## 3. 鎖定域名

### 接口描述

鎖定域名。

**默認接口請求頻率限制：** 20次/秒

### 輸入參數

| 參數名稱 | 必選 | 類型 | 描述 |
|---------|------|------|------|
| Action | 是 | String | 公共參數，本接口取值：ModifyDomainLock |
| Domain | 是 | String | 域名 |
| LockDays | 是 | Integer | 域名要鎖定的天數 |
| DomainId | 否 | Integer | 域名 ID |

### 輸出參數

| 參數名稱 | 類型 | 描述 |
|---------|------|------|
| LockInfo | LockInfo | 域名鎖定信息 |
| RequestId | String | 唯一請求 ID |

### 示例

**輸出示例：**

```json
{
    "Response": {
        "LockInfo": {
            "DomainId": 62,
            "LockCode": "M3xkbnNwb2Quc2l0ZXwxNjE3MzMwODMwfDhlMTky...",
            "LockEnd": "2021-05-02"
        },
        "RequestId": "ab4f1426-ea15-42ea-8183-dc1b44151166"
    }
}
```

---

## 4. 域名鎖定解鎖

### 接口描述

域名鎖定解鎖。

**默認接口請求頻率限制：** 20次/秒

### 輸入參數

| 參數名稱 | 必選 | 類型 | 描述 |
|---------|------|------|------|
| Action | 是 | String | 公共參數，本接口取值：ModifyDomainUnlock |
| Domain | 是 | String | 域名 |
| LockCode | 是 | String | 域名解鎖碼，鎖定的時候會返回 |
| DomainId | 否 | Integer | 域名 ID |

### 輸出參數

| 參數名稱 | 類型 | 描述 |
|---------|------|------|
| RequestId | String | 唯一請求 ID |

---

## 5. 獲取域名分組列表

### 接口描述

獲取域名分組列表。

**默認接口請求頻率限制：** 20次/秒

### 輸入參數

| 參數名稱 | 必選 | 類型 | 描述 |
|---------|------|------|------|
| Action | 是 | String | 公共參數，本接口取值：DescribeDomainGroupList |

### 輸出參數

| 參數名稱 | 類型 | 描述 |
|---------|------|------|
| GroupList | Array of GroupInfo | 分組列表 |
| RequestId | String | 唯一請求 ID |

### 示例

**輸出示例：**

```json
{
    "Response": {
        "GroupList": [
            {
                "GroupId": 1,
                "GroupName": "默認分組",
                "GroupType": "system",
                "Size": 3
            }
        ],
        "RequestId": "ab4f1426-ea15-42ea-8183-dc1b44151166"
    }
}
```

---

## 相關數據結構

### DomainCountInfo

列表頁分頁統計信息

| 名稱 | 類型 | 描述 |
|-----|------|------|
| DomainTotal | Integer | 符合條件的域名數量 |
| AllTotal | Integer | 用戶可以查看的所有域名數量 |
| MineTotal | Integer | 用戶賬號添加的域名數量 |
| ShareTotal | Integer | 共享給用戶的域名數量 |
| VipTotal | Integer | 付費域名數量 |
| PauseTotal | Integer | 暫停域名數量 |

### DomainListItem

域名列表元素

| 名稱 | 類型 | 描述 |
|-----|------|------|
| DomainId | Integer | 域名 ID |
| Name | String | 域名 |
| Punycode | String | 域名的 punycode 碼 |
| Grade | String | 域名套餐等級 |
| GradeLevel | Integer | 域名套餐等級對應的序號 |
| GradeTitle | String | 域名套餐名稱 |
| Status | String | 域名狀態：ENABLE（啟用）、PAUSE（暫停） |
| TTL | Integer | 域名的 TTL |
| CNAMESpeedup | String | CNAME 加速狀態：ENABLE（啟用）、DISABLE（禁用） |
| DNSStatus | String | DNS 設置狀態 |
| IsVip | String | 是否 VIP 域名：YES、NO |
| RecordCount | Integer | 解析記錄數量 |
| CreatedOn | String | 域名添加時間 |
| UpdatedOn | String | 域名最近修改時間 |
| Owner | String | 域名所有者 |
| GroupId | Integer | 域名所屬分組 ID |

### GroupInfo

分組信息

| 名稱 | 類型 | 描述 |
|-----|------|------|
| GroupId | Integer | 分組 ID |
| GroupName | String | 分組名稱 |
| GroupType | String | 分組類型：system（系統分組）、user（用戶分組） |
| Size | Integer | 該分組內的域名數量 |

### LockInfo

域名鎖定信息

| 名稱 | 類型 | 描述 |
|-----|------|------|
| DomainId | Integer | 域名 ID |
| LockCode | String | 域名解鎖碼 |
| LockEnd | String | 域名鎖定結束時間 |
