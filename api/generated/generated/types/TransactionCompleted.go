package types

// TransactionCompleted represents a TransactionCompleted model.
type TransactionCompleted struct {
	Id      string              `json:"id" mapstructure:"id"`
	Details *TransactionDetails `json:"details" mapstructure:"details"`
	Context *TxContext          `json:"context" mapstructure:"context"`
}
