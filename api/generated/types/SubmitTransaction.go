package types

// SubmitTransaction represents a SubmitTransaction model.
type SubmitTransaction struct {
	Context    *TxContext `json:"context" mapstructure:"context"`
	BinaryHex  string     `json:"binary_hex" mapstructure:"binary_hex"`
	Blockchain string     `json:"blockchain" mapstructure:"blockchain"`
	Priority   *Priority  `json:"priority" mapstructure:"priority"`
	Retry      bool       `json:"retry" mapstructure:"retry"`
}
