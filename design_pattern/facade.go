package main

import "fmt"

// --- Subsystem Components ---

// CPU represents the CPU component
type CPU struct{}

func (c *CPU) Start() {
	fmt.Println("CPU: Starting...")
}

func (c *CPU) Execute() {
	fmt.Println("CPU: Executing instructions...")
}

func (c *CPU) Stop() {
	fmt.Println("CPU: Stopping...")
}

// Memory represents the Memory component
type Memory struct{}

func (m *Memory) Load(address int, data string) {
	fmt.Printf("Memory: Loading data '%s' to address %d\n", data, address)
}

func (m *Memory) SelfTest() {
	fmt.Println("Memory: Running self-test...")
}

// HardDrive represents the HardDrive component
type HardDrive struct{}

func (hd *HardDrive) Read(lba int, size int) string {
	fmt.Printf("HardDrive: Reading %d bytes from LBA %d\n", size, lba)
	return "boot_sector_data" // Simulate reading data
}

func (hd *HardDrive) SpinUp() {
	fmt.Println("HardDrive: Spinning up...")
}

// --- Facade ---

// ComputerFacade provides a simplified interface to the computer's subsystem
type ComputerFacade struct {
	cpu       *CPU
	memory    *Memory
	hardDrive *HardDrive
}

// NewComputerFacade creates a new instance of the ComputerFacade
func NewComputerFacade() *ComputerFacade {
	return &ComputerFacade{
		cpu:       &CPU{},
		memory:    &Memory{},
		hardDrive: &HardDrive{},
	}
}

// StartComputer provides a simplified method to start the computer
func (cf *ComputerFacade) StartComputer() {
	fmt.Println("\nComputerFacade: Starting computer...")
	cf.cpu.Start()
	cf.memory.SelfTest()
	cf.hardDrive.SpinUp()
	bootData := cf.hardDrive.Read(0, 1024)
	cf.memory.Load(0, bootData)
	cf.cpu.Execute()
	fmt.Println("ComputerFacade: Computer started successfully!")
}

// ShutdownComputer provides a simplified method to shut down the computer
func (cf *ComputerFacade) ShutdownComputer() {
	fmt.Println("\nComputerFacade: Shutting down computer...")
	cf.cpu.Stop()
	fmt.Println("ComputerFacade: Computer shut down.")
}

// --- Client Code ---

func main() {
	// The client interacts only with the facade, not the complex subsystem components directly.
	computer := NewComputerFacade()

	computer.StartComputer()
	computer.ShutdownComputer()
}
