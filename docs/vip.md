# 套餐及增值服務相關接口

## 1. 獲取域名套餐列表

### 接口描述

獲取域名套餐列表。

### 輸入參數

| 參數名稱 | 必選 | 類型 | 描述 |
|---------|------|------|------|
| Action | 是 | String | 公共參數，本接口取值：DescribeDomainFilterList |
| Type | 否 | String | 域名分組類型 |
| Offset | 否 | Integer | 記錄開始的偏移 |
| Limit | 否 | Integer | 要獲取的域名數量 |

### 輸出參數

| 參數名稱 | 類型 | 描述 |
|---------|------|------|
| DomainCountInfo | DomainCountInfo | 列表頁統計信息 |
| DomainList | Array of DomainListItem | 域名列表 |
| RequestId | String | 唯一請求 ID |

---

## 2. 升級域名套餐

### 接口描述

升級域名套餐。

### 輸入參數

| 參數名稱 | 必選 | 類型 | 描述 |
|---------|------|------|------|
| Action | 是 | String | 公共參數，本接口取值：ModifyDomainOwner |
| Domain | 是 | String | 域名 |
| ToDomain | 是 | String | 目標域名 |
| DomainId | 否 | Integer | 域名 ID |

### 輸出參數

| 參數名稱 | 類型 | 描述 |
|---------|------|------|
| RequestId | String | 唯一請求 ID |

---

## 3. 創建訂單

### 接口描述

創建訂單。

### 輸入參數

| 參數名稱 | 必選 | 類型 | 描述 |
|---------|------|------|------|
| Action | 是 | String | 公共參數，本接口取值：CreateDeal |
| DealName | 是 | String | 交易名稱 |
| GoodsType | 是 | String | 商品類型 |
| GoodsChildType | 是 | String | 商品子類型 |
| GoodsNum | 是 | Integer | 商品數量 |
| GoodsPriceId | 是 | Integer | 商品價格 ID |

### 輸出參數

| 參數名稱 | 類型 | 描述 |
|---------|------|------|
| Deals | Array of Deals | 子訂單號列表 |
| BigDealId | String | 大訂單號 |
| RequestId | String | 唯一請求 ID |

---

## 相關數據結構

### Deals

子訂單號列表

| 名稱 | 類型 | 描述 |
|-----|------|------|
| DealId | String | 子訂單 ID |
| DealName | String | 子訂單號 |
