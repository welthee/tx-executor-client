
package types

// CommandKind represents an enum of CommandKind.
type CommandKind uint

const (
  CommandKindSubmitTx CommandKind = iota
  CommandKindInterrogateTx
  CommandKindSpeedUpTx
  CommandKindCreateAccount
  CommandKindListAccount
  CommandKindDisableAccount
)

// Value returns the value of the enum.
func (op CommandKind) Value() any {
	if op >= CommandKind(len(CommandKindValues)) {
		return nil
	}
	return CommandKindValues[op]
}

var CommandKindValues = []any{"submitTx","interrogateTx","speedUpTx","createAccount","listAccount","disableAccount"}
var ValuesToCommandKind = map[any]CommandKind{
  CommandKindValues[CommandKindSubmitTx]: CommandKindSubmitTx,
  CommandKindValues[CommandKindInterrogateTx]: CommandKindInterrogateTx,
  CommandKindValues[CommandKindSpeedUpTx]: CommandKindSpeedUpTx,
  CommandKindValues[CommandKindCreateAccount]: CommandKindCreateAccount,
  CommandKindValues[CommandKindListAccount]: CommandKindListAccount,
  CommandKindValues[CommandKindDisableAccount]: CommandKindDisableAccount,
}
