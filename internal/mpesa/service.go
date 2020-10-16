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
	"reflect"
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

			if req.Method == http.MethodGet {
				formData := make(map[string][]string{})

				for k, v := range body {
					formData[k] = {v}
				}

				req.Form = formData
			} else {				
				if jsonBody, err := json.Marshal(body); err == nil {
					req.Body = jsonBody
				}

				return nil, new(ValidationError)
			}

			res, err := httpClient.Do(req)
			if err != nil {
				// TODO
			}

			return s.buildResult(res)
		} else {
			return nil, errors.NewAuthenticationError()
		}
	} else {
		return nil, errors.NewInvalidHostError()
	}
}

func (s *Service) buildResult(res *http.Result) (*Result, error) {
	input := make(map[string]string)
	output := make(map[string]string)

	data, err := ioutil.ReadAll(res.Body)
	if err != nil {
		// TODO
	}
	
	res.Body.Close()

	json.Unmarshall(input, &data)

	properties := map[string]string{
		"output_ConversationID":      "conversation",
		"output_TransactionID":       "transaction",
		"output_ResponseDesc":        "response",
		"output_ResponseCode":        "code",
		"output_ThirdPartyReference": "reference",
	}

	for oldName, newName := range properties {
		if value, ok := input[oldName]; ok {
			output[newName] = value
		}
	}

	switch res.StatusCode {
	case StatusOk, StatusCreated:
		return s.NewResult(statusCode, output), nil
	case StatusBadRequest:
		if code, ok := output["code"]; ok {
			switch code {
			case INS13:
				return nil, new(InvalidShortcodeError)
			case INS14:
				return nil, new(InvalidReferenceError)
			case INS15:
				return nil, new(InvalidAmountError)
			case INS17:
				return nil, new(InvalidTransactionReferenceError)
			case INS18:
				return nil, new(InvalidTransactionIdError)
			case INS19:
				return nil, new(InvalidThirdPartyReferenceError)
			case INS20:
				return nil, new(InvalidMissingPropertiesError)
			case INS21:
				return nil, new(ValidationError)
			case INS22:
				return nil, new(InvalidOperationPartError)
			case INS23:
				return nil, new(UnknownStatusError)
			case INS24:
				return nil, new(InvalidInitiatorIdentifierError)
			case INS25:
				return nil, new(InvalidSecurityCredentialError)
			case INS993:
				return nil, new(DirectDebtMissingError)
			case INS994:
				return nil, new(DuplicatedDirectDebtError)
			case INS995:
				return nil, new(ProfileProblemsError)
			case INS996:
				return nil, new(InactiveAccountError)
			case INS997:
				return nil, new(InvalidLanguageCodeError)
			case INS998:
				return nil, new(InvalidMarketError)
			case INS2001:
				return nil, new(InitiatorAuthenticationError)
			case INS2002:
				return nil, new(InvalidReceiverError)
			case INS2051:
				return nil, new(InvalidMSISDNError)
			case INS2057:
				return nil, new(InvalidLanguageCodeError)
			}
		}
	case StatusUnauthorized:
		if code, ok := output["code"]; ok {
			switch code {
			case INS5:
				return nil, new(TransactionCancelledError)
			case INS6:
				return nil, new(TransactionFailedError)
			}
		}
	case StatusRequestTimeout:
		return nil, new(RequestTimeoutError)
	case StatusConflict:
		return nil, new(DuplicateTransactionError)
	case StatusUnprocessableEntity:
		return nil, new(InsufficientBalanceError)
	case StatusInternalServerError:
		return nil, new(InternalError)
	case StatusServiceUnavailable:
		return nil, new(UnavailableServerError)
	}

	return nil, new(UnkownError)
}

func (s *Service) requestBody(operation *Operation, intent Intent) map[string]string {
	output := map[string]string{}

	v := reflect.ValueOf(operation)
	for i := 0; i < v.NumField(); i++ {
		field = v.Field(i)

		for k, v := range operation.mapping {
			if newName, ok := operation.mapping[field.Name]; ok {
				switch v.Kind() {
				case reflect.String:
					output[newName] = fmt.Sprintf("%s", field.String())
				case reflect.Float64:
					output[newName] = fmt.Sprintf("%s", field.Float64())
				default:
					// TODO
				}
				
			}
		}
	}

	return output
}