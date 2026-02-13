// Copyright 2020 CleverGo. All rights reserved.
// Use of this source code is governed by a MIT style license that can be found
// in the LICENSE file.

package dnspodcn

import (
	"bytes"
	"context"
	"crypto/hmac"
	"crypto/sha256"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"io"
	"math/rand"
	"net/http"
	"strconv"
	"strings"
	"time"
)

const (
	name       = "Go DNSPod Client"
	version    = "dev"
	apiBaseURL = "https://dnspod.tencentcloudapi.com"
	apiVersion = "2021-03-23"
	service    = "dnspod"
)

// Client is an API client for Tencent Cloud DNSPod API 3.0.
type Client struct {
	// SecretId is the access key ID from Tencent Cloud CAM.
	SecretId string
	// SecretKey is the access key secret from Tencent Cloud CAM.
	SecretKey string
	// BaseURL is the API endpoint.
	BaseURL string
	// Region is the API region, e.g., ap-guangzhou, ap-shanghai, ap-beijing.
	Region string
}

// NewClient returns a client with the given secret ID, secret key and default settings.
func NewClient(secretId, secretKey string) *Client {
	return &Client{
		SecretId:  secretId,
		SecretKey: secretKey,
		BaseURL:   apiBaseURL,
		Region:    "",
	}
}

// sha256hex calculates SHA256 hash and returns hex string
func sha256hex(s string) string {
	b := sha256.Sum256([]byte(s))
	return hex.EncodeToString(b[:])
}

// hmacsha256 calculates HMAC-SHA256
func hmacsha256(key []byte, s string) []byte {
	h := hmac.New(sha256.New, key)
	h.Write([]byte(s))
	return h.Sum(nil)
}

// sign generates the TC3-HMAC-SHA256 signature for Tencent Cloud API 3.0
func (c *Client) sign(action string, payload string, timestamp int64) (string, string) {
	// Get current date in UTC
	date := time.Unix(timestamp, 0).UTC().Format("2006-01-02")

	// Step 1: Create Canonical Request
	httpMethod := "POST"
	canonicalURI := "/"
	canonicalQueryString := ""
	host := strings.TrimPrefix(c.BaseURL, "https://")
	canonicalHeaders := fmt.Sprintf("content-type:application/json; charset=utf-8\nhost:%s\nx-tc-action:%s\n",
		host, strings.ToLower(action))
	signedHeaders := "content-type;host;x-tc-action"
	hashedRequestPayload := sha256hex(payload)
	canonicalRequest := fmt.Sprintf("%s\n%s\n%s\n%s\n%s\n%s",
		httpMethod, canonicalURI, canonicalQueryString, canonicalHeaders, signedHeaders, hashedRequestPayload)

	// Step 2: Create String to Sign
	algorithm := "TC3-HMAC-SHA256"
	credentialScope := fmt.Sprintf("%s/%s/tc3_request", date, service)
	hashedCanonicalRequest := sha256hex(canonicalRequest)
	stringToSign := fmt.Sprintf("%s\n%d\n%s\n%s",
		algorithm, timestamp, credentialScope, hashedCanonicalRequest)

	// Step 3: Calculate Signature
	secretDate := hmacsha256([]byte("TC3"+c.SecretKey), date)
	secretService := hmacsha256(secretDate, service)
	secretSigning := hmacsha256(secretService, "tc3_request")
	signature := hex.EncodeToString(hmacsha256(secretSigning, stringToSign))

	// Step 4: Create Authorization Header
	authorization := fmt.Sprintf("%s Credential=%s/%s, SignedHeaders=%s, Signature=%s",
		algorithm, c.SecretId, credentialScope, signedHeaders, signature)

	return authorization, date
}

func (c *Client) newRequest(ctx context.Context, req Request) (*http.Request, error) {
	// Convert request to map for JSON encoding
	data := make(map[string]interface{})
	vs, err := toValues(req)
	if err != nil {
		return nil, err
	}
	for k, v := range vs {
		if v != "" {
			// Try to convert to int if possible
			if intVal, err := strconv.Atoi(v); err == nil {
				data[k] = intVal
			} else {
				data[k] = v
			}
		}
	}

	// Encode request body as JSON
	payloadBytes, err := json.Marshal(data)
	if err != nil {
		return nil, err
	}
	payload := string(payloadBytes)

	// Generate timestamp and signature
	timestamp := time.Now().Unix()
	action := req.GetEndpoint()
	authorization, _ := c.sign(action, payload, timestamp)

	// Create HTTP request
	httpReq, err := http.NewRequestWithContext(ctx, http.MethodPost, c.BaseURL, bytes.NewReader(payloadBytes))
	if err != nil {
		return nil, err
	}

	// Set headers
	httpReq.Header.Set("Content-Type", "application/json; charset=utf-8")
	httpReq.Header.Set("Host", strings.TrimPrefix(c.BaseURL, "https://"))
	httpReq.Header.Set("Authorization", authorization)
	httpReq.Header.Set("X-TC-Action", action)
	httpReq.Header.Set("X-TC-Version", apiVersion)
	httpReq.Header.Set("X-TC-Timestamp", strconv.FormatInt(timestamp, 10))
	httpReq.Header.Set("X-TC-Nonce", strconv.Itoa(rand.Intn(1000000)))
	if c.Region != "" {
		httpReq.Header.Set("X-TC-Region", c.Region)
	}

	return httpReq, nil
}

// Do makes a request and returns the response.
func (c *Client) Do(ctx context.Context, req Request, resp Response) error {
	httpReq, err := c.newRequest(ctx, req)
	if err != nil {
		return err
	}

	result, err := http.DefaultClient.Do(httpReq)
	if err != nil {
		return err
	}
	defer result.Body.Close()

	body, err := io.ReadAll(result.Body)
	if err != nil {
		return err
	}

	if err = json.Unmarshal(body, resp); err != nil {
		return err
	}
	if err = resp.Error(); err != nil {
		return err
	}
	return nil
}

// CreateRecord creates a DNS record.
func (c *Client) CreateRecord(ctx context.Context, req *CreateRecordRequest) (*CreateRecordResponse, error) {
	resp := &CreateRecordResponse{}
	err := c.Do(ctx, req, resp)
	return resp, err
}

// DescribeRecordList describes DNS record list.
func (c *Client) DescribeRecordList(ctx context.Context, req *DescribeRecordListRequest) (*DescribeRecordListResponse, error) {
	resp := &DescribeRecordListResponse{}
	err := c.Do(ctx, req, resp)
	return resp, err
}

// DeleteRecord deletes a DNS record.
func (c *Client) DeleteRecord(ctx context.Context, req *DeleteRecordRequest) (*DeleteRecordResponse, error) {
	resp := &DeleteRecordResponse{}
	err := c.Do(ctx, req, resp)
	return resp, err
}

// UpdateRecord updates a DNS record.
func (c *Client) UpdateRecord(ctx context.Context, req *UpdateRecordRequest) (*UpdateRecordResponse, error) {
	resp := &UpdateRecordResponse{}
	err := c.Do(ctx, req, resp)
	return resp, err
}
