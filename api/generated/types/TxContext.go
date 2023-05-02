package types

// TxContext represents a TxContext model.
type TxContext struct {
	ExternalRef string `json:"external_ref" mapstructure:"external_ref"`
	AccountRef  string `json:"account_ref" mapstructure:"account_ref"`
}
