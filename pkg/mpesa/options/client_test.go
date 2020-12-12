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

import "testing"

func TestWithApiKey(t *testing.T) {
	key := "ApiKey"
	o := WithApiKey(key)
	if k, ok := o.(apiKey); ok && string(k) != key {
		t.Errorf("different than expected")
	}
}

func TestWithUserAgent(t *testing.T) {
	subject := "UserAgent"
	object := WithUserAgent(subject)
	if k, ok := object.(userAgent); ok && string(k) != subject {
		t.Errorf("different than expected")
	}
}

func TestWithTimeout(t *testing.T) {
	subject := uint(5)
	object := WithTimeout(subject)
	if k, ok := object.(timeout); ok && uint(k) != subject {
		t.Errorf("different than expected")
	}
}

func TestWithOrigin(t *testing.T) {
	subject := "*"
	object := WithApiKey(subject)
	if k, ok := object.(origin); ok && string(k) != subject {
		t.Errorf("different than expected")
	}
}

func TestVerifySSL(t *testing.T) {
	subject := true
	object := VerifySSL(subject)
	if k, ok := object.(verifySSL); ok && bool(k) != subject {
		t.Errorf("different than expected")
	}
}

func TestWithDebbuging(t *testing.T) {
	subject := true
	object := WithDebugging(subject)
	if k, ok := object.(verifySSL); ok && bool(k) != subject {
		t.Errorf("different than expected")
	}
}

func TestWithServiceProviderCode(t *testing.T) {
	subject := "*"
	object := WithServiceProviderCode(subject)
	if k, ok := object.(serviceProviderCode); ok && string(k) != subject {
		t.Errorf("different than expected")
	}
}

func TestWithInitiatorIdentifier(t *testing.T) {
	subject := "*"
	object := WithInitiatorIdentifier(subject)
	if k, ok := object.(initiatorIdentifier); ok && string(k) != subject {
		t.Errorf("different than expected")
	}
}

func TestWithSecurityCredential(t *testing.T) {
	subject := "*"
	object := WithSecurityCredential(subject)
	if k, ok := object.(securityCredential); ok && string(k) != subject {
		t.Errorf("different than expected")
	}
}

func TestApplyForApiKey(t *testing.T) {
	keyToUse := "K"
	newApikey := WithApiKey(keyToUse)
	properties := ClientProperties{}
	newApikey.Apply(&properties)

	if properties.ApiKey != keyToUse {
		t.Errorf("different than expected")
	}
}

func TestApplyForUserAgent(t *testing.T) {
	userAgentProp := "K"
	newUserAgent := WithUserAgent(userAgentProp)
	properties := ClientProperties{}
	newUserAgent.Apply(&properties)

	if properties.UserAgent != userAgentProp {
		t.Errorf("different than expected")
	}
}

func TestApplyForTimeout(t *testing.T) {
	timeoutProp := uint(0)
	newTimeout := WithTimeout(timeoutProp)
	properties := ClientProperties{}
	newTimeout.Apply(&properties)

	if properties.Timeout != timeoutProp {
		t.Errorf("different than expected")
	}
}

func TestApplyForOrigin(t *testing.T) {
	originProp := "*"
	newOrigin := WithOrigin(originProp)
	properties := ClientProperties{}
	newOrigin.Apply(&properties)

	if properties.Origin != originProp {
		t.Errorf("different than expected")
	}
}
