package mpesa

type Service struct {
	configuration Configuration
}

func NewService(options ClientOptions) *Service {
	return &Service{
		//configuration: NewConfiguration(options)
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
		// Panic
	}

	intent := s.fillOptionalValues(operation, intent)

	if missingProperties, err := operation.detectMissingProperties(intent); err != nil {

	}

	if invalidProperties, err := operation.validateProperties(intent); err != nil {

	}

	return s.performRequest(operation, intent)
}

func (s *Service) detectOperation(intent Intent) (OperationCode, error) {

}

func (s *Service) fillOptionalValues(opcode Operationcode, intent Intent) Intent {

}

func (s *Service) performRequest(opcode OperationCode, intent Intent) (Result, error) {

}
