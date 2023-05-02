package types

// TransactionSubmitted represents a TransactionSubmitted model.
type TransactionSubmitted struct {
	Id      string     `json:"id" mapstructure:"id"`
	Context *TxContext `json:"context" mapstructure:"context"`
}
