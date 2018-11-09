package client

import (
	"context"
	"github.com/dayan-be/access-service/proto"
)

type methodCallback func(context.Context, *access.PkgRsp) error

type Client struct {
	ses *Session
	methodMap map[string]methodCallback
}

func NewClient() *Client {
	ses, _ := NewSession(Config().Srv.Addr)
	c := &Client{
		ses:ses,
	}
	return c
}

func (c *Client) Run() {
	c.methodMap["login.aaaa"] = c.LoginResp
	c.ses.Run()
}

func (c *Client) Login(phoneNum, password string) error {
	c.ses.Auth(phoneNum, password, "")
}

func (c *Client) LoginResp(ctx context.Context, rsp *access.PkgRsp) error {

}