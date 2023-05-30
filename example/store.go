package main

import (
	"github.com/Epur/ext-sdk/model"
	"github.com/Epur/ext-sdk/store"
	"github.com/Epur/ext-sdk/utils"
	"log"
)

func main() {
	s1 := store.New()
	s1.LoopWait = 3
	s1.Listen()

	if err := s1.AddJob(&store.Token{
		Id: "=",
		CallBack: func(e *store.Event) {
			log.Println(e.Success, e.Msg, utils.ToJson(e.Token))
		},
		Refresh: store.Refresh{
			Key:               "appxxxxxxxxxxxxxxxx",
			Secret:            "xxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxx",
			PlatformCode:      model.PFC_EPUR,
			RefreshToken:      "",
			AccessTokenExpire: 1684066892,
		},
	}); err != nil {
		panic(err.Error())
	}

	<-make(chan bool)
}
