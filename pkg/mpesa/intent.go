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
	intent "github.com/paymentsds/mpesa-go-sdk/pkg/mpesa/intent"
)

type Intent struct {
	properties intent.Properties
}

func NewIntent(intentOptions ...intent.Option) *Intent {
	i := new(Intent)

	for _, option := range intentOptions {
		option.Apply(&i.properties)
	}

	return i
}

//
