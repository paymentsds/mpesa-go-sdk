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
	errors "github.com/paymentsds/mpesa-go-sdk/pkg/mpesa/errors"
	"io/ioutil"
	"net/http"
	"time"
)

type Service struct {
	config Configuration
}

func NewService(options ClientOptions) *Service {
	return &Service{
		config: NewConfiguration(options),
	}
}

func (s *Service) HandleSend(intent Intent) (Result, error) {
	opcode := s.detectOperation(intent)
	return s.handleRequest(opcode, intent)
}

func (s *Service) HandleReceive(intent Intent) (Result, error) {
	return s.handleRequest(C2BPayment, intent)
}

func (s *Service) HandleRevert(intent Intent) (Result, error) {
	return s.handleRequest(Reversal, intent)
}

func (s *Service) HandleQuery(intent Intent) (Result, error) {
	return s.handleRequest(QueryTransactionStatus, intent)
}

func (s *Service) handleRequest(opcode OperationCode, intent Intent) (Result, error) {
	operation, ok := Operations[opcode]

	if !ok {
		// TODO
	}

	intent := s.fillOptionalValues(operation, intent)

	if missingProperties, err := operation.detectMissingValues(intent); err != nil {
		// TODO
	}

	if invalidProperties, err := operation.validateValues(intent); err != nil {
		// TODO
	}

	return s.performRequest(operation, intent)
}

func (s *Service) detectOperation(intent Intent) (OperationCode, error) {
	// TODO
}

func (s *Service) fillOptionalValues(opcode Operationcode, intent Intent) Intent {
	// TODO
}

func (s *Service) performRequest(operation *Operation, intent Intent) (Result, error) {
	s.config.GenerateAccessToken()

	if s.config.HasValidHost() {
		if s.config.HasToken() {
			headers := s.config.Headers()
			body := s.requestBody()

			httpClient = &http.Client{
				Timeout: s.config.Timeout * time.Second,
			}

			req := &http.Request{
				Method:  opetation.Method(),
				URL:     operation.URL(s.config),
				Headers: s.config.Headers(),
			}

			if req.Method == MethodGet {
				// TODO
			} else {
				// TODO
			}

			res, err := httpClient.Do(req)
			if err != nil {
				// TODO
			}

			data, err := ioutil.ReadAll(res.Body)
			if err != nil {
				// TODO
			}

			res.Body.Close()

			output := make(map[string]string)
			json.Unmarshall(output, &data)

			return NewResult(output), nil
		} else {
			return nil, errors.NewAuthenticationError()
		}
	} else {
		return nil, errors.NewInvalidHostError()
	}
}

func (s *Service) requestBody(operation *Operation, intent Intent) map[string]string {
	// TODO
}
