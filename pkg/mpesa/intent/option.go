package intent

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

type Option interface {
	Apply(intent *Properties)
}

func WithDestination(s string) Option {
	return destination(s)
}

func WithTransactionId(s string) Option {
	return transactionId(s)
}

func WithSource(s string) Option {
	return source(s)
}

func WithAmount(f float64) Option {
	return amount(f)
}

func WithThirdPartyReference(s string) Option {
	return thirdPartyReference(s)
}

func WithTransactionReference(s string) Option {
	return transactionReference(s)
}

func WithQueryReference(s string) Option {
	return queryReference(s)
}

func (d destination) Apply(i *Properties) {
	i.Destination = string(d)
}

func (t transactionId) Apply(i *Properties) {
	i.TransactionId = string(t)
}

func (s source) Apply(i *Properties) {
	i.Source = string(s)
}

func (a amount) Apply(i *Properties) {
	i.Amount = float64(a)
}

func (r thirdPartyReference) Apply(i *Properties) {
	i.ThirdPartyReference = string(r)
}

func (t transactionReference) Apply(i *Properties) {
	i.TransactionReference = string(t)
}

func (q queryReference) Apply(i *Properties) {
	i.QueryReference = string(q)
}
