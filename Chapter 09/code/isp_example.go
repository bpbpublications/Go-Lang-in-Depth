package main

import "fmt"

type vehicle interface {
	getSpeed() float64
	getDistanceTravelled() float64
	start()
}

type group interface {
	addToTheGroup()
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

func (Car car) start() {

	fmt.Printf(">>> starting the [%s] car ... \n", Car.brand)
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

func (Truck truck) start() {

	fmt.Printf(">>> starting the [%s] Truck ... \n", Truck.brand)
}

func (Truck truck) addToTheGroup(member truck) {

	fmt.Printf(">>> adding another  [%s] Truck to the group ... \n", Truck.brand)
}

func getVehicles() []vehicle {
	return []vehicle{&car{}, &truck{}}
}

func main() {

	var veh vehicle
	veh = car{"Toyota Tercel", 1997}
	fmt.Println("Car's Speed of the Vehicle :", veh.getSpeed())
	fmt.Println("Car's Distance Travelled by the Vehicle:", veh.getDistanceTravelled())

	var truck1 truck

	truck1 = truck{"Land Rover", 2001}
	fmt.Println("Truck's Speed of the Vehicle :", truck1.getSpeed())
	fmt.Println("Truck's Distance Travelled by the Vehicle:", truck1.getDistanceTravelled())

	var truck2 truck
	truck2 = truck{"Hyundai Trunk", 2011}
	fmt.Println("Truck's Speed of the Vehicle :", truck2.getSpeed())
	fmt.Println("Truck's Distance Travelled by the Vehicle:", truck2.getDistanceTravelled())

	truck2.addToTheGroup(truck1)

}
