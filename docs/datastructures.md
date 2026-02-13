# 數據結構

本文檔列出了 DNSPod API 使用的所有數據結構。

---

## 域名相關

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
| ErrorTotal | Integer | 錯誤域名數量 |
| LockTotal | Integer | 鎖定域名數量 |
| SpamTotal | Integer | 封禁域名數量 |
| VipExpire | Integer | 過期域名數量 |
| ShareOutTotal | Integer | 我共享出去的域名數量 |
| GroupTotal | Integer | 分組總數 |

### DomainListItem

域名列表元素

| 名稱 | 類型 | 描述 |
|-----|------|------|
| DomainId | Integer | 域名 ID |
| Name | String | 域名 |
| Punycode | String | 域名的 punycode 碼 |
| Grade | String | 域名套餐等級，如：DP_FREE、DP_PLUS、DP_EXTRA、DP_ULTRA |
| GradeLevel | Integer | 域名套餐等級對應的序號 |
| GradeTitle | String | 域名套餐名稱 |
| Status | String | 域名狀態：ENABLE（啟用）、PAUSE（暫停） |
| TTL | Integer | 域名的 TTL |
| CNAMESpeedup | String | CNAME 加速狀態：ENABLE、DISABLE |
| DNSStatus | String | DNS 設置狀態 |
| IsVip | String | 是否 VIP 域名：YES、NO |
| RecordCount | Integer | 解析記錄數量 |
| CreatedOn | String | 域名添加時間 |
| UpdatedOn | String | 域名最近修改時間 |
| Owner | String | 域名所有者 |
| GroupId | Integer | 域名所屬分組 ID |
| IsMark | String | 是否標記：YES、NO |
| Remark | String | 域名備註 |
| PunycodeSerial | String | 域名的 unicode 碼 |
| SearchEnginePush | String | 搜索引擎推送：YES、NO |
| TagList | Array of TagItem | 域名標籤列表 |
| VipAutoRenew | String | VIP 自動續費：YES、NO |
| VipEndAt | String | VIP 到期時間 |
| VipStartAt | String | VIP 開始時間 |
| EffectiveDNS | Array of String | 生效的 DNS 服務器列表 |

### GroupInfo

分組信息

| 名稱 | 類型 | 描述 |
|-----|------|------|
| GroupId | Integer | 分組 ID |
| GroupName | String | 分組名稱 |
| GroupType | String | 分組類型：system、user |
| Size | Integer | 該分組內的域名數量 |

### LockInfo

域名鎖定信息

| 名稱 | 類型 | 描述 |
|-----|------|------|
| DomainId | Integer | 域名 ID |
| LockCode | String | 域名解鎖碼 |
| LockEnd | String | 域名鎖定結束時間 |

---

## 記錄相關

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

---

## 批量操作相關

### AddRecordBatch

批量添加的記錄

| 名稱 | 類型 | 必選 | 描述 |
|-----|------|------|------|
| RecordType | String | 是 | 記錄類型 |
| Value | String | 是 | 記錄值 |
| SubDomain | String | 否 | 子域名，默認為 @ |
| RecordLine | String | 否 | 解析記錄的線路 |
| RecordLineId | String | 否 | 解析記錄的線路 ID |
| MX | Integer | 否 | MX 記錄值，非 MX 記錄默認為 0 |
| TTL | Integer | 否 | 記錄的 TTL 值，默認 600 |

### CreateDomainBatchDetail

批量添加域名返回結構

| 名稱 | 類型 | 描述 |
|-----|------|------|
| Id | Integer | 任務編號 |
| Domain | String | 域名 |
| DomainGrade | String | 域名等級 |
| Status | String | 該條任務運行狀態 |
| Operation | String | 操作類型 |
| ErrMsg | String | 錯誤信息 |
| RecordList | Array of CreateDomainBatchRecord | 記錄列表 |

### CreateRecordBatchDetail

批量添加記錄返回結構

| 名稱 | 類型 | 描述 |
|-----|------|------|
| Id | Integer | 任務編號 |
| Domain | String | 域名 |
| DomainId | Integer | 域名 ID |
| DomainGrade | String | 域名等級 |
| Status | String | 該條任務運行狀態 |
| Operation | String | 操作類型 |
| ErrMsg | String | 錯誤信息 |
| RecordList | Array of CreateRecordBatchRecord | 記錄列表 |

### BatchRecordInfo

批量任務中的記錄信息

| 名稱 | 類型 | 描述 |
|-----|------|------|
| RecordId | Integer | 記錄 ID |
| SubDomain | String | 子域名 |
| RecordType | String | 記錄類型 |
| RecordLine | String | 解析記錄的線路 |
| Value | String | 記錄值 |
| TTL | Integer | 記錄的 TTL 值 |
| Status | String | 記錄添加狀態 |
| Operation | String | 操作類型 |
| ErrMsg | String | 錯誤信息 |
| Id | Integer | 此條記錄在列表中的 ID |
| Enabled | Integer | 記錄生效狀態 |
| MX | Integer | 記錄的 MX 權重 |
| Weight | Integer | 記錄權重 |
| Remark | String | 備註信息 |

---

## 線路相關

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

---

## 解析量相關

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

---

## 別名相關

### DomainAliasInfo

域名別名信息

| 名稱 | 類型 | 描述 |
|-----|------|------|
| Id | Integer | 域名別名 ID |
| DomainAlias | String | 域名別名 |
| Status | Integer | 別名狀態：1-DNS 不正確；2-正常；3-封禁 |

---

## 套餐相關

### Deals

子訂單號列表

| 名稱 | 類型 | 描述 |
|-----|------|------|
| DealId | String | 子訂單 ID |
| DealName | String | 子訂單號 |

---

## 用戶相關

### UserInfo

用戶信息

| 名稱 | 類型 | 描述 |
|-----|------|------|
| Nickname | String | 用戶昵稱 |
| Email | String | 用戶郵箱 |
| Phone | String | 用戶手機號碼 |
| EmailVerified | String | 郵箱是否驗證 |
| PhoneVerified | String | 手機是否驗證 |
| AccountId | String | 賬戶 ID |
| RealName | String | 是否實名認證 |

---

## 快照相關

### SnapshotInfo

快照信息

| 名稱 | 類型 | 描述 |
|-----|------|------|
| SnapshotId | String | 快照 ID |
| Domain | String | 域名 |
| Status | String | 快照狀態 |
| CreatedOn | String | 創建時間 |
| UpdatedOn | String | 更新時間 |
| RecordCount | Integer | 記錄數量 |

---

## 其他

### TagItem

標籤項

| 名稱 | 類型 | 描述 |
|-----|------|------|
| TagKey | String | 標籤鍵 |
| TagValue | String | 標籤值 |

### TagItemFilter

標籤過濾項

| 名稱 | 類型 | 描述 |
|-----|------|------|
| TagKey | String | 標籤鍵 |
| TagValue | Array of String | 標籤值列表 |
