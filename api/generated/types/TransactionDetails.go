package types

// TransactionDetails represents a TransactionDetails model.
type TransactionDetails struct {
	Status string `json:"status" mapstructure:"status"`
	Hash   string `json:"hash" mapstructure:"hash"`
}
