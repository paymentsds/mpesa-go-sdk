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
	environment "github.com/paymentsds/mpesa-go-sdk/pkg/mpesa/environment"
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
	environment         Environment
}

func (c *Configuration) Headers() map[string][]string {
	c.generateToken()

	return map[string][]string{
		"Origin": []string{
			c.origin,
		},
		"User-Agent": []string{
			c.userAgent,
		},
		"Authorization": []string{
			fmt.Sprintf("Bearer %s", c.accessToken),
		},
		"Content-Type": []string{
			"application/json",
		}
	}
}

func (c *Configuration) generateToken() {
	// TODO
}

func (c *Configuration) HasValidHost() bool {
	if u, err := url.Parse(c.host); err == nil {
		return u.isAbs()
	}

	return false
}

func (c *Configuration) HasToken() bool {
	return c.accessToken != "" && c.accessToken != nil
}
