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
	"testing"

	client "github.com/paymentsds/mpesa-go-sdk/pkg/mpesa/client"
	"github.com/paymentsds/mpesa-go-sdk/pkg/mpesa/environment"
)

func TestClientPublicKey(t *testing.T) {
	publicKey := "X"

	client := NewClient(
		client.WithPublicKey(publicKey),
	)

	if publicKey != client.properties.PublicKey {
		t.Errorf("Failed to initialize publicKey")
	}
}

func TestClientEnvironment(t *testing.T) {

	client := NewClient()

	if client.properties.Environment != environment.Sandbox {
		t.Errorf("Failed to initialize Environment")
	}
}

func TestClientUserAgent(t *testing.T) {
	userAgent := "X"

	client := NewClient(
		client.WithUserAgent(userAgent),
	)

	if userAgent != client.properties.UserAgent {
		t.Errorf("Failed to initialize User Agent")
	}
}
