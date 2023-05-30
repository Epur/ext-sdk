package main

import (
	"errors"
	"fmt"
	"github.com/tangchen2018/eshop-sdk/model"
	"github.com/tangchen2018/eshop-sdk/pfc/epur"
)

type EpurTest struct {
	api *epur.Api
}

func main() {
	api := epur.New(
		new(model.Setting).
			SetKey("appxxxxxxxxxxxxxxxx").
			SetSecret("secret_xxxxxxxxxxxxx").
			SetServerUrl("http://dev.open.epur.cn").
			SetShopId("appxxxxxxxxxxxxx__&&@@%%__110").
			SetAuthCallbackUrl("https://www.baidu.com").
			SetAccessToken(`W8picM3Zs0YMA1g9KvJPgvP7kHbehwjkSTIDMUWLuaok9L5Ua7y999YUoalZ`),
	)

	testApi := EpurTest{api: api}
	testApi.RefreshToken()
}

func (p *EpurTest) GetAuthUrl() {
	result := p.api.GetAuthUrl("123")
	fmt.Println(result)
}

func (p *EpurTest) GetToken() {
	c := p.api.GetToken(model.BodyMap{"code": "PNYzImsQzisSBFzLY1ewE24p261IAcR6TqdKZDbGs0KrxC29ynVEydjAFmjp"})
	if c.Err != nil {
		panic(c.Err)
	}
	result := c.GetResponseTo().(epur.GetTokenResponse)
	fmt.Println(result)
}

func (p *EpurTest) RefreshToken() {
	c := p.api.RefreshToken(model.BodyMap{"refreshToken": "nzdvlgvCTXkNfvG5hVuyjIyE9DgOobVExQ0e5ecLztNbyHJxLgRZwIkrKnxz"})
	if c.Err != nil {
		panic(c.Err)
	}
	if !c.Response.Success {
		panic(errors.New(c.Response.Response.Message))
	}
	result := c.GetResponseTo().(epur.GetTokenResponse)
	fmt.Println(result)
}
