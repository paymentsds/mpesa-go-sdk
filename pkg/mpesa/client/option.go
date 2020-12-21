package client

import env "github.com/paymentsds/mpesa-go-sdk/pkg/mpesa/environment"

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

// Option export
type Option interface {
	Apply(c *Properties)
}

// WithApiKey export
// Option export
func WithApiKey(s string) Option {
	return apiKey(s)
}

func WithHost(s string) Option {
	return host(s)
}

func WithEnvironment(e env.Environment) Option {
	return environment(e)
}

func WithPublicKey(s string) Option {
	return publicKey(s)
}

func WithUserAgent(s string) Option {
	return userAgent(s)
}

func WithOrigin(s string) Option {
	return origin(s)
}

func WithReadTimeout(u uint) Option {
	return readTimeout(u)
}

func WithConnectionTimeout(u uint) Option {
	return connectionTimeout(u)
}

func VerifySSL(b bool) Option {
	return verifySSL(b)
}

func WithDebugging(b bool) Option {
	return debugging(b)
}

func WithServiceProviderCode(s string) Option {
	return serviceProviderCode(s)
}

func WithInitiatorIdentifier(s string) Option {
	return initiatorIdentifier(s)
}

func WithSecurityCredential(s string) Option {
	return securityCredential(s)
}

func (key apiKey) Apply(c *Properties) {
	c.ApiKey = string(key)
}

func (pb publicKey) Apply(c *Properties) {
	c.PublicKey = string(pb)
}

func (e environment) Apply(c *Properties) {
	c.Environment = env.Environment(e)
}

func (h host) Apply(c *Properties) {
	c.Host = string(h)
}

func (ua userAgent) Apply(c *Properties) {
	c.UserAgent = string(ua)
}

func (t readTimeout) Apply(c *Properties) {
	c.ReadTimeout = uint(t)
}

func (t connectionTimeout) Apply(c *Properties) {
	c.ConnectionTimeout = uint(t)
}

func (o origin) Apply(c *Properties) {
	c.Origin = string(o)
}

func (v verifySSL) Apply(c *Properties) {
	c.VerifySSL = bool(v)
}

func (d debugging) Apply(c *Properties) {
	c.Debugging = bool(d)
}

func (s serviceProviderCode) Apply(c *Properties) {
	c.ServiceProviderCode = string(s)
}

func (i initiatorIdentifier) Apply(c *Properties) {
	c.InitiatorIdentifier = string(i)
}

func (s securityCredential) Apply(c *Properties) {
	c.SecurityCredential = string(s)
}
