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
	"net/http"
)

var (
	UserAgent = fmt.Sprintf("Paymentsds Go/%s", Version)
)

const (
	C2BPayment OperationCode = iota
	B2BPayment
	B2CPayment
	Reversal
	QueryTransactionStatus
)

const (
	PatternPhoneNumber         = "^((00|\\+)?258)?8[45][0-9]{7}$"
	PatternServiceProviderCode = "^[0-9]{6}$"
	PatternWord                = "^\\w{20}$"
	PatternMoneyAmount         = "^[0-9]+(\\.[0-9]+)?$"
)

var (
	Operations = map[OperationCode]Operation{
		C2BPayment: &Operation{
			method:   http.MethodPost,
			path:     "",
			port:     "",
			required: []string{},
			optional: []string{},
			mapping:  map[string]string{},
			rules:    map[string]string{},
		},

		B2BPayment: &Operation{
			method:   http.MethodPost,
			path:     "",
			port:     "",
			required: []string{},
			optional: []string{},
			mapping:  map[string]string{},
			rules:    map[string]string{},
		},

		B2CPayment: &Operation{
			method:   http.MethodPost,
			path:     "",
			port:     "",
			required: []string{},
			optional: []string{},
			mapping:  map[string]string{},
			rules:    map[string]string{},
		},

		Reversal: &Operation{
			method:   http.MethodPut,
			path:     "",
			port:     "",
			required: []string{},
			optional: []string{},
			mapping:  map[string]string{},
			rules:    map[string]string{},
		},

		QueryTransactionStatus: &Operation{
			method:   http.MethodGet,
			path:     "",
			port:     "",
			required: []string{},
			optional: []string{},
			mapping:  map[string]string{},
			rules:    map[string]string{},
		},
	}
)
