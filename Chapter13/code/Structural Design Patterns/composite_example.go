package main

import "fmt"

type Car struct {
	wheels []Wheel
	name   string
	engine string
}

func (car *Car) AddWheel(wheel Wheel) {
	car.wheels = append(car.wheels, wheel)

	fmt.Print("adding :")
	wheel.GetType()
}

func (car *Car) GetModel() {
	fmt.Println("The Car model is", car.name)
}

type Wheel interface {
	GetType()
}

type CarWheel struct {
	wheelType string
}

func (wheel *CarWheel) GetType() {
	fmt.Println("The wheel type is", wheel.wheelType)
}

func main() {

	car := new(Car)

	car.name = "Toyota Tercel"

	wheel1 := &CarWheel{wheelType: "Good year"}

	car.AddWheel(wheel1)

	wheel2 := &CarWheel{wheelType: "Good year"}

	car.AddWheel(wheel2)

	wheel3 := &CarWheel{wheelType: "Good year"}

	car.AddWheel(wheel3)

	wheel4 := &CarWheel{wheelType: "Good year"}

	car.AddWheel(wheel4)

	car.GetModel()

	fmt.Println("The number of wheels in the car are :", len(car.wheels))

}
