const (
	C2BPayment OperationCode = iota
	B2BPayment
	B2CPayment
	Reversal
	QueryTransactionStatus
)

const (
	Get HttpMethod = iota
	Post
	Put
	Delete
)

const (
	PatternPhoneNumber := ""
	PatternServiceProviderCode := ""
	PatternWord := ""
	PatternMoneyAmount := ""
)

var (
	Operations := map[OperationCode]Operation{
		C2BPayment: &Operation{
		},

		B2BPayment: &Operation{
		},

		B2CPayment: &Operation{
		},

		Reversal: &Operation{
		},

		QueryTransactionStatus: &Operation{
		}
	}
)
