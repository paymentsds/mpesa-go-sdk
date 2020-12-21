package main

import (
	mpesa "github.com/paymentsds/mpesa-go-sdk/pkg/mpesa"
	client "github.com/paymentsds/mpesa-go-sdk/pkg/mpesa/client"
	env "github.com/paymentsds/mpesa-go-sdk/pkg/mpesa/env"
	intent "github.com/paymentsds/mpesa-go-sdk/pkg/mpesa/intent"
)

func main() {
	client := mpesa.NewClient(
		client.WithApiKey(""),
		client.WithPublicKey(""),
		client.WithConnectionTimeout(10),
		client.WithReadTimeout(10),
		client.WithUserAgent(""),
		client.WithOrigin(""),
		client.WithHost(""),
		client.WithEnvironment(env.Sandbox),
		client.WithInitiatorIdentifier(""),
		client.WithSecurityCredential(""),
		client.WithServiceProviderCode("")
	)

	intent := mpesa.NewIntent(
		intent.WithAmount(10.0),
		intent.WithReversalAmount(10.0),
		intent.WithThirdPartyReference(""),
		intent.WithTransactionReference(""),
		intent.Source(""),
		intent.Destination(""),
		intent.TransactionId(""),
	)

	if result, err := client.Receive(intent); err != nil {

	}
}
