// Copyright 2020 CleverGo. All rights reserved.
// Use of this source code is governed by a MIT style license that can be found
// in the LICENSE file.

package dnspodcn

import (
	"net/http"
	"strconv"
	"time"

	"github.com/libdns/libdns"
)

const (
	defaultRecordLine = "默认"
)

// CreateRecordRequest is a request for creating a DNS record.
type CreateRecordRequest struct {
	BasicRequest
	Domain             string `json:"Domain,omitempty"`             // 域名
	DomainId           int    `json:"DomainId,omitempty"`           // 域名 ID，優先級比 Domain 高
	SubDomain          string `json:"SubDomain,omitempty"`          // 主機記錄，如 www，默認為 @
	RecordType         string `json:"RecordType,omitempty"`         // 記錄類型，如：A、CNAME、MX 等
	RecordLine         string `json:"RecordLine,omitempty"`         // 記錄線路，如：默認
	RecordLineId       string `json:"RecordLineId,omitempty"`       // 線路的 ID
	Value              string `json:"Value,omitempty"`              // 記錄值
	MX                 int    `json:"MX,omitempty"`                 // MX 優先級，MX/HTTPS/SVCB 記錄時必填，範圍 1-65535
	TTL                int    `json:"TTL,omitempty"`                // TTL，範圍 1-604800，默認 600
	Weight             int    `json:"Weight,omitempty"`             // 權重信息，0 到 100 的整數
	Status             string `json:"Status,omitempty"`             // 記錄初始狀態：ENABLE、DISABLE
	Remark             string `json:"Remark,omitempty"`             // 備註
	GroupId            int    `json:"GroupId,omitempty"`            // 記錄分組 Id
	DnssecConflictMode string `json:"DnssecConflictMode,omitempty"` // 開啟DNSSEC時，強制添加CNAME/URL記錄
}

// NewCreateRecordRequest returns a request for creating a DNS record.
func NewCreateRecordRequest(domain, subdomain, recordType, value string) *CreateRecordRequest {
	return &CreateRecordRequest{
		BasicRequest: NewBasicRequest(http.MethodPost, "CreateRecord"),
		Domain:       domain,
		SubDomain:    subdomain,
		RecordType:   recordType,
		Value:        value,
		RecordLine:   defaultRecordLine,
	}
}

// CreateRecordResponse is a response of creating a DNS record.
type CreateRecordResponse struct {
	BasicResponse
	RecordId int `json:"RecordId"` // 記錄 ID
}

// RecordResponse contains the infomation of DNS record field.
type RecordResponse struct {
	ID      string `json:"id"`
	Type    string `json:"type"`
	Name    string `json:"name"`
	Status  string `json:"status"`
	Line    string `json:"line"`
	LineID  string `json:"line_id"`
	TTL     string `json:"ttl"`
	Value   string `json:"value"`
	Weight  string `json:"weight"`
	MX      string `json:"mx"`
	Enabled string `json:"enabled"`
	Remark  string `json:"remark"`
}

func (resp RecordResponse) libdnsRecord() libdns.Record {
	ttl, _ := strconv.Atoi(resp.TTL)
	return libdns.Record{
		ID:    resp.ID,
		Type:  resp.Type,
		Name:  resp.Name,
		Value: resp.Value,
		TTL:   time.Second * time.Duration(ttl),
	}
}

// DescribeRecordListRequest is a request for describing DNS record list.
type DescribeRecordListRequest struct {
	BasicRequest
	Domain       string `json:"Domain,omitempty"`       // 域名
	DomainId     int    `json:"DomainId,omitempty"`     // 域名 ID，優先級比 Domain 高
	Subdomain    string `json:"Subdomain,omitempty"`    // 解析記錄的主機頭
	RecordType   string `json:"RecordType,omitempty"`   // 獲取某種類型的解析記錄，如 A、CNAME、NS、AAAA 等
	RecordLine   string `json:"RecordLine,omitempty"`   // 獲取某條線路名稱的解析記錄
	RecordLineId string `json:"RecordLineId,omitempty"` // 獲取某個線路 ID 對應的解析記錄
	GroupId      int    `json:"GroupId,omitempty"`      // 獲取某個分組下的解析記錄
	Keyword      string `json:"Keyword,omitempty"`      // 通過關鍵字搜索解析記錄
	SortField    string `json:"SortField,omitempty"`    // 排序字段，支持 name、line、type、value、weight、mx、ttl、updated_on
	SortType     string `json:"SortType,omitempty"`     // 排序方式，正序：ASC，逆序：DESC
	Offset       int    `json:"Offset,omitempty"`       // 偏移量，默認值為 0
	Limit        int    `json:"Limit,omitempty"`        // 限制數量，最大支持 3000，默認值為 100
}

// NewDescribeRecordListRequest returns a request for describing DNS record list.
func NewDescribeRecordListRequest(domain string) *DescribeRecordListRequest {
	return &DescribeRecordListRequest{
		BasicRequest: NewBasicRequest(http.MethodPost, "DescribeRecordList"),
		Domain:       domain,
	}
}

// RecordCountInfo is the count info of DNS records.
type RecordCountInfo struct {
	TotalCount     int `json:"TotalCount"`     // 符合條件的記錄總數
	ListCount      int `json:"ListCount"`      // 本次返回的記錄數量
	SubdomainCount int `json:"SubdomainCount"` // 符合條件的子域名數量
}

// RecordListItem is the item of DNS record list.
type RecordListItem struct {
	RecordId      int    `json:"RecordId"`      // 記錄 ID
	Value         string `json:"Value"`         // 記錄值
	Status        string `json:"Status"`        // 記錄狀態：ENABLE、DISABLE
	UpdatedOn     string `json:"UpdatedOn"`     // 更新時間
	Name          string `json:"Name"`          // 主機記錄
	Line          string `json:"Line"`          // 線路
	LineId        string `json:"LineId"`        // 線路 ID
	Type          string `json:"Type"`          // 記錄類型
	Weight        *int   `json:"Weight"`        // 權重（可能為 null）
	MonitorStatus string `json:"MonitorStatus"` // 監控狀態
	Remark        string `json:"Remark"`        // 備註
	TTL           int    `json:"TTL"`           // TTL 值
	MX            int    `json:"MX"`            // MX 優先級
	DefaultNS     bool   `json:"DefaultNS"`     // 是否為默認 NS 記錄
	GroupId       int    `json:"GroupId"`       // 分組 ID
	Enabled       *int   `json:"Enabled"`       // 是否啟用（可能為 null）
}

func (item RecordListItem) libdnsRecord() libdns.Record {
	return libdns.Record{
		ID:    strconv.Itoa(item.RecordId),
		Type:  item.Type,
		Name:  item.Name,
		Value: item.Value,
		TTL:   time.Second * time.Duration(item.TTL),
	}
}

// GetWeight returns the weight value, or 0 if nil.
func (item RecordListItem) GetWeight() int {
	if item.Weight == nil {
		return 0
	}
	return *item.Weight
}

// GetEnabled returns the enabled value, or 0 if nil.
func (item RecordListItem) GetEnabled() int {
	if item.Enabled == nil {
		return 0
	}
	return *item.Enabled
}

// DescribeRecordListResponse is the response of describing DNS record list.
type DescribeRecordListResponse struct {
	BasicResponse
	RecordCountInfo RecordCountInfo  `json:"RecordCountInfo"` // 記錄的數量統計信息
	RecordList      []RecordListItem `json:"RecordList"`      // 獲取的記錄列表
}

// Error returns the error of response.
func (resp DescribeRecordListResponse) Error() error {
	if resp.Status.Code == CodeSuccess || resp.Status.Code == CodeNoRecords {
		return nil
	}
	return resp.Status.Error()
}

// DeleteRecordRequest is a request for deleting a DNS record.
type DeleteRecordRequest struct {
	BasicRequest
	Domain   string `json:"Domain,omitempty"`   // 域名
	DomainId int    `json:"DomainId,omitempty"` // 域名 ID
	RecordId int    `json:"RecordId,omitempty"` // 記錄 ID
}

// NewDeleteRecordRequest returns a request for deleting a DNS record.
func NewDeleteRecordRequest(domain string, recordId int) *DeleteRecordRequest {
	return &DeleteRecordRequest{
		BasicRequest: NewBasicRequest(http.MethodPost, "DeleteRecord"),
		Domain:       domain,
		RecordId:     recordId,
	}
}

// DeleteRecordResponse is the response of deleting a DNS record.
type DeleteRecordResponse struct {
	BasicResponse
}

// UpdateRecordRequest is a request for updating a DNS record.
type UpdateRecordRequest struct {
	BasicRequest
	Domain             string `json:"Domain,omitempty"`             // 域名
	DomainId           int    `json:"DomainId,omitempty"`           // 域名 ID
	SubDomain          string `json:"SubDomain,omitempty"`          // 主機記錄
	RecordId           int    `json:"RecordId,omitempty"`           // 記錄 ID
	RecordType         string `json:"RecordType,omitempty"`         // 記錄類型
	RecordLine         string `json:"RecordLine,omitempty"`         // 記錄線路
	RecordLineId       string `json:"RecordLineId,omitempty"`       // 線路的 ID
	Value              string `json:"Value,omitempty"`              // 記錄值
	MX                 int    `json:"MX,omitempty"`                 // MX 優先級
	TTL                int    `json:"TTL,omitempty"`                // TTL
	Weight             int    `json:"Weight,omitempty"`             // 權重信息
	Status             string `json:"Status,omitempty"`             // 記錄狀態：ENABLE、DISABLE
	Remark             string `json:"Remark,omitempty"`             // 備註信息
	DnssecConflictMode string `json:"DnssecConflictMode,omitempty"` // 開啟DNSSEC時，強制將其它記錄修改為CNAME/URL記錄
}

// NewUpdateRecordRequest returns a request for updating a DNS record.
func NewUpdateRecordRequest(domain string, recordId int, recordType, value string) *UpdateRecordRequest {
	return &UpdateRecordRequest{
		BasicRequest: NewBasicRequest(http.MethodPost, "ModifyRecord"),
		Domain:       domain,
		RecordId:     recordId,
		RecordType:   recordType,
		Value:        value,
		RecordLine:   defaultRecordLine,
	}
}

// UpdateRecordResponse is the response of updating a DNS record.
type UpdateRecordResponse struct {
	BasicResponse
	RecordId int `json:"RecordId"` // 記錄 ID
}
