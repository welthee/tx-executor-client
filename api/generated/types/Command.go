package types

// Command represents a Command model.
type Command struct {
	Kind    *CommandKind `json:"kind" mapstructure:"kind"`
	Details interface{}  `json:"details" mapstructure:"details"`
}
