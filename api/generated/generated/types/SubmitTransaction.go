
package types

// SubmitTransaction represents a SubmitTransaction model.
type SubmitTransaction struct {
  Context *TxContext
  BinaryHex string
  Blockchain string
  Priority *Priority
  Retry bool
}