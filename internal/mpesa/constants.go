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

const (
	C2BPayment OperationCode = iota
	B2BPayment
	B2CPayment
	Reversal
	QueryTransactionStatus
)

const (
	Get HttpMethod = iota
	Post
	Put
	Delete
)

const (
	PatternPhoneNumber         = ""
	PatternServiceProviderCode = ""
	PatternWord                = ""
	PatternMoneyAmount         = ""
)

var (
	Operations = map[OperationCode]Operation{
		C2BPayment: &Operation{},

		B2BPayment: &Operation{},

		B2CPayment: &Operation{},

		Reversal: &Operation{},

		QueryTransactionStatus: &Operation{},
	}
)
