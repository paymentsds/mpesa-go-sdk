package mpesa

import (
	"fmt"
	errors "github.com/paymentsds/mpesa-go-sdk/pkg/mpesa/errors"
	"net/http"
	"time"
	"io/ioutil"
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
				Method: opetation.Method(),
				URL: operation.URL(s.config),
				Headers: s.config.Headers(),
			}

			if req.Method == "GET" {
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
