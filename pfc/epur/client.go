package epur

import (
	"crypto/hmac"
	"crypto/sha256"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"github.com/google/uuid"
	"github.com/tangchen2018/eshop-sdk/model"
	"github.com/tangchen2018/eshop-sdk/utils"
	"github.com/tangchen2018/go-utils/http"
	"log"
	http2 "net/http"
	"sort"
	"strconv"
	"strings"
	"time"
)

type Client struct {
	model.Client
}

func NewClient(setting *model.Setting) *Client {
	return &Client{Client: model.Client{Setting: setting}}
}

func (p *Client) Execute() {

	p.Request.Method = utils.PString(http.POST)

	p.HttpReq = http.New(
		http.WithUrl(p.urlParse(p.Request.Path)),
		http.WithMethod(*p.Request.Method),
		http.WithRequestType(http.TypeJSON),
	)
	payload := p.requestPayload(p.Request.Body)

	p.HttpReq.Header.Set("host", strings.Replace(strings.Replace(*p.Setting.ServerUrl, "https://", "", -1), "http://", "", -1))
	p.HttpReq.Header.Set("user-agent", "AccOpenApiClient/1.0.2 (+https://www.epur.cn/)")
	p.HttpReq.Header.Set("x-acc-version", "1.0.2")
	p.HttpReq.Header.Set("x-acc-action", *p.Request.Path)
	p.HttpReq.Header.Set("x-acc-date", time.Now().Format("2006-01-02T15:04:05Z"))
	p.HttpReq.Header.Set("x-acc-signature-nonce", p.generateNonce())
	p.HttpReq.Header.Set("accept", "application/json")
	p.HttpReq.Header.Set("content-type", "application/json")
	p.HttpReq.Header.Set("x-acc-content-sha256", payload)

	if p.Setting.AccessToken != nil {
		p.HttpReq.Header.Set("x-acc-access-token", *p.Setting.AccessToken)
	}

	for key, item := range p.HttpReq.Header {
		fmt.Println(key, item)
	}
	log.Printf("%s", payload)

	p.HttpReq.Header.Set("s-authorization", p.sign(p.HttpReq.Header, payload))
	p.HttpReq.Body = p.Request.Body

	fmt.Println("s-authorization->", p.HttpReq.Header.Get("s-authorization"))

	if p.Err = p.Client.Execute(); p.Err != nil {
		return
	}

	result := new(Response)
	_ = json.Unmarshal(p.HttpReq.Result, &result)

	p.Response.Response.Code = fmt.Sprintf("%d", result.Code)
	p.Response.Response.Message = result.Message
	p.Response.Response.RequestId = ""
	p.Response.Response.Data = result.Data

	if result.Code == 200 {
		p.Response.Success = true
	}

	if p.Response.HttpStatus != 200 {
		p.Response.Success = false
	}
}

func (p *Client) urlParse(apiPath *string) string {

	if apiPath == nil {
		panic("apiPath is void...")
	}

	return fmt.Sprintf("%s%s", *p.Setting.ServerUrl, *apiPath)
}

func (p *Client) requestPayload(body model.BodyMap) string {

	if body != nil {
		return p.hexEncode(p.hash(body.JsonBody()))
	} else {
		return p.hexEncode(p.hash(""))
	}
}

func (p *Client) sign(headers http2.Header, payload string) string {

	canonicalURI := *p.Request.Path

	if utils.IsEmpty(canonicalURI) || strings.EqualFold(strings.Trim(canonicalURI, ""), "") {
		canonicalURI = "/"
	}
	method := p.HttpReq.Method
	cannoicalHeaders := p.canonicalizedHeadersMap(headers)
	signedHeaders := cannoicalHeaders["signedHeaders"]
	queryString := ""
	//sb := method
	var sb strings.Builder
	sb.WriteString(method)
	sb.WriteString("\n")
	sb.WriteString(canonicalURI)
	sb.WriteString("\n")
	sb.WriteString(queryString)
	sb.WriteString("\n")
	sb.WriteString(cannoicalHeaders["canonicalHeaders"])
	sb.WriteString("\n")
	sb.WriteString(signedHeaders)
	sb.WriteString("\n")
	sb.WriteString(payload)

	fmt.Println(*p.Setting.Secret)

	hex := p.hexEncode(p.hash(sb.String()))
	stringToSign := "ACS3-HMAC-SHA256\n" + hex
	signature := p.hexEncode(p.signatureHSHA256(stringToSign, *p.Setting.Secret))
	return "ACS3-HMAC-SHA256 Credential=" + *p.Setting.Key + ",SignedHeaders=" + signedHeaders + ",Signature=" + signature
}

func (p *Client) canonicalizedHeadersMap(headers http2.Header) map[string]string {
	result := make(map[string]string)
	prefix := "x-acc-"
	keys := make([]string, 0, len(headers))
	for val := range headers {
		keys = append(keys, val)
	}
	canonicalizedKeys := make([]string, 0)
	valueMap := make(map[string]string)
	for _, val := range keys {
		lowerKey := strings.ToLower(val)
		if strings.HasPrefix(lowerKey, prefix) || strings.EqualFold(lowerKey, "host") ||
			strings.EqualFold(lowerKey, "content-type") {
			canonicalizedKeys = append(canonicalizedKeys, lowerKey)
			valueMap[lowerKey] = headers.Get(val)
		}
	}
	sort.Strings(canonicalizedKeys)
	signedHeaders := strings.Join(canonicalizedKeys, ";")
	var sb string
	for _, val := range canonicalizedKeys {
		sb = sb + val
		sb = sb + ":"
		sb = sb + valueMap[val]
		sb = sb + "\n"
	}
	result["canonicalHeaders"] = sb
	result["signedHeaders"] = signedHeaders
	return result
}

func (p *Client) generateNonce() string {
	uid := uuid.New().String()
	uid = uid + strconv.FormatInt(time.Now().Unix(), 10)
	return uid
}

func (p *Client) hash(raw string) []byte {
	if raw == "" {
		return nil
	}
	hash := sha256.New()
	hash.Write([]byte(raw))
	return hash.Sum(nil)
}

// 二进制转HEX
func (p *Client) hexEncode(raw []byte) string {
	if raw == nil {
		return ""
	}
	return hex.EncodeToString(raw)
}

// HMAC-SHA256 加密
func (p *Client) signatureHSHA256(stringToSign, secret string) []byte {
	h := hmac.New(sha256.New, []byte(secret))
	h.Write([]byte(stringToSign))
	return h.Sum(nil)
}
