// Adapter Pattern allows incompatible interfaces to work together by wrapping an object
// in an adapter to make it compatible with another class.
//
// Use cases:
// - When you want to use an existing class that doesn't fit your interface
// - When you need to reuse existing classes with incompatible interfaces
// - When you want to create a reusable class that cooperates with classes that don't have compatible interfaces

package structural

// Target defines the domain-specific interface used by the client code
type Target interface {
	Request() string
}

// Adaptee contains some useful behavior, but its interface is incompatible
// with the existing client code
type Adaptee struct{}

// SpecificRequest provides the existing behavior with a different interface
func (a *Adaptee) SpecificRequest() string {
	return "Specific request from Adaptee"
}

// Adapter makes Adaptee compatible with the Target interface
type Adapter struct {
	*Adaptee
}

// NewAdapter creates a new Adapter
func NewAdapter() Target {
	return &Adapter{&Adaptee{}}
}

// Request is compatible with the Target interface
func (a *Adapter) Request() string {
	return "Adapter: " + a.SpecificRequest()
}
