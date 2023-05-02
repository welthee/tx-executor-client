
package types

// Event represents a Event model.
type Event struct {
  Kind *EventKind
  Details interface{}
}