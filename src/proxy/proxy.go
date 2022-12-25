package main

import "fmt"

type Vehicle interface {
	Drive()
}

type Car struct{}

func (c *Car) Drive() {
	fmt.Println("Car is being driven")
}

type Driver struct {
	Age int
}

type CarProxy struct {
	vehicle    Vehicle
	driver *Driver
}

func NewCarProxy(driver *Driver) *CarProxy {
	return &CarProxy{&Car{}, driver}
}

func (c *CarProxy) Drive() {
	if c.driver.Age >= 16 {
		c.vehicle.Drive()
	} else {
		fmt.Println("Driver too young!")
	}
}

func main() {
	car := NewCarProxy(&Driver{12})
	car.Drive()
	car2 := NewCarProxy(&Driver{22})
	car2.Drive()
}
