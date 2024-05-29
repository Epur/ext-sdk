package store

import (
	"errors"
	"fmt"
	"github.com/Epur/ext-sdk/api"
	"github.com/Epur/ext-sdk/model"
	"github.com/Epur/ext-sdk/utils"
	"log"
	"reflect"
	"runtime"
	"time"
)

func New() *Store {
	return &Store{JobChan: make(chan *Job, 2048), LoopWait: 60 * 2, TimeOut: 30}
}

func (p *Store) AddJob(token *Token) error {
	if len(token.Id) <= 0 {
		return errors.New("Id is void.")
	}
	if token.SecondsBeforeRefresh <= 0 {
		token.SecondsBeforeRefresh = 5 * 60 //默认5分钟
	}
	if len(token.Refresh.Key) <= 0 {
		return errors.New("Key is void.")
	}
	if len(token.Refresh.Secret) <= 0 {
		return errors.New("Secret is void.")
	}
	if len(token.Refresh.PlatformCode) <= 0 {
		return errors.New("PlatformCode is void.")
	}
	if len(token.Refresh.RefreshToken) <= 0 {
		return errors.New("RefreshToken is void.")
	}
	if token.Refresh.AccessTokenExpire <= 0 {
		return errors.New("AccessTokenExpire is void.")
	}

	if p != nil && p.JobChan != nil {
		p.JobChan <- &Job{Method: "add", Token: token}
	} else {
		log.Printf("指针未初始化，无法加入任务刷新Token，store.New()")
	}
	return nil
}

func (p *Store) RestartJob(token *Token) {
	p.JobChan <- &Job{Method: "add", Token: token}
}

func (p *Store) DelJob(token *Token) {
	p.ActionDel(token)
}

func (p *Store) ErrJob(token *Token) {
	//p.ActionDel(token)
	p.JobChan <- &Job{Method: "err", Token: token}
}

func (p *Store) ListenJob() {

	for job := range p.JobChan {

		switch job.Method {
		case "add":

			index := -1
			for i, j := range p.List {
				if j.Id == job.Token.Id {
					index = i
					break
				}
			}

			if index > -1 {
				p.List[index] = job.Token
			} else {
				p.List = append(p.List, job.Token)
			}
			p.List.Sort()
		case "err":

			index := -1
			for i, j := range p.ErrorList {
				if j.Id == job.Token.Id {
					index = i
					break
				}
			}
			if index > -1 {
				p.ErrorList[index] = job.Token
			} else {
				p.ErrorList = append(p.ErrorList, job.Token)
			}
			p.ErrorList.Sort()
		}
	}
}

func (p *Store) ActionDel(token *Token) {
	index := -1
	for i, j := range p.List {
		if j.Id == token.Id {
			index = i
			break
		}
	}
	if index > -1 && index < len(p.List) {
		p.List = append(p.List[:index], p.List[index+1:]...)
	}
}

func (p *Store) Listen() {

	go func() {
		p.ListenJob()
	}()

	go func() {
		for {
			if len(p.ErrorList) > 0 {
				log.Printf("Error Len[%d]", len(p.ErrorList))

				for _, item := range p.ErrorList {
					log.Printf("Error %s: Id [%s]",
						item.Refresh.PlatformCode, item.Id)
				}
			}

			time.Sleep(time.Duration(p.LoopWait) * time.Second)
		}
	}()

	go func() {
		for {

			if len(p.List) > 0 {
				t1 := p.List[0]

				log.Printf("%s: Id [%s] Trigger[%s] [%d-%d]",
					t1.Refresh.PlatformCode, t1.Id,
					time.Unix(t1.Refresh.AccessTokenExpire-t1.SecondsBeforeRefresh, 0).Format(
						"2006-01-02 15:04:05"),
					t1.Refresh.AccessTokenExpire, t1.SecondsBeforeRefresh)

				if t1.Refresh.AccessTokenExpire-t1.SecondsBeforeRefresh <= utils.TimestampSecond() {
					p.DelJob(t1)

					go func(t1 *Token) {
						e := &Event{Token: t1}
						p.RefreshRun(e)

						if e.Success {
							p.RestartJob(t1)
						} else {
							p.ErrJob(t1)
						}
					}(t1)

					continue
				}
			}

			time.Sleep(time.Duration(p.LoopWait) * time.Second)
		}
	}()
}

func (p *Store) RefreshRun(e *Event) {

	defer func() {
		if r := recover(); r != nil {
			const size = 64 << 10
			buf := make([]byte, size)
			buf = buf[:runtime.Stack(buf, false)]
			err, ok := r.(error)
			if !ok {
				err = fmt.Errorf("%v", r)
			}
			log.Println(err, "panic", "stack", "...\n"+string(buf))

			e.Success = false
			e.Msg = reflect.ValueOf(err).String()
		}
	}()

	c := api.New(e.Token.Refresh.PlatformCode, new(model.Setting).
		SetKey(e.Token.Refresh.Key).
		SetMerchantId(e.Token.Refresh.MerchantId).
		SetShopId(e.Token.Refresh.ShopId).
		SetIsMerchant(e.Token.Refresh.IsMerchant).
		SetServerUrl(e.Token.Refresh.ServerUrl).
		SetSecret(e.Token.Refresh.Secret)).StoreRefreshToken(model.BodyMap{
		"refresh_token": e.Token.Refresh.RefreshToken})

	if c == nil {
		e.Msg = fmt.Sprintf("平台[%s]不支持", e.Token.Refresh.PlatformCode)
		return
	}
	if c.Err != nil {
		e.Msg = c.Err.Error()
		return
	}
	if !c.Response.Success {
		e.Msg = c.Response.Response.Message
	} else {
		resp := c.Response.Response.DataTo.(model.StoreTokenResponse)
		e.Token.Refresh.AccessToken = resp.AccessToken
		e.Token.Refresh.AccessTokenExpire = resp.AccessTokenExpire
		e.Token.Refresh.RefreshToken = resp.RefreshToken
		e.Token.Refresh.RefreshTokenExpire = resp.RefreshTokenExpire

		e.Success = true
		e.Msg = "success"
	}

	go func(e *Event) {
		ch := make(chan string)

		go func(e *Event) {
			defer func() {
				if r := recover(); r != nil {
					const size = 64 << 10
					buf := make([]byte, size)
					buf = buf[:runtime.Stack(buf, false)]
					err, ok := r.(error)
					if !ok {
						err = fmt.Errorf("%v", r)
					}
					log.Println(err, "panic", "stack", "...\n"+string(buf))
				}
			}()
			e.Token.CallBack(e)
			ch <- "ok"
		}(e)

		select {
		case _ = <-ch:
			log.Printf("%s: ID [%s] 推送完成...", e.Token.Refresh.PlatformCode, e.Token.Id)
		case <-time.After(time.Second * time.Duration(p.TimeOut)):
			log.Printf("%s: ID [%s] 推送超时...", e.Token.Refresh.PlatformCode, e.Token.Id)
		}

		fmt.Println(2222)
	}(e)
}
