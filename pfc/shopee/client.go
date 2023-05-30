package shopee

import (
	"bytes"
	"crypto/hmac"
	"crypto/sha256"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"github.com/Epur/ext-sdk/model"
	"github.com/Epur/ext-sdk/utils"
	"github.com/tangchen2018/go-utils/http"
	"strings"
	"time"
)

type Client struct {
	timestamp string
	model.Client
	PathType string
}

func NewClient(setting *model.Setting) *Client {
	return &Client{Client: model.Client{Setting: setting}}
}

func (p *Client) Execute() {

	p.PathType = p.getPathType()

	if p.PathType == PATH_TYPE_MERCHANT && p.Client.Setting.MerchantId == nil {
		p.Err = utils.Err("merchantId is null..")
		return
	} else if p.PathType == PATH_TYPE_SHOP && p.Client.Setting.ShopId == nil {
		p.Err = utils.Err("shopId is null..")
		return
	}

	if p.Request.Method == nil {
		p.Request.Method = utils.PString("GET")
	}

	if p.Request.Params == nil {
		p.Request.Params = model.BodyMap{}
	}

	p.timestamp = fmt.Sprintf("%d", time.Now().Unix())

	p.Request.Params.Set("timestamp", p.timestamp).
		Set("partner_id", *p.Setting.Key)

	if p.PathType != PATH_TYPE_PUBLIC {
		if p.Client.Setting.AccessToken != nil {
			p.Request.Params.Set("access_token", *p.Client.Setting.AccessToken)
		}

		if p.Client.Setting.ShopId != nil {
			p.Request.Params.Set("shop_id", *p.Client.Setting.ShopId)
		}

		if p.Client.Setting.MerchantId != nil {
			p.Request.Params.Set("merchant_id", *p.Client.Setting.MerchantId)
		}
	}

	p.Request.Params.Set("sign", p.sign())

	p.HttpReq = http.New(
		http.WithUrl(p.urlParse()),
		http.WithMethod(*p.Request.Method),
	)

	for key, value := range p.Request.Params {
		p.HttpReq.QueryParams.Set(key, value.(string))
	}

	if *p.Client.Request.Method != http.GET {
		http.WithRequestType(http.TypeJSON)(p.HttpReq)
		p.HttpReq.Body = p.Request.Body
	}

	if p.Err = p.Client.Execute(); p.Err != nil {
		return
	}

	result := new(Response)
	_ = json.Unmarshal(p.HttpReq.Result, &result)

	p.Response.Response.Code = result.Error
	p.Response.Response.Message = result.Message
	p.Response.Response.RequestId = result.RequestId
	p.Response.Response.Data = result.Response

	if len(result.Error) <= 0 {
		p.Response.Success = true
	}

	if p.Response.HttpStatus != 200 {
		p.Response.Success = false
	}
}

func (p *Client) sign() string {

	/*
		签名
	*/

	var baseString bytes.Buffer
	baseString.WriteString(*p.Client.Setting.Key)
	baseString.WriteString(*p.Request.Path)
	baseString.WriteString(p.timestamp)

	if p.PathType == PATH_TYPE_MERCHANT {
		baseString.WriteString(*p.Client.Setting.AccessToken)
		baseString.WriteString(*p.Client.Setting.MerchantId)
	} else if p.PathType == PATH_TYPE_SHOP {
		baseString.WriteString(*p.Client.Setting.AccessToken)
		baseString.WriteString(*p.Client.Setting.ShopId)
	}

	hash := hmac.New(sha256.New, []byte(*p.Client.Setting.Secret))
	hash.Write(baseString.Bytes())
	return hex.EncodeToString(hash.Sum(nil))
}

func (p *Client) urlParse() string {

	if p.Client.Request.Path == nil {
		panic("apiPath is void...")
	}
	return fmt.Sprintf("%s%s", *p.Setting.ServerUrl, *p.Client.Request.Path)
}

func (p *Client) getPathType() string {
	// 0-Merchant 1-Shop 2-Public
	switch *p.Request.Path {
	case AUTH_ACCESSTOKEN, AUTH_PARTNER, AUTH_REFRESHTOKEN:
		return PATH_TYPE_PUBLIC
	case MERCHAT_URL:
		return PATH_TYPE_MERCHANT
	default:
		return PATH_TYPE_SHOP
	}
}

func (p *Client) getOrderFields() string {
	return strings.Join([]string{
		"item_list",
		"order_sn",
		"total_amount",
		"payment_method",
		"invoice_data",
		"estimated_shipping_fee",
		"actual_shipping_fee",
		"pay_time",
		"recipient_address",
		"dropshipper_phone",
		"fulfillment_flag",
		"pickup_done_time",
		"invoice_data",
		"checkout_shipping_carrier",
		"package_list"}, ",")
}
