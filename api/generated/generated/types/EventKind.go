package types

// EventKind represents an enum of EventKind.
type EventKind uint

const (
	EventKindSubmitted EventKind = iota
	EventKindInterrogated
	EventKindSpeededup
	EventKindCompleted
)

// Value returns the value of the enum.
func (op EventKind) Value() any {
	if op >= EventKind(len(EventKindValues)) {
		return nil
	}
	return EventKindValues[op]
}

var EventKindValues = []any{"submitted", "interrogated", "speededup", "completed"}
var ValuesToEventKind = map[any]EventKind{
	EventKindValues[EventKindSubmitted]:    EventKindSubmitted,
	EventKindValues[EventKindInterrogated]: EventKindInterrogated,
	EventKindValues[EventKindSpeededup]:    EventKindSpeededup,
	EventKindValues[EventKindCompleted]:    EventKindCompleted,
}
