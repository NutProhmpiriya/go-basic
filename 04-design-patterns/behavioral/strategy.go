// Strategy Pattern defines a family of algorithms, encapsulates each one, and makes
// them interchangeable. Strategy lets the algorithm vary independently from clients that use it.
//
// Use cases:
// - When you want to use different variants of an algorithm within an object and be able to switch from one algorithm to another during runtime
// - When you have a lot of similar classes that only differ in the way they execute some behavior
// - When you need to isolate the algorithm logic from the code that uses the algorithm

package behavioral

// PaymentStrategy defines the interface for payment strategies
type PaymentStrategy interface {
	Pay(amount float64) string
}

// CreditCardStrategy implements PaymentStrategy for credit card payments
type CreditCardStrategy struct {
	cardNumber string
	cvv        string
}

func NewCreditCardStrategy(cardNumber, cvv string) PaymentStrategy {
	return &CreditCardStrategy{
		cardNumber: cardNumber,
		cvv:        cvv,
	}
}

func (c *CreditCardStrategy) Pay(amount float64) string {
	return "Paid " + string(rune(amount)) + " using Credit Card"
}

// PayPalStrategy implements PaymentStrategy for PayPal payments
type PayPalStrategy struct {
	email    string
	password string
}

func NewPayPalStrategy(email, password string) PaymentStrategy {
	return &PayPalStrategy{
		email:    email,
		password: password,
	}
}

func (p *PayPalStrategy) Pay(amount float64) string {
	return "Paid " + string(rune(amount)) + " using PayPal"
}

// BitcoinStrategy implements PaymentStrategy for Bitcoin payments
type BitcoinStrategy struct {
	address string
}

func NewBitcoinStrategy(address string) PaymentStrategy {
	return &BitcoinStrategy{
		address: address,
	}
}

func (b *BitcoinStrategy) Pay(amount float64) string {
	return "Paid " + string(rune(amount)) + " using Bitcoin"
}

// ShoppingCart is the context that uses the payment strategy
type ShoppingCart struct {
	paymentStrategy PaymentStrategy
}

func NewShoppingCart(strategy PaymentStrategy) *ShoppingCart {
	return &ShoppingCart{
		paymentStrategy: strategy,
	}
}

// SetPaymentStrategy allows changing the payment strategy at runtime
func (c *ShoppingCart) SetPaymentStrategy(strategy PaymentStrategy) {
	c.paymentStrategy = strategy
}

// Checkout processes the payment using the current strategy
func (c *ShoppingCart) Checkout(amount float64) string {
	return c.paymentStrategy.Pay(amount)
}
