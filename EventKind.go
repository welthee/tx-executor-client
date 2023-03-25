
package types

// EventKind represents an enum of EventKind.
type EventKind uint

const (
  EventKindSubmitted EventKind = iota
  EventKindInterrogated
  EventKindSpeededup
  EventKindCompleted
  EventKindAccountCreated
  EventKindAccountDetailed
  EventKindAccountDisabled
)

// Value returns the value of the enum.
func (op EventKind) Value() any {
	if op >= EventKind(len(EventKindValues)) {
		return nil
	}
	return EventKindValues[op]
}

var EventKindValues = []any{"submitted","interrogated","speededup","completed","accountCreated","accountDetailed","accountDisabled"}
var ValuesToEventKind = map[any]EventKind{
  EventKindValues[EventKindSubmitted]: EventKindSubmitted,
  EventKindValues[EventKindInterrogated]: EventKindInterrogated,
  EventKindValues[EventKindSpeededup]: EventKindSpeededup,
  EventKindValues[EventKindCompleted]: EventKindCompleted,
  EventKindValues[EventKindAccountCreated]: EventKindAccountCreated,
  EventKindValues[EventKindAccountDetailed]: EventKindAccountDetailed,
  EventKindValues[EventKindAccountDisabled]: EventKindAccountDisabled,
}
