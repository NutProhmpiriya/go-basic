// Decorator Pattern attaches additional responsibilities to an object dynamically.
// Decorators provide a flexible alternative to subclassing for extending functionality.
//
// Use cases:
// - When you need to add responsibilities to objects dynamically
// - When extension by subclassing is impractical
// - When you want to add behaviors without altering existing code

package structural

// Coffee defines the interface for coffee types
type Coffee interface {
	GetCost() float64
	GetDescription() string
}

// SimpleCoffee implements the base coffee
type SimpleCoffee struct{}

func (c *SimpleCoffee) GetCost() float64 {
	return 1.0
}

func (c *SimpleCoffee) GetDescription() string {
	return "Simple coffee"
}

// CoffeeDecorator provides the base decorator structure
type CoffeeDecorator struct {
	coffee Coffee
}

// MilkDecorator adds milk to coffee
type MilkDecorator struct {
	CoffeeDecorator
}

func NewMilkDecorator(c Coffee) Coffee {
	return &MilkDecorator{CoffeeDecorator{coffee: c}}
}

func (m *MilkDecorator) GetCost() float64 {
	return m.coffee.GetCost() + 0.5
}

func (m *MilkDecorator) GetDescription() string {
	return m.coffee.GetDescription() + ", milk"
}

// SugarDecorator adds sugar to coffee
type SugarDecorator struct {
	CoffeeDecorator
}

func NewSugarDecorator(c Coffee) Coffee {
	return &SugarDecorator{CoffeeDecorator{coffee: c}}
}

func (s *SugarDecorator) GetCost() float64 {
	return s.coffee.GetCost() + 0.2
}

func (s *SugarDecorator) GetDescription() string {
	return s.coffee.GetDescription() + ", sugar"
}

// WhipDecorator adds whipped cream to coffee
type WhipDecorator struct {
	CoffeeDecorator
}

func NewWhipDecorator(c Coffee) Coffee {
	return &WhipDecorator{CoffeeDecorator{coffee: c}}
}

func (w *WhipDecorator) GetCost() float64 {
	return w.coffee.GetCost() + 0.7
}

func (w *WhipDecorator) GetDescription() string {
	return w.coffee.GetDescription() + ", whip"
}
