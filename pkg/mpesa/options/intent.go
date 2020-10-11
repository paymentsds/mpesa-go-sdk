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
	mpesa "github.com/paymentsds/mpesa-go-sdk/pkg/mpesa"
)

type (
	to          string
	from        string
	amount      float64
	reference   string
	transaction string
	subject     string
)

type IntentOption interface {
	Apply(intent *mpesa.Intent)
}

func To(s string) IntentOption {
	return to(s)
}

func From(s string) IntentOption {
	return from(s)
}

func WithAmount(f float64) IntentOption {
	return amount(f)
}

func WithReference(s string) IntentOption {
	return reference(s)
}

func WithTransaction(s string) IntentOption {
	return transaction(s)
}

func WithSubject(s string) IntentOption {
	return subject(s)
}

func (t to) Apply(i *IntentOption) {
	i.To = string(t)
}

func (f from) Apply(i *IntentOption) {
	i.From = string(f)
}

func (a amount) Apply(i *IntentOption) {
	i.Amount = float64(a)
}

func (r reference) Apply(i *IntentOption) {
	i.Reference = string(r)
}

func (t transaction) Apply(i *IntentOption) {
	i.Transaction = string(t)
}

func (s subject) Apply(i *IntentOption) {
	i.Subject = string(s)
}
