package types

// ModelError represents a ModelError model.
type ModelError struct {
	Code        int    `json:"code" mapstructure:"code"`
	Key         string `json:"key" mapstructure:"key"`
	Message     string `json:"message" mapstructure:"message"`
	ExternalRef string `json:"external_ref" mapstructure:"external_ref"`
}
