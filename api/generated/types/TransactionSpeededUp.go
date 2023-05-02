package types

// TransactionSpeededUp represents a TransactionSpeededUp model.
type TransactionSpeededUp struct {
	Id      string     `json:"id" mapstructure:"id"`
	Context *TxContext `json:"context" mapstructure:"context"`
}
