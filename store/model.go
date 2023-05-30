package store

import (
	"sort"
)

var TokenStore *Store

type Tokens []*Token

func (x Tokens) Len() int {
	return len(x)
}

func (x Tokens) Less(i, j int) bool {
	return x[i].Refresh.AccessTokenExpire < x[j].Refresh.AccessTokenExpire
}
func (x Tokens) Swap(i, j int) {
	x[i], x[j] = x[j], x[i]
}

func (x Tokens) Sort() {
	sort.Sort(x)
}

type Refresh struct {
	Key               string `json:"key"`               // Key
	Secret            string `json:"secret"`            // 密钥
	PlatformCode      string `json:"platformCode"`      // 平台
	RefreshToken      string `json:"refreshToken"`      // 刷新令牌
	AccessTokenExpire int64  `json:"accessTokenExpire"` // 访问令牌有效期(时间戳)
	IsMerchant        bool   `json:"isMerchant"`        // 是否商户
	ServerUrl         string `json:"serverUrl"`
	//下面2个参数只有返回的时候有 传入可以不传
	MerchantId         string `json:"merchantId"`
	ShopId             string `json:"shopId"`
	AccessToken        string `json:"accessToken"`        //刷新令牌
	RefreshTokenExpire int64  `json:"refreshTokenExpire"` // 刷新令牌有效期(时间戳)
}

type Token struct {
	Id                   string         `json:"id"`
	Refresh              Refresh        `json:"refresh"`
	Body                 []byte         `json:"-"`                    // 请求的附带参数 回调时原路返回
	SecondsBeforeRefresh int64          `json:"secondsBeforeRefresh"` // 到期提前多少秒刷新accessToken
	RetryCount           int64          `json:"retryCount"`           // 错误重试次数
	CallBack             func(e *Event) `json:"-"`                    // AccessToken 刷新回调
}

type Event struct {
	Success bool
	Msg     string
	Token   *Token
}

type Store struct {
	Map       map[string]*Token
	List      Tokens
	ErrorList Tokens
	JobChan   chan *Job
	LoopWait  int64 // 秒
	TimeOut   int64 // 超时时间 默认30s
}

type Job struct {
	Method string // add del
	Token  *Token
}
