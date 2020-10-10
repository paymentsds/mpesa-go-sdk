package mpesa

type Configuration struct {
	timeout             int
	apiKey              string
	publicKkey          string
	accessToken         string
	securityCredential  string
	initiatorIdentifier string
	verifySSL           string
	serviceProviderCode string
	origin              string
}

func (c *Configuration) Headers() map[string][]string {
	// TODO
}

func (c *Configuration) GenerateToken() {
	// TODO
}

func (c *Configuration) HasValidHost() bool {
	// TODO
}

func (c *Configuration) HasToken() bool {
	// TODO
}
