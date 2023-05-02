package types

// InterrogateTransaction represents a InterrogateTransaction model.
type InterrogateTransaction struct {
	Blockchain string     `json:"blockchain" mapstructure:"blockchain"`
	Context    *TxContext `json:"context" mapstructure:"context"`
}
