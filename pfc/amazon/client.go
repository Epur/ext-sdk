package amazon

import (
	"encoding/json"
	"fmt"
	"github.com/Epur/ext-sdk/model"
	"github.com/Epur/ext-sdk/utils"
	"github.com/tangchen2018/go-utils/http"
	http2 "net/http"
	"runtime"
	"strings"
)

type Client struct {
	SignMethod string
	model.Client
}

func NewClient(setting *model.Setting) *Client {
	return &Client{Client: model.Client{
		Setting: setting}, SignMethod: "HmacSHA256"}
}

func (c *Client) Execute() {
	if c.Request.Method == nil || len(*c.Request.Method) < 0 {
		c.Request.Method = utils.PString(http.GET)
	}
	if c.Request.Body == nil {
		c.SetBody(make(model.BodyMap))
	}
	if c.Setting.Key == nil {
		c.Err = utils.Err("Key is null..")
		return
	}
	if c.Request.Path == nil {
		c.Err = utils.Err("Path is null..")
		return
	}

	c.HttpReq = http.New(
		http.WithUrl(c.urlParse()),
		http.WithMethod(*c.Request.Method),
	)
	c.HttpReq.Header.Set("User-Agent", fmt.Sprintf(
		"ext-sdk-amz/v1.0.0 (Language=%s; Platform=%s-%s)",
		strings.Replace(runtime.Version(), "go", "go/", -1), runtime.GOOS, runtime.GOARCH))

	if c.Setting.AccessToken != nil && *c.Request.Path != TokenURL {
		c.HttpReq.Header.Set("X-Amz-Access-Token", *c.Setting.AccessToken)
	}

	//if c.Err = c.signRequest(); c.Err != nil {
	//	return
	//}

	for key, value := range c.Request.Params {
		// 使用类型 switch
		switch v := value.(type) {
		//数组类型
		case []string:
			for _, vv := range v {
				c.HttpReq.QueryParams.Add(key, vv)
			}
		default:
			c.HttpReq.QueryParams.Set(key, value.(string))
		}
	}

	if strings.ToUpper(*c.Request.Method) == http2.MethodPost ||
		strings.ToUpper(*c.Request.Method) == http2.MethodPut {
		http.WithRequestType(http.TypeJSON)(c.HttpReq)
		c.HttpReq.Body = c.Request.Body
	}

	if c.Err = c.Client.Execute(); c.Err != nil {
		return
	}

	result := new(Response)
	_ = json.Unmarshal(c.HttpReq.Result, &result)

	code := "Success"
	msg := "success"
	if result.Errors != nil {
		code = result.Errors[0].Code
		msg = result.Errors[0].Message
	}

	c.Response.Response.Code = code
	c.Response.Response.Message = msg
	c.Response.Response.RequestId = ""
	c.Response.Response.Data = result.Payload

	if result.Errors == nil {
		c.Response.Success = true
	}

	if c.Response.HttpStatus != 200 {
		c.Response.Success = false
	}
}

//func (c *Client) signRequest() error {
//	timestamp := time.Now().UTC().Format("20060102T150405Z")
//	date := time.Now().UTC().Format("20060102")
//	c.HttpReq.Header.Set("X-Amz-Date", timestamp)
//
//	// Step 1: Create canonical request
//	canonicalURI := c.Request.Path
//	canonicalQueryString := c.Request.Params
//
//	var canonicalHeaders string
//	var signedHeaders string
//	var headers []string
//
//	for key := range c.HttpReq.Header {
//		headers = append(headers, strings.ToLower(key))
//	}
//	sort.Strings(headers)
//
//	for _, key := range headers {
//		canonicalHeaders += fmt.Sprintf("%s:%s\n", key, c.HttpReq.Header.Get(key))
//		signedHeaders += key + ";"
//	}
//	signedHeaders = strings.TrimSuffix(signedHeaders, ";")
//
//	payloadHash := c.getPayloadHash()
//	canonicalRequest := fmt.Sprintf("%s\n%s\n%s\n%s\n%s\n%s",
//		c.HttpReq.Method,
//		canonicalURI,
//		canonicalQueryString,
//		canonicalHeaders,
//		signedHeaders,
//		payloadHash,
//	)
//
//	// Step 2: Create string to sign
//	algorithm := "AWS4-HMAC-SHA256"
//	credentialScope := fmt.Sprintf("%s/%s/%s/%s",
//		date,
//		*c.Setting.SiteNo,
//		"execute-api",
//		"aws4_request",
//	)
//	stringToSign := fmt.Sprintf("%s\n%s\n%s\n%s",
//		algorithm,
//		timestamp,
//		credentialScope,
//		sha256Hex(canonicalRequest),
//	)
//
//	// Step 3: Calculate signature
//	signingKey := c.getSigningKey(date)
//	h := hmac.New(sha256.New, signingKey)
//	h.Write([]byte(stringToSign))
//	signature := hex.EncodeToString(h.Sum(nil))
//
//	// Step 4: Add authorization header
//	authorizationHeader := fmt.Sprintf("%s Credential=%s/%s, SignedHeaders=%s, Signature=%s",
//		algorithm,
//		*c.Setting.Key,
//		credentialScope,
//		signedHeaders,
//		signature,
//	)
//	c.HttpReq.Header.Set("Authorization", authorizationHeader)
//
//	return nil
//}
//
//func (c *Client) getPayloadHash() string {
//	var payload []byte
//	if c.HttpReq.Body != nil {
//		bodyBytes, _ := json.Marshal(c.HttpReq.Body)
//		payload, _ = ioutil.ReadAll(bytes.NewReader(bodyBytes))
//		json.Unmarshal(bodyBytes, &c.HttpReq.Body) // Reset the body after reading
//	}
//	h := sha256.New()
//	h.Write(payload)
//	return hex.EncodeToString(h.Sum(nil))
//}
//
//func sha256Hex(data string) string {
//	h := sha256.New()
//	h.Write([]byte(data))
//	return hex.EncodeToString(h.Sum(nil))
//}
//
//func (c *Client) getSigningKey(date string) []byte {
//	kDate := hmacSHA256([]byte("AWS4"+*c.Setting.Secret), date)
//	kRegion := hmacSHA256(kDate, *c.Setting.SiteNo)
//	kService := hmacSHA256(kRegion, "execute-api")
//	kSigning := hmacSHA256(kService, "aws4_request")
//	return kSigning
//}
//
//func hmacSHA256(key []byte, data string) []byte {
//	h := hmac.New(sha256.New, key)
//	h.Write([]byte(data))
//	return h.Sum(nil)
//}

func (c *Client) urlParse() string {
	urlPath := ""
	if strings.HasPrefix(*c.Request.Path, "http") {
		urlPath = *c.Request.Path
	} else {
		urlPath = c.getServerURL() + *c.Request.Path
	}
	return urlPath
}

func (c *Client) getServerURL() string {
	switch strings.ToUpper(*c.Setting.SiteNo) {
	case "US":
	case "MX":
	case "CA":
	case "BR":
		return APIGatewayUS
	case "ES":
	case "UK":
	case "FR":
	case "BE":
	case "NL":
	case "DE":
	case "IT":
	case "SE":
	case "ZA":
	case "PL":
	case "EG":
	case "TR":
	case "SA":
	case "AE":
	case "IN":
		return APIGatewayEU
	case "SG":
	case "AU":
	case "JP":
		return APIGatewayFE
	case "refresh":
		return TokenURL
	}
	return ""
}
