package main

import "fmt"

type Vehicle interface {
	fly()
}

type LandVehicle struct {
	Vehicle
	name string
}

func (dr *LandVehicle) init() {
	dr.checkBattery()
	dr.checkEngine()
}

func (dr *LandVehicle) drive() {
	fmt.Printf(">>> driving Car[%s] ... \n", dr.name)
	dr.init()
	dr.start()
	dr.monitorHealth()
}

func (dr *LandVehicle) checkBattery() {
	fmt.Println(" checking battery's status ... ")
}

func (dr *LandVehicle) checkEngine() {
	fmt.Println(" checking engine's status ... ")
}

func (dr *LandVehicle) start() {
	fmt.Println("starting now ... ")
}

func (dr *LandVehicle) monitorHealth() {
	fmt.Println("[ everything is ok ... ")
}

func main() {
	dr := &LandVehicle{name: "Car"}
	dr.drive()
}
