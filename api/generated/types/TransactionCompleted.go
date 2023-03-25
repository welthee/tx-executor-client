
package types

// TransactionCompleted represents a TransactionCompleted model.
type TransactionCompleted struct {
  Id string
  Details *TransactionDetails
  Context *TxContext
}