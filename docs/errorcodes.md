# 錯誤碼

本文檔列出了 DNSPod API 的錯誤碼。

## 錯誤返回格式

當 API 調用失敗時，返回結果中會包含 `Error` 字段：

```json
{
    "Response": {
        "Error": {
            "Code": "AuthFailure",
            "Message": "CAM簽名/鑒權錯誤。"
        },
        "RequestId": "eac6b301-a322-493a-8e36-83b295459397"
    }
}
```

## 公共錯誤碼

| 錯誤碼 | 描述 |
|-------|------|
| AuthFailure | CAM 簽名/鑒權錯誤 |
| AuthFailure.InvalidSecretId | 密鑰非法 |
| AuthFailure.SecretIdNotFound | 密鑰不存在 |
| AuthFailure.SignatureExpire | 簽名過期 |
| AuthFailure.SignatureFailure | 簽名錯誤 |
| AuthFailure.TokenFailure | token 錯誤 |
| DryRunOperation | DryRun 操作 |
| FailedOperation | 操作失敗 |
| InternalError | 內部錯誤 |
| InvalidAction | 接口不存在 |
| InvalidParameter | 參數錯誤 |
| InvalidParameterValue | 參數取值錯誤 |
| LimitExceeded | 超過配額限制 |
| MissingParameter | 缺少參數 |
| OperationDenied | 操作被拒絕 |
| RequestLimitExceeded | 請求的次數超過了頻率限制 |
| ResourceInUse | 資源被占用 |
| ResourceInsufficient | 資源不足 |
| ResourceNotFound | 資源不存在 |
| UnauthorizedOperation | 未授權操作 |
| UnknownParameter | 未知參數 |
| UnsupportedOperation | 操作不支持 |
| UnsupportedProtocol | HTTP(S) 請求協議錯誤 |
| UnsupportedRegion | 接口不支持所傳地域 |

## 業務錯誤碼

### 賬戶相關

| 錯誤碼 | 描述 |
|-------|------|
| FailedOperation.AccountIsLocked | 抱歉，該賬戶已經被鎖定 |
| FailedOperation.NotRealNamedUser | 未實名認證用戶，請先完成實名認證再操作 |
| FailedOperation.UnknowError | 操作未響應，請稍後重試 |
| InvalidParameter.AccountIsBanned | 您的賬號已被系統封禁 |
| InvalidParameter.UnrealNameUser | 未實名認證用戶，請先完成實名認證再操作 |
| InvalidParameter.UserNotExists | 用戶不存在 |
| InvalidParameterValue.UserIdInvalid | 用戶編號不正確 |
| LimitExceeded.FailedLoginLimitExceeded | 登錄失敗次數過多已被系統封禁 |
| FailedOperation.LoginAreaNotAllowed | 賬號異地登錄，請求被拒絕 |
| FailedOperation.LoginFailed | 登錄失敗，請檢查賬號和密碼是否正確 |
| FailedOperation.LoginTimeout | 登錄已經超時，請重新登錄 |
| InvalidParameter.EmailNotVerified | 抱歉，您的賬戶還沒有通過郵箱驗證 |
| InvalidParameter.MobileNotVerified | 抱歉，您的賬戶還沒有通過手機驗證 |

### 域名相關

| 錯誤碼 | 描述 |
|-------|------|
| FailedOperation.DomainExists | 該域名已在您的列表中，無需重複添加 |
| FailedOperation.DomainIsLocked | 鎖定域名不能進行此操作 |
| FailedOperation.DomainIsSpam | 封禁域名不能進行此操作 |
| FailedOperation.DomainIsVip | VIP 域名不能進行此操作 |
| FailedOperation.DomainIsKeyDomain | 該域名為騰訊雲 DNSPod 重點保護資源，禁止自行操作刪除 |
| FailedOperation.NotDomainOwner | 域名不在您的名下 |
| FailedOperation.DomainOwnedByOtherUser | 該域名已被其他賬號添加，可在域名列表中添加取回 |
| FailedOperation.DomainInEffectOrInvalidated | 不允許操作生效中或失效中的域名 |
| InvalidParameter.DomainIdInvalid | 域名編號不正確 |
| InvalidParameter.DomainInvalid | 域名不正確，請輸入主域名，如 dnspod.cn |
| InvalidParameter.DomainIsAliaser | 此域名是其它域名的別名 |
| InvalidParameter.DomainIsMyAlias | 此域名是自己域名的別名 |
| InvalidParameter.DomainIsNotlocked | 域名沒有鎖定 |
| InvalidParameter.DomainNotAllowedLock | 暫停域名不支持鎖定 |
| InvalidParameter.DomainNotAllowedModifyRecords | 處於生效中/失效中的域名，不允許變更解析記錄 |
| InvalidParameter.DomainNotBeian | 該域名未備案，無法添加 URL 記錄 |
| InvalidParameter.DomainRecordExist | 記錄已經存在，無需再次添加 |
| InvalidParameterValue.DomainNotExists | 當前域名有誤，請返回重新操作 |
| InvalidParameter.LockDaysInvalid | 鎖定天數不正確 |
| InvalidParameter.UnLockCodeExpired | 解鎖代碼已失效 |
| InvalidParameter.UnLockCodeInvalid | 解鎖代碼不正確 |
| InvalidParameter.AliasIsMyDomain | 該域名已在您的域名列表中，請刪除後再添加到別名列表 |
| InvalidParameter.DomainAliasExists | 別名已經存在 |
| InvalidParameter.DomainAliasIdInvalid | 別名編號錯誤 |
| InvalidParameter.GroupIdInvalid | 分組編號不正確 |
| InvalidParameter.GroupNameEmpty | 分組名為空 |
| InvalidParameter.GroupNameExists | 同名分組已經存在 |
| InvalidParameter.GroupNameInvalid | 分組名為 1-17 個字符 |
| InvalidParameter.GroupNameOccupied | 指定的分組名已存在，或為系統內置線路或自定義線路 |

### 記錄相關

| 錯誤碼 | 描述 |
|-------|------|
| InvalidParameter.RecordIdInvalid | 記錄編號錯誤 |
| InvalidParameter.RecordLineInvalid | 記錄線路不正確 |
| InvalidParameter.RecordTypeInvalid | 記錄類型不正確 |
| InvalidParameter.RecordValueInvalid | 記錄的值不正確 |
| InvalidParameter.RecordValueLengthInvalid | 解析記錄值過長 |
| InvalidParameter.SubdomainInvalid | 子域名不正確 |
| InvalidParameter.MxInvalid | MX 優先級不正確 |
| InvalidParameter.InvalidWeight | 權重不合法。請輸入 0~100 的整數 |
| LimitExceeded.AAAACountLimit | AAAA 記錄數量超出限制 |
| LimitExceeded.NsCountLimit | NS 記錄數量超出限制 |
| LimitExceeded.SrvCountLimit | SRV 記錄數量超出限制 |
| LimitExceeded.UrlCountLimit | 該域名的顯性 URL 轉發數量已達上限 |
| LimitExceeded.HiddenUrlExceeded | 該域名使用的套餐不支持隱性 URL 轉發或數量已達上限 |
| LimitExceeded.SubdomainLevelLimit | 子域名級數超出限制 |
| LimitExceeded.SubdomainRollLimit | 子域名負載均衡數量超出限制 |
| LimitExceeded.SubdomainWcardLimit | 泛解析級數超出限制 |
| LimitExceeded.RecordTtlLimit | 記錄的 TTL 值超出了限制 |
| LimitExceeded.AtNsRecordLimit | @ 的 NS 記錄只能設置為默認線路 |
| ResourceNotFound.NoDataOfRecord | 記錄列表為空 |
| FailedOperation.MustAddDefaultLineFirst | 請先添加默認線路的解析記錄 |
| InvalidParameter.DomainSelfNoCopy | 域名自己無需進行複製 |
| InvalidParameter.GradeNotCopy | 當前域名等級低於源域名的等級，無法進行複製 |

### DNSSEC 相關

| 錯誤碼 | 描述 |
|-------|------|
| FailedOperation.DNSSECIncompleteClosed | DNSSEC 未完全關閉，不允許添加 @ 子域名 CNAME、顯性 URL 或者隱性 URL 記錄 |
| InvalidParameter.DnssecAddCnameError | 該域名開啟了 DNSSEC，不允許添加 @ 子域名 CNAME、顯性 URL 或者隱性 URL 記錄 |

### 批量操作相關

| 錯誤碼 | 描述 |
|-------|------|
| FailedOperation.NotBatchTaskOwner | 權限錯誤，您無法查看該任務的詳情 |
| InvalidParameter.BatchDomainCreateActionError | 創建批量域名任務失敗 |
| InvalidParameter.BatchDomainNotAuth | 列表中存在您沒有權限的域名 |
| InvalidParameter.BatchLimitUndo | 您有批量任務未執行完成，請等待完成後繼續添加 |
| InvalidParameter.BatchRecordCreateActionError | 創建批量記錄任務失敗 |
| InvalidParameter.BatchRecordModifyActionError | 批量修改記錄任務失敗 |
| InvalidParameter.BatchRecordRemoveActionError | 批量刪除記錄任務失敗 |
| InvalidParameter.BatchTaskCountLimit | 超過單個賬號的批量任務數並發上限 4 個 |
| InvalidParameter.BatchTaskNotExist | 任務不存在，無法獲取任務詳情 |
| InvalidParameter.JobGreaterThanLimit | 單次任務記錄數量超過上限 5000 條 |

### 線路相關

| 錯誤碼 | 描述 |
|-------|------|
| InvalidParameter.CopiedLineGroupDuplicated | 您複製的線路已存在，無需重複複製 |
| InvalidParameter.DefaultLineNotSelfdefined | 默認線路無法進行自定義線路分組 |
| InvalidParameter.LineFormatInvalid | 線路格式不正確 |
| InvalidParameter.LineGroupNotSupported | 線路不存在，或者線路不支持自定義分組 |
| InvalidParameter.LineGroupOverCounted | 當前套餐的線路分組已達到數量上限 |
| InvalidParameter.LineGroupUpdateFailed | 線路分組更新失敗 |
| InvalidParameter.LineInAnotherGroup | 線路已存在於其他分組中 |
| InvalidParameter.LineInUse | 線路正在使用當中，無法修改名稱 |
| InvalidParameter.LineInUseNotDelete | 線路正在使用當中，無法刪除名稱 |
| InvalidParameter.LineNameInvalid | 線路名稱的長度不能超過 17 個字符 |
| InvalidParameter.LineNameInvalidCharacter | 線路名稱包含不被接受的字符 |
| InvalidParameter.LineNameOccupied | 線路名是系統內置線路或用戶自定義分組線路 |
| InvalidParameter.LineNotExist | 分組不存在 |
| InvalidParameter.LineNotSelected | 您至少需要選擇一個線路 |
| InvalidParameter.LineOverCounted | 最多只能選擇 120 個線路 |
| InvalidParameter.IpArea | 刪除自定義線路失敗，原因：線路不存在或者已刪除 |
| InvalidParameter.InvalidIp | 不是合法的 IP 段 |
| InvalidParameter.IpAlreadyExist | IP 已經存在 |
| InvalidParameter.IpsExceedLimit | IP 段過長 |

### 權限相關

| 錯誤碼 | 描述 |
|-------|------|
| OperationDenied | 操作被拒絕 |
| OperationDenied.AccessDenied | 您沒有權限執行此操作 |
| OperationDenied.DomainOwnerAllowedOnly | 僅域名所有者可進行此操作 |
| OperationDenied.NoPermissionToOperateDomain | 當前域名無權限，請返回域名列表 |
| OperationDenied.NotAdmin | 您不是管理用戶 |
| OperationDenied.NotAgent | 您不是代理用戶 |
| OperationDenied.NotManagedUser | 不是您名下用戶 |
| UnauthorizedOperation | 未授權操作 |
| FailedOperation.NotResourceOwner | 您沒有權限操作此資源 |
| InvalidParameter.RequestIpLimited | 您的 IP 非法，請求被拒絕 |
| OperationDenied.IPInBlacklistNotAllowed | 抱歉，不允許添加黑名單中的 IP |

### 套餐相關

| 錯誤碼 | 描述 |
|-------|------|
| FailedOperation.ContainsPersonalVip | 您的賬戶下包含個人豪華域名，不能直接升級 |
| FailedOperation.CouponForFreeDomain | 此優惠券只能被免費域名使用 |
| FailedOperation.CouponNotSupported | 您的賬戶不滿足使用此優惠券的條件 |
| FailedOperation.CouponTypeAlreadyUsed | 域名已經使用過該類型的禮券 |
| FailedOperation.OrderCanNotPay | 您不能付款此訂單 |
| FailedOperation.OrderHasPaid | 此訂單已經付過款 |
| FailedOperation.InsufficientBalance | 賬戶餘額不足 |
| FailedOperation.DomainAlreadyVip | 目標域名已經是 VIP 域名，無法替換 |
| FailedOperation.DomainNotVip | 原域名不是 VIP 域名，無法替換 |
| FailedOperation.DomainIsEnterpriseType | 域名已升級為企業套餐 |
| FailedOperation.DomainIsPersonalType | 域名已升級為個人套餐 |
| InvalidParameter.InvalidCoupon | 禮券代碼無效 |
| InvalidParameter.GoodsChildTypeInvalid | 商品子類型無效 |
| InvalidParameter.GoodsNumInvalid | 商品數量無效 |
| InvalidParameter.GoodsTypeInvalid | 商品類型無效 |
| InvalidParameter.DealTypeInvalid | 訂單類型無效 |

### 解析量統計相關

| 錯誤碼 | 描述 |
|-------|------|
| FailedOperation.DomainNotInService | 當前域名還未使用 DNSPod 的解析服務，無法獲取解析量數據 |
| FailedOperation.TemporaryError | 請求量統計數據暫時不可用，請稍後再試 |
| FailedOperation.AuthLogUnsupport | 當前套餐版本不支持流量分析 |
| InvalidParameter.AuthLogInvalidRetCode | 解析狀態不正確 |
| InvalidParameter.AuthLogInvalidScope | 地域不正確 |
| InvalidParameter.EndDateBeyondRange | 結束時間超出範圍 |
| InvalidParameter.InvalidEndDate | 無效的結束時間 |
| InvalidParameter.InvalidStartDate | 無效的開始時間 |
| InvalidParameter.InvalidTime | 無效的時間 |

### 其他

| 錯誤碼 | 描述 |
|-------|------|
| FailedOperation.FrequencyLimit | 您操作過於頻繁，請稍後重試 |
| FailedOperation.GetWhoisFailed | 獲取不到域名信息，可能域名非法或網絡故障 |
| FailedOperation.TencentCloudForbid | 按騰訊雲域名新規範，tencentyun.com 不允許新增子域名和修改存量解析 |
| InvalidParameter.InvalidSignature | 無效簽名 |
| InvalidParameter.InvalidSecretId | 無效密鑰 ID |
| InvalidParameter.ParamInvalid | param 格式錯誤 |
| InvalidParameter.ResultMoreThan500 | 搜索結果大於 500 條，請增加關鍵字 |
| InvalidParameter.OperateFailed | 操作失敗，請稍後再試 |
| InvalidParameterValue.LimitInvalid | 分頁長度數量錯誤 |
| InvalidParameter.OffsetInvalid | 分頁起始數量錯誤 |
| RequestLimitExceeded.RequestLimitExceeded | API 請求次數超出限制 |
