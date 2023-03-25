
package types

// AccountKind represents an enum of AccountKind.
type AccountKind uint

const (
  AccountKindPk AccountKind = iota
  AccountKindKms
)

// Value returns the value of the enum.
func (op AccountKind) Value() any {
	if op >= AccountKind(len(AccountKindValues)) {
		return nil
	}
	return AccountKindValues[op]
}

var AccountKindValues = []any{"pk","kms"}
var ValuesToAccountKind = map[any]AccountKind{
  AccountKindValues[AccountKindPk]: AccountKindPk,
  AccountKindValues[AccountKindKms]: AccountKindKms,
}
