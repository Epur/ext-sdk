package model

import (
	"encoding/json"
	"github.com/Epur/ext-sdk/logger"
	"github.com/Epur/ext-sdk/utils"
	"github.com/tangchen2018/go-utils/http"
)

type Response struct {
	Success    bool // 是否成功
	HttpStatus int  // http状态码
	Response   struct {
		RequestId string          // 请求ID
		Code      string          // 平台code信息
		Message   string          // 平台msg信息
		Data      json.RawMessage // 数据
		Result    json.RawMessage // Lazada业务数据
		DataTo    interface{}     // 结构转化后的数据
	}
}

func (p *Response) ToMap() BodyMap {
	result := make(BodyMap)
	_ = json.Unmarshal(p.Response.Data, &result)
	return result
}

func (p *Response) To(row interface{}) error {
	if p.Response.Data != nil {
		return json.Unmarshal(p.Response.Data, row)
	}
	return nil
}

type Request struct {
	Path      *string
	Method    *string
	Params    BodyMap
	Body      BodyMap
	Protected BodyMap
}

type Client struct {
	Request  Request
	Response Response
	Setting  *Setting
	Err      error
	HttpReq  *http.Request
}

func (p *Client) Execute() error {
	logger.SdkLogger.Infof("请求URL:%s\n请求报文:%v", p.HttpReq.Url, utils.ToJson(p.HttpReq.Body))

	p.Response.Success = false
	if err := p.HttpReq.Do(); err != nil {
		logger.SdkLogger.Error("ERROR:", err.Error())
		return err
	} else {
		logger.SdkLogger.Infof("响应报文:%v", string(p.HttpReq.Result))
		//fmt.Println("Response:", p.Response)
		p.Response.HttpStatus = p.HttpReq.Response.StatusCode
	}

	return nil
}

func (c *Client) GetResponseTo() interface{} {
	return c.Response.Response.DataTo
}

func (c *Client) SetPath(data string) *Client {
	c.Request.Path = &data
	return c
}

func (c *Client) SetMethod(data string) *Client {
	c.Request.Method = &data
	return c
}

func (c *Client) SetParams(data BodyMap) *Client {
	c.Request.Params = data
	return c
}

func (c *Client) SetBody(data BodyMap) *Client {
	c.Request.Body = data
	return c
}

func (c *Client) SetProtected(data BodyMap) *Client {
	c.Request.Protected = data
	return c
}
