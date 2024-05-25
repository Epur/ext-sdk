package api

import (
	"errors"
	"github.com/Epur/ext-sdk/model"
	"github.com/Epur/ext-sdk/pfc/amazon"
	"github.com/Epur/ext-sdk/pfc/epur"
	"github.com/Epur/ext-sdk/pfc/go1688"
	"github.com/Epur/ext-sdk/pfc/lazada"
	"github.com/Epur/ext-sdk/pfc/shopee"
	"github.com/Epur/ext-sdk/pfc/tiktokv2"
)

type PfcApi interface {
	StoreRefreshToken(Body model.BodyMap) *model.Client
}

type Api struct {
	Pfc     string
	Setting *model.Setting
	api     PfcApi
}

func New(pfc string, setting *model.Setting) *Api {

	api := &Api{Pfc: pfc, Setting: setting}

	switch pfc {
	case model.PFC_TIKTOK:
		api.api = tiktok.New(setting)
		//api.api = tiktokv2.New(setting)  去掉旧版本  逐渐使用最新版本
	case model.PFC_EPUR:
		api.api = epur.New(setting)
	case model.PFC_LAZADA:
		api.api = lazada.New(setting)
	case model.PFC_SHOPEE:
		api.api = shopee.New(setting)
	case model.PFC_AMAZON:
		api.api = amazon.New(setting)
	case model.PFC_1688:
		api.api = go1688.New(setting)
	default:
		panic(errors.New("平台不支持!"))
	}
	return api
}

func (p *Api) StoreRefreshToken(Body model.BodyMap) *model.Client {
	return p.api.StoreRefreshToken(Body)
}
