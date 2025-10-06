package main

import (
	"errors"
	"fmt"
)

// buildDesign pattern is a design pattern that allows for flexible configuration of functions or methods.

type personalComputer struct {
	CPU     string
	RAM     string
	Storage string
}

type PcBuilder struct {
	pc personalComputer
}

func NewBuilder() *PcBuilder {
	return &PcBuilder{
		pc: personalComputer{
			CPU:     "default CPU",
			RAM:     "default RAM",
			Storage: "default Storage",
		},
	}
}

func (pcb *PcBuilder) WithCPU(cpu string) *PcBuilder {
	pcb.pc.CPU = cpu
	return pcb
}

func (pcb *PcBuilder) WithRAM(ram string) *PcBuilder {
	pcb.pc.RAM = ram
	return pcb
}

func (pcb *PcBuilder) WithStorage(storage string) *PcBuilder {
	pcb.pc.Storage = storage
	return pcb
}

func (pcb *PcBuilder) Build() (*personalComputer, error) {
	if pcb.pc.CPU == "" || pcb.pc.RAM == "" || pcb.pc.Storage == "" {
		return nil, errors.New("missing required fields")
	}

	return &pcb.pc, nil
}

func main() {
	pcBuilder := NewBuilder().
		WithCPU("Intel i7").
		WithRAM("16GB").
		WithStorage("512GB SSD")
	pc, err := pcBuilder.Build()
	if err != nil {
		fmt.Println("Error building PC:", err)
		return
	}

	fmt.Printf("Built PC: %+v\n", pc)
	fmt.Printf("Type of pc: %T\n", pc)
}
