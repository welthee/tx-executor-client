package types

// SpeedUpTransaction represents a SpeedUpTransaction model.
type SpeedUpTransaction struct {
	Context    *TxContext `json:"context" mapstructure:"context"`
	Blockchain string     `json:"blockchain" mapstructure:"blockchain"`
	Priority   *Priority  `json:"priority" mapstructure:"priority"`
}
