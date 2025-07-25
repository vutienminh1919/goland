package vietqr

import (
	"bytes"
	"encoding/json"
	"errors"
	"io"
	"net/http"
	"time"
)

type VietQr struct {
	ClientID string
	ApiKey   string
	BaseUrl  string
	Inputs   map[string]interface{}
	ApiUrl   string
	Response map[string]interface{}
}

func NewVietQr() *VietQr {
	return &VietQr{
		ClientID: "5e319e0e-eb5a-4cdc-bef9-8464474a7f46",
		ApiKey:   "ae2fdd49-88b6-4da0-b2a1-cb3381d13786",
		BaseUrl:  "https://api.vietqr.io/v2/",
	}
}

func (v *VietQr) Call(params map[string]interface{}, endpoint string, method string) (map[string]interface{}, error) {
	if method == "" {
		method = http.MethodGet
	}

	// Gán inputs
	v.Inputs = params

	// Tạo URL
	if endpoint != "" {
		v.ApiUrl = v.BaseUrl + endpoint
	} else {
		return nil, errors.New("endpoint không được để trống")
	}

	// Encode params thành JSON
	bodyBytes, err := json.Marshal(params)
	if err != nil {
		return nil, err
	}

	// Tạo request
	req, err := http.NewRequest(method, v.ApiUrl, bytes.NewBuffer(bodyBytes))
	if err != nil {
		return nil, err
	}

	// Header
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("x-client-id", v.ClientID)
	req.Header.Set("x-api-key", v.ApiKey)

	// Tạo HTTP client
	client := &http.Client{Timeout: 15 * time.Second}

	// Gửi request
	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	defer func(Body io.ReadCloser) {
		err := Body.Close()
		if err != nil {

		}
	}(resp.Body)

	// Đọc response body
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	// Kiểm tra status code
	if resp.StatusCode != http.StatusOK {
		return nil, errors.New("HTTP status: " + resp.Status)
	}

	// Decode JSON
	err = json.Unmarshal(body, &v.Response)
	if err != nil {
		return nil, err
	}

	return v.Response, nil
}
