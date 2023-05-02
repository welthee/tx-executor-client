package types

// TransactionCanceled represents a TransactionCanceled model.
type TransactionCanceled struct {
	Id      string     `json:"id" mapstructure:"id"`
	Context *TxContext `json:"context" mapstructure:"context"`
}
