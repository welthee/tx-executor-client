
package types

// Priority represents an enum of Priority.
type Priority uint

const (
  PriorityLow Priority = iota
  PriorityStandard
  PriorityFast
)

// Value returns the value of the enum.
func (op Priority) Value() any {
	if op >= Priority(len(PriorityValues)) {
		return nil
	}
	return PriorityValues[op]
}

var PriorityValues = []any{"low","standard","fast"}
var ValuesToPriority = map[any]Priority{
  PriorityValues[PriorityLow]: PriorityLow,
  PriorityValues[PriorityStandard]: PriorityStandard,
  PriorityValues[PriorityFast]: PriorityFast,
}
