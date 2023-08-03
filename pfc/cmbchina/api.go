package cmbchina

import "github.com/Epur/ext-sdk/model"

type Api struct {
	Setting *model.Setting
}

func New(setting *model.Setting) *Api {
	return &Api{Setting: setting}
}

/*
	可经办业务模式查询DCLISMOD
	Repsonse : DCLISMODResponse
*/

func (p *Api) DCLISMOD(Body model.BodyMap) *model.Client {

	c := NewClient(p.Setting)
	c.SetPath(`DCLISMOD`).
		SetMethod("POST").
		SetBody(Body)

	if c.Err = Body.CheckEmptyError("buscod"); c.Err != nil {
		return &c.Client
	}

	c.Execute()
	if c.Err != nil {
		return &c.Client
	}

	response := DCLISMODResponse{}
	if c.Err = c.Client.Response.To(&response); c.Err != nil {
		return &c.Client
	}
	c.Response.Response.DataTo = response
	return &c.Client
}
