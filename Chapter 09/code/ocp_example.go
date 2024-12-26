package main

import "fmt"

type vehicle interface {
	getSpeed() float64
	getDistanceTravelled() float64
}

type car struct {
	brand              string
	manufacturing_year int
}

func (Car car) getSpeed() float64 {

	return 120
}

func (Car car) getDistanceTravelled() float64 {

	return 4000
}

type truck struct {
	brand              string
	manufacturing_year int
}

func (Truck truck) getSpeed() float64 {

	return 180
}

func (Truck truck) getDistanceTravelled() float64 {

	return 40000
}

func main() {

	var veh vehicle
	veh = car{"Toyota Tercel", 1997}
	fmt.Println("Car's Speed of the Vehicle :", veh.getSpeed())
	fmt.Println("Car's Distance Travelled by the Vehicle:", veh.getDistanceTravelled())

	veh = truck{"Land Rover", 2001}
	fmt.Println("Truck's Speed of the Vehicle :", veh.getSpeed())
	fmt.Println("Truck's Distance Travelled by the Vehicle:", veh.getDistanceTravelled())
}
