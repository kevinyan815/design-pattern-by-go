package main

import (
	"fmt"
)

const (
	BOOT_ADDRESS = 0
	BOOT_SECTOR  = 0
	SECTOR_SIZE  = 0
)

type CPU struct{}

func (c *CPU) Freeze() {
	fmt.Println("CPU.Freeze()")
}

func (c *CPU) Jump(position int) {
	fmt.Println("CPU.Jump()")
}

func (c *CPU) Execute() {
	fmt.Println("CPU.Execute()")
}

type Memory struct{}

func (m *Memory) Load(position int, data []byte) {
	fmt.Println("Memory.Load()")
}

type HardDrive struct{}

func (hd *HardDrive) Read(lba int, size int) []byte {
	fmt.Println("HardDrive.Read()")
	return make([]byte, 0)
}

type ComputerFacade struct {
	processor *CPU
	ram       *Memory
	hd        *HardDrive
}

func NewComputerFacade() *ComputerFacade {
	return &ComputerFacade{new(CPU), new(Memory), new(HardDrive)}
}

func (c *ComputerFacade) start() {
	c.processor.Freeze()
	c.ram.Load(BOOT_ADDRESS, c.hd.Read(BOOT_SECTOR, SECTOR_SIZE))
	c.processor.Jump(BOOT_ADDRESS)
	c.processor.Execute()
}

func main() {
	computer := NewComputerFacade()
	computer.start()
}
