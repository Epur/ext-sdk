package go1688

import "encoding/json"

const (
	AccessTokenURL = "/http/1/system.oauth2/getToken"
	AuthURL        = "https://auth.1688.com/oauth/authorize"
)

type Response struct {
	Error            string          `json:"error"`
	ErrorDescription string          `json:"error_description"`
	ErrorCode        string          `json:"error_code"`
	ErrorMessage     string          `json:"error_message"`
	RequestId        string          `json:"request_id"`
	Success          *bool           `json:"success"`
	Message          string          `json:"message"`
	Data             json.RawMessage `json:"data"`
}
