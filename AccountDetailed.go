
package types

// AccountDetailed represents a AccountDetailed model.
type AccountDetailed struct {
  Id string
  Context *AccountContext
  Kind *AccountKind
  Address string
  Active bool
}