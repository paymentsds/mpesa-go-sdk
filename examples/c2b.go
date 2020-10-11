package main

import (
	mpesa "github.com/paymentsds/mpesa-go-sdk/pkg/mpesa"
	environment "github.com/paymentsds/mpesa-go-sdk/pkg/mpesa/environment"
	options "github.com/paymentsds/mpesa-go-sdk/pkg/mpesa/options"
	errors "github.com/paymentsds/mpesa-go-sdk/pkg/mpesa/errors"
)

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
	switch err {

	}
}