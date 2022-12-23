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
	"fmt"
	environment "github.com/paymentsds/mpesa-go-sdk/pkg/mpesa/env"
	"net/url"
)

type Configuration struct {
	timeout             int
	apiKey              string
	publickey           string
	accessToken         string
	securityCredential  string
	initiatorIdentifier string
	verifySSL           string
	serviceProviderCode string
	origin              string
	userAgent           string
	host                string
	environment         environment.Environment
}

func (c *Configuration) Headers() map[string][]string {
	c.generateToken()

	return map[string][]string{
		"Origin": {
			c.origin,
		},
		"User-Agent": {
			c.userAgent,
		},
		"Authorization": {
			fmt.Sprintf("Bearer %s", c.accessToken),
		},
		"Content-Type": {
			"application/json",
		},
	}
}

func (c *Configuration) generateToken() {
	tok := NewTokenGenerator()
	tok.SetApiKey(c.apiKey)
	tok.SetPublicKey(c.publickey)
	err := tok.GenerateAccessToken()
	if err != nil {

		return
	}
	c.accessToken = tok.GetAccessToken()
}

func (c *Configuration) HasValidHost() bool {
	if u, err := url.Parse(c.host); err == nil {
		return u.IsAbs()
	}

	return false
}

func (c *Configuration) HasToken() bool {
	return c.accessToken != "" && &c.accessToken != nil
}
