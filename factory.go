/*
The factory pattern allow us to use a single method to create different types, as long
as those types implement the same methods as the factory interface; in this example,
we have two different types of PCs (laptops and desktops), and each one of them wraps
a 'Computer' type.
*/
package main

import (
	"fmt"
)

type IProduct interface {
	setStock(stock int)
	getStock() int
	setName(name string)
	getName() string
}

type Computer struct {
	name  string
	stock int
}

func (c *Computer) setStock(stock int) {
	c.stock = stock
}

func (c *Computer) setName(name string) {
	c.name = name
}

func (c *Computer) getStock() int {
	return c.stock
}

func (c *Computer) getName() string {
	return c.name
}

type Laptop struct {
	Computer
}

func newLaptop() IProduct {
	return &Laptop{
		Computer: Computer{
			name:  "laptop computer",
			stock: 25,
		},
	}
}

type Desktop struct {
	Computer
}

func newDesktop() IProduct {
	return &Desktop{
		Computer: Computer{
			name:  "desktop computer",
			stock: 35,
		},
	}
}

func GetComputerFactory(computerType string) (IProduct, error) {
	switch computerType {
	case "laptop":
		return newLaptop(), nil
	case "desktop":
		return newDesktop(), nil
	default:
		return nil, fmt.Errorf("invalid %s type (not a computer)", computerType)
	}
}

func printNameAndStock(p IProduct) {
	fmt.Printf("Product name: %s, with stock %d\n", p.getName(), p.getStock())
}

func main() {
	const (
		laptop  = "laptop"
		desktop = "desktop"
	)
	// we can tell there won't be any errors in this example, so we can discard those returns
	laptopPC, _ := GetComputerFactory(laptop)
	desktopPC, _ := GetComputerFactory(desktop)

	printNameAndStock(laptopPC)
	printNameAndStock(desktopPC)
}
