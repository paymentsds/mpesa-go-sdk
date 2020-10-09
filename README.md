# M-Pesa SDK for Go

M-Pesa SDK for Go is an unofficial library aiming to help businesses integrating every [M-Pesa](https://developer.mpesa.vm.co.mz) operations to their Go applications.

## Contents

- [Features](#features)
- [Usage](#usage)
   - [Quickstart](#usage/scenario-1)
   - [Receive Money from a Mobile Account](#usage/scenario-1)
   - [Send Money to a Mobile Account](#usage/scenario-2)
   - [Send Money to a Business Account](#usage/scenario-3)
   - [Revert a Transaction](#usage/scenario-4)
   - [Query the Status of a Transaction](#usage/scenario-5)
   - [Examples](#usage/scenario-6)
- [Prerequisites](#prerequisites)
- [Installation](#installation)
   - [Using RubyGems](#installation/scenario-1)
   - [Using Bundler](#installation/scenario-2)
   - [Installation Scenario 3](#installation/scenario-3)
   - [Installation Scenario 4](#installation/scenario-4)
- [Configuration](#configuration)
   - [Configuration Scenario 1](#configuration/scenario-1)
   - [Configuration Scenario 2](#configuration/scenario-2)
   - [Configuration Scenario 3](#configuration/scenario-3)
   - [Configuration Scenario 4](#configuration/scenario-4)
- [Related Projects](#related-projects)
   - [Dependencies](#related-projects/dependencies)
   - [Friends](#related-projects/friends)
   - [Alternatives](#related-projects/alternatives)
- [Contributing](#contributing)
- [Changelog](#changelog)
- [Authors](#authors)
- [Credits](#credits)
- [License](#license)

## Features <a name="features"></a>

- Receive money from a mobile account to a business account
- Send money from a business account to a mobile account
- Send money from a business account to a another business account
- Revert a transaction
- Query the status of a transaction

## Usage <a name="usage"></a>

### Quickstart <a name="#usage/scenario-1"></a>

### Receive Money from a Mobile Account <a name="#usage/scenario-2"></a>

```go

import (
    mpesa "github.com/paymentsds/mpesa-go-sdk/pkg/mpesa"
)

var client mpesa.Client = mpesa.NewClient(mpesa.Configuration{
	ApiKey: "<REPLACE>"
	PublicKey: "<REPLACE>"
	Timeout: 0
	AccessToken: "<REPLACE>"
	ServiceProviderCode: "<REPLACE>"
	SecurityCredential: "<REPLACE>"
	InitiatorIdentifier: "<REPLACE>"
})

var intent mpesa.Intent = mpesa.Intent{
	To: "<REPLACE>"
	From: "<REPLACE>"
	Reference: "<REPLACE>"
	Transaction: "<REPLACE>"
	Amount: 10
}

response, err := client.Receive(intent)
if err != nil {
	// Handle failure scenarion
}

// Handle success scenarion
```

### Query the Status of a Transaction <a name="#usage/scenario-6"></a>

### Examples <a name="usage/scenario-7"></a>

## Prerequisites <a name="prerequisites"></a>

## Installation <a name="installation"></a>

### Using RubyGems <a name="installation/scenario-1"></a>

```bash
go get -u github.com/paymentsds/mpesa-go-sdk
```

## Configuration <a name="configuration"></a>

### Configuration Scenario 1 <a name="configuration/scenario-1"></a>

### Configuration Scenario 2 <a name="configuration/scenario-2"></a>

### Configuration Scenario 3 <a name="configuration/scenario-3"></a>

## Related Projects <a name="related-projects"></a>

### Dependencies <a name="related-projects/dependencies"></a>

#### Production Dependencies

#### Development Dependencies

### Friends <a name="related-projects/friends"></a>

- [M-Pesa SDK for Javascript](https://github.com/paymentsds/mpesa-js-sdk)
- [M-Pesa SDK for PHP](https://github.com/paymentsds/mpesa-php-sdk)
- [M-Pesa SDK for Ruby](https://github.com/paymentsds/mpesa-ruby-sdk)
- [M-Pesa SDK for Python](https://github.com/paymentsds/mpesa-python-sdk)

### Alternatives <a name="related-projects/alternatives"></a>

- [Alternative 1](https://github.com/<username>/<project>)
- [Alternative 2](https://github.com/<username>/<project>)
- [Alternative 3](https://github.com/<username>/<project>)
- [Alternative 4](https://github.com/<username>/<project>)

## Contributing <a name="contributing"></a>

## Changelog <a name="changelog"></a>

## Authors <a name="authors"></a>

## Credits <a name="credits"></a>

## License <a name="license"></a>

Copyright 2020 Paymentsds Developers

Licensed under the Apache License, Version 2.0 (the "License"); you may not use this file except in compliance with the License. You may obtain a copy of the License at

http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software distributed under the License is distributed on an "AS IS" BASIS, WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied. See the License for the specific language governing permissions and limitations under the License.

