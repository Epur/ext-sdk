package cmbchina

import (
	"github.com/Epur/ext-sdk/model"
)

type Client struct {
	model.Client
}

func NewClient(setting *model.Setting) *Client {
	return &Client{Client: model.Client{Setting: setting}}
}

func (p *Client) Execute() {

}
