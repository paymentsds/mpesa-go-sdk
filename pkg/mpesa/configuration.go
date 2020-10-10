package mpesa

type Configuration struct {
	Timeout             int
	ApiKey              string
	PublicKkey          string
	AccessToken         string
	SecurityCredential  string
	InitiatorIdentifier string
	VerifySSL           string
	ServiceProviderCode string
	Origin              string
}
