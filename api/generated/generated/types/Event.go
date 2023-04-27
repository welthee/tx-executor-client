package types

// Event represents a Event model.
type Event struct {
	Kind    *EventKind  `json:"kind" mapstructure:"kind"`
	Details interface{} `json:"details" mapstructure:"details"`
}
