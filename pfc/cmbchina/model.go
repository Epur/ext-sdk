package cmbchina

import "encoding/json"

type Response struct {
	Code      string          `json:"code"`
	Message   string          `json:"message"`
	RequestID string          `json:"request_id"`
	Data      json.RawMessage `json:"data"`
	Result    json.RawMessage `json:"result"`
}
