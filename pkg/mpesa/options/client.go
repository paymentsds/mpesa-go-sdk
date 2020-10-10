package options

import (
	environment "github.com/paymentsds/mpesa-go-sdk/pkg/mpesa/environment"
)

type ClientOptions struct {
	ApiKey string
	PublicKey string
	Timeout string
	AccessToken string
	Origin string
	VerifySSL string
	Debugging string
	Host string
	Env Environment string
	ServiceProviderCode string
	SecurityCredential string
	InitiatorIdentifier string
}