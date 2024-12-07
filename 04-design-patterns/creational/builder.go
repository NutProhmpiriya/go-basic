// Builder Pattern separates the construction of a complex object from its representation.
//
// Use cases:
// - When object needs to be created with lots of optional parameters
// - When you need different representations of the same construction process
// - When you want to build objects step by step

package creational

// Computer represents the complex object we're building
type Computer struct {
	CPU       string
	RAM       int
	Storage   int
	GPU       string
	Bluetooth bool
}

// ComputerBuilder interface defines the building steps
type ComputerBuilder interface {
	SetCPU(cpu string) ComputerBuilder
	SetRAM(ram int) ComputerBuilder
	SetStorage(storage int) ComputerBuilder
	SetGPU(gpu string) ComputerBuilder
	SetBluetooth(hasBluetooth bool) ComputerBuilder
	Build() *Computer
}

// concreteComputerBuilder implements ComputerBuilder
type concreteComputerBuilder struct {
	computer *Computer
}

// NewComputerBuilder creates a new builder instance
func NewComputerBuilder() ComputerBuilder {
	return &concreteComputerBuilder{computer: &Computer{}}
}

func (b *concreteComputerBuilder) SetCPU(cpu string) ComputerBuilder {
	b.computer.CPU = cpu
	return b
}

func (b *concreteComputerBuilder) SetRAM(ram int) ComputerBuilder {
	b.computer.RAM = ram
	return b
}

func (b *concreteComputerBuilder) SetStorage(storage int) ComputerBuilder {
	b.computer.Storage = storage
	return b
}

func (b *concreteComputerBuilder) SetGPU(gpu string) ComputerBuilder {
	b.computer.GPU = gpu
	return b
}

func (b *concreteComputerBuilder) SetBluetooth(hasBluetooth bool) ComputerBuilder {
	b.computer.Bluetooth = hasBluetooth
	return b
}

func (b *concreteComputerBuilder) Build() *Computer {
	return b.computer
}

// Director helps in using the builder
type Director struct {
	builder ComputerBuilder
}

// NewDirector creates a new Director
func NewDirector(builder ComputerBuilder) *Director {
	return &Director{builder: builder}
}

// BuildGamingPC builds a gaming PC configuration
func (d *Director) BuildGamingPC() *Computer {
	return d.builder.
		SetCPU("Intel i9").
		SetRAM(32).
		SetStorage(2000).
		SetGPU("RTX 4080").
		SetBluetooth(true).
		Build()
}

// BuildOfficePC builds a basic office PC configuration
func (d *Director) BuildOfficePC() *Computer {
	return d.builder.
		SetCPU("Intel i5").
		SetRAM(16).
		SetStorage(512).
		SetGPU("Integrated").
		SetBluetooth(true).
		Build()
}
