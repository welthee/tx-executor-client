
package types

// TransactionInterrogated represents a TransactionInterrogated model.
type TransactionInterrogated struct {
  Id string
  Details *TransactionDetails
  Context *TxContext
}