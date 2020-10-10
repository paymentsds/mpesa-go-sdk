package mpesa

import (
	options "github.com/paymentsds/mpesa-go-sdk/pkg/mpesa/options"
)

type Client struct {
	service Service
}

func NewClient(options *options.ClientOptions) Client {
	return &Client{
		service: NewService(options)
	}
}

func (c *Client) Send(intent Intent) (Result, error) {
	return c.service.HandleSend(intent)
}

func (c *Client) Receive(intent Intent) (Result, error) {
	return c.service.HandleReceive(intent)
}

func (c *Client) Revert(intent Intent) (Result, error) {
	return c.service.HandleRevert(intent)
}

func (c *Client) Query(intent Intent) (Result, error) {
	return c.service.HandleQuery(intent)
}
