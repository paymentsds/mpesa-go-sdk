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
	options "github.com/paymentsds/mpesa-go-sdk/pkg/mpesa/options"
)

type Intent struct {
	To          string
	From        string
	Reference   string
	Transaction string
	Subject     string
}

func NewIntent(opts ...options.IntentOption) *Intent {
	intent := make(Intent)

	for _, option := range opts {
		option.Apply(intent)
	}

	return intent
}
