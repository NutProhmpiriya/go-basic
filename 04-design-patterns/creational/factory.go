// Factory Pattern provides an interface for creating objects but lets subclasses
// decide which class to instantiate.
//
// Use cases:
// - When object creation logic is complex
// - When you need to create objects based on conditions
// - When you want to work with objects through a common interface

package creational

// PaymentMethod interface defines the contract for different payment methods
type PaymentMethod interface {
	Pay(amount float64) string
}

// CreditCard implements PaymentMethod
type CreditCard struct{}

func (c *CreditCard) Pay(amount float64) string {
	return "Paid using Credit Card"
}

// DebitCard implements PaymentMethod
type DebitCard struct{}

func (d *DebitCard) Pay(amount float64) string {
	return "Paid using Debit Card"
}

// PayPal implements PaymentMethod
type PayPal struct{}

func (p *PayPal) Pay(amount float64) string {
	return "Paid using PayPal"
}

// PaymentType represents different types of payment methods
type PaymentType int

const (
	CreditCardType PaymentType = iota
	DebitCardType
	PayPalType
)

// PaymentFactory creates payment methods based on the type
func PaymentFactory(paymentType PaymentType) PaymentMethod {
	switch paymentType {
	case CreditCardType:
		return &CreditCard{}
	case DebitCardType:
		return &DebitCard{}
	case PayPalType:
		return &PayPal{}
	default:
		return nil
	}
}
