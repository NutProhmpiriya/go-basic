// Facade Pattern provides a unified interface to a set of interfaces in a subsystem.
// Facade defines a higher-level interface that makes the subsystem easier to use.
//
// Use cases:
// - When you need to provide a simple interface to a complex subsystem
// - When there are many dependencies between clients and implementation classes
// - When you want to layer your subsystems

package structural

// Complex subsystem components
type CPU struct{}

func (c *CPU) Freeze() string {
	return "CPU: Freezing..."
}

func (c *CPU) Jump(position string) string {
	return "CPU: Jumping to " + position
}

func (c *CPU) Execute() string {
	return "CPU: Executing..."
}

type Memory struct{}

func (m *Memory) Load(position string, data string) string {
	return "Memory: Loading " + data + " to " + position
}

type HardDrive struct{}

func (h *HardDrive) Read(position string, size int) string {
	return "HardDrive: Reading data of size " + string(rune(size)) + " from " + position
}

// ComputerFacade provides a unified interface to a set of interfaces in the subsystem
type ComputerFacade struct {
	cpu       *CPU
	memory    *Memory
	hardDrive *HardDrive
}

// NewComputerFacade creates a new facade
func NewComputerFacade() *ComputerFacade {
	return &ComputerFacade{
		cpu:       &CPU{},
		memory:    &Memory{},
		hardDrive: &HardDrive{},
	}
}

// Start provides a simple interface to the complex subsystem
func (c *ComputerFacade) Start() []string {
	results := make([]string, 0)
	
	results = append(results, c.cpu.Freeze())
	results = append(results, c.memory.Load("0x00", "BOOT_SECTOR"))
	results = append(results, c.cpu.Jump("0x00"))
	results = append(results, c.cpu.Execute())
	
	return results
}
