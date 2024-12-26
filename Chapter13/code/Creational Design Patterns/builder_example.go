package main

import (
	"fmt"
	"strconv"
)

type Truck struct {
	Tyres    int
	Capacity int
}

type TruckBuildProcess interface {
	SetTyreNumber() TruckBuildProcess
	SetCapacity() TruckBuildProcess
	GetVehicle() Truck
}

type TruckShopFloor struct {
	tbuilder TruckBuildProcess
}

func (f *TruckShopFloor) SetBuilder(tb TruckBuildProcess) {
	f.tbuilder = tb
}

func (f *TruckShopFloor) Construct() {
	f.tbuilder.SetCapacity().SetTyreNumber()
}

type TruckBuilder struct {
	t Truck
}

func (c *TruckBuilder) SetTyreNumber() TruckBuildProcess {
	c.t.Tyres = 6
	return c
}

func (c *TruckBuilder) SetCapacity() TruckBuildProcess {
	c.t.Capacity = 4
	return c
}

func (c *TruckBuilder) GetVehicle() Truck {
	return c.t
}

func (c *TruckBuilder) Build() Truck {
	return c.t
}

func main() {

	shopFloor := TruckShopFloor{}

	truckBuilder := &TruckBuilder{}
	shopFloor.SetBuilder(truckBuilder)
	shopFloor.Construct()
	truck := truckBuilder.Build()
	if truck.Tyres != 6 {
		fmt.Errorf("It is wrong that : " + strconv.Itoa(truck.Tyres) + " tyres found")
	} else {

		fmt.Printf("Truck has " + strconv.Itoa(truck.Tyres) + "  tyres\n")
	}

}
