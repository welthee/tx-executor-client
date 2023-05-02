
package types

// CreateAccount represents a CreateAccount model.
type CreateAccount struct {
  Context *AccountContext
  Blockchain string
  Kind *AccountKind
  Details interface{}
}