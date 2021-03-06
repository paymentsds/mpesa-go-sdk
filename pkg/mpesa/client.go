package mpesa

// Copyright 2020 Paymentds Developers

// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at

//     http://www.apache.org/licenses/LICENSE-2.0

// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

import (
	client "github.com/paymentsds/mpesa-go-sdk/pkg/mpesa/client"
	// "github.com/paymentsds/mpesa-go-sdk/pkg/mpesa/environment"
)

type Client struct {
	properties client.Properties
}

func NewClient(clientOptions ...client.Option) *Client {

	client := new(Client)

	for _, clientOption := range clientOptions {
		clientOption.Apply(&client.properties)
	}

	if client.properties.UserAgent == "" {
		client.properties.UserAgent = "Paymentsds Go/0.1.0"
	}

	if client.properties.Origin == "" {
		client.properties.Origin = "*"
	}

	return client
}

// func (c *Client) Send(intent Intent) (Result, error) {
// 	return c.service.HandleSend(intent)
// }

// func (c *Client) Receive(intent Intent) (Result, error) {
// 	return c.service.HandleReceive(intent)
// }

// func (c *Client) Revert(intent Intent) (Result, error) {
// 	return c.service.HandleRevert(intent)
// }

// func (c *Client) Query(intent Intent) (Result, error) {
// 	return c.service.HandleQuery(intent)
// }
