package tikbc

import (
	"github.com/T4t4KAU/tikbc/errno"
	"github.com/T4t4KAU/tikbc/iface"
	"github.com/T4t4KAU/tikbc/resp"
	"github.com/T4t4KAU/tikbc/tiko"
	"github.com/T4t4KAU/tikbc/utils"
	"strings"
)

var protos = map[string]struct{}{
	"tiko": {},
	"resp": {},
}

func newClient(proto string, addr string) (iface.Client, error) {
	proto = strings.ToLower(proto)

	switch proto {
	case "tiko":
		return tiko.NewClient(addr)
	case "resp":
		return resp.NewClient(addr)
	default:
		panic("invalid protocol")
	}
}

type Client struct {
	Protocol string
	Addr     string
	client   iface.Client
}

func New(addr, proto string) (*Client, error) {
	if !utils.ValidateAddress(addr) {
		return &Client{}, errno.ErrInvalidAddress
	}
	if _, ok := protos[proto]; !ok {
		return &Client{}, errno.ErrInvalidProtocol
	}

	cli, err := newClient(proto, addr)
	if err != nil {
		return nil, err
	}

	return &Client{
		Protocol: proto,
		Addr:     addr,
		client:   cli,
	}, nil
}

func (c *Client) Set(key, value string) error {
	return c.client.Set(key, value)
}

func (c *Client) Get(key string) (string, error) {
	return c.client.Get(key)
}

func (c *Client) Del(key string) error {
	return c.client.Del(key)
}

func (c *Client) Expire(key string, ttl int64) error {
	return c.client.Expire(key, ttl)
}
