# 記錄相關接口

## 1. 獲取域名的解析記錄列表

### 接口描述

獲取某個域名下的解析記錄列表。

**備註：**
1. 新添加的解析記錄存在短暫的索引延遲，如果查詢不到新增記錄，請在 30 秒後重試
2. API 獲取的記錄總條數會比控制台多 2 條，原因是：為了防止用戶誤操作導致解析服務不可用，對 2021-10-29 14:24:26 之後添加的域名，在控制台都不顯示這 2 條 NS 記錄。

**默認接口請求頻率限制：** 100次/秒

### 輸入參數

| 參數名稱 | 必選 | 類型 | 描述 |
|---------|------|------|------|
| Action | 是 | String | 公共參數，本接口取值：DescribeRecordList |
| Domain | 是 | String | 域名 |
| DomainId | 否 | Integer | 域名 ID，優先級比 Domain 高 |
| Subdomain | 否 | String | 解析記錄的主機頭 |
| RecordType | 否 | String | 獲取某種類型的解析記錄，如 A、CNAME、NS、AAAA 等 |
| RecordLine | 否 | String | 獲取某條線路名稱的解析記錄 |
| RecordLineId | 否 | String | 獲取某個線路 ID 對應的解析記錄，如果傳 RecordLineId，系統會忽略 RecordLine 參數。可以通過接口 DescribeRecordLineList 查看當前域名允許的線路信息 |
| GroupId | 否 | Integer | 獲取某個分組下的解析記錄時，傳這個分組 Id。可通過 DescribeRecordGroupList 接口獲取所有分組 |
| Keyword | 否 | String | 通過關鍵字搜索解析記錄，當前支持搜索主機頭和記錄值 |
| SortField | 否 | String | 排序字段，支持 name、line、type、value、weight、mx、ttl、updated_on 幾個字段 |
| SortType | 否 | String | 排序方式，正序：ASC，逆序：DESC。默認值為 ASC |
| Offset | 否 | Integer | 偏移量，默認值為 0 |
| Limit | 否 | Integer | 限制數量，當前 Limit 最大支持 3000。默認值為 100 |

### 輸出參數

| 參數名稱 | 類型 | 描述 |
|---------|------|------|
| RecordCountInfo | RecordCountInfo | 記錄的數量統計信息 |
| RecordList | Array of RecordListItem | 獲取的記錄列表 |
| RequestId | String | 唯一請求 ID |

### 示例

**輸入示例：**

```json
{
    "Offset": 0,
    "Limit": 2,
    "Domain": "dnspod.cn",
    "Subdomain": "www",
    "RecordType": "NS",
    "RecordLine": "默認"
}
```

**輸出示例：**

```json
{
    "Response": {
        "RecordCountInfo": {
            "SubdomainCount": 2,
            "TotalCount": 2,
            "ListCount": 2
        },
        "RecordList": [
            {
                "RecordId": 556507778,
                "Value": "f1g1ns1.dnspod.net.",
                "Status": "ENABLE",
                "Name": "@",
                "Line": "默認",
                "LineId": "0",
                "Type": "NS",
                "TTL": 86400,
                "MX": 0,
                "DefaultNS": true
            }
        ],
        "RequestId": "561cdfcb-37a6-47de-b3c5-2b038e2c2276"
    }
}
```

---

## 2. 添加記錄

### 接口描述

添加解析記錄。

**備註：** 新添加的解析記錄存在短暫的索引延遲，如果查詢不到新增記錄，請在 30 秒後重試

### 輸入參數

| 參數名稱 | 必選 | 類型 | 描述 |
|---------|------|------|------|
| Action | 是 | String | 公共參數，本接口取值：CreateRecord |
| Domain | 是 | String | 域名 |
| RecordType | 是 | String | 記錄類型，可通過接口 DescribeRecordType 獲得，大寫英文，比如：A |
| RecordLine | 是 | String | 記錄線路，可以通過接口 DescribeRecordLineList 查看當前域名允許的線路信息，中文，比如：默認 |
| Value | 是 | String | 記錄值，如 IP : 200.200.200.200， CNAME : cname.dnspod.com.， MX : mail.dnspod.com. |
| DomainId | 否 | Integer | 域名 ID |
| SubDomain | 否 | String | 主機記錄，如 www，默認為 @ |
| RecordLineId | 否 | String | 線路的 ID，可以通過接口 DescribeRecordLineList 查看當前域名允許的線路信息，英文字符串，比如：10=1。參數 RecordLineId 優先級高於 RecordLine，如果同時傳遞二者，優先使用 RecordLineId 參數 |
| MX | 否 | Integer | MX 優先級，當記錄類型是 MX、HTTPS、SVCB 時必填，範圍 1-65535 |
| TTL | 否 | Integer | TTL，範圍 1-604800，不同套餐域名最小值不同 |
| Weight | 否 | Integer | 權重信息，0 到 100 的整數。0 表示關閉，不傳該參數，表示不設置權重信息 |
| Status | 否 | String | 記錄初始狀態，取值範圍為 ENABLE 和 DISABLE。默認為 ENABLE，如果傳入 DISABLE，解析不會生效，也不會驗證負載均衡的限制 |
| Remark | 否 | String | 備註 |
| GroupId | 否 | Integer | 記錄分組 Id。可以通過接口 DescribeRecordGroupList 接口 GroupId 字段獲取 |
| DnssecConflictMode | 否 | String | 開啟 DNSSEC 時，強制添加 CNAME/URL 記錄 |

### 輸出參數

| 參數名稱 | 類型 | 描述 |
|---------|------|------|
| RecordId | Integer | 記錄 ID |
| RequestId | String | 唯一請求 ID |

### 示例

**輸入示例：**

```json
{
    "Domain": "dnspod.site",
    "SubDomain": "bbbb",
    "RecordType": "A",
    "RecordLine": "默認",
    "Value": "129.23.32.32",
    "TTL": 600,
    "Weight": 10,
    "Status": "DISABLE"
}
```

**輸出示例：**

```json
{
    "Response": {
        "RecordId": 162,
        "RequestId": "ab4f1426-ea15-42ea-8183-dc1b44151166"
    }
}
```

---

## 3. 修改記錄

### 接口描述

修改解析記錄。

### 輸入參數

| 參數名稱 | 必選 | 類型 | 描述 |
|---------|------|------|------|
| Action | 是 | String | 公共參數，本接口取值：ModifyRecord |
| Domain | 是 | String | 域名 |
| RecordId | 是 | Integer | 記錄 ID |
| RecordType | 是 | String | 記錄類型 |
| RecordLine | 是 | String | 記錄線路 |
| Value | 是 | String | 記錄值 |
| DomainId | 否 | Integer | 域名 ID。參數 DomainId 優先級比參數 Domain 高，如果傳遞參數 DomainId 將忽略參數 Domain。可以通過接口 DescribeDomainList 查到所有的 Domain 以及 DomainId |
| SubDomain | 否 | String | 主機記錄，如 www，如果不傳，默認為 @ |
| RecordLineId | 否 | String | 線路的 ID，可以通過接口 DescribeRecordLineList 查看當前域名允許的線路信息，比如：10=1。參數 RecordLineId 優先級高於 RecordLine，如果同時傳遞二者，優先使用 RecordLineId 參數 |
| MX | 否 | Integer | MX 優先級，當記錄類型是 MX、HTTPS、SVCB 時必填，範圍 1-65535 |
| TTL | 否 | Integer | TTL，範圍 1-604800，不同等級域名最小值不同 |
| Weight | 否 | Integer | 權重信息，0 到 100 的整數。0 表示關閉，不傳該參數，表示不設置權重信息 |
| Status | 否 | String | 記錄初始狀態，取值範圍為 ENABLE 和 DISABLE。默認為 ENABLE，如果傳入 DISABLE，解析不會生效，也不會驗證負載均衡的限制 |
| Remark | 否 | String | 記錄的備註信息。傳空刪除備註 |
| DnssecConflictMode | 否 | String | 開啟 DNSSEC 時，強制將其它記錄修改為 CNAME/URL 記錄 |

### 輸出參數

| 參數名稱 | 類型 | 描述 |
|---------|------|------|
| RecordId | Integer | 記錄 ID |
| RequestId | String | 唯一請求 ID |

---

## 4. 刪除記錄

### 接口描述

刪除解析記錄。

### 輸入參數

| 參數名稱 | 必選 | 類型 | 描述 |
|---------|------|------|------|
| Action | 是 | String | 公共參數，本接口取值：DeleteRecord |
| Domain | 是 | String | 域名 |
| RecordId | 是 | Integer | 記錄 ID。可以通過接口 DescribeRecordList 查到所有的解析記錄列表以及對應的 RecordId |
| DomainId | 否 | Integer | 域名 ID。參數 DomainId 優先級比參數 Domain 高，如果傳遞參數 DomainId 將忽略參數 Domain。可以通過接口 DescribeDomainList 查到所有的 Domain 以及 DomainId |

### 輸出參數

| 參數名稱 | 類型 | 描述 |
|---------|------|------|
| RequestId | String | 唯一請求 ID |

### 示例

**輸入示例：**

```json
{
    "Domain": "dnspod.cn",
    "RecordId": 162
}
```

**輸出示例：**

```json
{
    "Response": {
        "RequestId": "ab4f1426-ea15-42ea-8183-dc1b44151166"
    }
}
```

---

## 相關數據結構

### RecordCountInfo

記錄數量統計信息

| 名稱 | 類型 | 描述 |
|-----|------|------|
| TotalCount | Integer | 符合條件的記錄總數 |
| ListCount | Integer | 本次返回的記錄數量 |
| SubdomainCount | Integer | 符合條件的子域名數量 |

### RecordListItem

解析記錄列表元素

| 名稱 | 類型 | 描述 |
|-----|------|------|
| RecordId | Integer | 記錄 ID |
| Value | String | 記錄值 |
| Status | String | 記錄狀態：ENABLE、DISABLE |
| UpdatedOn | String | 更新時間 |
| Name | String | 主機記錄 |
| Line | String | 線路 |
| LineId | String | 線路 ID |
| Type | String | 記錄類型 |
| Weight | Integer | 權重（可能為 null） |
| MonitorStatus | String | 監控狀態 |
| Remark | String | 備註 |
| TTL | Integer | TTL 值 |
| MX | Integer | MX 優先級 |
| DefaultNS | Boolean | 是否為默認 NS 記錄 |
| GroupId | Integer | 分組 ID |
| Enabled | Integer | 是否啟用（可能為 null） |
