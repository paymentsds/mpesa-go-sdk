package main

import (
	mpesa "github.com/paymentsds/mpesa-go-sdk/pkg/mpesa"
	environment "github.com/paymentsds/mpesa-go-sdk/pkg/mpesa/environment"
	errors "github.com/paymentsds/mpesa-go-sdk/pkg/mpesa/errors"
	options "github.com/paymentsds/mpesa-go-sdk/pkg/mpesa/options"
)

func main() {
	client := mpesa.NewClient(
		options.WithApiKey("<REPLACE>"),
		options.WithPublicKey("<REPLACE>"),
		options.WithTimeout(10),
		options.WithUserAgent("Paymentsds/M-Pesa"),
		options.WithOrigin("*"),
		options.WithHost(""),
		options.WithEnvironment(environment.SANDBOX),
	)

	intent := mpesa.NewIntent(
		options.WithAmount(10),
		options.WithReference("<REPLACE>"),
		options.WithTransaction("<REPLACE>"),
		options.From("<REPLACE>"),
		options.To("<REPLACE>"),
	)

	if result, err := client.Receive(intent); err != nil {

	}
}
