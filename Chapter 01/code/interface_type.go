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

func main() {

	var veh vehicle
	veh = car{"Toyota Tercel", 1997}
	fmt.Println("Speed of the Vehicle :", veh.getSpeed())
	fmt.Println("Distance Travelled by the Vehicle:", veh.getDistanceTravelled())
}
