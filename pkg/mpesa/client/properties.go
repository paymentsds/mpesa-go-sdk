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

// Properties export
type Properties struct {
	// ApiKey export
	ApiKey string

	// PublicKey export
	PublicKey           string
	ReadTimeout         uint
	ConnectionTimeout   uint
	AccessToken         string
	Origin              string
	VerifySSL           bool
	Debugging           bool
	Host                string
	Environment         env.Environment
	ServiceProviderCode string
	SecurityCredential  string
	InitiatorIdentifier string
	UserAgent           string
}
