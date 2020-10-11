package options

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

type (
	userAgent           string
	apiKey              string
	publicKey           string
	timeout             uint
	origin              string
	verifySSL           bool
	debugging           bool
	serviceProviderCode string
	initiatorIdentifier string
	securityCredential  string
)

type ClientOption interface {
	Appl(c *ClientProperties)
}

type ClientProperties struct {
	ApiKey              string
	PublicKey           string
	Timeout             string
	AccessToken         string
	Origin              string
	VerifySSL           string
	Debugging           string
	Host                string
	Environment         Environment
	ServiceProviderCode string
	SecurityCredential  string
	InitiatorIdentifier string
	UserAgent           string
}

func WithApiKey(s string) ClientOption {
	return apiKey(s)
}

func WithUserAgent(s string) ClientOption {
	return userAgent(s)
}

func WithTimeout(u uint) ClientOption {
	return timeout(u)
}

func WithOrigin(s string) ClientOption {
	return origin(s)
}

func WithTimeout(u uint) ClientOption {
	return timeout(u)
}

func VerifySSL(b bool) ClientOption {
	return verifySSL(b)
}

func WithDebugging(b bool) ClientOption {
	return debugging(b)
}

func WithServiceProviderCode(s string) ClientOption {
	return serviceProviderCode(s)
}

func WithInitiatorIdentifier(s string) ClientOption {
	return initiatorIdentifier(s)
}

func WithSecurityCredential(s string) ClientOption {
	return securityCredential(s)
}

func (key apiKey) Apply(c *ClientProperties) {
	c.ApiKey = string(key)
}

func (ua userAgent) Apply(c *ClientProperties) {
	c.UserAgent = string(ua)
}
