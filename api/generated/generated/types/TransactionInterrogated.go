package types

// TransactionInterrogated represents a TransactionInterrogated model.
type TransactionInterrogated struct {
	Id      string              `json:"id" mapstructure:"id"`
	Details *TransactionDetails `json:"details" mapstructure:"details"`
	Context *TxContext          `json:"context" mapstructure:"context"`
}
