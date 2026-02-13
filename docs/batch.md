# 批量操作相關接口

## 1. 批量添加域名

### 接口描述

批量添加域名。

**默認接口請求頻率限制：** 20次/秒

### 輸入參數

| 參數名稱 | 必選 | 類型 | 描述 |
|---------|------|------|------|
| Action | 是 | String | 公共參數，本接口取值：CreateDomainBatch |
| DomainList | 是 | Array of String | 域名數組，最多支持 5000 個 |
| RecordList | 否 | Array of AddRecordBatch | 添加域名時同時添加的記錄列表 |

### 輸出參數

| 參數名稱 | 類型 | 描述 |
|---------|------|------|
| JobId | Integer | 批量任務 ID |
| DetailList | Array of CreateDomainBatchDetail | 批量添加域名返回結構列表 |
| RequestId | String | 唯一請求 ID |

---

## 2. 批量刪除域名

### 接口描述

批量刪除域名。

**默認接口請求頻率限制：** 20次/秒

### 輸入參數

| 參數名稱 | 必選 | 類型 | 描述 |
|---------|------|------|------|
| Action | 是 | String | 公共參數，本接口取值：DeleteDomainBatch |
| DomainIdList | 是 | Array of Integer | 域名 ID 列表，最多支持 5000 個 |

### 輸出參數

| 參數名稱 | 類型 | 描述 |
|---------|------|------|
| JobId | Integer | 批量任務 ID |
| DetailList | Array of DeleteDomainBatchDetail | 批量刪除域名詳情列表 |
| RequestId | String | 唯一請求 ID |

---

## 3. 批量添加記錄

### 接口描述

批量添加記錄。

**默認接口請求頻率限制：** 20次/秒

### 輸入參數

| 參數名稱 | 必選 | 類型 | 描述 |
|---------|------|------|------|
| Action | 是 | String | 公共參數，本接口取值：CreateRecordBatch |
| DomainIdList | 是 | Array of Integer | 域名 ID 列表，最多支持 5000 個 |
| RecordList | 是 | Array of AddRecordBatch | 批量添加的記錄列表 |

### 輸出參數

| 參數名稱 | 類型 | 描述 |
|---------|------|------|
| JobId | Integer | 批量任務 ID |
| DetailList | Array of CreateRecordBatchDetail | 批量添加記錄返回結構列表 |
| RequestId | String | 唯一請求 ID |

---

## 4. 批量修改記錄

### 接口描述

批量修改記錄。

**默認接口請求頻率限制：** 20次/秒

### 輸入參數

| 參數名稱 | 必選 | 類型 | 描述 |
|---------|------|------|------|
| Action | 是 | String | 公共參數，本接口取值：ModifyRecordBatch |
| RecordList | 是 | Array of BatchRecordInfo | 批量任務中的記錄信息列表 |

### 輸出參數

| 參數名稱 | 類型 | 描述 |
|---------|------|------|
| JobId | Integer | 批量任務 ID |
| DetailList | Array of BatchRecordInfo | 批量任務中的記錄信息列表 |
| RequestId | String | 唯一請求 ID |

---

## 5. 批量刪除記錄

### 接口描述

批量刪除解析記錄。

**默認接口請求頻率限制：** 2次/秒

### 輸入參數

| 參數名稱 | 必選 | 類型 | 描述 |
|---------|------|------|------|
| Action | 是 | String | 公共參數，本接口取值：DeleteRecordBatch |
| DomainId | 否 | Integer | 域名 ID |
| Domain | 否 | String | 域名 |
| RecordIdList | 是 | Array of Integer | 解析記錄 ID 列表，最多支持 2000 個 |

### 輸出參數

| 參數名稱 | 類型 | 描述 |
|---------|------|------|
| JobId | Integer | 批量任務 ID |
| DetailList | Array of DeleteRecordBatchDetail | 批量刪除記錄詳情列表 |
| RequestId | String | 唯一請求 ID |

---

## 6. 查看任務詳情

### 接口描述

查看批量任務詳情。

### 輸入參數

| 參數名稱 | 必選 | 類型 | 描述 |
|---------|------|------|------|
| Action | 是 | String | 公共參數，本接口取值：DescribeBatchTask |
| JobId | 是 | Integer | 任務 ID |

### 輸出參數

| 參數名稱 | 類型 | 描述 |
|---------|------|------|
| DetailList | Array of DescribeBatchTaskDetail | 查看任務詳情返回結構列表 |
| RequestId | String | 唯一請求 ID |

---

## 相關數據結構

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

### DeleteDomainBatchDetail

批量刪除域名詳情

| 名稱 | 類型 | 描述 |
|-----|------|------|
| DomainId | Integer | 域名 ID |
| Domain | String | 域名 |
| Error | String | 錯誤信息 |
| Status | String | 刪除狀態 |
| Operation | String | 操作 |

### DeleteRecordBatchDetail

批量刪除記錄詳情

| 名稱 | 類型 | 描述 |
|-----|------|------|
| DomainId | Integer | 域名 ID |
| Domain | String | 域名 |
| Error | String | 錯誤信息 |
| Status | String | 刪除狀態 |
| Operation | String | 操作 |
| RecordList | String | 解析記錄列表 |
