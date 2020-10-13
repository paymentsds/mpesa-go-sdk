package errors

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

const (
	internalError              = "INS-1: Internal Error"
	transactionCancelled       = "INS-5: Transaction cancelled by customer"
	transactionFailed          = "INS-6: Transaction Failed"
	requestTimeout             = "INS-9: Request timeout"
	duplicateTransaction       = "INS-10:	Duplicate Transaction"
	invalidShortCode           = "INS-13: Invalid Shortcode Used"
	invalidReference           = "INS-14: Invalid Reference Used"
	invalidAmount              = "INS-15: Invalid Amount Used"
	unavailableServer          = "INS-16: Unable to handle the request due to a temporary overloading"
	invalidTransactionRef      = "INS-17: Invalid Transaction Reference. Length Should Be Between 1 and 20."
	invalidTransactionId       = "INS-18: Invalid TransactionID Used"
	invalidThirdPartyReference = "INS-19: Invalid ThirdPartyReference Used"
	missingParameters          = "INS-20: Not All Parameters Provided. Please try again."
	invalidParameters          = "INS-21: Parameter validations failed. Please try again."
	invalidOperationType       = "INS-22: Invalid Operation Type"
	unknownStatus              = "INS-23: Unknown Status. Contact M-Pesa Support"
	invalidInitiatorIdentifier = "INS-24: Invalid InitiatorIdentifier Used"
	invalidSecurityCredential  = "INS-25: Invalid SecurityCredential Used"
	directDebtMissing          = "INS-993: Direct Debit Missing"
	duplicatedDirectDebt       = "INS-994: Direct Debit Already Exists"
	profileProblems            = "INS-995: Customer's Profile Has Problems"
	inactiveAccount            = "INS-996: Customer Account Status Not Active"
	invalidLinkingTransaction  = "INS-997: Linking Transaction Not Found"
	invalidMarket              = "INS-998: Invalid Market"
	initiatorAuthentication    = "INS-2001: Initiator authentication error."
	invalidReceiver            = "INS-2002: Receiver invalid."
	insufficientBalance        = "INS-2006: Insufficient balance"
	invalidMSISDN              = "INS-2051: MSISDN invalid."
	invalidLanguageCode        = "INS-2057: Language code invalid."

	INS0    = "INS-0"
	INS1    = "INS-1"
	INS5    = "INS-5"
	INS6    = "INS-6"
	INS9    = "INS-9"
	INS10   = "INS-10"
	INS13   = "INS-13"
	INS14   = "INS-14"
	INS15   = "INS-15"
	INS16   = "INS-16"
	INS17   = "INS-17"
	INS18   = "INS-18"
	INS19   = "INS-19"
	INS20   = "INS-20"
	INS21   = "INS-21"
	INS22   = "INS-22"
	INS23   = "INS-23"
	INS24   = "INS-24"
	INS25   = "INS-25"
	INS993  = "INS-993"
	INS994  = "INS-994"
	INS995  = "INS-995"
	INS996  = "INS-996"
	INS997  = "INS-997"
	INS998  = "INS-998"
	INS2001 = "INS-2001"
	INS2002 = "INS-2002"
	INS2006 = "INS-2006"
	INS2051 = "INS-2051"
	INS2057 = "INS-2057"
)
