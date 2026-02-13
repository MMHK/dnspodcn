# 快照相關接口

## 1. 回滾域名解析記錄快照

### 接口描述

回滾域名解析記錄快照。

### 輸入參數

| 參數名稱 | 必選 | 類型 | 描述 |
|---------|------|------|------|
| Action | 是 | String | 公共參數，本接口取值：RollbackRecordSnapshot |
| Domain | 是 | String | 域名 |
| SnapshotId | 是 | String | 快照 ID |
| Record | 否 | String | 需要回滾的記錄，為空時回滾全部記錄 |
| DomainId | 否 | Integer | 域名 ID |

### 輸出參數

| 參數名稱 | 類型 | 描述 |
|---------|------|------|
| RequestId | String | 唯一請求 ID |

---

## 2. 查詢快照列表

### 接口描述

查詢快照列表。

### 輸入參數

| 參數名稱 | 必選 | 類型 | 描述 |
|---------|------|------|------|
| Action | 是 | String | 公共參數，本接口取值：DescribeSnapshotList |
| Domain | 是 | String | 域名 |
| DomainId | 否 | Integer | 域名 ID |

### 輸出參數

| 參數名稱 | 類型 | 描述 |
|---------|------|------|
| SnapshotList | Array of SnapshotInfo | 快照列表 |
| RequestId | String | 唯一請求 ID |

---

## 3. 下載快照

### 接口描述

下載快照。

### 輸入參數

| 參數名稱 | 必選 | 類型 | 描述 |
|---------|------|------|------|
| Action | 是 | String | 公共參數，本接口取值：DownloadSnapshot |
| Domain | 是 | String | 域名 |
| SnapshotId | 是 | String | 快照 ID |
| DomainId | 否 | Integer | 域名 ID |

### 輸出參數

| 參數名稱 | 類型 | 描述 |
|---------|------|------|
| CosUrl | String | 快照的對象存儲地址 |
| RequestId | String | 唯一請求 ID |

---

## 相關數據結構

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
